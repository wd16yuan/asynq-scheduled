package hook

import (
	"padmin/scheduled/global"
	"padmin/scheduled/model"
	"strings"

	"github.com/hibiken/asynq"
	"gorm.io/gorm"
)

func PostEnqueueFunc(info *asynq.TaskInfo, err error) {
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		return
	}
	e := strings.Split(info.ID, "||")
	if len(e) != 3 {
		return
	}
	// 禁止更新updated_at
	err = global.GVA_DB.Model(&model.SysPeriodicTask{}).Select("total_run_count", "last_run_at").Omit("updated_at").Where("id = ? AND name = ?", e[0], e[1]).
		Updates(map[string]interface{}{
			"total_run_count": gorm.Expr("total_run_count + 1"),
			"last_run_at":     info.NextProcessAt,
		}).Error
	if err != nil {
		global.GVA_LOG.Error(err.Error())
	}
}
