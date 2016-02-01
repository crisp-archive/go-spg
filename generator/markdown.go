package generator

import (
    "github.com/russross/blackfriday"
    "github.com/crispgm/go-spg/variables"
)

type MarkdownGenerator struct {
    template []byte
    content []byte
}

func NewMarkdown() Generator {
    return &MarkdownGenerator{[]byte(``), []byte(``)}
}

func (sg *MarkdownGenerator) SetTemplate(tpl []byte) bool {
    sg.template = tpl
    return true
}

func (sg *MarkdownGenerator) Render(v variables.Variables) bool {
    sg.content = blackfriday.MarkdownBasic(sg.template)
    return true
}

func (sg *MarkdownGenerator) GetContent() []byte {
    return sg.content
}
