/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/giatrakosg/supreme-spoon/core"
	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for a torrent",
	Long: `We search the YTS api for a movie matching the provided string.
	
	eg. supreme-spoon search "The Killing of a Chinese Bookie"`,
	Run: func(cmd *cobra.Command, args []string) {
		searchMovie(args[0])
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func searchMovie(movieString string) {
	core.SearchMovie(movieString)
}
