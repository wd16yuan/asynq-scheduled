package global

import (
	"padmin/scheduled/config"

	"github.com/hibiken/asynq"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	GVA_DB        *gorm.DB
	GVA_CONFIG    config.Server
	GVA_VP        *viper.Viper
	GVA_LOG       *zap.Logger
	GVA_INSPECTOR *asynq.Inspector // 查询任务结果，前提未超过任务配置的保留时间
)
