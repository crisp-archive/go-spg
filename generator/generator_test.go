package generator

import (
    "testing"
    "fmt"

    "github.com/crispgm/go-spg/variables"
)

func TestStaticGenerator(t *testing.T){
    var jsonBlob = []byte(`{"Value": "Monotremata"}`)
    vrb := variables.New(jsonBlob)

    Register("static", NewStatic)
    err, g := NewGenerator("static")
    if err != nil {
        t.Error("Test New Static Generator Failed")
    }
    g.SetTemplate([]byte(`hello, world`))
    g.Render(vrb)
    cnt := g.GetContent()
    fmt.Println(string(cnt))
    if string(cnt) != "hello, world" {
        t.Error("Test Static Render failed")
    }
}

func TestMarkdownGenerator(t *testing.T){
    var jsonBlob = []byte(`{"Value": "Monotremata"}`)
    vrb := variables.New(jsonBlob)

    Register("markdown", NewMdx)
    err, g := NewGenerator("markdown")
    if err != nil {
        t.Error("Test New Markdown Generator Failed")
    }
    g.SetTemplate([]byte(`### hello, world`))
    g.Render(vrb)
    cnt := g.GetContent()
    fmt.Println(string(cnt))
    if string(cnt) != "<h3>hello, world</h3>\n" {
        t.Error("Test Markdown Render failed")
    }
}

func TestMdxGenerator(t *testing.T){
    var jsonBlob = []byte(`{"Value": "Monotremata"}`)
    vrb := variables.New(jsonBlob)

    Register("mdx", NewMdx)
    err, g := NewGenerator("mdx")
    if err != nil {
        t.Error("Test New Mdx Generator Failed")
    }
    g.SetTemplate([]byte(`{{include ../testing/head.html}}### hello, world{{include ../testing/foot.html}}`))
    g.Render(vrb)
    cnt := g.GetContent()
    fmt.Println(string(cnt))
    // if cnt != "hello, world" {
    //     t.Error("Test Mdx Render failed")
    // }
}
