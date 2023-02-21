package utils

import (
	"strings"

	"github.com/josa42/go-ls/lsp"
)

func LineRange(no int, line string) lsp.Range {
	return lsp.Range{
		Start: lsp.Position{Line: no, Character: 0},
		End:   lsp.Position{Line: no, Character: lastChar([]string{line})},
	}
}

func Range(str string) lsp.Range {

	lines := strings.Split(str, "\n")

	return lsp.Range{
		Start: lsp.Position{Line: 0, Character: 0},
		End:   lsp.Position{Line: lastLine(lines), Character: lastChar(lines)},
	}
}

func lastLine(lines []string) int {
	idx := len(lines) - 1
	if idx >= 0 {
		return idx
	}
	return 0
}

func lastChar(lines []string) int {
	if len(lines) == 0 {
		return 0
	}

	line := []rune(lines[len(lines)-1])
	idx := len([]rune(line))
	if idx >= 0 {
		return idx
	}
	return 0
}
