package cmd

import (
	"log"
	"net/url"

	"github.com/spf13/viper"
)

// isErrNotFound checks if the error passed is Viper's Config File Not Found
// error.
func isErrNotFound(err error) bool {
	_, ok := err.(viper.ConfigFileNotFoundError)
	return ok
}

func urlMustParse(path string) *url.URL {
	source, err := url.Parse(path)
	if err != nil {
		log.Fatal("failed to parse source url: ", err)
	}

	return source
}
