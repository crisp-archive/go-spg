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

func (sg *MarkdownGenerator) SetTemplate(tpl []byte) error {
    sg.template = tpl
    return nil
}

func (sg *MarkdownGenerator) Render(v variables.Variables) error {
    sg.content = blackfriday.MarkdownBasic(sg.template)
    return nil
}

func (sg *MarkdownGenerator) GetContent() []byte {
    return sg.content
}
