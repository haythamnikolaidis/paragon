package main

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"os"
	"paragon/analysis"
	"paragon/lsp"
	"paragon/rpc"
	"paragon/dictionary"
)

func main() {
    logger := getLogger("/Users/ivankruger/Documents/Development/open_source/paragon/log.txt")
    logger.Println("Starting Paragon...")

    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(rpc.Split)

    state := analysis.NewState()

    writer := os.Stdout

    for scanner.Scan() {
        msg := scanner.Bytes()
        method, content, err := rpc.DecodeMessage(msg)
        if err != nil {
            logger.Println("Error decoding message:", err)
            continue
        }
        handleMessage(logger, writer, state, method, content)
    }
}

func handleMessage(logger *log.Logger,
    writer io.Writer, state analysis.State,
    method string, content []byte) {

    logger.Printf("Received message: %s", method)

    switch method {
    case "initialize":
        var request lsp.InitializeRequest
        if err := json.Unmarshal(content, &request); err != nil {
            logger.Println("Error unmarshalling initialize request:", err)
            return
        }

        logger.Printf("Connected to client: %s, version: %s", 
            request.Params.ClientInfo.Name, 
            request.Params.ClientInfo.Version)

        msg := lsp.NewInitializeResponse(request.ID)
        response := rpc.EncodeMessage(msg)

        WriteResponse(writer, logger, response)

    case "textDocument/didOpen":
        var notification lsp.DidOpenTextDocumentNotification
        if err := json.Unmarshal(content, &notification); err != nil {
            logger.Println("Error unmarshalling did open request:", err)
            return
        }

        logger.Printf("Did Open: %s", notification.Params.TextDocument.URI)
        state.OpenDocument(notification.Params.TextDocument.URI, notification.Params.TextDocument.Text)
    case "textDocument/didChange":
        var notification lsp.DidChangeTextDocumentNotification
        if err := json.Unmarshal(content, &notification); err != nil {
            logger.Println("Error unmarshalling did change request:", err)
            return
        }

        for _, change := range notification.Params.ContentChanges {
            logger.Printf("Did Change: %s", notification.Params.TextDocument.URI)
            state.UpdateDocument(notification.Params.TextDocument.URI, change.Text)
        }
    case "textDocument/hover":
        var request lsp.TextDocumentHoverRequest
        if err := json.Unmarshal(content, &request); err != nil {
            logger.Println("Error unmarshalling hover request:", err)
            sendErrorResponse(writer, logger, request.ID, "Invalid hover request")
            return
        }

        uri := request.Params.TextDocument.URI
        position := request.Params.Position
        content := state.Documents[uri]

        word, err := analysis.FindWordAtPosition(content,
            position.Line,
            position.Character)

        if err != nil {
            logger.Println("Error finding word at position:", err)
            sendErrorResponse(writer, logger, request.ID, "Could not find word at position")
            return
        }

        definition, err := dictionary.GetDefinition(word)
        if err != nil {
            logger.Println("Error querying Ollama API:", err)
            sendErrorResponse(writer, logger, request.ID, "Failed to retrieve definition")
            return
        }

        msg := lsp.NewTextDocumentHoverResponse(request.ID, definition)
        response := rpc.EncodeMessage(msg)

        WriteResponse(writer, logger, response)
    }

}

func sendErrorResponse(writer io.Writer, logger *log.Logger, id int, message string) {
    errorResponse := rpc.EncodeMessage(lsp.ErrorResponse{
        ID:    id,
        Error: lsp.ResponseError{Code: -32602, Message: message},
    })

    if _, err := writer.Write([]byte(errorResponse)); err != nil {
        logger.Println("Error sending error response:", err)
    }
}

func WriteResponse(writer io.Writer, logger *log.Logger, response string) {
    _, err := writer.Write([]byte(response))
    if err != nil {
        logger.Println("Error writing response:", err)
    }
}

func getLogger(filename string) *log.Logger {
    file, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
    if err != nil {
        panic(err)
    }
    logger := log.New(file, "paragon ", log.Ldate|log.Ltime|log.Lshortfile)
    return logger
}
