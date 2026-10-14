package models

import (
	"time"

	"github.com/goravel/framework/database/orm"
)

type Result struct {
	orm.Model
	TargetID  uint
	Target    *Target
	HttpCode  int
	Latency   time.Duration
	Err       string `gorm:"type:text;column:err"`
	CheckedAt time.Time
}
