package config

type TopicConfig struct {
	Name string `yaml:"name"`
}

func NewTopicConfig(name string) TopicConfig {
	return TopicConfig{
		Name: name,
	}
}
