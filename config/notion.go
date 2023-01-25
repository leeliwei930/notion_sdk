package config

type NotionConfig struct {
	NotionVersion string
	AccessToken   string
}

func GetDefaultNotionConfig() NotionConfig {
	return NotionConfig{
		NotionVersion: "2022-06-28",
	}
}
