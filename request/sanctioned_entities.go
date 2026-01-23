package request

// GetSanctionedEntities represents request parameters for sanctioned entities data.
type GetSanctionedEntities struct {
	APIKey
	Source string `schema:"-"`
}
