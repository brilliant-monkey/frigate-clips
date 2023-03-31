package events

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"git.brilliantmonkey.net/frigate/frigate-clips/config"
	"git.brilliantmonkey.net/frigate/frigate-clips/ffmpeg"
	"git.brilliantmonkey.net/frigate/frigate-clips/types"
	"github.com/brilliant-monkey/go-kafka-client"
)

type FrigateEventConsumer struct {
	config *config.AppConfig
	client *kafka.KafkaClient
}

func NewFrigateEventConsumer(config *config.AppConfig, client *kafka.KafkaClient) FrigateEventConsumer {
	return FrigateEventConsumer{
		config,
		client,
	}
}

func decodeMQTTPayload(message []byte) (event types.FrigateMQTTEvent, err error) {
	err = json.Unmarshal(message, &event)
	return
}

func (consumer *FrigateEventConsumer) Consume() error {
	consumer.client.Consume(func(message []byte) (err error) {
		event, err := decodeMQTTPayload(message)
		if err != nil {
			log.Printf("Couldn't deserialize payload. %s", err)
			return
		}

		if event.Type != "end" {
			log.Printf("Skipping incomplete event %s.", event.After.Id)
			return
		}

		after := event.After
		startTime := after.StartTime * float64(time.Microsecond)
		endTime := after.EndTime * float64(time.Microsecond)
		duration := (endTime - startTime) / float64(time.Microsecond)
		camera := after.Camera

		log.Printf("Processing event %s with duration of %f seconds.", after.Id, duration)
		url := fmt.Sprintf("%s/%s/start/%f/end/%f/index.m3u8", consumer.config.Frigate.BaseUrl, camera, startTime/float64(time.Microsecond), endTime/float64(time.Microsecond))

		const TRIM_SECONDS = 60 * 5
		inputArgs := []string{
			"-t", fmt.Sprint(TRIM_SECONDS),
		}

		const CLIP_LENGTH = 5 // TODO: dynamic clip length
		timeScale := CLIP_LENGTH / duration
		scaleFilter := "scale=320:-1"
		ptsFilter := fmt.Sprintf("setpts=%f*PTS", timeScale)
		outputArgs := []string{
			"-vf", strings.Join([]string{
				scaleFilter,
				ptsFilter,
			}, ","),
		}

		outputPath := fmt.Sprintf("%s/%s.mp4", consumer.config.FFMPEG.OutputDir, after.Id)
		err = ffmpeg.Transcode(url, outputPath, inputArgs, outputArgs)
		if err != nil {
			log.Printf("An error occurred during conversion. %s \n", err.Error())
			return
		}

		clipsEvent := types.FrigateClipsEvent{
			ClipUri: fmt.Sprintf("/clips/v1/%s.mp4", event.After.Id),
		}

		clipsEvent.After = event.After
		clipsEvent.Before = event.Before
		clipsEvent.Type = event.Type

		clipsEventBytes, err := json.Marshal(clipsEvent)
		if err != nil {
			log.Println("Failed to marshal clip payload JSON.")
		}

		return consumer.client.Produce(clipsEventBytes)
	})
	return nil
}

// func (consumer *FrigateEventConsumer) Stop() error {
// 	return consumer.client.Stop()
// }
