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
	"os/user"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/spf13/viper"
)

func GetMusicDir() string {
	return os.ExpandEnv(TildeExpand(viper.GetString("files.music_dir")))
}

func TildeExpand(path string) string {
	currentUser, err := user.Current()
	var homedir string
	if err != nil {
		homedir = os.Getenv("$HOME")
		if homedir == "" {
			panic("Failed to get user home directory.")
		}
	} else {
		homedir = currentUser.HomeDir
	}
	return strings.Replace(path, "~", homedir, -1)
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

func StripMusicDirPath(path string) (string, error) {
	cleanedPath := filepath.Clean(path)
	musicDir := GetMusicDir()
	if !strings.HasPrefix(cleanedPath, musicDir) {
		return "", errors.New(viper.GetString("files.non_music_dir_prefix_error"))
	}
	return cleanedPath[len(musicDir):], nil
}
