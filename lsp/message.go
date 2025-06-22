package lsp

type Request struct {
    RPC string `json:"jsonrpc"`
    ID  int    `json:"id"`
    Method string `json:"method"`

    //we will specify param types for each request type
}

type Response struct {
    RPC string `json:"jsonrpc"`
    ID  *int    `json:"id,omitempty"`

    //we will specify result types for each response type 
}

type Notification struct {
    RPC string `json:"jsonrpc"`
    Method string `json:"method"`
}

// ResponseError represents an error object in the JSON-RPC 2.0 specification.
type ResponseError struct {
	Code    int    `json:"code"`    // Error code (e.g., -32602 for invalid parameters).
	Message string `json:"message"` // Human-readable error message.
}

// ErrorResponse represents an error response in the JSON-RPC 2.0 specification.
type ErrorResponse struct {
	JSONRPC string        `json:"jsonrpc"` // JSON-RPC version (always "2.0").
	ID      int           `json:"id"`      // The ID of the request that caused the error.
	Error   ResponseError `json:"error"`   // The error object containing details about the failure.
}
