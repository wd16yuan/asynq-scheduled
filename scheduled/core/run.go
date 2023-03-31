package core

import (
	"padmin/scheduled/core/internal"
	"padmin/scheduled/global"
	"padmin/scheduled/hook"
	"padmin/scheduled/initialize"
	"time"

	"github.com/hibiken/asynq"
)

func Run() {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}
	clientOpt := asynq.RedisClientOpt{Addr: global.GVA_CONFIG.Redis.Addr, DB: global.GVA_CONFIG.Redis.DB}
	provider := &initialize.DBBasedConfigProvider{}
	mgr, err := asynq.NewPeriodicTaskManager(
		asynq.PeriodicTaskManagerOpts{
			RedisConnOpt:               clientOpt,
			PeriodicTaskConfigProvider: provider,                                        // 配置源的接口
			SyncInterval:               global.GVA_CONFIG.System.Interval * time.Second, // 指定同步频率（同步配置源）
			SchedulerOpts: &asynq.SchedulerOpts{
				Location:        loc,
				Logger:          internal.NewLogger(global.GVA_LOG),
				PostEnqueueFunc: hook.PostEnqueueFunc,
				PreEnqueueFunc:  hook.PreEnqueueFunc},
		})
	if err != nil {
		panic(err)
	}
	global.GVA_INSPECTOR = asynq.NewInspector(clientOpt)
	if err := mgr.Run(); err != nil {
		global.GVA_LOG.Error(err.Error())
	}
}
