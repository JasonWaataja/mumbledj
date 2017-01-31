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
	if fs.isTrack(url) {
		path, err := bot.PathForFileURL(url)
		if err != nil {
			return tracks, err
		}
		track, err := fs.createTrackForFile(path, submitter)
		if err == nil {
			tracks = append(tracks, track)
		} else {
			return tracks, err
		}
	} else if fs.isPlaylist(url) {
		path, err := bot.PathForFileURL(url)
		if err != nil {
			return tracks, err
		}
		reader, err := os.Open(path)
		if err != nil {
			return tracks, err
		}
		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			line := scanner.Text()
			if len(line) > 0 && !strings.HasPrefix(line, "#") {
				track, err := fs.createTrackForFile(line, submitter)
				if err != nil {
					return tracks, err
				}
				tracks = append(tracks, track)
			}
		}
	}
	return tracks, nil
}

// createTrackForFile extracts the info from an mp3 file. Returns a tracks on
// success and false on failure.
func (fs *Filesystem) createTrackForFile(path string, submitter *gumble.User) (*bot.Track, error) {
	mp3file, err := id3.Open(path)
	defer mp3file.Close()
	if err != nil {
		return nil, err
	}
	basename := filepath.Base(path)
	track := bot.Track{
		ID:        basename,
		URL:       "file://" + path,
		Title:     mp3file.Artist(),
		Submitter: submitter.Name,
		Service:   fs.GetReadableName(),
		Filename:  basename,
	}
	return &track, nil
}
