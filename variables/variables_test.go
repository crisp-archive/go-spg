package variables

import (
    "testing"
)

func TestGetStringVar(t *testing.T) {
    var jsonBlob = []byte(`{"Value": "Monotremata"}`)
    vrb := Variables{jsonBlob, nil}
    vrb.loadVar()
    b, v := vrb.getVar("Value")
    if b != true || v != "Monotremata" {
        t.Error("Test GetStringVar failed")
    }
}

func TestGetIntVar(t *testing.T) {
    var jsonBlob = []byte(`{"Value": 123}`)
    var expectV float64
    expectV = 123

    vrb := Variables{jsonBlob, nil}
    vrb.loadVar()
    b, v := vrb.getVar("Value")
    v = v.(float64)
    if b != true || v != expectV {
        t.Error("Test GetIntVar failed")
    }
}

// func TestGetArrayVar(t *testing.T) {
//     var jsonBlob = []byte(`{"Value": []}`)
//     vrb := Variables{jsonBlob, nil}
//     vrb.loadVar()
//     b, v := vrb.getVar("Value")
//     fmt.Println(v)
//     if b != true || v != 123 {
//         // t.Error("Test GetIntVar failed")
//     }
// }