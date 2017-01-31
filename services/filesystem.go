/*
 * MumbleDJ
 * By Matthieu Grieger
 * services/filesystem.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package services

import (
	"regexp"
	"time"

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
			},
			PlaylistRegex: []*regexp.Regexp{
				regexp.MustCompile(`.*`),
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

// CreateTrackForPath creates a bot.Track for the given localPath which is
// interpreted relative to the music directory. Returns the track on success and
// nil and an error on failure.
func (fs *Filesystem) CreateTrackForLocalFile(localPath string, submitter *gumble.User) (*bot.Track, error) {
	path := bot.GetPathForLocalFile(localPath)
	mp3Reader, err := id3.Open(path)
	if err != nil {
		return nil, err
	}
	defer mp3Reader.Close()

	// Leaving out some fields for their zero values.
	track := bot.Track{
		Local:          true,
		ID:             localPath,
		Title:          mp3Reader.Title(),
		Author:         mp3Reader.Artist(),
		Submitter:      submitter.Name,
		Service:        fs.GetReadableName(),
		Filename:       localPath,
		Duration:       time.Duration(0),
		PlaybackOffset: 0,
	}
	return &track, nil
}

// CreateTracksForLocalFile scans the localPath and creates a corresponding list
// of tracks, assuming that the file is a playlist file.
func (fs *Filesystem) CreateTracksForLocalFile(localPath string, submitter *gumble.User) ([]bot.Track, error) {
	return nil, nil
}
