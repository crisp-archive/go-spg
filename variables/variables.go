package variables

import (
    "encoding/json"
    "fmt"
)

type (
    Variables struct {
        rawJSONVars []byte
        parsedVars interface{}
    }
)

func New(rawJSONVars []byte) Variables {
    vrb := Variables{rawJSONVars, nil}
    return vrb
}

func (vrb *Variables) loadVar() bool {
    err := json.Unmarshal(vrb.rawJSONVars, &vrb.parsedVars)
    if err != nil {
        fmt.Println("error:", err)
        return false
    }
    return true
}

func (vrb *Variables) getVar(varname string) (bool, interface{}) {
    if (vrb.parsedVars == nil) {
        return false, ""
    }
    
    for key, val := range vrb.parsedVars.(map[string]interface{}) {
        if key == varname {
            return true, val
        }
    }
    
    return false, ""
}
