package page

import (
    "github.com/crispgm/go-spg/generator"
    "github.com/crispgm/go-spg/variables"
)

type Page struct {
    src string
    dst string

    template []byte
    content []byte
    variables variables.Variables

    gtype string
    engine generator.Generator
}

type PageError struct {
    s string
}

func (e *PageError) Error() string {
    return e.s
}

func NewError(text string) error {
    return &PageError{text}
}

func New(src string, dst string, vrb variables.Variables, gtype string) (error, Page) {
    p := Page{"", "", []byte(``), []byte(``), vrb, gtype, nil}
    p.src = src
    p.dst = dst

    switch p.gtype {
    // case generator.G_STATIC:
    // case generator.G_MARKDOWN:
    case generator.G_MDX:
        generator.Register(p.gtype, generator.NewMdx)
        err, engine := generator.NewGenerator(p.gtype)
        if err != nil {
            return err, p
        }
        p.engine = engine
        return nil, p
    case generator.G_ERROR:
        return NewError("Invalid Generator Type"), p
    default:
        return NewError("Invalid Generator Type"), p
    }

    return nil, p
}

func (page *Page) SetEngine(g generator.Generator) {
    page.engine = g
}

func (page *Page) LoadTemplate() {
    page.template = []byte(`# hello`)
    page.engine.SetTemplate(page.template)
}

func (page *Page) Generate() (error, []byte) {
    err := page.engine.Render(page.variables)
    if err != nil {
        return NewError("Render Failed"), []byte(``)
    }
    return nil, page.engine.GetContent()
}
