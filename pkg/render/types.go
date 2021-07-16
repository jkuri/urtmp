package render

// Empty represents an empty response.
// swagger:response emptyResponse
type Empty struct{}

// Error represents a JSON encoded API error.
// swagger:response errorResponse
type Error struct {
	Message string `json:"message"`
}

// BoolResponse represents a JSON encoded API for bool responses.
// swagger:response boolResponse
type BoolResponse struct {
	Status bool `json:"status"`
}
