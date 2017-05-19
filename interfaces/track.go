/*
 * MumbleDJ
 * By Matthieu Grieger
 * interfaces/track.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package interfaces

import "time"

// Track is an interface of methods that must be implemented by tracks.
type Track interface {
	IsLocal() bool
	GetID() string
	GetURL() string
	GetTitle() string
	GetAuthor() string
	GetAuthorURL() string
	GetSubmitter() string
	GetService() string
	DownloadIfNeeded() error
	GetFilename() string
	GetFullPath() string
	GetThumbnailURL() string
	GetDuration() time.Duration
	GetPlaybackOffset() time.Duration
	GetPlaylist() Playlist
}
