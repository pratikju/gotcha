package session

import (
	"github.com/astaxie/beego/session"
)

var (
	Manager *session.Manager
)

func Init() {
	Manager, _ = session.NewManager("memory", `{"cookieName":"gosessionid","gclifetime":3600}`)
	go Manager.GC()
}
