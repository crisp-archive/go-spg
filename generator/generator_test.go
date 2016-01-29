package generator

import (
    "testing"

    "github.com/crispgm/go-spg/variables"
)

func TestGenerator(t *testing.T){
    var jsonBlob = []byte(`{"Value": "Monotremata"}`)
    vrb := variables.New(jsonBlob)

    Register("static", NewStatic)
    err, g := NewGenerator("static")
    if err != nil {
        t.Error("Test New Static Generator Failed")
    }
    g.SetTemplate("hello, world")
    g.Render(vrb)
    cnt := g.GetContent()
    if cnt != "hello, world" {
        t.Error("Test Static Render failed")
    }
}
