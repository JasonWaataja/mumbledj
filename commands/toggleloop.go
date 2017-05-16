/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/toggleloop.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package commands

import (
	"github.com/layeh/gumble/gumble"
	"github.com/spf13/viper"
)

// ToggleLoopCommand is a command... (put a short description of the command here)
type ToggleLoopCommand struct{}

// Aliases returns the current aliases for the command.
func (c *ToggleLoopCommand) Aliases() []string {
	return viper.GetStringSlice("commands.toggleloop.aliases")
}

// Description returns the description for the command.
func (c *ToggleLoopCommand) Description() string {
	return viper.GetString("commands.toggleloop.description")
}

// IsAdminCommand returns true if the command is only for admin use, and
// returns false otherwise.
func (c *ToggleLoopCommand) IsAdminCommand() bool {
	return viper.GetBool("commands.toggleloop.is_admin")
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
func (c *ToggleLoopCommand) Execute(user *gumble.User, args ...string) (string, bool, error) {
	new_value := !DJ.Loop
	DJ.Loop = new_value
	if new_value {
		return viper.GetString("commands.toggleloop.messages.toggled_on"), false, nil
	}
	return viper.GetString("commands.toggleloop.messages.toggled_off"), false, nil
}
