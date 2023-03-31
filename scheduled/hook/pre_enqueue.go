package hook

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/hibiken/asynq"
)

func PreEnqueueFunc(task *asynq.Task, opts []asynq.Option) {
	// 每次运行重新生成taskID
	for i, opt := range opts {
		if opt.Type() == asynq.TaskIDOpt {
			prefix := opt.Value().(string)
			e := strings.Split(prefix, "||")
			option := asynq.TaskID(fmt.Sprintf("%s||%s||%s", e[0], e[1], uuid.NewString()))
			opts[i] = option
		}
	}
}
