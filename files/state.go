package files

import (
	"bytes"
	"errors"
	"log"
	"strings"
	"sync"
	"unicode/utf8"

	"github.com/sourcegraph/go-lsp"
)

type State struct {
	mu sync.Mutex

	data map[lsp.DocumentURI]*file
}

func NewState() *State {
	return &State{
		data: map[lsp.DocumentURI]*file{},
	}
}

func (s *State) GetText(uri lsp.DocumentURI) (string, error) {

	if f, ok := s.data[uri]; ok {
		return f.Text, nil
	}

	return "", errors.New("Unknown file")
}

func (s *State) GetLine(uri lsp.DocumentURI, line int) (string, error) {

	txt, err := s.GetText(uri)
	if err != nil {
		return txt, err
	}

	// TODO handle windows \n\r
	lines := strings.Split(txt, "\n")

	if len(lines) > line {
		return lines[line], nil
	}

	return "", errors.New("No such line")
}

func (s *State) SetText(uri lsp.DocumentURI, content string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[uri] = &file{
		Text: content,
	}
}

func (s *State) SetDocument(doc lsp.TextDocumentItem) {
	s.SetText(doc.URI, doc.Text)
}

func (s *State) Remove(uri lsp.DocumentURI) {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.data, uri)
}

func (s *State) ApplyCanges(uri lsp.DocumentURI, changes []lsp.TextDocumentContentChangeEvent) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if f, ok := s.data[uri]; ok {
		return f.ApplyCanges(changes)
	}

	return errors.New("Unknown file")
}

type file struct {
	Text string
}

func (s *file) ApplyCanges(changes []lsp.TextDocumentContentChangeEvent) error {
	log.Printf("applyChanges: %d", len(changes))

	if len(changes) == 1 && changes[0].Range == nil {
		// If range is empty, we expect the full content of file, i.e. a single change with no range.
		change := changes[0]
		if change.RangeLength != 0 {
			return errors.New("unexpected change range provided")
		}

		s.Text = change.Text

		return nil
	}

	content := []byte(s.Text)

	for _, change := range changes {
		start := bytesOffset(content, change.Range.Start)
		if start == -1 {
			return errors.New("invalid range for content change")
		}
		end := bytesOffset(content, change.Range.End)
		if end == -1 {
			return errors.New("invalid range for content change")
		}
		var buf bytes.Buffer
		buf.Write(content[:start])
		buf.WriteString(change.Text)
		buf.Write(content[end:])
		content = buf.Bytes()
	}

	s.Text = string(content)

	return nil
}

func bytesOffset(content []byte, pos lsp.Position) int {
	var line, char, offset int

	for len(content) > 0 {
		if line == int(pos.Line) && char == int(pos.Character) {
			return offset
		}
		r, size := utf8.DecodeRune(content)
		char++
		// The offsets are based on a UTF-16 string representation.
		// So the rune should be checked twice for two code units in UTF-16.
		if r >= 0x10000 {
			if line == int(pos.Line) && char == int(pos.Character) {
				return offset
			}
			char++
		}
		offset += size
		content = content[size:]
		if r == '\n' {
			line++
			char = 0
		}
	}

	if line == int(pos.Line) && char == int(pos.Character) {
		return offset
	}

	return -1
}
