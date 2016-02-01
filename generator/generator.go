package generator

import (
    "fmt"

    "github.com/crispgm/go-spg/variables"
)

const (
    G_ERROR    = "error"
    G_STATIC   = "static"
    G_MARKDOWN = "markdown"
    G_MDT      = "mdx" // markdown with simple template
)

type Generator interface {
    SetTemplate(content []byte) bool
    Render(v variables.Variables) bool
    GetContent() []byte
}

type Instance func() Generator

var adapters = make(map[string]Instance)

func Register(name string, adapter Instance) {
    if (adapter == nil) {
        panic("Generator: Register adapter is nil")
    }
    if _, ok := adapters[name]; ok {
        panic("Generator: Register called twice for adapter " + name)
    }
    adapters[name] = adapter
}

func NewGenerator(gName string) (err error, g Generator) {
    instanceFunc, ok := adapters[gName]
    if !ok {
        err = fmt.Errorf("Generator: unknown generator name %q", gName)
        return
    }
    g = instanceFunc()
    return
}
