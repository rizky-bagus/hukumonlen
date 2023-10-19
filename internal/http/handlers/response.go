package handlers

// Success represents success response.
type Success struct {
	Data interface{} `json:"data,omitempty"`
	Meta interface{} `json:"meta,omitempty"`
}

// SuccessResponse creates an instance of Success response.
func SuccessResponse(data, meta interface{}) *Success {
	return &Success{
		Data: data,
		Meta: meta,
	}
}
