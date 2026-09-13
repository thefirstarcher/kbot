/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/spf13/cobra"
	tele "gopkg.in/telebot.v3"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Запустити Telegram-бота",
	Long: `Запускає Telegram-бота

Токен читається зі змінної середовища TELE_TOKEN.

Команди бота:
  /add <текст>   додати задачу
  /list          показати список задач
  /done <номер>  закрити задачу за номером`,
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
			return c.Send("Це бот для трекінгу списку задач.\n\n" +
				"/add <текст> - додати задачу\n" +
				"/list - показати список задач\n" +
				"/done <номер> - закрити задачу за номером")
		})

		b.Handle("/add", func(c tele.Context) error {
			payload := strings.TrimSpace(c.Message().Payload)
			if payload == "" {
				return c.Send("Текст не може бути порожнім")
			}
			chatID := c.Chat().ID

			mu.Lock()
			todos[chatID] = append(todos[chatID], payload)
			n := len(todos[chatID])
			mu.Unlock()

			return c.Send(fmt.Sprintf("Додано %d. %s", n, payload))
		})
		b.Handle("/list", func(c tele.Context) error {
			chatID := c.Chat().ID

			mu.Lock()
			lines := make([]string, 0, len(todos[chatID]))
			for i, v := range todos[chatID] {
				lines = append(lines, fmt.Sprintf("%d. %s", i+1, v))
			}
			mu.Unlock()

			if len(lines) == 0 {
				return c.Send("Список порожній")
			}
			return c.Send(strings.Join(lines, "\n"))
		})
		b.Handle("/done", func(c tele.Context) error {
			payload := strings.TrimSpace(c.Message().Payload)
			n, err := strconv.Atoi(payload)
			if err != nil {
				return c.Send("Має бути число, наприклад /done 1")
			}
			chatID := c.Chat().ID
			mu.Lock()
			list := todos[chatID]
			if n < 1 || n > len(list) {
				mu.Unlock()
				return c.Send("Немає задачі з таким номером")
			}
			i := n - 1
			text := list[i]
			todos[chatID] = append(list[:i], list[i+1:]...)
			mu.Unlock()
			return c.Send(fmt.Sprintf("Готово: %s", text))
		})
		b.Handle(tele.OnPhoto, func(c tele.Context) error { return c.Send("Це фото") })
		b.Handle(tele.OnLocation, func(c tele.Context) error { return c.Send("Це геолокація") })
		b.Handle(tele.OnSticker, func(c tele.Context) error { return c.Send("Це стікер") })
		b.Handle(tele.OnDocument, func(c tele.Context) error { return c.Send("Це файл") })
		b.Handle(tele.OnText, func(c tele.Context) error {
			if strings.HasPrefix(c.Text(), "/") {
				return c.Send("Невідома команда")
			}
			return c.Send("Бот працює тільки з командами. /start — список")
		})

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
