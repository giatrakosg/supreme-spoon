package core

import (
	"log"

	"github.com/anacrolix/torrent"
	"github.com/schollz/progressbar/v3"
)

func DownloadTorrent(path string) {
	// Download torrent in path
	cfg := torrent.NewDefaultClientConfig()
	cfg.DataDir = "./data/"
	cfg.Debug = false

	c, _ := torrent.NewClient(cfg)
	defer c.Close()
	t, _ := c.AddTorrentFromFile(path)
	<-t.GotInfo()
	t.DownloadAll()

	bar := progressbar.DefaultBytes(t.BytesMissing(), "Dowloading torrent")
	for t.BytesMissing() > 0 {
		err := bar.Set(int(t.BytesCompleted()))
		if err != nil {
			log.Fatal(err)
		}
	}

	c.WaitAll()
}

type TorrentInfo struct {
	Name       string
	Downloaded int
	Total      int
}

func ListTorrents() {
	// torrent.DisableLogging()
	// session, error := torrent.NewSession(torrent.DefaultConfig)
	// if error != nil {
	// 	fmt.Println("Error creating session")
	// }
	// torrents := session.ListTorrents()
	// var torrentInfos []TorrentInfo
	// for _, tor := range torrents {
	// 	torrentInfos = append(torrentInfos, TorrentInfo{tor.Name(), int(tor.Stats().Bytes.Completed), int(tor.Stats().Bytes.Total)})
	// }
	// ViewTorrents(torrentInfos)
}
