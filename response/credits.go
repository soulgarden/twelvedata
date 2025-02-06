package response

type Credits interface {
	GetCreditsLeft() int64
	GetCreditsUsed() int64
	SetCreditsLeft(val int64)
	SetCreditsUsed(val int64)
}

type CreditsImpl struct {
	CreditsLeft int64
	CreditsUsed int64
}

func NewCreditsImpl(creditsLeft int64, creditsUsed int64) Credits {
	return &CreditsImpl{CreditsLeft: creditsLeft, CreditsUsed: creditsUsed}
}

func (c *CreditsImpl) SetCreditsLeft(val int64) {
	c.CreditsLeft = val
}

func (c *CreditsImpl) SetCreditsUsed(val int64) {
	c.CreditsUsed = val
}

func (c *CreditsImpl) GetCreditsLeft() int64 {
	return c.CreditsLeft
}

func (c *CreditsImpl) GetCreditsUsed() int64 {
	return c.CreditsUsed
}
