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
	fmt.Printf("Starting vault logrotation with \"%s\"\n", crontab)
	crond.Cron(crontab).Do(run_logrotate)
	crond.StartBlocking()
}

func run_logrotate() {
	cmd := exec.Command("/usr/sbin/logrotate", "--state=/tmp/logrotate.status", "/etc/logrotate.conf")
	fmt.Printf("Starting logrotation\n")
	if err := cmd.Run(); err != nil {
		fmt.Printf("Finished logrotation with error: %v", err)
	} else {
		fmt.Println("Finished logrotation")
	}
}
