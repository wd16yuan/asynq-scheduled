package initialize

import (
	"padmin/worker/task/system"

	"github.com/hibiken/asynq"
)

func Handles(mux *asynq.ServeMux) {
	mux.HandleFunc(system.CleanJwt, system.HandleCleanJwtTask)

	mux.HandleFunc(system.CleanLog, system.HandleCleanOperationLogTask)
	mux.HandleFunc(system.CollectTapData, system.HandleCollectTapDataTask)
	mux.HandleFunc(system.CollectActiveData, system.HandleActiveDataTask)
}
