package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/azimjohn/jprq/server/config"
	"github.com/azimjohn/jprq/server/github"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	var (
		conf config.Config
		jprq Jprq
	)

	err := conf.Load()
	if err != nil {
		log.Fatalf("failed to load conf: %v", err)
	}

	oauth := github.New(conf.GithubClientID, conf.GithubClientSecret)

	err = jprq.Init(conf, oauth)
	if err != nil {
		log.Fatalf("failed to init jprq %v", err)
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	jprq.Start()
	defer jprq.Stop()

	<-signalChan
}
