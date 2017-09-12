/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/addlocal_test.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package commands

import (
	"testing"

	"github.com/layeh/gumble/gumbleffmpeg"
	"github.com/matthieugrieger/mumbledj/bot"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
)

type AddLocalCommandTestSuite struct {
	Command AddLocalCommand
	suite.Suite
}

func (suite *AddLocalCommandTestSuite) SetupSuite() {

	// Trick the tests into thinking audio is already playing to avoid
	// attempting to play tracks that don't exist.
	DJ = bot.NewMumbleDJ()
	bot.DJ = DJ
	// Note: Copying this from the add command test. Making the audio stream
	// shouldn't be necessary for this test but not setting the stream here
	// fails other tests so I'm putting it here.
	DJ.AudioStream = new(gumbleffmpeg.Stream)

	viper.Set("commands.addlocal.aliases", []string{"addlocal", "al"})
	viper.Set("commands.addlocal.description", "addlocal")
	viper.Set("commands.addlocal.is_admin", false)
}

func (suite *AddLocalCommandTestSuite) SetupTest() {
	DJ.Queue = bot.NewQueue()
}

func (suite *AddLocalCommandTestSuite) TestAliases() {
	suite.Equal([]string{"addlocal", "al"}, suite.Command.Aliases())
}

func (suite *AddLocalCommandTestSuite) TestDescription() {
	suite.Equal("addlocal", suite.Command.Description())
}

func (suite *AddLocalCommandTestSuite) TestIsAdminCommand() {
	suite.False(suite.Command.IsAdminCommand())
}

// Implement more tests here as necessary! It may be helpful to take a look
// at the stretchr/testify documentation:
// https://github.com/stretchr/testify
// Remove this comment before sending a pull request.

func TestAddLocalCommandTestSuite(t *testing.T) {
	suite.Run(t, new(AddLocalCommandTestSuite))
}
