package page

import (
    "github.com/crispgm/go-spg/generator"
    "github.com/crispgm/go-spg/variables"
)

type Page struct {
    src string
    dst string

    content []byte
    variables variables.Variables

    gtype string
    generator Generator
}

func New(src string, dst string, vrb variables.Variables, gtype string) (bool, Page) {
    page := Page{"", "", []byte(``), vrb, G_STATIC, nil}
    page.src = src
    page.dst = dst
    page.gtype = gtype
    switch page.gtype {
    case G_STATIC:
        page.generator = static.NewStatic("", "")
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
    page.generator = g
}

func (page *Page) generate() (bool, []byte) {
    page.generator.SetTemplate(page.content)
    if !page.generator.Render(page.variables) {
        return false, ""
    }
    return true, page.generator.GetContent()
}
