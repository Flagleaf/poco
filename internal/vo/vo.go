package vo

import "time"

type LogVo struct {
	ID       int
	Content  string
	Pictures string
	Date     time.Time
	Location string
}
