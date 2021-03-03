package job

import (
	"fmt"
	"local/halodoc/config"
	"local/halodoc/log"

	cron "github.com/robfig/cron/v3"
)

func InitAlert(cfg config.Config) {
	c := cron.New()

	for i := range cfg.LogTypes {
		logType := cfg.LogTypes[i]
		logType.Users = getUsers(cfg.Users, logType.Name)

		checkForAlert(&logType)
	}

	c.Start()
}

func checkForAlert(logType *log.Type) {
	if err := logType.Read(); err != nil {
		fmt.Println(err.Error())
	}
}
