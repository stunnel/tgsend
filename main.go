package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"os"
	"strings"
)

var _version_ = "0.1"
var (
	version		bool
	format 		string
	pre 		bool	// preformatted fixed-width
	preview 	bool  	// disable link previews in the message(s)
	debug		bool
	token 		string
	timeout 	int
	ChatID 		int64
	ChannelName string
	message 	string
	filename 	string
	filetype 	string
	caption 	string
	location 	bool
	longitude 	float64
	latitude 	float64
	disNotice	bool
)

func Bot(token string, timeout int, debug bool) *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = debug
	u := tgbotapi.NewUpdate(0)
	u.Timeout = timeout

	return bot
}

func errorExit(err string) {
	log.Panic("send message failed. error msg: ", err)
}

func sendMessage(bot *tgbotapi.BotAPI, message string) {
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

	if pre == true || format == "markdown" {
		msg.ParseMode = tgbotapi.ModeMarkdown
	} else if format == "html" {
		msg.ParseMode = tgbotapi.ModeHTML
	}

	_, err := bot.Send(msg)
	if err != nil {
		errorExit(err.Error())
	}
}

func sendPhoto(bot *tgbotapi.BotAPI, filename string, caption string) {
	var msg tgbotapi.PhotoConfig
	msg = tgbotapi.NewPhotoUpload(ChatID, filename)
	if len(caption) != 0 {
		msg.Caption = caption
	}
	_, err := bot.Send(msg)
	if err != nil {
		errorExit(err.Error())
	}
}

func sendVideo(bot *tgbotapi.BotAPI, filename string, caption string) {
	var msg tgbotapi.VideoConfig
	msg = tgbotapi.NewVideoUpload(ChatID, filename)
	if len(caption) != 0 {
		msg.Caption = caption
	}
	_, err := bot.Send(msg)
	if err != nil {
		errorExit(err.Error())
	}
}

func sendDocument(bot *tgbotapi.BotAPI, filename string, caption string) {
	var msg tgbotapi.DocumentConfig
	msg = tgbotapi.NewDocumentUpload(ChatID, filename)
	if len(caption) != 0 {
		msg.Caption = caption
	}
	_, err := bot.Send(msg)
	if err != nil {
		errorExit(err.Error())
	}
}

func sendLocation(bot *tgbotapi.BotAPI, latitude float64, longitude float64) {
	var msg tgbotapi.LocationConfig
	msg = tgbotapi.NewLocation(ChatID, latitude, longitude)
	_, err := bot.Send(msg)
	if err != nil {
		errorExit(err.Error())
	}
}

func init() {
	flag.BoolVar(&version, "version", false, "Print version information.")
	flag.StringVar(&format, "format", "text", "How to format the message(s). " +
		"Choose from ['text', 'markdown', 'html']")
	flag.BoolVar(&pre, "pre", false, "Send preformatted fixed-width (monospace) text.")
	flag.BoolVar(&preview, "preview", false, "disable link previews in the message(s)")
	flag.BoolVar(&debug, "debug", false, "Show debug message.")
	flag.StringVar(&token, "token", "", "Set the bot token.")
	flag.IntVar(&timeout, "timeout", 30, "Set the read timeout for network operations(in seconds).")
	flag.Int64Var(&ChatID, "chatid", 0, "Send message to this chatID.")
	flag.StringVar(&ChannelName, "channel", "", "Send message to the public channel.")
	flag.StringVar(&message, "message", "", "The message to sent.")
	flag.StringVar(&filename, "filename", "", "The file to sent.")
	flag.StringVar(&filetype, "filetype", "document", "Set the file type, " +
		"Choose from ['photo', 'video', 'document']")
	flag.StringVar(&caption, "caption", "","Set the photo/video/document caption")
	flag.BoolVar(&location, "location", false, "Send location")
	flag.Float64Var(&longitude, "longitude", 0, "Set longitude, value valid [-180, 180]")
	flag.Float64Var(&latitude, "latitude", 0, "Set latitude, value valid [-90, 90]")
	flag.BoolVar(&disNotice, "disable_notification", false, "Disable notification")
}

func main() {
	flag.Parse()

	if flag.NFlag() == 0 {
		version = true
	}

	if version {
		fmt.Println("telegram send version:", _version_)
		os.Exit(0)
	}

	if ChatID == 0 && len(ChannelName) == 0 {
		log.Panic("Does not provide chatid or channel name.")
	}
	if len(token) < 45 {
		log.Panic("token is invalid.")
	}

	if message == "-" {	// stdin
		scanner := bufio.NewScanner(os.Stdin)
		message = ""
		for scanner.Scan() {
			line := scanner.Text()
			if len(line) == 1 && line[0] == '\x1D' {
				break
			}
			message = strings.Join([]string{message, line, "\n"}, "")
		}
		if err := scanner.Err(); err != nil {
			log.Panic("The input was invalid.", message)
		}
	}

	bot := Bot(token, timeout, debug)

	if len(message) != 0 {
		sendMessage(bot, message)
	}
	if location {
		if longitude < -180 || longitude > 180 || latitude < -90 || latitude > 90 {
			log.Panic("Longitude or latitude value invalid.")
		}
		sendLocation(bot, latitude, longitude)
	}
	if _, err := os.Stat(filename); err == nil {
		switch filetype {
		case "photo":
			sendPhoto(bot, filename, caption)
		case "video":
			sendVideo(bot, filename, caption)
		default:
			sendDocument(bot, filename, caption)
		}
	}

}
