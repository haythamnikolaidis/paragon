package analysis

type State struct {
    Documents map[string]string
}

func NewState() State {
    return State{
        Documents: make(map[string]string),
    }
}

func (s *State) OpenDocument(uri, content string) {
    s.Documents[uri] = content
}

func (s *State) UpdateDocument(uri, content string) {
    s.Documents[uri] = content
}
