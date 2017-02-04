/*
 * MumbleDJ
 * By Matthieu Grieger
 * bot/filesystem.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package bot

import (
	"errors"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	id3 "github.com/mikkyang/id3-go"
	"github.com/spf13/viper"
)

func GetMusicDir() string {
	return os.ExpandEnv(viper.GetString("files.music_directory"))
}

// PathIsSong takes a path and returns true if it ends with an mp3 extension.
// This could be implemented with strings.HasSuffix but this regex allows for
// capital letters. Also returns false if there was an error with the regex.
func PathIsSong(path string) bool {
	isSong, err := regexp.MatchString(`^.+\.[mM][pP]3$`, path)
	return err == nil && isSong
}

// PathIsPlaylist takes a path and returns true if it ends with an m3u
// extension. Also returns false if there was an error with the regex.
func PathIsPlaylist(path string) bool {
	isPlaylist, err := regexp.MatchString(`^.+\.[mM]3[uU]$`, path)
	return err == nil && isPlaylist
}

func GetPathForLocalFile(localPath string) string {
	return GetMusicDir() + "/" + localPath
}

// IsSafePath checks the path to make sure it is in the music directory. Returns
// the cleaned path on success and a blank string on failure.
func IsSafePath(path string) (string, error) {
	cleanedPath := filepath.Clean(path)
	musicDir := GetMusicDir()
	if !strings.HasPrefix(cleanedPath, musicDir) {
		return "", errors.New(viper.GetString("files.messages.non_music_dir_prefix_error"))
	}
	return cleanedPath, nil
}

func StripMusicDirPath(path string) (string, error) {
	cleanedPath := filepath.Clean(path)
	musicDir := GetMusicDir()
	if !strings.HasPrefix(cleanedPath, musicDir) {
		return "", errors.New(viper.GetString("files.messages.non_music_dir_prefix_error"))
	}
	return cleanedPath[len(musicDir):], nil
}

func ReadMP3Duration(mp3Reader *id3.File) (time.Duration, error) {
	durationFrame := mp3Reader.Frame("TLEN")
	if durationFrame == nil {
		return 0, errors.New("Failed to read mp3 duration.")
	}
	durationMilliseconds, err := strconv.Atoi(durationFrame.String())
	if err != nil {
		return 0, err
	}
	return time.Duration(durationMilliseconds * 1000), nil
}
