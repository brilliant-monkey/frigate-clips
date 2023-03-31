package config

type AppConfig struct {
	FFMPEG  FFMPEGConfig  `yaml:"ffmpeg"`
	Frigate FrigateConfig `yaml:"frigate"`
	Kafka   KafkaConfig   `yaml:"kafka"`
	MQTT    MQTTConfig    `yaml:"mqtt"`
}
