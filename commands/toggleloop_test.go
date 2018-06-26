/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/toggleloop_test.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package commands

import (
	"testing"

	"github.com/matthieugrieger/mumbledj/bot"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
)

type ToggleLoopCommandTestSuite struct {
	Command ToggleLoopCommand
	suite.Suite
}

func (suite *ToggleLoopCommandTestSuite) SetupSuite() {
	viper.Set("commands.toggleloop.aliases", []string{"toggleloop", "tl"})
	viper.Set("commands.toggleloop.description", "toggleloop")
	viper.Set("commands.toggleloop.is_admin", false)
}

func (suite *ToggleLoopCommandTestSuite) SetupTest() {
	DJ.Queue = bot.NewQueue()
}

func (suite *ToggleLoopCommandTestSuite) TestAliases() {
	suite.Equal([]string{"toggleloop", "tl"}, suite.Command.Aliases())
}

func (suite *ToggleLoopCommandTestSuite) TestDescription() {
	suite.Equal("toggleloop", suite.Command.Description())
}

func (suite *ToggleLoopCommandTestSuite) TestIsAdminCommand() {
	suite.False(suite.Command.IsAdminCommand())
}

func (suite *ToggleLoopCommandTestSuite) TestTogglesLoop() {
	originalValue := DJ.Loop
	suite.Command.Execute(nil)
	suite.Equal(DJ.Loop, !originalValue)
	suite.Command.Execute(nil)
	suite.Equal(DJ.Loop, originalValue)
}

// Implement more tests here as necessary! It may be helpful to take a look
// at the stretchr/testify documentation:
// https://github.com/stretchr/testify
// Remove this comment before sending a pull request.

func TestToggleLoopCommandTestSuite(t *testing.T) {
	suite.Run(t, new(ToggleLoopCommandTestSuite))
}
