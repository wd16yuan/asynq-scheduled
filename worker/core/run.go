package core

import (
	"padmin/worker/core/internal"
	"padmin/worker/global"
	"padmin/worker/initialize"
	"padmin/worker/middleware"

	"github.com/hibiken/asynq"
)

func Run() {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: global.GVA_CONFIG.Redis.Addr, DB: global.GVA_CONFIG.Redis.DB},
		asynq.Config{Concurrency: global.GVA_CONFIG.System.Concurrency, Logger: internal.NewLogger(global.GVA_LOG)},
	)
	mux := asynq.NewServeMux()
	mux.Use(middleware.LoggingMiddleware)
	initialize.Handles(mux)

	if err := srv.Run(mux); err != nil {
		global.GVA_LOG.Error(err.Error())
	}
}
