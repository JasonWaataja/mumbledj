/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/addlocal.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package commands

import (
	"errors"
	"os"
	"strings"

	"github.com/layeh/gumble/gumble"
	"github.com/matthieugrieger/mumbledj/bot"
	"github.com/matthieugrieger/mumbledj/services"
	"github.com/spf13/viper"
)

// AddlocalCommand is a command that adds a tracks from the local filesystem. It
// searches the music directory if it exists and adds the track if the file is
// there.
type AddlocalCommand struct{}

// Aliases returns the current aliases for the command.
func (c *AddlocalCommand) Aliases() []string {
	return viper.GetStringSlice("commands.addlocal.aliases")
}

// Description returns the description for the command.
func (c *AddlocalCommand) Description() string {
	return viper.GetString("commands.addlocal.description")
}

// IsAdminCommand returns true if the command is only for admin use, and
// returns false otherwise.
func (c *AddlocalCommand) IsAdminCommand() bool {
	return viper.GetBool("commands.addlocal.is_admin")
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
func (c *AddlocalCommand) Execute(user *gumble.User, args ...string) (string, bool, error) {
	if len(args) == 0 {
		return "", true, errors.New(viper.GetString("commands.addlocal.messages.no_argument_error"))
	}

	// If arguments were split around spaces, put them back together
	// separated by spaces.
	arg := strings.Join(args, " ")

	path := bot.GetPathForLocalFile(arg)

	if bot.PathIsSong(path) {
		for _, service := range DJ.AvailableServices {
			fs, ok := service.(services.Filesystem)
			if ok {
				track, err := fs.CreateTrackForLocalFile(path, user)
				if err != nil {
					return viper.GetString("commands.addlocal.track_creation_error", true, err)
				} else {
					return "", true, nil
				}
			}
		}
		return "", true, errors.New(viper.GetString("commands.addlocal.no_filesystem_service_error"))
	} else if bot.PathIsPlaylist(path) {
		for _, service := DJ.AvailableServices {
			fs, ok := service.(services.Filesystem)
			if ok {
			}
		}

	}

	if _, err := os.Stat(path); err != nil {
		return "", true, errors.New(viper.GetString("commands.addlocal.messages.no_matching_song_error"))
	}

}
