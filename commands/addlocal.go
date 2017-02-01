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
	"os"
	"strings"

	"github.com/layeh/gumble/gumble"
	"github.com/matthieugrieger/mumbledj/bot"
	"github.com/matthieugrieger/mumbledj/interfaces"
	"github.com/matthieugrieger/mumbledj/services"
	"github.com/spf13/viper"
)

// AddLocalCommand is a command that adds a tracks from the local filesystem. It
// searches the music directory if it exists and adds the track if the file is
// there.
type AddLocalCommand struct{}

// Aliases returns the current aliases for the command.
func (c *AddLocalCommand) Aliases() []string {
	return viper.GetStringSlice("commands.addlocal.aliases")
}

// Description returns the description for the command.
func (c *AddLocalCommand) Description() string {
	return viper.GetString("commands.addlocal.description")
}

// IsAdminCommand returns true if the command is only for admin use, and
// returns false otherwise.
func (c *AddLocalCommand) IsAdminCommand() bool {
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
func (c *AddLocalCommand) Execute(user *gumble.User, args ...string) (string, bool, error) {
	if len(args) == 0 {
		return "", true, errors.New(viper.GetString("commands.addlocal.messages.no_argument_error"))
	}

	// If arguments were split around spaces, put them back together
	// separated by spaces.
	localPath := strings.Join(args, " ")

	path := bot.GetPathForLocalFile(localPath)
	path, err := bot.IsSafePath(path)

	if err != nil {
		return "", true, err
	}

	if _, err := os.Stat(path); err != nil {
		return "", true, errors.New(viper.GetString("commands.addlocal.messages.no_matching_song_error"))
	}

	tracks := make([]interfaces.Track, 0)

	if bot.PathIsSong(path) {
		found := false
		for _, service := range DJ.AvailableServices {
			fs, ok := service.(*services.Filesystem)
			if ok {
				track, err := fs.CreateTrackForLocalFile(localPath, user)
				if err != nil {
					return "", true, err
				} else {
					tracks = append(tracks, track)
					found = true
				}
			}
		}
		if !found {
			return "", true, errors.New(viper.GetString("commands.addlocal.messages.no_filesystem_service_error"))
		}
	} else if bot.PathIsPlaylist(path) {
		for _, service := range DJ.AvailableServices {
			fs, ok := service.(*services.Filesystem)
			if ok {
				newTracks, err := fs.CreateTracksForLocalFile(localPath, user)
				if err != nil {
					for _, track := range newTracks {
						tracks = append(tracks, track)
					}
					return "", true, err
				}
			}
		}
	} else {
		return "", true, errors.New(viper.GetString("commands.addlocal.messages.unrecognized_filetype_error"))
	}

	if len(tracks) == 0 {
		return "", true, errors.New(viper.GetString("commands.addlocal.messages.no_valid_tracks_error"))
	}

	numTooLong := 0
	numAdded := 0
	var lastTrackAdded interfaces.Track
	for _, track := range tracks {
		if err := DJ.Queue.AppendTrack(track); err != nil {
			numTooLong++
		} else {
			numAdded++
			lastTrackAdded = track
		}
	}

	if numAdded == 0 {
		return "", true, errors.New(viper.GetString("commands.addlocal.messages.tracks_too_long_error"))
	} else if numAdded == 1 {
		return fmt.Sprintf(viper.GetString("commands.addlocal.messages.one_track_added"), user.Name, lastTrackAdded.GetTitle(), lastTrackAdded.GetService()), false, nil
	}

	retString := fmt.Sprintf(viper.GetString("commands.addlocal.messages.many_tracks_added"), user.Name, numAdded)
	if numTooLong != 0 {
		retString += fmt.Sprintf(viper.GetString("commands.addlocal.messages.num_tracks_too_long"), numTooLong)
	}
	return retString, false, nil
}
