# tgsend in golang

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)


tgsend is a command-line tool to send messages and files over Telegram to your account, to a group or to a channel.

<!-- markdown-toc start - Don't edit this section. Run M-x markdown-toc-generate-toc again -->
**Table of Contents**

- [tgsend in golang](#tgsend-in-golang)
- [Usage](#usage)
  - [Command parameter](#command-parameter)
  - [Start](#start)
  - [Example](#example)

<!-- markdown-toc end -->

# Usage

## Command parameter
```shell
tgsend -h

  -token string
    	Set the bot token.
  -chatid int
    	Send message to this User ID or Channel ID or Group ID.

  -message string
    	The message to sent.
  -pre
    	Send preformatted fixed-width (monospace) text.
  -format string
    	How to format the message(s). Choose from ['text', 'markdown', 'html'] (default 'text')
  -preview
    	disable link previews in the message(s)

  -filename path
    	The file to sent.
  -filetype string
    	Set the file type, Choose from ['photo', 'video', 'document'] (default 'document')
  -caption string
    	Set the photo/video/document caption

  -location
    	Send location
  -latitude float
    	Set latitude, value valid [-90, 90]
  -longitude float
    	Set longitude, value valid [-180, 180]

  -timeout int
    	Set the read timeout for network operations(in seconds). (default 30)
  -debug
    	Show debug message.
  -version
    	Print version information.
```

### Message type priority

**message > location > file**  
It means that if you provide a `message`, parameters such as `location`, `filetype` and `caption` will not be used.

## Start

1. Create a bot.  
   Chat with BotFather https://telegram.me/botfather, send /newbot command to create a new bot, and get the token, for example `110201543:AAHdqTcvCH1vGWJxfSeofSAs0K5PALDsaw`
2. Chat with your bot, use your Telegram account.
3. Get your telegram account's ID. Send any message to https://t.me/RawDataBot, This bot will return the raw message from Telegram(json format). `["message"]["from"]["id"]` is your account's ID. etc 12345678.
4. Afterward, now you can send a message via `tgsend`.  
   `tgsend -token 110201543:AAHdqTcvCH1vGWJxfSeofSAs0K5PALDsaw -chatid 12345678 -message "Hello World."`

There is a maximum message length of 4096 characters, larger messages will be automatically split up into smaller ones and sent separately.

## Example

Send a message:

```shell
tgsend -token 110201543:AAHdqTcvCH1vGWJxfSeofSAs0K5PALDsaw -chatid 12345678 \
       -message "Hello World."
```

Send stdin:
```shell
echo "now: $(date)" | tgsend -chatid 12345678 \
    -token 110201543:AAHdqTcvCH1vGWJxfSeofSAs0K5PALDsaw \
    -message -
```

To send a message using Markdown or HTML formatting:
```shell
tgsend -token 110201543:AAHdqTcvCH1vGWJxfSeofSAs0K5PALDsaw -chatid 12345678 \
       -format markdown -message "*bold text* _italic text_"

tgsend -token 110201543:AAHdqTcvCH1vGWJxfSeofSAs0K5PALDsaw -chatid 12345678 \
       -format html -message '<b>bold</b>, <strong>bold</strong>
<i>italic</i>, <em>italic</em>
<a href="https://www.example.com/">inline URL</a>
<a href="tg://user?id=123456789">inline mention of a user</a>
<code>inline fixed-width code</code>
<pre>pre-formatted fixed-width code block</pre>'
```

send a message using Markdown with multi lines:
```shell
tgsend -token 110201543:AAHdqTcvCH1vGWJxfSeofSAs0K5PALDsaw -chatid 12345678 \
       -format markdown -message '*bold text*
_italic text_
[inline URL](https://www.example.com/)
[inline mention of a user](tg://user?id=123456789)
`inline fixed-width code`
 ```block_language
pre-formatted fixed-width code block
 ```'
```

For more information on supported formatting, see the [Formatting options](https://core.telegram.org/bots/api#formatting-options).

The `--pre` flag formats messages as fixed-width text:
```shell
tgsend -token 110201543:AAHdqTcvCH1vGWJxfSeofSAs0K5PALDsaw -chatid 12345678 \
       -pre -message "monospace"
```

To send a message without link previews:
```shell
tgsend -token 110201543:AAHdqTcvCH1vGWJxfSeofSAs0K5PALDsaw -chatid 12345678 \
       -preview -message "https://github.com/"
```

Send a file (maximum file size of 50 MB):
```shell
tgsend -token 110201543:AAHdqTcvCH1vGWJxfSeofSAs0K5PALDsaw -chatid 12345678 \
       -filetype document -filename document.pdf
```

To send an image with an optional caption (maximum file size of 10 MB):
```shell
tgsend -token 110201543:AAHdqTcvCH1vGWJxfSeofSAs0K5PALDsaw -chatid 12345678 \
       -filetype photo -filename photo.jpg -caption "The Moon at night"
```

To send a location via latitude and longitude:
```shell
tgsend -token 110201543:AAHdqTcvCH1vGWJxfSeofSAs0K5PALDsaw -chatid 12345678 \
       -location -latitude 35.5398033 -longitude -79.7488965
```
