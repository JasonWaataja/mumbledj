/*
 * MumbleDJ
 * By Matthieu Grieger
 * bot/filesystem.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package bot

import (
	"fmt"
	"io"
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
	homeDir := user.Current().HomeDir
	return strings.Replace(path, "~", homeDir, -1)
}

func GetPathForLocalFile(localPath string) string {
	return GetMusicDir() + "/" + localPath
}
