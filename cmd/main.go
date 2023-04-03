package main

import (
	"os"

	"git.brilliantmonkey.net/frigate/frigate-clips/config"
	"git.brilliantmonkey.net/frigate/frigate-clips/http"
	"github.com/brilliant-monkey/go-app"
)

func createOutputDirectory(config *config.AppConfig) {
	err := os.MkdirAll(config.FFMPEG.OutputDir, 0750)
	if err != nil {
		panic(err)
	}
}

func main() {
	a := app.NewApp()

	var appConfig config.AppConfig
	a.LoadConfig("FRIGATE_CLIPS_CONFIG_PATH", &appConfig)

	createOutputDirectory(&appConfig)

	apiServer := http.NewAPIServer(&appConfig)
	a.Go(func() error {
		return apiServer.Start()
	})

	// kafkaClient := kafka.NewKafkaClient(&appConfig.Kafka)
	// consumer := events.NewFrigateEventConsumer(&appConfig, kafkaClient)
	// a.Go(consumer.Consume)

	a.Start(func() error {
		return apiServer.Stop()
	})
}
