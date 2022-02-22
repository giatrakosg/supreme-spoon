package core

import (
	"fmt"
	"log"

	"github.com/manifoldco/promptui"
)

func SelectMovie(movies []Movie) {

	movieTemplate := &promptui.SelectTemplates{
		Label:    "{{ .Title | cyan}} {{ .Year | cyan}}",
		Active:   "{{ .Title | red }} {{ .Year | red}}",
		Inactive: "{{ .Title | cyan}} {{ .Year | cyan}}",
		Selected: "{{ .Title | cyan}} {{ .Year | cyan}}",
	}

	movieSelector := promptui.Select{
		Label:     "Select Movie",
		Items:     movies,
		Size:      20,
		Templates: movieTemplate,
	}

	pos, _, err := movieSelector.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	selectedMovie := movies[pos]
	torrentTemplate := &promptui.SelectTemplates{
		Label:    "Quality: {{ .Quality | cyan}} {{ .Seeds | green}}/{{ .Peers | red}} (s/p)",
		Active:   "Quality: {{ .Quality | red}} {{ .Seeds | green}}/{{ .Peers | red}} (s/p)",
		Inactive: "Quality: {{ .Quality | cyan}} {{ .Seeds | green}}/{{ .Peers | red}} (s/p)",
		Selected: "Quality: {{ .Quality | cyan}} {{ .Seeds | green}}/{{ .Peers | red}} (s/p)",
	}

	torrentSelector := promptui.Select{
		Label:     "Select Torrent",
		Items:     selectedMovie.Torrents,
		Size:      20,
		Templates: torrentTemplate,
	}

	posTorrent, _, err := torrentSelector.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	selectedTorrent := selectedMovie.Torrents[posTorrent]

	filePath := fmt.Sprintf("./data/%s.torrent", selectedMovie.Slug)

	err = downloadFile(filePath, selectedTorrent.Url)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("Downloading torrent")
	DownloadTorrent(filePath)
}

func ViewTorrents(torrents []TorrentInfo) {
	movieTemplate := &promptui.SelectTemplates{
		Label:    "{{ .Title | cyan}} {{ .Year | cyan}}",
		Active:   "{{ .Title | red }} {{ .Year | red}}",
		Inactive: "{{ .Title | cyan}} {{ .Year | cyan}}",
		Selected: "{{ .Title | cyan}} {{ .Year | cyan}}",
	}

	movieSelector := promptui.Select{
		Label:     "Select Movie",
		Items:     torrents,
		Size:      20,
		Templates: movieTemplate,
	}

	_, _, err := movieSelector.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

}
