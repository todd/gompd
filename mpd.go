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

	status.cStatus = mpdStatus
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

func (mpd Client) Play() (success bool, cerr error) {
	cSuccess, cerr := C.mpd_send_play(mpd.connection)
	success = bool(cSuccess)
	return
}

func (mpd Client) Close() {
	C.mpd_connection_free(mpd.connection)
}
