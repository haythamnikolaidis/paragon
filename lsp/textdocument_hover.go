package lsp

type TextDocumentHoverRequest struct {
    Request
    Params TextDocumentHoverParams `json:"params"`
}

type TextDocumentHoverParams struct {
    TextDocumentPositionParams
}

type TextDocumentHoverResponse struct {
    Response
    Result TextDocumentHoverResult `json:"result"`
}

type TextDocumentHoverResult struct {
    Contents string `json:"contents"`
}

func NewTextDocumentHoverResponse(id int, contents string) TextDocumentHoverResponse {
    return TextDocumentHoverResponse{
        Response: Response{
            RPC: "2.0",
            ID:  &id,
        },
        Result: TextDocumentHoverResult{
            Contents: contents,
        },
    }
}

