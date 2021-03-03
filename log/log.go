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

	file, err := os.Open(t.File)
	if err != nil {
		return err
	}
	defer file.Close()

	startTime := time.Now()
	count := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		logTimeStr := strings.Split(scanner.Text(), " ")[0]
		logTime, err := time.Parse(time.RFC3339, logTimeStr)

		if err != nil {
			log.Println("cannot read time")
		}

		if startTime.Sub(logTime) < time.Second*time.Duration(t.Duration) {
			count++
		}

		if count >= t.Frequency {
			fmt.Println("notify")
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	t.LastLine++

	return nil
}
