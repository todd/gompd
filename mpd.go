package mpd

/*
#cgo pkg-config: libmpdclient

#include <stdlib.h>
#include <mpd/client.h>
*/
import "C"

import (
	"unsafe"
)

type Client struct {
	Connection connection
}

type connection struct {
	connection   *C.struct_mpd_connection
	ErrorMessage string
}

type Status struct {
	Status *C.struct_mpd_status
}

func InitClient(host string, port int, timeoutMS int) (mpd Client) {
	chost := C.CString(host)
	cport := C.uint(port)
	ctimeoutMS := C.uint(timeoutMS)

	connection, _ := C.mpd_connection_new(chost, cport, ctimeoutMS)

	C.free(unsafe.Pointer(chost))

	if connection != nil {
		mpd.Connection.connection = connection
	} else {
		errorMessage := C.mpd_connection_get_error_message(connection)
		mpd.Connection.ErrorMessage = C.GoString(errorMessage)
	}

	return
}

func (mpd Client) Status() (status Status, cerr error) {
	mpdStatus, cerr := C.mpd_run_status(mpd.Connection.connection)

	status.Status = mpdStatus
	return
}

func (status Status) GetSongID() (id int, cerr error) {
	cid, cerr := C.mpd_status_get_song_id(status.Status)
	id = int(cid)
	return
}

func (mpd Client) Close() {
	C.mpd_connection_free(mpd.Connection.connection)
}
