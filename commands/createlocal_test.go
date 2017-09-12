/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/createlocal_test.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package commands

import (
	"testing"

	"github.com/matthieugrieger/mumbledj/bot"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
)

type CreateLocalCommandTestSuite struct {
	Command CreateLocalCommand
	suite.Suite
}

func (suite *CreateLocalCommandTestSuite) SetupSuite() {
	viper.Set("commands.createlocal.aliases", []string{"createlocal", "cl"})
	viper.Set("commands.createlocal.description", "createlocal")
	viper.Set("commands.createlocal.is_admin", false)
}

func (suite *CreateLocalCommandTestSuite) SetupTest() {
	DJ.Queue = bot.NewQueue()
}

func (suite *CreateLocalCommandTestSuite) TestAliases() {
	suite.Equal([]string{"createlocal", "cl"}, suite.Command.Aliases())
}

func (suite *CreateLocalCommandTestSuite) TestDescription() {
	suite.Equal("createlocal", suite.Command.Description())
}

func (suite *CreateLocalCommandTestSuite) TestIsAdminCommand() {
	suite.False(suite.Command.IsAdminCommand())
}

// Implement more tests here as necessary! It may be helpful to take a look
// at the stretchr/testify documentation:
// https://github.com/stretchr/testify
// Remove this comment before sending a pull request.

func TestCreateLocalCommandTestSuite(t *testing.T) {
	suite.Run(t, new(CreateLocalCommandTestSuite))
}
