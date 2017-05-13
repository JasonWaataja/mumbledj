/*
 * MumbleDJ
 * By Matthieu Grieger
 * services/filesystem.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package services

import (
	"bufio"
	"errors"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/layeh/gumble/gumble"
	"github.com/matthieugrieger/mumbledj/bot"
	"github.com/matthieugrieger/mumbledj/interfaces"
	id3 "github.com/mikkyang/id3-go"
	"github.com/spf13/viper"
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
	path, err := bot.GetSafePath(bot.GetPathForLocalFile(localPath))
	if err != nil {
		return nil, err
	}
	mp3Reader, err := id3.Open(path)
	if err != nil {
		return nil, err
	}
	defer mp3Reader.Close()
	// This function returns 0 on failure, which is the desired behavior.
	duration, _ := bot.ReadMP3Duration(mp3Reader)
	// Leave out some fields for their zero values.
	track := bot.Track{
		Local:          true,
		ID:             localPath,
		Title:          mp3Reader.Title(),
		Author:         mp3Reader.Artist(),
		Submitter:      submitter.Name,
		Service:        fs.GetReadableName(),
		Filename:       filepath.Base(path),
		Duration:       duration,
		PlaybackOffset: 0,
	}
	return &track, nil
}

func (fs *Filesystem) CreateTrackForAbsFile(absPath string, submitter *gumble.User) (*bot.Track, error) {
	path, err := bot.GetSafePath(absPath)
	if err != nil {
		return nil, err
	}
	mp3Reader, err := id3.Open(path)
	if err != nil {
		return nil, err
	}
	defer mp3Reader.Close()
	// This function returns 0 on failure, which is the desired behavior.
	duration, _ := bot.ReadMP3Duration(mp3Reader)
	songID, err := bot.StripMusicDirPath(path)
	if err != nil {
		songID = filepath.Base(path)
	}
	// Leaving out some fields for their zero values.
	track := bot.Track{
		Local:          true,
		ID:             songID,
		Title:          mp3Reader.Title(),
		Author:         mp3Reader.Artist(),
		Submitter:      submitter.Name,
		Service:        fs.GetReadableName(),
		Filename:       filepath.Base(path),
		Duration:       duration,
		PlaybackOffset: 0,
	}
	return &track, nil
}

// CreateTracksForLocalFile scans the localPath and creates a corresponding list
// of tracks, assuming that the file is a playlist file.
func (fs *Filesystem) CreateTracksForLocalFile(localPath string, submitter *gumble.User) ([]interfaces.Track, error) {
	if !bot.PathIsPlaylist(localPath) {
		return nil, errors.New(viper.GetString("files.messages.non_playlist_error"))
	}
	fullPath, err := bot.GetSafePath(bot.GetPathForLocalFile(localPath))
	if err != nil {
		return nil, err
	}
	reader, err := os.Open(fullPath)
	if err != nil {
		return nil, errors.New(viper.GetString("files.messages.file_open_error"))
	}
	defer reader.Close()
	scanner := bufio.NewScanner(reader)
	tracks := make([]interfaces.Track, 0)
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) == 0 || strings.HasPrefix(text, "#") {
			continue
		}
		// This conditional doesn't assume that it could be a url. Just
		// ignoring that option for now.
		if filepath.IsAbs(text) {
			track, err := fs.CreateTrackForAbsFile(text, submitter)
			if err != nil {
				return nil, err
			}
			tracks = append(tracks, track)
		} else {
			track, err := fs.CreateTrackForLocalFile(text, submitter)
			if err != nil {
				return nil, err
			}
			tracks = append(tracks, track)
		}

	}
	return tracks, nil
}
