/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/listlocal_test.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package commands

import (
	"testing"

	"github.com/matthieugrieger/mumbledj/bot"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
)

type ListLocalCommandTestSuite struct {
	Command ListLocalCommand
	suite.Suite
}

func (suite *ListLocalCommandTestSuite) SetupSuite() {
	viper.Set("commands.listlocal.aliases", []string{"listlocal", "ll"})
	viper.Set("commands.listlocal.description", "listlocal")
	viper.Set("commands.listlocal.is_admin", false)
}

func (suite *ListLocalCommandTestSuite) SetupTest() {
	DJ.Queue = bot.NewQueue()
}

func (suite *ListLocalCommandTestSuite) TestAliases() {
	suite.Equal([]string{"listlocal", "ll"}, suite.Command.Aliases())
}

func (suite *ListLocalCommandTestSuite) TestDescription() {
	suite.Equal("listlocal", suite.Command.Description())
}

func (suite *ListLocalCommandTestSuite) TestIsAdminCommand() {
	suite.False(suite.Command.IsAdminCommand())
}

// Implement more tests here as necessary! It may be helpful to take a look
// at the stretchr/testify documentation:
// https://github.com/stretchr/testify
// Remove this comment before sending a pull request.

func TestListLocalCommandTestSuite(t *testing.T) {
	suite.Run(t, new(ListLocalCommandTestSuite))
}
