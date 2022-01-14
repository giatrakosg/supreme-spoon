package main

import (
	"fmt"
	"github.com/cenkalti/rain/torrent"
	"time"
	"log"
	"os/signal"
	"os"
	"syscall"
)

func sessionTorrentInfo(ses *torrent.Session) {
	torrents := ses.ListTorrents()
	for i, tor := range torrents {
		stats := tor.Stats()
		fmt.Printf("%d: Name: %s, Bytes Completed: %d, Bytes Total %d \n", i, stats.Name, stats.Bytes.Completed, stats.Bytes.Total)
	}
}

func main()  {
	session, error := torrent.NewSession(torrent.DefaultConfig);
	if(error != nil) {
		fmt.Println("Error creating session");
	}

	fmt.Println("Torrents: ", session.ListTorrents())

	sigs := make(chan os.Signal, 1)
    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	
    go func() {
        sig := <-sigs
        fmt.Println()
        fmt.Println(sig)
		session.Close()
		os.Exit(1)
    }()

	file, _ := os.Open("./data/killing_chinese_bookie.torrent")

	sessionTorrentInfo(session)

	tor, _ := session.AddTorrent(file, nil)
	// Watch the progress
	for range time.Tick(time.Second) {
		s := tor.Stats()
		log.Printf("Status: %s, Downloaded: %d, Total %d, Peers: %d", s.Status.String(), s.Bytes.Completed , s.Bytes.Total, s.Peers.Total)
	}


}