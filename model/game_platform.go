package model

import "time"

// LogLogin [...]
type LogLogin struct {
	ID     int       `gorm:"primaryKey;column:id;type:int(11);not null"`
	RoleID int       `gorm:"column:role_id;type:int(11);not null"`
	Time   time.Time `gorm:"column:time;type:datetime;not null"`
	IP     string    `gorm:"column:ip;type:varchar(255)"`
	Mac    string    `gorm:"column:mac;type:varchar(255)"`
	State  int8      `gorm:"column:state;type:tinyint(4);not null"`
	Region string    `gorm:"column:region;type:varchar(200)"`
}
