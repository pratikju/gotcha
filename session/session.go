package session

import (
	"github.com/astaxie/beego/session"
)

var (
	Manager *session.Manager
)

func Init() {
	Manager, _ = session.NewManager("memory",
		&session.ManagerConfig{CookieName: "gosessionid", Gclifetime: 3600})

	go Manager.GC()
}
