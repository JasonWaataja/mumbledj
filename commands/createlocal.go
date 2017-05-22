/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/createlocal.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package commands

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/layeh/gumble/gumble"
	"github.com/matthieugrieger/mumbledj/bot"
	"github.com/matthieugrieger/mumbledj/interfaces"
	"github.com/spf13/viper"
)

// CreateLocalCommand is a command... (put a short description of the command here)
type CreateLocalCommand struct{}

// Aliases returns the current aliases for the command.
func (c *CreateLocalCommand) Aliases() []string {
	return viper.GetStringSlice("commands.createlocal.aliases")
}

// Description returns the description for the command.
func (c *CreateLocalCommand) Description() string {
	return viper.GetString("commands.createlocal.description")
}

// IsAdminCommand returns true if the command is only for admin use, and
// returns false otherwise.
func (c *CreateLocalCommand) IsAdminCommand() bool {
	return viper.GetBool("commands.createlocal.is_admin")
}

func createDownloadDirIfNeeded() error {
	path := bot.GetPathForLocalFile(viper.GetString("files.download_directory"))
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return os.MkdirAll(path, 0777)
		}
		return err
	}
	if !info.IsDir() {
		return errors.New("Download directory is not a directory.")
	}
	return nil
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
func (c *CreateLocalCommand) Execute(user *gumble.User, args ...string) (string, bool, error) {
	allTracks := make([]interfaces.Track, 0)
	for _, arg := range args {
		if service, err := DJ.GetService(arg); err == nil {
			tracks, err := service.GetTracks(arg, user)
			if err == nil {
				allTracks = append(allTracks, tracks...)
			}
		}
	}
	if len(allTracks) == 0 {
		return "", true, errors.New(viper.GetString("commands.createlocal.messages.no_valid_tracks_error"))
	}
	// Store the tracks that we've copied for the user later.
	addedNames := make([]string, 0)
	for _, track := range allTracks {
		relPath := filepath.Join(viper.GetString("files.download_directory"), track.GetTitle()+".mp3")
		err := bot.DownloadMP3To(track, bot.GetPathForLocalFile(relPath))
		if err == nil {
			addedNames = append(addedNames, relPath)
		}
	}
	if len(addedNames) == 0 {
		return viper.GetString("commands.createlocal.messages.no_tracks_copied"), true, nil
	}
	message := "<h3>Tracks Added</h3>\n"
	message += "<ul>\n"
	for _, localPath := range addedNames {
		message += "<li>" + localPath + "</li>\n"
	}
	message += "</ul>"
	return message, true, nil
}
