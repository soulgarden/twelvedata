package dictionary

import "time"

const (
	// PongWait is the maximum time to wait for a pong message from the server.
	PongWait = 60 * time.Second
	// PingPeriod is the interval for sending ping messages to the server.
	// Updated to 10 seconds as recommended by the API documentation for heartbeat events.
	PingPeriod = 10 * time.Second
	// WriteWait is the maximum time allowed for writing a message to the connection.
	WriteWait = 10 * time.Second
	// EventsChSize is the buffer size for the events channel.
	EventsChSize = 1024
	// HeartbeatPeriod is the interval for sending heartbeat messages to maintain connection stability.
	// As per API documentation, heartbeat should be sent every 10 seconds.
	HeartbeatPeriod = 10 * time.Second
)
