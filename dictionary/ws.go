package dictionary

import "time"

const (
	PongWait     = 60 * time.Second
	PingPeriod   = (PongWait * 9) / 10 //nolint:gomnd
	WriteWait    = 10 * time.Second
	EventsChSize = 1024
)

const PriceEventType = "price"
