package model

import "log"

func ListNullRegion() []*LogLogin {
	logins := make([]*LogLogin, 0)
	if tx := db.Debug().Table("log_login").Where("region IS NULL").Limit(100).Find(&logins); tx.Error != nil {
		log.Print(tx.Error)
		return nil
	}
	return logins
}

func (l *LogLogin) UpdateRegion() {
	db.Debug().Table("log_login").Where("id = ?", l.ID).UpdateColumn("region", l.Region)
}
