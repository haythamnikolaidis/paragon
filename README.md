# Overview
The Paragon LSP is a Language Server Protocol (LSP) implementation in Go, designed to act as a writing assistant by integrating with the Ollama API. Its primary features include:
1. **Completions API**: Suggests words to complete a sentence.
2. **Hover API**: Displays the definition of a word when hovering over it.

This document provides a detailed and complete roadmap for implementing the Paragon LSP, including edge cases, concurrency handling, and compliance with the LSP specification.

---

## Architecture

1. **Input Handling**:
   - The server reads JSON-RPC messages from `os.Stdin` using a `bufio.Scanner`.
   - Messages are parsed and dispatched to appropriate handlers based on their method (e.g., `initialize`, `textDocument/hover`, `textDocument/completion`).

2. **State Management**:
   - The `analysis.State` object tracks the state of open documents and their content.
   - This is used to provide context for hover and completion requests.
   - Concurrency is managed using a `sync.RWMutex` to ensure thread-safe access to the state.

3. **Message Handling**:
   - Each LSP method (e.g., `initialize`, `textDocument/hover`) has a dedicated case in the `handleMessage` function.
   - Requests are unmarshalled into Go structs, processed, and responses are sent back to the client.

4. **Logging**:
   - A logger writes diagnostic information to a file for debugging purposes.

5. **Protocol Compliance**:
   - The server adheres to the LSP specification, ensuring proper handling of required and optional fields in requests and responses.
   - Unsupported methods or invalid requests are handled gracefully with appropriate error responses.

---

## Key Features

### 1. **Initialize Request**
   - **Purpose**: Establishes a connection with the client and exchanges basic information.
   - **Implementation**:
     - Parse the `initialize` request using `lsp.InitializeRequest`.
     - Log client details (e.g., name, version).
     - Respond with an `lsp.InitializeResponse`.
   - **Edge Cases**:
     - If required fields (e.g., `clientInfo`) are missing, respond with an error.

### 2. **Hover API**
   - **Purpose**: Displays the definition of a word at a given position in the document.
   - **Implementation**:
     - Parse the `textDocument/hover` request using `lsp.TextDocumentHoverRequest`.
     - Extract the word at the specified position using `analysis.FindWordAtPosition`.
     - Query the Ollama API for the word's definition.
     - Respond with an `lsp.TextDocumentHoverResponse` containing the definition.
   - **Edge Cases**:
     - If the position is out of bounds, respond with an error.
     - If the word is not found in the Ollama API, return a fallback message (e.g., "No definition available").
     - If the document is not open, respond with an error.

### 3. **Completions API**
   - **Purpose**: Suggests words to complete a sentence.
   - **Implementation**:
     - Add a new case for `textDocument/completion` in `handleMessage`.
     - Parse the `textDocument/completion` request using `lsp.TextDocumentCompletionRequest`.
     - Extract the current sentence context from the document state.
     - Query the Ollama API for word suggestions based on the context.
     - Respond with an `lsp.TextDocumentCompletionResponse` containing the suggestions.
   - **Edge Cases**:
     - If the context is ambiguous or incomplete, return a fallback response (e.g., "No suggestions available").
     - If the Ollama API returns an error, log the error and respond with a fallback message.
     - If the document is not open, respond with an error.

---

## Implementation Details

### 1. **Ollama API Integration**
   - **Purpose**: Use the Ollama API to fetch word definitions and sentence completions.
   - **Steps**:
     - Create a new package `ollama` to handle API requests.
     - Define functions:
       - `GetDefinition(word string) (string, error)`: Fetches the definition of a word.
       - `GetCompletions(context string) ([]string, error)`: Fetches word suggestions for a given sentence context.
     - Use Go's `net/http` package to make HTTP requests to the Ollama API.
   - **Error Handling**:
     - If the API request fails, log the error and return a fallback response.

### 2. **State Management**
   - **Purpose**: Track open documents and their content.
   - **Steps**:
     - Extend the `analysis.State` struct to include:
       - `Documents map[string]string`: Maps document URIs to their content.
       - `sync.RWMutex`: Ensures thread-safe access to the state.
     - Implement methods:
       - `OpenDocument(uri string, content string)`: Adds a new document to the state.
       - `UpdateDocument(uri string, content string)`: Updates the content of an existing document.
   - **Concurrency**:
     - Use `state.Lock()` and `state.Unlock()` for write operations.
     - Use `state.RLock()` and `state.RUnlock()` for read operations.

### 3. **Error Handling**
   - **Purpose**: Ensure the server remains robust and logs all errors.
   - **Steps**:
     - Log all errors using the `logger`.
     - Respond with appropriate error messages to the client using LSP's `ErrorResponse` format.

### 4. **Testing**
   - **Purpose**: Validate the functionality of the LSP.
   - **Steps**:
     - Write unit tests for:
       - `handleMessage` function.
       - `ollama` package functions.
       - `analysis.State` methods.
     - Use table-driven tests in Go to cover:
       - Malformed requests (e.g., invalid JSON, missing fields).
       - Edge cases (e.g., out-of-bounds positions, empty documents).
       - API failures (e.g., timeouts, invalid responses).

### 5. **Performance**
   - **Caching**:
     - Cache Ollama API responses to reduce latency for repeated requests.
   - **Optimization**:
     - Minimize redundant operations when updating document state.

---

## Example Workflows

### Hover Request
**Request**:
```json
{
  "jsonrpc": "2.0",
  "method": "textDocument/hover",
  "params": {
    "textDocument": { "uri": "file://example.txt" },
    "position": { "line": 10, "character": 5 }
  },
  "id": 1
}
```

**Response**:
```json
{
  "jsonrpc": "2.0",
  "result": {
    "contents": "Definition of the word at the given position."
  },
  "id": 1
}
```

**Edge Case**:
- If the position is out of bounds:
```json
{
  "jsonrpc": "2.0",
  "error": {
    "code": -32602,
    "message": "Position out of bounds."
  },
  "id": 1
}
```

### Completion Request
**Request**:
```json
{
  "jsonrpc": "2.0",
  "method": "textDocument/completion",
  "params": {
    "textDocument": { "uri": "file://example.txt" },
    "position": { "line": 10, "character": 5 }
  },
  "id": 2
}
```

**Response**:
```json
{
  "jsonrpc": "2.0",
  "result": {
    "items": [
      { "label": "suggestion1" },
      { "label": "suggestion2" }
    ]
  },
  "id": 2
}
```

---

## Next Steps

1. Implement the `ollama` package for API integration.
2. Add the `textDocument/completion` handler in `handleMessage`.
3. Extend the `analysis.State` struct to support document updates with concurrency.
4. Write unit tests for all new functionality.
5. Optimize performance with caching and efficient state management.
 
