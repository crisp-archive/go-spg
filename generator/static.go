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

func (sg *StaticGenerator) SetTemplate(tpl []byte) error {
    sg.template = tpl
    return nil
}

func (sg *StaticGenerator) Render(v variables.Variables) error {
    sg.content = sg.template
    return nil
}

func (sg *StaticGenerator) GetContent() []byte {
    return sg.content
}
