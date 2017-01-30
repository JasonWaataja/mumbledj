/*
 * MumbleDJ
 * By Matthieu Grieger
 * bot/queue.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package bot

import (
	"io"
	"os"
	"regexp"

	"github.com/spf13/viper"
)

// URLError is an error that results from parsing a URL.
type URLError struct {
	URL string
}

func (err URLError) Error() string {
	return "Failed to extract information from URL: " + err.URL
}

// PathForFileURL parses a url and extract the file it points to of the form
// file:///path. Returns the extracted path on success and the original url on
// failure.
func PathForFileURL(url string) (string, error) {
	re := regexp.MustCompile(`file:\/\/\/(.*\.[mM][pP]3)`)
	matches := re.FindStringSubmatch(url)
	if matches == nil {
		return url, URLError{url}
	}
	return "/" + matches[1], nil
}

// CacheMP3File copies the given path into the cache directory. Returns an error
// if there was one, nil otherwise.
func CacheMP3File(path string) error {
	destinationPath := os.ExpandEnv(viper.GetString("cache.directory") + "/" + path)
	reader, err := os.Open(path)
	if err != nil {
		return err
	}
	defer reader.Close()
	writer, err := os.Create(destinationPath)
	if err != nil {
		return err
	}
	defer writer.Close()
	_, err = io.Copy(writer, reader)
	return err
}
