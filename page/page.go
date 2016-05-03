package page

import (
    "io/ioutil"
    "github.com/russross/blackfriday"
)

type Page struct {
    src string
    dst string
    input_file string
    head_file string
    foot_file string
    template []byte
    content []byte
    html []byte
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

func New(src string, dst string, input_file string, tpl []byte, cnt []byte, h []byte) (error, Page) {
    p := Page{"", "", "", []byte(``), []byte(``), []byte(``)}

    p.src = src
    p.dst = dst
    p.input_file = input_file
    p.template = tpl
    p.content = cnt
    p.html = h

    return nil, p
}

func (page *Page) LoadTemplate() error {
    dat, err := ioutil.ReadFile(page.src + page.input_file)
    if err != nil {
        panic(err)
    }

    page.template = dat

    return nil
}

func (page *Page) Compile() error {
    page.content = blackfriday.MarkdownBasic(page.template)
    return nil
}

func (page *Page) Build() error {
    //page.html = page.head + page.content + page.foot
    page.html = page.content
    return nil
}

func (page *Page) Output() error {
    err := ioutil.WriteFile(page.dst, page.html, 0644)
    if err != nil {
        panic(err)
    }
    return nil
}
