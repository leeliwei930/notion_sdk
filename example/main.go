package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/leeliwei930/notion_sdk/client"
	"github.com/leeliwei930/notion_sdk/config"
)

func main() {
	// pageID, _ := uuid.Parse("341b32cf0b41465db8d40b5ca9235d4a")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	client.InitializeNotionConfig(&config.NotionConfig{
		NotionVersion: "2022-06-28",
		AccessToken:   os.Getenv("NOTION_ACCESS_TOKEN"),
	})
	// _, err = CreatePage(pageID)
	QueryDatabase()
	if err != nil {
		fmt.Print(err.Error())
	}

}
