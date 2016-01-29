package page

import (
    "fmt"
    "testing"
    
    "github.com/crispgm/go-spg/variables"
)

func TestNewWithStatic(t *testing.T){
    var jsonBlob = []byte(`{"Value": "Monotremata"}`)
    vrb := variables.New(jsonBlob)

    page := Page{"", "", "", vrb, G_STATIC, nil}
    if page.gtype != G_STATIC {
        t.Error("Test GetIntVar failed")
    }
    fmt.Println(page.generate())
}
