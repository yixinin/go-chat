package models

import (
	"go-lib/db"

	log "github.com/sirupsen/logrus"
)

func SyncTables() {
	err := db.Mysql.Sync2(
		new(User),
		new(Contact),
		new(Group),
	)
	if err != nil {
		log.Errorf("sync tables error:%v", err)
	}
}
