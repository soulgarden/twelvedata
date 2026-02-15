package response

// Credits represents the interface for handling API credit information.
type Credits interface {
	GetCreditsLeft() int64
	GetCreditsUsed() int64
	SetCreditsLeft(val int64)
	SetCreditsUsed(val int64)
}

// CreditsImpl is the concrete implementation of the Credits interface.
type CreditsImpl struct {
	CreditsLeft int64
	CreditsUsed int64
}

// NewCreditsImpl creates a new CreditsImpl instance with the specified credits left and used values.
func NewCreditsImpl(creditsLeft int64, creditsUsed int64) Credits {
	return &CreditsImpl{CreditsLeft: creditsLeft, CreditsUsed: creditsUsed}
}

// SetCreditsLeft sets the number of API credits remaining.
func (c *CreditsImpl) SetCreditsLeft(val int64) {
	c.CreditsLeft = val
}

// SetCreditsUsed sets the number of API credits consumed by the current request.
func (c *CreditsImpl) SetCreditsUsed(val int64) {
	c.CreditsUsed = val
}

// GetCreditsLeft returns the number of API credits remaining.
func (c *CreditsImpl) GetCreditsLeft() int64 {
	return c.CreditsLeft
}

// GetCreditsUsed returns the number of API credits consumed by the current request.
func (c *CreditsImpl) GetCreditsUsed() int64 {
	return c.CreditsUsed
}
