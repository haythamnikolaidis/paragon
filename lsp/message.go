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
