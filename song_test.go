package mpd_test

import (
	. "github.com/todd/gompd"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Song", func() {
	var client Client
	var song Song

	BeforeEach(func() {
		client, _ = Init(host, port, timeoutms)
		song, _ = client.GetCurrentSong()
	})

	AfterEach(func() {
		song.Free()
		client.Close()
	})

	Describe("GetUri", func() {
		It("returns the song's URI", func() {
			uri, err := song.GetUri()

			Expect(err).ToNot(HaveOccurred())
			Expect(uri).To(Equal("Hydrate-Kenny_Beltrey.ogg"))
		})
	})
})
