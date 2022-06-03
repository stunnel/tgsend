package main

import (
	"testing"
)

const (
	testToken   = "0:AAHderikGFEtiTE4f"
	testTimeout = 60
	testDebug   = true
)

func init() {
	ChatID = 0
	initBot(testToken, testTimeout, testDebug)
}

func TestSendMessage(t *testing.T) {
	message = "test message"
	sendMessage(message)
}

func TestSendPhoto(t *testing.T) {
	caption = "Fuji Mountain"
	filename = "test/fuji.jpg"
	filetype = "photo"
	sendFile(filename, filetype, caption)
}

func TestSendVideo(t *testing.T) {
	caption = "animation.mp4"
	filename = "test/sunrise.mp4"
	filetype = "video"
	sendFile(filename, filetype, caption)
}

func TestSendAnimation(t *testing.T) {
	caption = "animation"
	filename = "test/sunrise.mp4"
	filetype = "animation"
	sendFile(filename, filetype, caption)
}

func TestSendSticker(t *testing.T) {
	caption = "sticker.mp4"
	filename = "test/sunrise.mp4"
	filetype = "sticker"
	sendFile(filename, filetype, caption)
}

func TestSendDocument(t *testing.T) {
	caption = "King Lear.pdf"
	filename = "test/king-lear.pdf"
	filetype = "document"
	sendFile(filename, filetype, caption)
}

func TestSendLocation(t *testing.T) {
	location = true
	latitude = -33.8688
	longitude = 151.2093
	sendLocation(latitude, longitude)
}
