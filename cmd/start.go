/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"
	"sync"
	"time"

	"github.com/spf13/cobra"
	tele "gopkg.in/telebot.v3"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var mu sync.Mutex
		todos := make(map[int64][]string)

		pref := tele.Settings{
			Token:  os.Getenv("TELE_TOKEN"),
			Poller: &tele.LongPoller{Timeout: 10 * time.Second},
		}

		b, err := tele.NewBot(pref)
		if err != nil {
			log.Fatal(err)
			return
		}

		b.Handle("/start", func(c tele.Context) error {
			return c.Send("Hello, it is basic telebot todo list bot, that supports multiple commands:\n/add <text> - adds new todo task\n/list - shows list of current unfinished tasks\n/done <number> - closes task by number")
		})

		b.Handle("/add", func(c tele.Context) error { mu.Lock(); defer mu.Unlock(); return c.Send("Not implemented") })
		b.Handle("/list", func(c tele.Context) error { mu.Lock(); defer mu.Unlock(); return c.Send("Not implemented") })
		b.Handle("/done", func(c tele.Context) error { mu.Lock(); defer mu.Unlock(); return c.Send("Not implemented") })
		b.Handle(tele.OnPhoto, func(c tele.Context) error { return c.Send("Not implemented") })
		b.Handle(tele.OnLocation, func(c tele.Context) error { return c.Send("Not implemented") })
		b.Handle(tele.OnSticker, func(c tele.Context) error { return c.Send("Not implemented") })
		b.Handle(tele.OnDocument, func(c tele.Context) error { return c.Send("Not implemented") })

		b.Start()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
