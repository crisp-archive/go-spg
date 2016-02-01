package generator

import (
    "github.com/crispgm/go-spg/variables"
)

type StaticGenerator struct {
    template []byte
    content []byte
}

func NewStatic() Generator {
    return &StaticGenerator{[]byte(``), []byte(``)}
}

func (sg *StaticGenerator) SetTemplate(tpl []byte) bool {
    sg.template = tpl
    return true
}

func (sg *StaticGenerator) Render(v variables.Variables) bool {
    sg.content = sg.template
    return true
}

func (sg *StaticGenerator) GetContent() []byte {
    return sg.content
}
