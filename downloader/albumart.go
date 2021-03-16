package downloader

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/lakiluki1/lfmbg/api/types"
)

func DownloadAlbumArt(tracks *types.RecentTracksType, location string) error {

	track := tracks.Recenttracks.Track[0]

	url := strings.Replace(track.Image[0].Text, "34s", "700x700", -1)

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(location)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err

}
