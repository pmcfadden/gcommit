package models

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Format string
	Story  string
	Pair   string
	Name   string
	Email  string
}

func ReadAndSetConfig(stdin io.Reader) Config {
	reader := bufio.NewReader(stdin)
	format := viper.GetString("format")
	story := viper.GetString("story")
	pair := viper.GetString("pair")
	name := viper.GetString("author.name")
	email := viper.GetString("author.email")
	fmt.Printf("Enter Story [%s]: ", story)
	maybeStory, _ := reader.ReadString('\n')
	if maybeStory != "\n" {
		story = strings.TrimSuffix(maybeStory, "\n")
	}
	fmt.Printf("Enter Pair [%s]: ", pair)
	maybePair, _ := reader.ReadString('\n')
	if maybePair != "\n" {
		pair = strings.TrimSuffix(maybePair, "\n")
	}
	viper.Set("story", story)
	viper.Set("pair", pair)
	viper.WriteConfig()
	return Config{Format: format, Story: story, Pair: pair, Name: name, Email: email}
}
