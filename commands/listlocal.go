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
	"time"

	"github.com/layeh/gumble/gumble"
	id3 "github.com/mikkyang/id3-go"
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
func (c *ListlocalCommand) IsAdminCommand() bool {
	return viper.GetBool("commands.listlocal.is_admin")
}

type MP3Info struct {
	SongName string
	Artist   string
	Duration time.Duration
}

func NewMP3Info(path string) *MP3Info {
	var songInfo MP3Info
	reader, err := id3.Open(path)
	if err != nil {
		return nil
	}
	defer reader.Close()
	songInfo.SongName = reader.Title()
	songInfo.Artist = reader.Artist()
	// TODO: Add duration.
	return &songInfo
}

type SongDirectory struct {
	Name    string
	Entries []interface{}
}

func NewSongDirectory(name string) *SongDirectory {
	var songDir SongDirectory
	songDir.Name = name
	return &songDir
}

func (songInfo *MP3Info) CreateInfo(indentation int) string {
	var infoString string
	for i := 0; i < indentation; i++ {
		infoString += " "
	}
	infoString += songInfo.SongName + " " + songInfo.Artist
	infoString += " (" + songInfo.Duration.String() + ")\n"

	return infoString
}

func (songDir *SongDirectory) CreateInfo(indentation int) string {
	var infoString string
	for i := 0; i < indentation; i++ {
		infoString += " "
	}
	infoString += songDir.Name + "\n"
	for _, entry := range songDir.Entries {
		switch t := entry.(type) {
		case *MP3Info:
			infoString += t.CreateInfo(indentation + 1)
		case *SongDirectory:
			infoString += t.CreateInfo(indentation + 1)
		}
	}
	return infoString
}

func (songDir *SongDirectory) ScanDirectory(path string) error {
	dirInfo, err := os.Stat(path)
	if err != nil {
		return err
	}
	if !dirInfo.IsDir() {
		// TODO: Replace this with a config message.
		return errors.New("Attemping to scan non-directory for entries")
	}
	entries, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		entryPath := path + "/" + entry.Name()
		switch {
		case entry.Mode().IsRegular():
			songInfo := NewMP3Info(entryPath)
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
func (c *ListlocalCommand) Execute(user *gumble.User, args ...string) (string, bool, error) {
	return "", true, nil
}
