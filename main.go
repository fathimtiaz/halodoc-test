package main

import (
	"fmt"

	"local/halodoc/config"
	"local/halodoc/job"
)

func main() {
	var (
		cfg config.Config
		err error
	)

	if cfg, err = config.Get(); err != nil {
		fmt.Println(err.Error())
	}

	job.InitAlert(cfg)
}
