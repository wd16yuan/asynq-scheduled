package middleware

import (
	"context"
	"fmt"
	"padmin/worker/global"

	"github.com/hibiken/asynq"
	"go.uber.org/zap"
)

func LoggingMiddleware(h asynq.Handler) asynq.Handler {
	return asynq.HandlerFunc(func(ctx context.Context, t *asynq.Task) error {
		name := t.Type()
		global.GVA_LOG.Info(fmt.Sprintf("Start processing %s", name))
		err := h.ProcessTask(ctx, t)
		if err != nil {
			global.GVA_LOG.Error(fmt.Sprintf("Failure processing %s", name), zap.Any("err", err))
			return err
		}
		global.GVA_LOG.Info(fmt.Sprintf("Finished processing %s", name))
		return nil
	})
}
