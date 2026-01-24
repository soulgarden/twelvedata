package request

// GetSanctionedEntities represents request parameters for sanctioned entities data.
type GetSanctionedEntities struct {
	APIKey
	Source string `schema:"-"`
}

// PathParams returns URL path parameters for the sanctioned entities endpoint.
func (req GetSanctionedEntities) PathParams() map[string]string {
	return map[string]string{
		"source": req.Source,
	}
}
