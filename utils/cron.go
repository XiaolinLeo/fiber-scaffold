package utils

import (

	"github.com/robfig/cron/v3"
)

func SetupCron() {
	c := cron.New()
	c.AddFunc("@every 100s", func() {
		return
	})

	c.AddFunc("0 * * * * *", func() {
		return
	})

	c.Start()

}
