package mpd_test

import (
	. "github.com/todd/gompd"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"reflect"
)

var _ = Describe("Mpd", func() {
	Describe("Initializing MPD Client", func() {
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

		Describe("Getting the current song", func() {
			Context("With currently playing song", func() {
				It("returns the currently playing song", func() {
					song, err := client.GetCurrentSong()

					Expect(err).ToNot(HaveOccurred())
					Expect(song.Artist).To(Equal("Kenny Beltrey"))
					Expect(song.Title).To(Equal("Hydrate - Kenny Beltrey"))
					song.Free()
				})
			})
		})

		Describe("Getting the current status", func() {
			Context("With the current status", func() {
				It("returns the current status", func() {
					status, err := client.GetStatus()
					objectType := reflect.TypeOf(status)

					Expect(err).ToNot(HaveOccurred())
					Expect(objectType.String()).To(Equal("mpd.Status"))
					status.Free()
				})
			})
		})
	})
})
