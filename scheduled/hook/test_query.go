package hook

import (
	"fmt"
	"padmin/scheduled/global"

	"github.com/hibiken/asynq"
)

func TestQueryTaskResult(info *asynq.TaskInfo, err error) {
	// 仅限设置过保留时间的任务
	inerInfo, _ := global.GVA_INSPECTOR.GetTaskInfo(info.Queue, info.ID)
	fmt.Println("Result: ", string(inerInfo.Result))
}
