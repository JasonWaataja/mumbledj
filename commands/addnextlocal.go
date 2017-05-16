/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/addnextlocal.go
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

// AddNextLocalCommand is a command... (put a short description of the command here)
type AddNextLocalCommand struct{}

// Aliases returns the current aliases for the command.
func (c *AddNextLocalCommand) Aliases() []string {
	return viper.GetStringSlice("commands.addnextlocal.aliases")
}

// Description returns the description for the command.
func (c *AddNextLocalCommand) Description() string {
	return viper.GetString("commands.addnextlocal.description")
}

// IsAdminCommand returns true if the command is only for admin use, and
// returns false otherwise.
func (c *AddNextLocalCommand) IsAdminCommand() bool {
	return viper.GetBool("commands.addnextlocal.is_admin")
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
func (c *AddNextLocalCommand) Execute(user *gumble.User, args ...string) (string, bool, error) {
	if len(args) == 0 {
		return "", true, errors.New(viper.GetString("commands.addlocal.messages.no_argument_error"))
	}
	// If arguments were split around spaces, put them back together
	// separated by spaces.
	localPath := strings.Join(args, " ")

	path := bot.GetPathForLocalFile(localPath)
	path, err := bot.GetSafePath(path)
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
					return "", true, err
				}
				tracks = append(tracks, newTracks...)
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
	for i := len(tracks) - 1; i >= 0; i++ {
		if err := DJ.Queue.InsertTrack(1, tracks[i]); err != nil {
			numTooLong++
		} else {
			numAdded++
			lastTrackAdded = tracks[i]
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
