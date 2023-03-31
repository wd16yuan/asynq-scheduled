package model

import (
	"padmin/scheduled/global"
	"time"
)

type SysCrontab struct {
	global.GVA_MODEL
	Minute      string `json:"minute" gorm:"comment:分"`
	Hour        string `json:"hour" gorm:"comment:时"`
	Day         string `json:"day" gorm:"comment:天"`
	Month       string `json:"month" gorm:"comment:月"`
	Week        string `json:"week" gorm:"comment:周"`
	Description string `json:"description" gorm:"comment:描述"`
}

type SysInterval struct {
	global.GVA_MODEL
	Every       int    `json:"every" gorm:"comment:时间值"`
	Period      string `json:"period" gorm:"comment:时间单位"` // 时间单位s,m,h
	Description string `json:"description" gorm:"comment:描述"`
}

type SysPeriodicTask struct {
	global.GVA_MODEL
	Name          string        `json:"name" gorm:"comment:名称"`
	TaskFunc      string        `json:"taskFunc" gorm:"comment:任务方法"`
	Payload       string        `json:"payload" gorm:"comment:参数"`
	OneOff        bool          `json:"oneOff" gorm:"comment:是否只运行一次"`
	TotalRunCount int           `json:"totalRunCount" gorm:"comment:已运行次数;default:0"`
	LastRunAt     *time.Time    `json:"lastRunAt" gorm:"comment:上次运行时间"`
	Expires       *time.Time    `json:"expires" gorm:"comment:过期时间"`
	MaxRetry      int           `json:"maxRetry" gorm:"comment:重试次数"`
	Retention     time.Duration `json:"retention" gorm:"comment:任务保留时间（单位：分钟）"`
	Enabled       bool          `json:"enabled" gorm:"comment:启用"`
	Prompt        string        `json:"prompt" gorm:"comment:提示"`
	Description   string        `json:"description" gorm:"comment:描述"`
	SysCrontab    SysCrontab    `json:"crontab"`
	SysInterval   SysInterval   `json:"interval"`
	SysCrontabID  *uint         `json:"crontabID"`
	SysIntervalID *uint         `json:"intervalID"`
}
