package entity

import "time"

type LogEntity struct {
	ID       int
	Content  string
	Pictures string
	Date     time.Time
	Location string
}

func (LogEntity) TableName() string {
	return "log"
}
