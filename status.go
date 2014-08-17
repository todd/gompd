package mpd

/*
#cgo pkg-config: libmpdclient

#include <stdlib.h>
#include <mpd/client.h>
*/
import "C"

type Status struct {
	cStatus *C.struct_mpd_status
}

func (status Status) GetSongID() (id int, cerr error) {
	cid, cerr := C.mpd_status_get_song_id(status.cStatus)
	id = int(cid)
	return
}

func (status Status) StatusFree() {
	C.mpd_status_free(status.cStatus)
}
