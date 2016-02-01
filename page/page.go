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
    engine generator.Generator
}

func New(src string, dst string, vrb variables.Variables, gtype string) (bool, Page) {
    page := Page{"", "", []byte(``), vrb, generator.G_STATIC, nil}
    page.src = src
    page.dst = dst
    page.gtype = gtype
    switch page.gtype {
    case generator.G_STATIC:
        generator.Register("static", generator.NewStatic)
        var err error
        err, page.engine = generator.NewGenerator("static")
        if err != nil {
            return false, page
        }
        return true, page
    case generator.G_MDX:
        return true, page
    case generator.G_MARKDOWN:
        return true, page
    case generator.G_ERROR:
        return false, page
    default:
        return false, page
    }

    return true, page
}

func (page *Page) setEngine(g generator.Generator) {
    page.engine = g
}

func (page *Page) generate() (bool, []byte) {
    page.engine.SetTemplate(page.content)
    if !page.engine.Render(page.variables) {
        return false, []byte(``)
    }
    return true, page.engine.GetContent()
}
