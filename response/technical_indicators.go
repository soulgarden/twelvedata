package response

// TechnicalIndicators represents a map of technical indicators by name.
type TechnicalIndicators map[string]*TechnicalIndicator

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
