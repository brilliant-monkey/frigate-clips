package config

type MQTTConfig struct {
	Endpoint string       `yaml:"endpoint"`
	ClientId string       `yaml:"client_id"`
	Username string       `yaml:"username"`
	Password string       `yaml:"password"`
	Producer *TopicConfig `yaml:"producer,omitempty"`
	Consumer *TopicConfig `yaml:"consumer,omitempty"`
}

func (c *MQTTConfig) GetEndpoint() string {
	return c.Endpoint
}

func (c *MQTTConfig) GetClientId() string {
	return c.ClientId
}

func (c *MQTTConfig) GetUsername() string {
	return c.Username
}

func (c *MQTTConfig) GetPassword() string {
	return c.Password
}

func (c *MQTTConfig) GetProducerTopic() *string {
	if c.Producer != nil {
		return &c.Producer.Topic
	}
	return nil
}

func (c *MQTTConfig) GetConsumerTopic() *string {
	if c.Consumer != nil {
		return &c.Consumer.Topic
	}
	return nil
}
