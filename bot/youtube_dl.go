/*
 * MumbleDJ
 * By Matthieu Grieger
 * bot/youtube_dl.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package bot

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/matthieugrieger/mumbledj/interfaces"
	id3 "github.com/mikkyang/id3-go"
	"github.com/spf13/viper"
)

// YouTubeDL is a struct that gathers all methods related to the youtube-dl
// software.
// youtube-dl: https://rg3.github.io/youtube-dl/
type YouTubeDL struct{}

// Download downloads the audio associated with the incoming `track` object
// and stores it `track.Filename`.
func (yt *YouTubeDL) Download(t interfaces.Track) error {
	player := "--prefer-ffmpeg"
	if viper.GetString("defaults.player_command") == "avconv" {
		player = "--prefer-avconv"
	}

	filepath := t.GetFullPath()

	// Determine which format to use.
	format := "bestaudio"
	for _, service := range DJ.AvailableServices {
		if service.GetReadableName() == t.GetService() {
			format = service.GetFormat()
		}
	}

	// Check to see if track is already downloaded.
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		var cmd *exec.Cmd
		if t.GetService() == "Mixcloud" {
			cmd = exec.Command("youtube-dl", "--verbose", "--no-mtime", "--output", filepath, "--formmat", format, "-x", "--external-downloader", "aria2c", player, t.GetURL())
		} else {
			cmd = exec.Command("youtube-dl", "--verbose", "--no-mtime", "--output", filepath, "--format", format, "-x", player, t.GetURL())
		}
		output, err := cmd.CombinedOutput()
		if err != nil {
			args := ""
			for s := range cmd.Args {
				args += cmd.Args[s] + " "
			}
			logrus.Warnf("%s\n%s\nyoutube-dl: %s", args, string(output), err.Error())
			return errors.New("Track download failed")
		}

		if viper.GetBool("cache.enabled") {
			DJ.Cache.CheckDirectorySize()
		}
	}

	return nil
}

// DownloadMP3To downloads the track like it normally would but instead then
// takes the audio and stores it in a file at path.
func DownloadMP3To(t interfaces.Track, path string) error {
	player := "--prefer-ffmpeg"
	if viper.GetString("defaults.player_command") == "avconv" {
		player = "--prefer-avconv"
	}

	// Determine which format to use.
	format := "bestaudio"
	for _, service := range DJ.AvailableServices {
		if service.GetReadableName() == t.GetService() {
			format = service.GetFormat()
		}
	}

	// Check to see if track is already downloaded.
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return nil
	}

	var cmd *exec.Cmd
	noExt := strings.TrimSuffix(path, filepath.Ext(path))
	outputPath := noExt + ".track"
	if t.GetService() == "Mixcloud" {
		cmd = exec.Command("youtube-dl", "--verbose", "--no-mtime", "--output", outputPath, "--format", format, "--extract-audio", "--audio-format", "mp3", "--external-downloader", "aria2c", player, t.GetURL())
	} else {
		cmd = exec.Command("youtube-dl", "--verbose", "--no-mtime", "--output", outputPath, "--format", format, "--extract-audio", "--audio-format", "mp3", player, t.GetURL())
	}
	output, err := cmd.CombinedOutput()
	if err != nil {
		args := ""
		for s := range cmd.Args {
			args += cmd.Args[s] + " "
		}
		logrus.Warnf("%s\n%s\nyoutube-dl: %s", args, string(output), err.Error())
		return errors.New("Track download failed")
	}
	asMP3Path := noExt + ".mp3"
	reader, err := id3.Open(asMP3Path)
	if err != nil {
		return errors.New("Failed to open mp3 file.")
	}
	reader.SetTitle(t.GetTitle())
	reader.SetArtist(t.GetAuthor())
	reader.Close()
	if asMP3Path != path {
		err := os.Rename(asMP3Path, path)
		if err != nil {
			return errors.New("Failed to rename to correct file.")
		}
	}
	return nil
}

// Delete deletes the audio file associated with the incoming `track` object.
func (yt *YouTubeDL) Delete(t interfaces.Track) error {
	if !viper.GetBool("cache.enabled") {
		filePath := t.GetFullPath()
		if _, err := os.Stat(filePath); err == nil {
			if err := os.Remove(filePath); err == nil {
				return nil
			}
			return errors.New("An error occurred while deleting the audio file")
		}
	}
	return nil
}
