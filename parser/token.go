package parser

import (
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Token is an abstraction from each lexical element, literal, etc.
type Token interface {
	GetLine() int
	GetColumn() int
	GetText() string
	SetText(s string)
}

type parseOption func(text string) string

func parseToken(t antlr.Token, option ...parseOption) string {
	text := t.GetText()
	for _, o := range option {
		text = o(text)
	}
	return text
}

func parseTerminalNode(t antlr.TerminalNode, option ...parseOption) string {
	text := t.GetText()
	for _, o := range option {
		text = o(text)
	}
	return text
}

func withTrim(c string) parseOption {
	return func(text string) string {
		return strings.Trim(text, c)
	}
}

func withUpperCase() parseOption {
	return func(text string) string {
		return strings.ToUpper(text)
	}
}

func withReplacer(oldnew ...string) parseOption {
	return func(text string) string {
		return strings.NewReplacer(oldnew...).Replace(text)
	}
}
