/*
 * MumbleDJ
 * By Matthieu Grieger
 * services/filesystem.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package services

import (
	"bufio"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	// id3 "github.com/ascherkus/go-id3"
	"github.com/layeh/gumble/gumble"
	"github.com/matthieugrieger/mumbledj/bot"
	"github.com/matthieugrieger/mumbledj/interfaces"
	id3 "github.com/mikkyang/id3-go"
)

// Filesystem is a services that reads files from your local filesystem. It
// currently does not support hosts, so the url must begin with file:///. It can
// read mp3 files and m3u playlists.
type Filesystem struct {
	*GenericService
}

// NewFilesystem returns an initialized Filesystem service object.
func NewFilesystemService() *Filesystem {
	return &Filesystem{
		&GenericService{
			ReadableName: "Filesystem",
			Format:       "bestaudio",
			TrackRegex: []*regexp.Regexp{
				regexp.MustCompile(`.*`),
				// regexp.MustCompile(`.*`),
			},
			PlaylistRegex: []*regexp.Regexp{
				regexp.MustCompile(`file:\/\/\/.*\.[mM]3[uU]`),
			},
		},
	}
}

// CheckAPIKey performs a test API call with the API key
// provided in the configuration file to determine if the
// service should be enabled.
func (fs *Filesystem) CheckAPIKey() error {
	// There is no API key required to use local files, so the error will
	// always be nil.
	return nil
}

// GetTracks uses the passed URL to find and return
// tracks associated with the URL. An error is returned
// if any error occurs during the API call.
func (fs *Filesystem) GetTracks(url string, submitter *gumble.User) ([]interfaces.Track, error) {
	tracks := make([]interfaces.Track, 0)
	return tracks, nil
}

// CreateTrackForLocalFile creates a track that represents a file in the music
// directory. Returns the track on success an nil on failure with an error.
func CreateTrackForLocalFile(localPath string) (*bot.Track, error) {

}
