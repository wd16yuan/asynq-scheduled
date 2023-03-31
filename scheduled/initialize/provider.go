package initialize

import (
	"encoding/json"
	"errors"
	"fmt"
	"padmin/scheduled/global"
	"padmin/scheduled/model"
	"time"

	"github.com/hibiken/asynq"
	"github.com/robfig/cron"
	"gorm.io/gorm"
)

type DBBasedConfigProvider struct {
	Total               int64
	UpdatedAt           time.Time
	CacheTaskConfigList []*asynq.PeriodicTaskConfig
}

func (p *DBBasedConfigProvider) GetConfigs() ([]*asynq.PeriodicTaskConfig, error) {
	updated, tasks, err := p.getFromDB()
	// 发生错误及未触发更新，都从缓存获取
	if err != nil || !updated {
		return p.CacheTaskConfigList, err
	}
	var taskConfigList []*asynq.PeriodicTaskConfig
	var total int64
	var index int
	for i, t := range tasks {
		// 检查定时任务表达式格式，任务参数
		spec, err := p.check(&t)
		if err != nil {
			p.disable(&t, err)
			continue
		}
		taskConfigList = append(taskConfigList, &asynq.PeriodicTaskConfig{
			Cronspec: spec,
			Task:     p.newTask(&t),
			Opts:     []asynq.Option{asynq.TaskID(fmt.Sprintf("%d||%s", t.ID, t.Name))},
		})
		total += 1
		index = i
	}
	p.Total = total
	p.UpdatedAt = tasks[index].UpdatedAt
	p.CacheTaskConfigList = taskConfigList
	return p.CacheTaskConfigList, nil
}

func (p *DBBasedConfigProvider) getFromDB() (bool, []model.SysPeriodicTask, error) {
	var updatedTask model.SysPeriodicTask
	var tasks []model.SysPeriodicTask
	var total int64
	var updated bool

	// 满足启用状态
	// 满足未到期，或未设置过期时间
	// 满足未设置执行一次，否则运行次数小于一次
	currentDB := global.GVA_DB.Model(&model.SysPeriodicTask{}).Where("enabled = ?", true).
		Where(global.GVA_DB.Where("expires > ?", time.Now()).Or(global.GVA_DB.Where("expires is null"))).
		Where(global.GVA_DB.Where("one_off = ?", false).Or(global.GVA_DB.Where("one_off is null")).
			Or(global.GVA_DB.Where("one_off = ?", true).Where("total_run_count < ?", 1)))

	err := currentDB.Count(&total).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return updated, tasks, err
	}
	err = currentDB.Session(&gorm.Session{}).Order("updated_at desc").Take(&updatedTask).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return updated, tasks, err
	}

	if total != p.Total || updatedTask.UpdatedAt.After(p.UpdatedAt) {
		err = currentDB.Order("updated_at").Preload("SysCrontab").Preload("SysInterval").Find(&tasks).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return updated, tasks, err
		}
		updated = true
	}
	return updated, tasks, nil
}

func (p *DBBasedConfigProvider) check(t *model.SysPeriodicTask) (string, error) {
	var spec string
	if t.SysCrontabID != nil {
		spec = fmt.Sprintf("%s %s %s %s %s", t.SysCrontab.Minute, t.SysCrontab.Hour, t.SysCrontab.Day, t.SysCrontab.Month, t.SysCrontab.Week)
	} else if t.SysIntervalID != nil {
		spec = fmt.Sprintf("@every %d%s", t.SysInterval.Every, t.SysInterval.Period)
	} else {
		return spec, errors.New("未知的定时任务表达式")
	}
	_, err := cron.ParseStandard(spec)
	if err != nil {
		global.GVA_LOG.Error(fmt.Sprintf("定时任务表达式错误：%s %s（%s）", err.Error(), spec, t.Name))
		return spec, errors.New("定时任务表达式错误")
	}
	if t.Payload != "" {
		if !json.Valid([]byte(t.Payload)) {
			global.GVA_LOG.Error(fmt.Sprintf("参数不是有效JSON格式：%s（%s）", t.Payload, t.Name))
			return spec, errors.New("参数不是有效JSON格式")
		}
	}
	return spec, nil
}

func (p *DBBasedConfigProvider) disable(t *model.SysPeriodicTask, err error) {
	gErr := global.GVA_DB.Model(&model.SysPeriodicTask{}).Select("enabled").Omit("updated_at").Where("id = ? AND name = ?", t.ID, t.Name).
		Updates(map[string]interface{}{
			"enabled": false,
			"prompt":  err.Error(),
		}).Error
	if gErr != nil {
		global.GVA_LOG.Error(gErr.Error())
	}
}

func (p *DBBasedConfigProvider) newTask(t *model.SysPeriodicTask) *asynq.Task {
	var payload []byte
	var retry int
	var retention time.Duration
	if t.Payload != "" {
		payload = []byte(t.Payload)
	}
	if t.MaxRetry != 0 {
		retry = t.MaxRetry
	} else {
		retry = global.GVA_CONFIG.System.Retry
	}
	if t.Retention != 0 {
		retention = t.Retention
	} else {
		retention = global.GVA_CONFIG.System.Retention
	}
	return asynq.NewTask(t.TaskFunc, payload, asynq.MaxRetry(retry), asynq.Retention(retention*time.Minute))
}
