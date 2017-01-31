/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/yournewcommand.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package commands

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/layeh/gumble/gumble"
	"github.com/matthieugrieger/mumbledj/interfaces"
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

	arg := strings.Join(args, " ")

	path := viper.GetString("files.music_directory") + "/" + arg
	if _, err := os.Stat(path); err != nil {
		return "", true, errors.New(viper.GetSTring("commands.addlocal.messages.no_matching_song_error"))
	}
}
