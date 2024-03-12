package config

import "time"

type NotionConfig struct {
	NotionVersion   string
	AccessToken     string
	TimeoutDuration time.Duration
}

func GetDefaultNotionConfig() NotionConfig {
	return NotionConfig{
		NotionVersion: "2022-06-28",
	}
}
