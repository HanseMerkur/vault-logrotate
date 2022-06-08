package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/go-co-op/gocron"
)

func main() {
	crond := gocron.NewScheduler(time.UTC)
	var crontab string
	var ok bool
	if crontab, ok = os.LookupEnv("CRONTAB"); !ok {
		crontab = "0 * * * *"
	}
	fmt.Printf("Starting logrotate with \"%s\"", crontab)
	crond.Cron(crontab).Do(run_logrotate)
	crond.StartBlocking()
}

func run_logrotate() {
	cmd := exec.Command("/usr/sbin/logrotate", "/etc/logrotate.conf")
	fmt.Printf("Starting logration")
	if err := cmd.Run(); err != nil {
		fmt.Printf("Finished logration with error: %v", err)
	}
}
