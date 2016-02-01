package page

import (
    "fmt"
    "testing"
    
    "github.com/crispgm/go-spg/generator"
    "github.com/crispgm/go-spg/variables"
)

func TestNewWithMdx(t *testing.T){
    var jsonBlob = []byte(`{"Value": "Monotremata"}`)
    vrb := variables.New(jsonBlob)

    err, page := New("", "", vrb, generator.G_MDX)
    if err != nil && page.gtype != generator.G_MDX {
        t.Error("Test NewMdx failed")
    }
    page.LoadTemplate()
    err, cnt := page.Generate()
    fmt.Println(string(cnt))
}
