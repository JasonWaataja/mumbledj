/*
 * MumbleDJ
 * By Matthieu Grieger
 * bot/filesystem.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package bot

import (
	"errors"
	"fmt"
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
	ext := filepath.Ext(path)
	isSong, err := regexp.MatchString(`[mM][pP]3`, ext)
	return err == nil && isSong
}

// PathIsPlaylist takes a path and returns true if it ends with an m3u
// extension. Also returns false if there was an error with the regex.
func PathIsPlaylist(path string) bool {
	ext := filepath.Ext(path)
	isPlaylist, err := regexp.MatchString(`[mM]3[uU]`, ext)
	return err == nil && isPlaylist
}

// GetPathForLocalFile returns a full file path with the given local path
// relative to the music directory concatenated with a forward slash.
func GetPathForLocalFile(localPath string) string {
	return filepath.Join(GetMusicDir(), localPath)
}

// GetPathElements slipts a path into its components. The first elements is a
// slash if it is an absolute path.
func GetPathElements(path string) []string {
	elements := make([]string, 0)
	path = filepath.Clean(path)
	dir, file := filepath.Split(path)
	dir = filepath.Clean(dir)
	separator := fmt.Sprintf("%c", filepath.Separator)
	// This doesn't assume Microsoft Windows paths that could have a
	// backslash instead.
	for dir != "" && dir != "." && !strings.HasSuffix(dir, separator) {
		if len(file) > 0 {
			elements = append(elements, file)
		}
		path = filepath.Clean(dir)
		dir, file = filepath.Split(path)
		dir = filepath.Clean(dir)
	}
	if file != "" {
		elements = append(elements, file)
	}
	if strings.HasSuffix(dir, separator) {
		elements = append(elements, dir)
	}
	for i := 0; i < len(elements)/2; i++ {
		opp := len(elements) - i - 1
		elements[i], elements[opp] = elements[opp], elements[i]
	}
	return elements
}

// GetSafePath checks the path to make sure it is in the music directory. Returns
// the cleaned path on success and a blank string on failure.
func GetSafePath(path string) (string, error) {
	cleanedPath := filepath.Clean(path)
	pathElements := GetPathElements(cleanedPath)
	musicDirElements := GetPathElements(GetMusicDir())
	if len(pathElements) < len(musicDirElements) {
		return "", errors.New(viper.GetString("files.messages.non_music_dir_prefix_error"))
	}
	for i, element := range musicDirElements {
		if pathElements[i] != element {
			return "", errors.New(viper.GetString("files.messages.non_music_dir_prefix_error"))
		}
	}
	return cleanedPath, nil
}

// StripMusicDirPath takes a path that begins with the music directory and
// returns a path with the music directory stripped out. In other words, the
// local file path. Returns a blank string and error on failuer or on non-safe
// path.
func StripMusicDirPath(path string) (string, error) {
	cleanedPath, err := GetSafePath(path)
	if err != nil {
		return "", err
	}
	relPath, err := filepath.Rel(GetMusicDir(), cleanedPath)
	if err != nil {
		return "", err
	}
	return relPath, nil
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
