package main

import (
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/robfig/cron"
	"github.com/urfave/cli"
)

const (
	configFile = "/usr/local/bin/curator-config.yaml"
	actionFile = "/usr/local/bin/clean-up-indices.yaml"
)

var (
	cronDesc string
)

func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "cron-desc",
			EnvVar:      "CRON_DESC",
			Value:       "0 0 * * * * ",
			Destination: &cronDesc,
		},
	}

	app.Action = action
	app.Run(os.Args)
}

func action(ctx *cli.Context) (err error) {
	log.Println("es-curator executed")

	defer log.Println("es-curator exit")

	c := cron.New()
	if err := c.AddFunc(cronDesc, cleanUp); err != nil {
		panic(err)
	}
	c.Start()

	log.Println("es-curator started")

	for {
		time.Sleep(24 * time.Hour)
	}
}

func cleanUp() {

	if _, err := os.Stat("/usr/bin/curator"); err != nil {
		log.Println(err)
	}

	if err := exec.Command("/usr/bin/curator", "--config", configFile, actionFile).Run(); err != nil {
		log.Println(err)
	}
}
