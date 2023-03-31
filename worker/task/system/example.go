package system

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/hibiken/asynq"
)

const (
	CleanJwt          = "clean:JWT"
	CleanLog          = "clean:sys_operation_records"
	CollectTapData    = "collect:tap_data"
	CollectActiveData = "collect:active_data"
)

func HandleCleanJwtTask(ctx context.Context, t *asynq.Task) error {
	res := []byte("task result data")
	_, err := t.ResultWriter().Write(res)
	fmt.Println(err)

	var parameter map[string]interface{}
	err = json.Unmarshal(t.Payload(), &parameter)
	if err != nil {
		return err
	}
	fmt.Printf("清理Jwt中... %s  %s \n", string(t.Payload()), time.Now())
	fmt.Println(parameter)
	return nil
}

func HandleCleanOperationLogTask(ctx context.Context, t *asynq.Task) error {
	fmt.Println("执行清理操作日志...")
	return nil
}

func HandleCollectTapDataTask(ctx context.Context, t *asynq.Task) error {
	fmt.Println("执行收集Tap榜单数据...")
	return nil
}

func HandleActiveDataTask(ctx context.Context, t *asynq.Task) error {
	fmt.Println("执行统计项目活跃数据...")
	return nil
}
