/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the desired Gopher",
	Long: `This command will get the desired Gopher image at a github repo`,
  Run: func(cmd *cobra.Command, args []string) {
    var gopherName = "GOOGLE"

    if len(args) >= 1 && args[0] != "" {
      gopherName = args[0]
    }

    URL := "https://raw.githubusercontent.com/vitortrimer/gophers/master/" + gopherName + ".png"
    fmt.Println("Yeah, ok... I will get the " + gopherName + " for you!")

		response, err := http.Get(URL)
    if err != nil {
			fmt.Println(err)
    }
		defer response.Body.Close()

		if response.StatusCode == 200 {
      out, err := os.Create(gopherName + ".png")
      if err != nil {
        fmt.Println(err)
      }
      defer out.Close()

      _, err = io.Copy(out, response.Body)
      if err != nil {
        fmt.Println(err)
      }

      fmt.Println("Nice one! Saved your Gopher image at " + out.Name())
    } else {
      fmt.Println("Error: " + gopherName + " does not exist! 😞")
    }
  },
}

func init() {
	rootCmd.AddCommand(getCmd)
}
