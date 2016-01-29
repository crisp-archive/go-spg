package page

import (
    //"static"
    //"mdx"
    "../variables"
    "./static"
)

const (
    G_ERROR  = iota
    G_STATIC = iota
    G_MDT    = iota // markdown with simple template
)

type Generator interface {
    SetTemplate(content string) bool
    Render(v variables.Variables) bool
    GetContent() string
}

type Page struct {
    src string
    dst string

    content string
    variables variables.Variables

    gtype int
    generator Generator
}

func New(src string, dst string, vrb variables.Variables, gtype int) (bool, Page) {
    page := Page{"", "", "", vrb, G_STATIC, nil}
    page.src = src
    page.dst = dst
    page.gtype = gtype
    switch page.gtype {
    case G_STATIC:
        page.generator = static.NewStatic{"", ""}
        return true, page
    case G_MDT:
        return true, page
    case G_ERROR:
        return false, page
    default:
        return false, page
    }

    return true, page
}

func (page *Page) setGenerator(g Generator) {
}

func (page *Page) generate() (bool, string) {
    page.generator.SetTemplate(page.content)
    if !page.generator.Render(page.variables) {
        return false, ""
    }
    return true, page.generator.GetContent()
}
