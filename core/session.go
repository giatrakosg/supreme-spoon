package core

import (
	"fmt"
	"log"

	"github.com/anacrolix/torrent"
	"github.com/dustin/go-humanize"
)

const MEGABYTE = 1e-6

func DownloadTorrent(path string) {
	c, _ := torrent.NewClient(nil)
	defer c.Close()
	t, _ := c.AddTorrentFromFile(path)
	<-t.GotInfo()
	t.DownloadAll()
	for t.BytesMissing() > 0 {
		fmt.Printf("Downloaded %s pieces, %s pieces are left \n", humanize.Bytes(uint64(t.BytesCompleted())), humanize.Bytes(uint64(t.BytesMissing())))
	}
	c.WaitAll()
	log.Print("ermahgerd, torrent downloaded")
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
