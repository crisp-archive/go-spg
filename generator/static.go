package generator

import (
    "github.com/crispgm/go-spg/variables"
)

type StaticGenerator struct {
    template string
    content string
}

func NewStatic() Generator {
    return &StaticGenerator{"", ""}
}

func (sg *StaticGenerator) SetTemplate(tpl string) bool {
    sg.template = tpl
    return true
}

func (sg *StaticGenerator) Render(v variables.Variables) bool {
    sg.content = sg.template
    return true
}

func (sg *StaticGenerator) GetContent() string {
    return sg.content
}
