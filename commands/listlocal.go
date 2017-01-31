/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/addlocal.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package commands

import (
	"errors"
	"fmt"

	"github.com/layeh/gumble/gumble"
	"github.com/matthieugrieger/mumbledj/interfaces"
	"github.com/spf13/viper"
)

// Listlocal is a command that adds a tracks from the local filesystem. It
// searches the music directory if it exists and adds the track if the file is
// there.
type ListlocalCommand struct{}

// Aliases returns the current aliases for the command.
func (c *ListlocalCommand) Aliases() []string {
	return viper.GetStringSlice("commands.listlocal.aliases")
}

// Description returns the description for the command.
func (c *ListlocalCommand) Description() string {
	return viper.GetString("commands.listlocal.description")
}

// IsAdminCommand returns true if the command is only for admin use, and
// returns false otherwise.
func (c *YourNewCommand) IsAdminCommand() bool {
	return viper.GetBool("commands.listlocal.is_admin")
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
func (c *ListlocalCommand) Execute(user *gumble.User, args ...string) (string, bool, error) {
	musicDir := viper.GetString(
}
