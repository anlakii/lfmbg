package main

import (
	"fmt"
	"time"

	"github.com/lakiluki1/lfmbg/api"
	"github.com/lakiluki1/lfmbg/config"
	"github.com/lakiluki1/lfmbg/downloader"
	"github.com/lakiluki1/lfmbg/process"
)

func main() {

	var lastImg string

	conf, err := config.GetConfig()
	if err != nil {
		panic(err)
	}

	for {
		tracks, err := api.GetRecentTracks(1)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if lastImg == tracks.Recenttracks.Track[0].Image[0].Text {
			continue
		}

		fmt.Println("[NEW ALBUM]", tracks.Recenttracks.Track[0].Album.Text)

		err = downloader.DownloadAlbumArt(tracks, "/tmp/cover.jpg")
		lastImg = tracks.Recenttracks.Track[0].Image[0].Text

		if err != nil {
			fmt.Println(err)
		}

		err = process.Process("/tmp/cover.jpg", conf)

		if err != nil {
			panic(err)
		}

		time.Sleep(time.Duration(conf.Interval) * time.Second)

		conf, err = config.GetConfig()
	}

}
