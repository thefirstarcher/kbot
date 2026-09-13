# kbot

A Telegram bot for keeping a to-do list. Course project for **DevOps and Kubernetes. Practical Intensive+** (Prometheus and GlobalLogic), module 2 "Version Control Systems".

Bot: [t.me/devopscoursetelebot](https://t.me/devopscoursetelebot)

## Features

- keeps a separate task list for every chat
- handles messages by type: text, photo, location, sticker, document
- replies to unknown commands

Tasks are kept in process memory, so the list is empty after a restart.

## Commands

| Command        | Description                    |
| -------------- | ------------------------------ |
| `/start`       | greeting and list of commands  |
| `/add <text>`  | add a task                     |
| `/list`        | show the task list             |
| `/done <n>`    | close a task by its number     |

Usage example:

```
/add купити хліб
→ Додано 1. купити хліб

/add подзвонити мамі
→ Додано 2. подзвонити мамі

/list
→ 1. купити хліб
  2. подзвонити мамі

/done 1
→ Готово: купити хліб
```

The bot replies in Ukrainian.

## Stack

- [Go](https://go.dev/) 1.26
- [spf13/cobra](https://github.com/spf13/cobra) for the CLI
- [telebot.v3](https://gopkg.in/telebot.v3) for the Telegram Bot API

## Installation

Requires Go 1.26 or newer.

```bash
git clone https://github.com/thefirstarcher/kbot.git
cd kbot
go build
```

## Configuration

1. Create a bot with [@BotFather](https://t.me/BotFather) using `/newbot`
2. Get the token and store it in an environment variable:

```bash
export TELE_TOKEN='<your token>'
```

## Running

```bash
./kbot start
```

Or without building:

```bash
go run . start
```

## CLI commands

```bash
./kbot          # help
./kbot start    # run the bot
./kbot version  # print version
```
