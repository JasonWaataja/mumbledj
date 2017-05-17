/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/addlocal.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package commands

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/layeh/gumble/gumble"
	"github.com/matthieugrieger/mumbledj/bot"
	"github.com/matthieugrieger/mumbledj/interfaces"
	"github.com/matthieugrieger/mumbledj/services"
	id3 "github.com/mikkyang/id3-go"
	"github.com/spf13/viper"
)

// IndentationString is the character used to indent string in CreateInfo. It is
// meant to be used with Mumble's HTML>
const IndentationString = "&nbsp;"

// LineBreakString is the string used to break lines in CreateInfo. It is meant
// to be used with Mumble's HTML.
const LineBreakString = "<br>"

// ListLocal is a command that adds a tracks from the local filesystem. It
// searches the music directory if it exists and adds the track if the file is
// there.
type ListLocalCommand struct{}

// Aliases returns the current aliases for the command.
func (c *ListLocalCommand) Aliases() []string {
	return viper.GetStringSlice("commands.listlocal.aliases")
}

// Description returns the description for the command.
func (c *ListLocalCommand) Description() string {
	return viper.GetString("commands.listlocal.description")
}

// IsAdminCommand returns true if the command is only for admin use, and
// returns false otherwise.
func (c *ListLocalCommand) IsAdminCommand() bool {
	return viper.GetBool("commands.listlocal.is_admin")
}

// createInfoForFile creates a small piece of information for the given path.
// The second return value indicates whether it should be displayed.
func createInfoForFile(path, relPath string, info os.FileInfo) (string, bool) {
	if info.Mode().IsRegular() {
		if bot.PathIsSong(path) {
			reader, err := id3.Open(path)
			if err != nil {
				return "", false
			}
			defer reader.Close()
			duration, _ := bot.ReadMP3Duration(reader)
			return "<b>" + relPath + "</b> " + reader.Title() +
				", " + reader.Artist() + " (" + duration.String() + ")", true
		} else if bot.PathIsPlaylist(path) {
			return "(Playlist) <b>" + relPath + "</b>", true
		}
		return "", false
	}
	return "(Directory) <b>" + relPath + "</b>", true
}

func createInfoForTrack(track interfaces.Track) (string, bool) {
	if track.IsLocal() {
		fullPath := bot.GetPathForLocalFile(track.GetFilename())
		info, err := os.Stat(fullPath)
		if err != nil {
			return "", false
		}
		return createInfoForFile(fullPath, track.GetFilename(), info)
	}
	message := "(" + track.GetService() + ") <b>"
	message += track.GetTitle() + "</b> "
	message += track.GetAuthor() + ", ("
	message += track.GetDuration().String() + ")"
	return message, true
}

func createMessageForDir(path, relPath string) (string, error) {
	dirInfo, err := os.Stat(path)
	if err != nil {
		return "", errors.New(viper.GetString("commands.listlocal.messages.read_failure_error"))
	}
	if !dirInfo.IsDir() {
		return "", errors.New(viper.GetString("commands.listlocal.messages.scan_non_directory_error"))
	}
	entries, err := ioutil.ReadDir(path)
	if err != nil {
		return "", errors.New(viper.GetString("commands.listlocal.messages.read_failure_error"))
	}
	message := "<h3>"
	if relPath == "." {
		message += "Music"
	} else {
		message += relPath
	}
	message += "</h3>\n"
	message += "<ul>\n"
	for _, entry := range entries {
		entryPath := filepath.Join(path, entry.Name())
		entryRelPath := filepath.Join(relPath, entry.Name())
		entryMessage, ok := createInfoForFile(entryPath, entryRelPath, entry)
		if ok && len(entryMessage) > 0 {
			message += "<li>" + entryMessage + "</li>\n"
		}
	}
	message += "</ul>"
	return message, nil
}

func createMessageForPlaylist(path, relPath string, submitter *gumble.User) (string, error) {
	var fs *services.Filesystem
	for _, service := range DJ.AvailableServices {
		asFilesystem, ok := service.(*services.Filesystem)
		if ok {
			fs = asFilesystem
		}
	}
	if fs == nil {
		return "", errors.New("No filesystem service.")
	}
	tracks, err := fs.CreateTracksForLocalFile(relPath, submitter)
	if err != nil {
		return "", err
	}
	message := "<h3>" + relPath + "</h3>"
	message += "<ol>\n"
	for _, track := range tracks {
		info, ok := createInfoForTrack(track)
		if ok {
			message += "<li>" + info + "</li>\n"
		}
	}
	message += "</ol>"
	return message, nil
}

// Execute executes the command with the given user and arguments.
// Return value descriptions:
//    string: A message to be returned to the user upon successful execution.
//    bool:   Whether the message should be private or not. true = private,
//            false = public (sent to whole channel).
//    error:  An error message to be returned upon unsuccessful execution.
//            If no error has occurred, pass nil instead.
// Example return statement:
//    return "This is a private message!", true, nil
func (c *ListLocalCommand) Execute(user *gumble.User, args ...string) (string, bool, error) {
	path := bot.GetMusicDir()
	if len(args) > 0 {
		argPath := strings.Join(args, " ")
		argPath, err := bot.GetSafePath(bot.GetPathForLocalFile(argPath))
		if err != nil {
			return "", true, errors.New(viper.GetString("commands.listlocal.messages.read_failure_error"))
		}
		path = argPath
	}
	relPath, err := bot.StripMusicDirPath(path)
	if err != nil {
		return "", true, errors.New(viper.GetString("commands.listlocal.messages.read_failure_error"))
	}
	if bot.PathIsPlaylist(path) {
		message, err := createMessageForPlaylist(path, relPath, user)
		if err != nil {
			return "", true, errors.New(viper.GetString("commands.listlocal.messages.read_failure_error"))
		}
		return message, true, nil
	}
	message, err := createMessageForDir(path, relPath)
	if err != nil {
		return "", true, errors.New(viper.GetString("commands.listlocal.messages.read_failure_error"))
	}
	return message, true, nil
}
