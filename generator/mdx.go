package generator

import (
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

func (sg *MdxGenerator) SetTemplate(tpl []byte) bool {
    sg.template = tpl
    return true
}

func (sg *MdxGenerator) Render(v variables.Variables) bool {
    // parser include
    // do include
    sg.content = blackfriday.MarkdownBasic(sg.template)
    return true
}

func (sg *MdxGenerator) GetContent() []byte {
    return sg.content
}
