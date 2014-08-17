package mpd_test

import (
	. "github.com/todd/gompd"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Mpd", func() {
	Describe("Initializing MPD Client", func() {
		It("returns a client", func() {
			client, err := Init(host, port, timeoutms)
			Expect(err).ToNot(HaveOccurred())
			client.Close()
		})
	})

	Describe("Getting the current song", func() {
		var client Client

		BeforeEach(func() {
			client, _ = Init(host, port, timeoutms)
		})

		AfterEach(func() {
			client.Close()
		})

		Context("With currently playing song", func() {
			It("returns the currently playing song", func() {
				song, err := client.GetCurrentSong()

				Expect(err).ToNot(HaveOccurred())
				Expect(song.Artist).To(Equal("Kenny Beltrey"))
				Expect(song.Title).To(Equal("Hydrate - Kenny Beltrey"))
			})
		})
	})
})
