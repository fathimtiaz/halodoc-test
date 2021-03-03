package main

import (
	"fmt"
	"sync"

	"local/halodoc/config"
	"local/halodoc/job"
)

func main() {
	var (
		cfg config.Config
		err error
	)

	wg := sync.WaitGroup{}
	wg.Add(1)

	if cfg, err = config.Get(); err != nil {
		fmt.Println(err.Error())
	}

	job.InitAlert(cfg)

	wg.Wait()

}
