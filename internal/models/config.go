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
	format := viper.GetString("format")
	story := viper.GetString("story")
	pair := viper.GetString("pair")
	name := viper.GetString("author.name")
	email := viper.GetString("author.email")
	fmt.Printf("Enter Story [%s]: ", story)
	story = readInputWithDefault(stdin, story)
	fmt.Printf("Enter Pair [%s]: ", pair)
	pair = readInputWithDefault(stdin, pair)
	viper.Set("story", story)
	viper.Set("pair", pair)
	viper.WriteConfig()
	return Config{Format: format, Story: story, Pair: pair, Name: name, Email: email}
}

func readInputWithDefault(stdin io.Reader, original string) string {
	reader := bufio.NewReader(stdin)
	maybeChange, _ := reader.ReadString('\n')
	if maybeChange != "\n" {
		original = strings.TrimSuffix(maybeChange, "\n")
	}
	return original
}
