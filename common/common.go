package common

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"os"
	"path"
	"time"
)

var file, _ = os.Getwd()

// defaultUserAgent default value for ua.
const defaultUserAgent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.101 Safari/537.36"

// Config contains information we need to process a novel
type Config struct {
	DomainFlagM    string `json:"DomainFlagM"`
	DomainFlagL    string `json:"DomainFlagL"`
	SoURL          string `json:"SoURL"`
	LatestChapeter struct {
		LatestChapterName string `json:"LatestChapterName"`
		LatestChapterURL  string `json:"LatestChapterURL"`
	}
	Sites []string `json:"Sites"`
}

// LoadConfiguration Get the configuration from the rules.json
func LoadConfiguration() (Config, error) {
	var config Config
	fileName := path.Join(file, "data", "rules.json")
	configFile, err := os.Open(fileName)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
		return config, err
	}
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	return config, err
}

// GetUserAgent Get a random user agent string.
func GetUserAgent() string {
	fileName := path.Join(file, "data", "user_agents.txt")
	fileData, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return defaultUserAgent
	}
	defer fileData.Close()
	userAgents := make([]string, 0)
	br := bufio.NewReader(fileData)
	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		userAgents = append(userAgents, string(line))
	}
	rand.Seed(time.Now().Unix())
	n := rand.Int() % len(userAgents)
	return userAgents[n]
}
