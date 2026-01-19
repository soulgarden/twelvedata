package response

// TechnicalIndicators represents the response structure for technical indicators list data.
type TechnicalIndicators struct {
	Data   map[string]*TechnicalIndicator `json:"data"`
	Status string                         `json:"status"`
}

// TechnicalIndicator represents a single technical indicator with its configuration and details.
type TechnicalIndicator struct {
	Enable         bool                   `json:"enable"`
	FullName       string                 `json:"full_name"`
	Description    string                 `json:"description"`
	Type           string                 `json:"type"`
	Overlay        bool                   `json:"overlay"`
	Parameters     map[string]interface{} `json:"parameters"`
	OutputValues   map[string]interface{} `json:"output_values"`
	TintingDetails map[string]interface{} `json:"tinting_details"`
}
