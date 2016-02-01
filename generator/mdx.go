package generator

import (
    "fmt"

    "github.com/russross/blackfriday"
    "github.com/crispgm/go-spg/variables"
)

type MdxGenerator struct {
    template []byte
    content []byte
}

func NewMdx() Generator {
    return &MdxGenerator{[]byte(``), []byte(``)}
}

func (sg *MdxGenerator) SetTemplate(tpl []byte) error {
    sg.template = tpl
    return nil
}

func (sg *MdxGenerator) Render(v variables.Variables) error {
    // parser include
    // do include
    sg.content = blackfriday.MarkdownBasic(sg.template)
    return nil
}

func (sg *MdxGenerator) GetContent() []byte {
    return sg.content
}
