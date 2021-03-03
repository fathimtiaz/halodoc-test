package log

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type Type struct {
	Name      string
	File      string
	Frequency int
	Duration  int
	WaitTime  int
	LastLine  int
	Users     []string
}

func (t *Type) Read() error {

	now := time.Now()
	intervalTime := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Now().Location())

	for intervalTime.Before(time.Now()) {

		file, err := os.Open(t.File)
		if err != nil {
			return err
		}
		defer file.Close()

		count := 0

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			logTimeStr := strings.Split(scanner.Text(), " ")[0]
			logTime, err := time.Parse(time.RFC3339, logTimeStr)
			if err != nil {
				log.Println("cannot read time")
			}

			logTypeStr := strings.Split(scanner.Text(), " ")[1]
			sub := intervalTime.Sub(logTime)

			if logTypeStr == t.Name && sub > 0 && sub < time.Second*time.Duration(t.Duration) {
				count++
			}
		}

		if count >= t.Frequency {
			fmt.Println("notify", t.Users)
		}

		time.Sleep(time.Second * time.Duration(t.Duration))

		intervalTime = intervalTime.Add(time.Second * time.Duration(t.Duration))
	}

	return nil
}
