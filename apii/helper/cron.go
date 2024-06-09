package helper

import (
	"github.com/robfig/cron/v3"
)

func NewCronHelper() *cron.Cron {
	return cron.New()
}
