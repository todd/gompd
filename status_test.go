package mpd_test

import (
	. "github.com/todd/gompd"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Status", func() {
  var client Client
  var status Status

  BeforeEach(func() {
    client, _ = Init(host, port, timeoutms)
    status, _ = client.GetStatus()
  })

  AfterEach(func() {
    status.Free()
    client.Close()
  })

  Describe("GetSongID", func() {
    It("Returns the ID of the current song", func() {
      id, err := status.GetSongID()

      Expect(err).ToNot(HaveOccurred())
      Expect(id).ToNot(Equal(-1)) // MPD returns -1 if no current song
    })
  })
})
