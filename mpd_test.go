package mpd_test

import (
	. "github.com/todd/gompd"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"reflect"
)

var _ = Describe("Mpd", func() {
	Describe("Init", func() {
		It("returns a client", func() {
			client, err := Init(host, port, timeoutms)
			Expect(err).ToNot(HaveOccurred())
			client.Close()
		})
	})

	Context("With initialized client", func() {
		var client Client

		BeforeEach(func() {
			client, _ = Init(host, port, timeoutms)
		})

		AfterEach(func() {
			client.Close()
		})

		Describe("GetCurrentSong", func() {
			It("returns the currently playing song", func() {
				song, err := client.GetCurrentSong()

				Expect(err).ToNot(HaveOccurred())
				Expect(song.Artist).To(Equal("Kenny Beltrey"))
				Expect(song.Title).To(Equal("Hydrate - Kenny Beltrey"))
				song.Free()
			})
		})

		Describe("GetStatus", func() {
			It("returns the current status", func() {
				status, err := client.GetStatus()
				objectType := reflect.TypeOf(status)

				Expect(err).ToNot(HaveOccurred())
				Expect(objectType.String()).To(Equal("mpd.Status"))
				status.Free()
			})
		})

		Describe("Play", func() {
			It("returns true if the song starts playing", func() {
				success, err := client.Play()

				Expect(err).ToNot(HaveOccurred())
				Expect(success).To(BeTrue())
			})
		})
	})
})
