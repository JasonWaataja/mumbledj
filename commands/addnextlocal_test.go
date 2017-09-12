/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/addnextlocal_test.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package commands

import (
	"testing"

	"github.com/matthieugrieger/mumbledj/bot"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
)

type AddNextLocalCommandTestSuite struct {
	Command AddNextLocalCommand
	suite.Suite
}

func (suite *AddNextLocalCommandTestSuite) SetupSuite() {
	viper.Set("commands.addnextlocal.aliases", []string{"addnextlocal", "anl"})
	viper.Set("commands.addnextlocal.description", "addnextlocal")
	viper.Set("commands.addnextlocal.is_admin", false)
}

func (suite *AddNextLocalCommandTestSuite) SetupTest() {
	DJ.Queue = bot.NewQueue()
}

func (suite *AddNextLocalCommandTestSuite) TestAliases() {
	suite.Equal([]string{"addnextlocal", "anl"}, suite.Command.Aliases())
}

func (suite *AddNextLocalCommandTestSuite) TestDescription() {
	suite.Equal("addnextlocal", suite.Command.Description())
}

func (suite *AddNextLocalCommandTestSuite) TestIsAdminCommand() {
	suite.False(suite.Command.IsAdminCommand())
}

// Implement more tests here as necessary! It may be helpful to take a look
// at the stretchr/testify documentation:
// https://github.com/stretchr/testify
// Remove this comment before sending a pull request.

func TestAddNextLocalCommandTestSuite(t *testing.T) {
	suite.Run(t, new(AddNextLocalCommandTestSuite))
}
