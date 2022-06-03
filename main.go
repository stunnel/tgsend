package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	VERSION = "0.1.3"
	AUTHOR  = "travis"
	LICENSE = "MIT license"
	WEBSITE = "https://github.com/travislee8964/tgsend"
)

var (
	version     bool
	format      string
	pre         bool // preformatted fixed-width.
	preview     bool // disable link previews in the message(s).
	debug       bool
	token       string
	timeout     int
	ChatID      int64
	ChannelName string
	message     string
	filename    string
	filetype    string
	caption     string
	location    bool
	longitude   float64
	latitude    float64
	disNotice   bool
)

var (
	bot *tgbotapi.BotAPI
	err error
)

func initBot(token string, timeout int, debug bool) {
	bot, err = tgbotapi.NewBotAPI(token)
	if err != nil {
		fmt.Printf("Create bot failed. error msg: %v\n", err.Error())
		os.Exit(1)
	}

	bot.Debug = debug
	u := tgbotapi.NewUpdate(0)
	u.Timeout = timeout
}

func errorExit(errMsg string) {
	fmt.Printf("send message failed. error msg: %v\n", errMsg)
	os.Exit(1)
}

func sendMessage(message string) {
	var msg tgbotapi.MessageConfig
	if ChatID != 0 {
		msg = tgbotapi.NewMessage(ChatID, message)
	} else if len(ChannelName) != 0 {
		msg = tgbotapi.NewMessageToChannel(ChannelName, message)
	} else {
		os.Exit(1)
	}

	msg.DisableWebPagePreview = preview
	msg.DisableNotification = disNotice

	if pre || format == "markdown" {
		msg.ParseMode = tgbotapi.ModeMarkdown
	} else if format == "html" {
		msg.ParseMode = tgbotapi.ModeHTML
	}

	_, err := bot.Send(msg)
	if err != nil {
		errorExit(err.Error())
	}
}

func sendLocation(latitude float64, longitude float64) {
	if longitude < -180 || longitude > 180 || latitude < -90 || latitude > 90 {
		fmt.Printf("Longitude or latitude value invalid: %v, %v\n", latitude, longitude)
		os.Exit(1)
	}

	msg := tgbotapi.NewLocation(ChatID, latitude, longitude)

	_, err = bot.Send(msg)
	if err != nil {
		errorExit(err.Error())
	}
}

func sendFile(filename string, filetype string, caption string) {
	stat, err := os.Stat(filename)
	if err != nil {
		fmt.Printf("Reading file %v error: %v\n", filename, err.Error())
		os.Exit(1)
	}

	if stat.IsDir() {
		fmt.Printf("%v is a directory.\n", filename)
		os.Exit(1)
	}

	if stat.Size() > 1024*1024*50 {
		fmt.Printf("File %v is too large, size: %v\n", filename, stat.Size())
		os.Exit(1)
	}

	if caption == "" {
		caption = path.Base(filename)
	}

	reader, _ := os.Open(filename)
	file := tgbotapi.FileReader{
		Name:   caption,
		Reader: reader,
	}

	var msg tgbotapi.Chattable
	switch filetype {
	case "photo":
		msg = tgbotapi.NewPhoto(ChatID, file)
	case "video":
		msg = tgbotapi.NewVideo(ChatID, file)
	case "audio":
		msg = tgbotapi.NewAudio(ChatID, file)
	case "sticker":
		msg = tgbotapi.NewSticker(ChatID, file)
	case "animation":
		msg = tgbotapi.NewAnimation(ChatID, file)
	default:
		msg = tgbotapi.NewDocument(ChatID, file)
	}

	_, err = bot.Send(msg)
	if err != nil {
		errorExit(err.Error())
	}
}

func checkParam() {
	if ChatID == 0 && len(ChannelName) == 0 {
		fmt.Println("ChatID or ChannelName must be set.")
		os.Exit(1)
	}

	if len(token) < 45 {
		fmt.Printf("Token is invalid: %v\n", token)
		os.Exit(1)
	}
}

func getMessage() {
	if message == "-" { // stdin
		var (
			scanner *bufio.Scanner
			line    string
		)

		scanner = bufio.NewScanner(os.Stdin)
		message = ""

		for scanner.Scan() {
			line = scanner.Text()
			if len(line) == 1 && line[0] == '\x1D' {
				break
			}
			message = strings.Join([]string{message, line, "\n"}, "")
		}

		if scanner.Err() != nil {
			fmt.Printf("Read stdin error: %v\n", scanner.Err())
			os.Exit(1)
		}
	}
}

func versionInfo() {
	fmt.Println("Send message via Telegram Bot.")
	fmt.Println("Version:", VERSION)
	fmt.Println("Author: ", AUTHOR)
	fmt.Println("License:", LICENSE)
	fmt.Println("Website:", WEBSITE)
	os.Exit(0)
}

func main() {
	initFlag()
	if flag.NFlag() == 0 || version {
		versionInfo()
	}

	checkParam()
	getMessage()
	initBot(token, timeout, debug)

	if len(message) != 0 {
		sendMessage(message)
		os.Exit(0)
	}

	if location {
		sendLocation(latitude, longitude)
		os.Exit(0)
	}

	if len(filename) != 0 {
		sendFile(filename, filetype, caption)
		os.Exit(0)
	}
}

func initFlag() {
	flag.BoolVar(&version, "version", false, "Print version information.")
	flag.StringVar(&format, "format", "text", "How to format the message(s). "+
		"Choose from ['text', 'markdown', 'html']")
	flag.BoolVar(&pre, "pre", false, "Send preformatted fixed-width (monospace) text.")
	flag.BoolVar(&preview, "preview", false, "disable link previews in the message(s)")
	flag.BoolVar(&debug, "debug", false, "Show debug message.")
	flag.StringVar(&token, "token", "", "Set the bot token.")
	flag.IntVar(&timeout, "timeout", 30, "Set the read timeout for network operations(in seconds).")
	flag.Int64Var(&ChatID, "chatid", 0, "Send message to this chatID.")
	flag.StringVar(&ChannelName, "channel", "", "Send message to the public channel.")
	flag.StringVar(&message, "message", "", "The message to sent.")
	flag.StringVar(&filename, "filename", "", "The file to sent. images up to 10 MiB, files up to 50 MiB.")
	flag.StringVar(&filetype, "filetype", "document",
		"Set the file type, Choose from ['photo', 'video', 'document', 'audio', 'sticker', 'animation']")
	flag.StringVar(&caption, "caption", "", "Set the photo/video/document caption")
	flag.BoolVar(&location, "location", false, "Send location")
	flag.Float64Var(&longitude, "longitude", 0, "Set longitude, value valid [-180, 180]")
	flag.Float64Var(&latitude, "latitude", 0, "Set latitude, value valid [-90, 90]")
	flag.BoolVar(&disNotice, "disable_notification", false, "Disable notification")
	flag.Parse()
}
