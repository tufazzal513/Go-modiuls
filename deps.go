package main

import (
	_ "github.com/gotd/td/session"
	_ "github.com/gotd/td/telegram"
	_ "github.com/gotd/td/telegram/auth"
	_ "github.com/gotd/td/telegram/downloader"
	_ "github.com/gotd/td/telegram/message"
	_ "github.com/gotd/td/telegram/uploader"
	_ "github.com/gotd/td/tg"
	_ "modernc.org/sqlite"
)

func main() {}
