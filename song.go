package mpd

/*
#cgo pkg-config: libmpdclient

#include <stdlib.h>
#include <mpd/client.h>
*/
import "C"

type Song struct {
	cSong       *C.struct_mpd_song
	Artist      string
	Album       string
	AlbumArtist string
	Title       string
	Track       string
	Name        string
}

// MPD Tag Types
const (
	MPD_TAG_UNKNOWN                   = C.MPD_TAG_UNKNOWN
	MPD_TAG_ARTIST                    = C.MPD_TAG_ARTIST
	MPD_TAG_ALBUM                     = C.MPD_TAG_ALBUM
	MPD_TAG_ALBUM_ARTIST              = C.MPD_TAG_ALBUM_ARTIST
	MPD_TAG_TITLE                     = C.MPD_TAG_TITLE
	MPD_TAG_TRACK                     = C.MPD_TAG_TRACK
	MPD_TAG_NAME                      = C.MPD_TAG_NAME
	MPD_TAG_GENRE                     = C.MPD_TAG_GENRE
	MPD_TAG_DATE                      = C.MPD_TAG_DATE
	MPD_TAG_COMPOSER                  = C.MPD_TAG_COMPOSER
	MPD_TAG_PERFORMER                 = C.MPD_TAG_PERFORMER
	MPD_TAG_COMMENT                   = C.MPD_TAG_COMMENT
	MPD_TAG_DISC                      = C.MPD_TAG_DISC
	MPD_TAG_MUSICBRAINZ_ARTISTID      = C.MPD_TAG_MUSICBRAINZ_ARTISTID
	MPD_TAG_MUSICBRAINZ_ALBUMID       = C.MPD_TAG_MUSICBRAINZ_ALBUMID
	MPD_TAG_MUSICBRAINZ_ALBUMARTISTID = C.MPD_TAG_MUSICBRAINZ_ALBUMARTISTID
	MPD_TAG_MUSICBRAINZ_TRACKID       = C.MPD_TAG_MUSICBRAINZ_TRACKID
	MPD_TAG_COUNT                     = C.MPD_TAG_COUNT
)

func (song Song) GetUri() (uri string, cerr error) {
	songUri, cerr := C.mpd_song_get_uri(song.cSong)
	uri = C.GoString(songUri)
	return
}

func (song Song) getArtist() (artist string, cerr error) {
	songArtist, cerr := C.mpd_song_get_tag(song.cSong, MPD_TAG_ARTIST, 0)
	artist = C.GoString(songArtist)
	return
}

func (song Song) getAlbum() (album string, cerr error) {
	songAlbum, cerr := C.mpd_song_get_tag(song.cSong, MPD_TAG_ALBUM, 0)
	album = C.GoString(songAlbum)
	return
}

func (song Song) getTitle() (title string, cerr error) {
	songTitle, cerr := C.mpd_song_get_tag(song.cSong, MPD_TAG_TITLE, 0)
	title = C.GoString(songTitle)
	return
}

func (song Song) getTrack() (track string, cerr error) {
	songTrack, cerr := C.mpd_song_get_tag(song.cSong, MPD_TAG_TRACK, 0)
	track = C.GoString(songTrack)
	return
}

func (song Song) getName() (name string, cerr error) {
	songName, cerr := C.mpd_song_get_tag(song.cSong, MPD_TAG_NAME, 0)
	name = C.GoString(songName)
	return
}

func (song Song) Free() {
	C.mpd_song_free(song.cSong)
}
