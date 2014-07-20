package mpd

/*
#cgo pkg-config: libmpdclient

#include <stdlib.h>
#include <mpd/client.h>
*/
import "C"

import (
	"errors"
	"unsafe"
)

type Client struct {
	connection *C.struct_mpd_connection
}

type Status struct {
	Status *C.struct_mpd_status
}

func Init(host string, port int, timeoutMS int) (mpd Client, err error) {
	chost := C.CString(host)
	cport := C.uint(port)
	ctimeoutMS := C.uint(timeoutMS)

	connection, err := C.mpd_connection_new(chost, cport, ctimeoutMS)

	C.free(unsafe.Pointer(chost))

	connection_error := C.mpd_connection_get_error(connection)

	if connection_error == 0 {
		mpd.connection = connection
		return mpd, nil
	} else {
		error_message := C.mpd_connection_get_error_message(connection)
		err = errors.New(C.GoString(error_message))
	}

	return
}

func (mpd Client) GetStatus() (status Status, cerr error) {
	mpdStatus, cerr := C.mpd_run_status(mpd.connection)

	status.Status = mpdStatus
	return
}

func (status Status) GetSongID() (id int, cerr error) {
	cid, cerr := C.mpd_status_get_song_id(status.Status)
	id = int(cid)
	return
}

func (mpd Client) GetCurrentSong() (song Song, cerr error) {
	mpdSong, cerr := C.mpd_run_current_song(mpd.connection)
	song.cSong = mpdSong
	song.Artist, _ = song.getArtist()
	song.Album, _ = song.getAlbum()
	song.Title, _ = song.getTitle()
	song.Track, _ = song.getTrack()
	song.Name, _ = song.getName()
	return
}

func (mpd Client) Close() {
	C.mpd_connection_free(mpd.connection)
}

func (status Status) StatusFree() {
	C.mpd_status_free(status.Status)
}
