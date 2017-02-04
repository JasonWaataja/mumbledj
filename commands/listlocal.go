/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/addlocal.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package commands

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/layeh/gumble/gumble"
	"github.com/matthieugrieger/mumbledj/bot"
	id3 "github.com/mikkyang/id3-go"
	"github.com/spf13/viper"
)

// IndentationString is the character used to indent string in CreateInfo. It is
// meant to be used with Mumble's HTML>
const IndentationString = "&nbsp"

// LineBreakString is the string used to break lines in CreateInfo. It is meant
// to be used with Mumble's HTML.
const LineBreakString = "<br>"

// ListLocal is a command that adds a tracks from the local filesystem. It
// searches the music directory if it exists and adds the track if the file is
// there.
type ListLocalCommand struct{}

// Aliases returns the current aliases for the command.
func (c *ListLocalCommand) Aliases() []string {
	return viper.GetStringSlice("commands.listlocal.aliases")
}

// Description returns the description for the command.
func (c *ListLocalCommand) Description() string {
	return viper.GetString("commands.listlocal.description")
}

// IsAdminCommand returns true if the command is only for admin use, and
// returns false otherwise.
func (c *ListLocalCommand) IsAdminCommand() bool {
	return viper.GetBool("commands.listlocal.is_admin")
}

// CreateInfoer is a type that can create information for itself. It is expected to be indented with the
type CreateInfoer interface {
	CreateInfo(indentation int) string
}

// SongInfo represents the basic inforamtion for an MP3 file that is used to
// print information on the song.
type SongInfo struct {
	SongName string
	Artist   string
	Duration time.Duration
}

// NewSongInfo creates an SongInfo for the given path. Returns a new SongInfo on
// success and nil on failure.
func NewSongInfo(path string) *SongInfo {
	var songInfo SongInfo
	reader, err := id3.Open(path)
	if err != nil {
		return nil
	}
	defer reader.Close()
	songInfo.SongName = reader.Title()
	songInfo.Artist = reader.Artist()
	songInfo.Duration, _ = bot.ReadMP3Duration(reader)
	return &songInfo
}

// SongDirectory represents the information on a directory that is used to print
// information about its contents. It does not know its path, scanning specific
// paths is left to the user.
type SongDirectory struct {
	Name    string
	Entries []interface{}
}

// NewSongDirectory creates a new SongDirectory with the given name.
func NewSongDirectory(name string) *SongDirectory {
	var songDir SongDirectory
	songDir.Name = name
	return &songDir
}

// CreateInfo returns a string representing songInfo. It contains the title,
// artist, and duration of the song. This string ends with a newline character.
// The string is indented with one tab character indentation times.
func (songInfo *SongInfo) CreateInfo(indentation int) string {
	var infoString string
	for i := 0; i < indentation; i++ {
		infoString += IndentationString
	}
	infoString += songInfo.SongName + " " + songInfo.Artist
	infoString += " (" + songInfo.Duration.String() + ")" + LineBreakString

	return infoString
}

// CreateInfo returns a string representing songDir. It incorporates the name of
// the directory and lists its contents with an additional indentation. The
// string is indented with one tab character indentation times.
//
// WARNING: THE ERROR MAY REVEAL INFORMATION ABOUT THE SONG DIRECTORY LOCATION.
// THIS SHOULD BE FIXED LATER.
func (songDir *SongDirectory) CreateInfo(indentation int) string {
	var infoString string
	for i := 0; i < indentation; i++ {
		infoString += IndentationString
	}
	infoString += songDir.Name + LineBreakString
	for _, entry := range songDir.Entries {
		switch t := entry.(type) {
		case *SongInfo:
			infoString += t.CreateInfo(indentation + 1)
		case *SongDirectory:
			infoString += t.CreateInfo(indentation + 1)
		}
	}
	return infoString
}

// ScanDirectory scans the given path and appends its contents to the entries of songDir.
//
// WARNING: THE ERROR MAY REVEAL INFORMATION ABOUT THE SONG DIRECTORY LOCATION.
// THIS SHOULD BE FIXED LATER.
func (songDir *SongDirectory) ScanDirectory(path string) error {
	dirInfo, err := os.Stat(path)
	if err != nil {
		return err
	}
	if !dirInfo.IsDir() {
		return errors.New(viper.GetString("commands.listlocal.messages.scan_non_directory_error"))
	}
	entries, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		entryPath := path + "/" + entry.Name()
		switch {
		case entry.Mode().IsRegular():
			songInfo := NewSongInfo(entryPath)
			if songInfo != nil {
				songDir.Entries = append(songDir.Entries, songInfo)
			}
		case entry.IsDir():
			newDir := NewSongDirectory(entry.Name())
			songDir.Entries = append(songDir.Entries, newDir)
			go newDir.ScanDirectory(entryPath)
		}
	}
	return nil
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
func (c *ListLocalCommand) Execute(user *gumble.User, args ...string) (string, bool, error) {
	// TODO: Fix the fact it sometimes directly returns the error. This may
	// reveal to those sending messages information about the filesystem.

	// If arguments were split around spaces, put them back together
	// separated by spaces.
	localPath := strings.Join(args, " ")

	path := bot.GetPathForLocalFile(localPath)
	path, err := bot.IsSafePath(path)

	if err != nil {
		return "", true, err
	}

	songDir := NewSongDirectory("Music Directory")
	err = songDir.ScanDirectory(path)

	if err != nil {
		return "", true, errors.New(viper.GetString("commands.listlocal.messages.scan_failure_error"))
	}
	infoString := songDir.CreateInfo(0)

	return infoString, true, nil
}
