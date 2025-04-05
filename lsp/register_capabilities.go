package lsp

type RegisterCapabilitiesRequest struct {
    Request
    Params RegisterCapabilitiesParams `json:"params"`
}

type RegisterCapabilitiesParams struct {
    ID int `json:"id"`
    Method string `json:"method"`
}

func NewRegisterCapabilitiesRequest(id int) RegisterCapabilitiesRequest {
    return RegisterCapabilitiesRequest{
        Request: Request{
            RPC: "2.0",
            ID:  id,
            Method: "client/registerCapability",
        },
        Params: RegisterCapabilitiesParams{
            ID: id,
            Method: "client/registerCapability",
        },
    }
}
