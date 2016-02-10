package main

import (
	"github.com/astaxie/beego/session"
)

var (
	GoChatManager *session.Manager
)

func InitSession() {
	GoChatManager, _ = session.NewManager("memory", `{"cookieName":"gosessionid","gclifetime":3600}`)
	go GoChatManager.GC()
}
