package models

import (
	"time"

	"github.com/goravel/framework/database/orm"
)

type Target struct {
	orm.Model

	Name     string        `json:"name"`
	Url      string        `json:"url"`
	Timeout  time.Duration `json:"timeout"`
	Interval time.Duration `json:"interval"`
	Results  []*Result
}
