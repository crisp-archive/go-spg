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

    page := Page{"", "", []byte(``), vrb, generator.G_MDX, nil}
    if page.gtype != generator.G_MDX {
        t.Error("Test NewMdx failed")
    }
    fmt.Println(page.generate())
}
