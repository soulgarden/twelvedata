package request

// GetSanctionedEntities represents request parameters for sanctioned entities data.
type GetSanctionedEntities struct {
	APIKey
	Source     string `schema:"source,omitempty"`
	EntityType string `schema:"entity_type,omitempty"`
	Country    string `schema:"country,omitempty"`
	Format     string `schema:"format,omitempty"`
	Page       string `schema:"page,omitempty"`
	PageSize   string `schema:"page_size,omitempty"`
}
