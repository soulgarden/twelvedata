package twelvedata

import (
	"context"
	"sync"
)

// EventChannel provides a generic, thread-safe event channel with graceful closing.
// It follows Go best practices: the creator (owner) is responsible for closing the channel.
type EventChannel[T any] struct {
	ch     chan T
	closed chan struct{}
	once   sync.Once
}

// NewEventChannel creates a new generic event channel with the specified buffer size.
func NewEventChannel[T any](size int) *EventChannel[T] {
	return &EventChannel[T]{
		ch:     make(chan T, size),
		closed: make(chan struct{}),
	}
}

// Send attempts to send an event to the channel in a non-blocking way.
// Returns true if the event was sent successfully, false otherwise.
// This method respects context cancellation and channel closure.
func (ec *EventChannel[T]) Send(ctx context.Context, event T) bool {
	select {
	case <-ctx.Done():
		return false
	case <-ec.closed:
		return false
	case ec.ch <- event:
		return true
	default:
		// Channel is full, drop the event
		return false
	}
}

// Channel returns a read-only channel for consuming events.
// Consumers can range over this channel safely, as it will be closed
// when the EventChannel is closed.
func (ec *EventChannel[T]) Channel() <-chan T {
	return ec.ch
}

// Close closes the event channel safely. Can be called multiple times.
// This method should only be called by the channel owner.
func (ec *EventChannel[T]) Close() {
	ec.once.Do(func() {
		close(ec.closed) // Signal that we're closing
		close(ec.ch)     // Close the actual channel (releases range loops)
	})
}

// IsClosed returns true if the channel has been closed.
func (ec *EventChannel[T]) IsClosed() bool {
	select {
	case <-ec.closed:
		return true
	default:
		return false
	}
}
