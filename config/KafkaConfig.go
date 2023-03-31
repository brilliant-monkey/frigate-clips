package config

import "log"

type KafkaConfig struct {
	Brokers  []string     `yaml:"brokers"`
	GroupId  string       `yaml:"group_id,omitempty"`
	Producer *TopicConfig `yaml:"producer,omitempty"`
	Consumer *TopicConfig `yaml:"consumer,omitempty"`
}

func (config *KafkaConfig) GetBrokers() []string {
	if len(config.Brokers) < 1 {
		log.Print("Missing Kafka brokers.")
		return nil
	}
	return config.Brokers
}

func (config *KafkaConfig) GetGroupId() string {
	return config.GroupId
}

func (config *KafkaConfig) GetProducerTopic() *string {
	if config.Producer != nil {
		return &config.Producer.Topic
	}
	return nil
}

func (config *KafkaConfig) GetConsumerTopic() *string {
	if config.Consumer != nil {
		return &config.Consumer.Topic
	}
	return nil
}
