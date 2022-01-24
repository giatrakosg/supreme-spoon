package core

import (
	"fmt"
	"os"
	"time"

	"github.com/cenkalti/rain/torrent"
	"github.com/schollz/progressbar/v3"
)

const MEGABYTE = 1e-6

func DownloadTorrent(path string) {
	torrent.DisableLogging()

	cfg := torrent.DefaultConfig

	cfg.DataDirIncludesTorrentID = false
	// cfg.DataDirIncludesTorrentID = false
	session, error := torrent.NewSession(cfg)
	if error != nil {
		fmt.Println("Error creating session")
	}

	file, _ := os.Open(path)
	tor, _ := session.AddTorrent(file, nil)

	s := tor.Stats()

	totalMb := float64(s.Bytes.Total) * MEGABYTE
	bar := progressbar.Default(int64(totalMb))

	// Watch the progress
	for range time.Tick(time.Second) {
		s := tor.Stats()
		bar.Set64(int64(float64(s.Bytes.Completed) * MEGABYTE))
	}
}

type TorrentInfo struct {
	Name       string
	Downloaded int
	Total      int
}

func ListTorrents() {
	torrent.DisableLogging()
	session, error := torrent.NewSession(torrent.DefaultConfig)
	if error != nil {
		fmt.Println("Error creating session")
	}
	torrents := session.ListTorrents()
	var torrentInfos []TorrentInfo
	for _, tor := range torrents {
		torrentInfos = append(torrentInfos, TorrentInfo{tor.Name(), int(tor.Stats().Bytes.Completed), int(tor.Stats().Bytes.Total)})
	}
	ViewTorrents(torrentInfos)
}
