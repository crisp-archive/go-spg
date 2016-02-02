package generator

import (
    // "fmt"
    "io/ioutil"
    "strings"

    "github.com/golang-commonmark/markdown"
    "github.com/crispgm/go-spg/variables"
)

const (
    LEFT_DELIMITER = "{{"
    RIGHT_DELIMITER = "}}"
    SPACE_DELIMITER = " "

    TOKEN_INCLUDE = "include"
    TOKEN_INCLUDE_RECURSIVE = "include_recursive"
    TOKEN_VARIABLE = "var"

    MAX_PARAM_NUM = 5
)

type MdxGenerator struct {
    template []byte
    parsedTemplate []byte
    content []byte
}

type MdxToken struct {
    token string
    param []string
    startPos int
    endPos int
}

type MdxTokenItem struct {
    tokenName string
    paramNum int
}

type MdxError struct {
    s string
}

func (e *MdxError) Error() string {
    return e.s
}

func NewError(text string) error {
    return &MdxError{text}
}

func NewMdx() Generator {
    return &MdxGenerator{[]byte(``), []byte(``), []byte(``)}
}

func (sg *MdxGenerator) SetTemplate(tpl []byte) error {
    sg.template = tpl
    return nil
}

func (sg *MdxGenerator) Render(v variables.Variables) error {
    // parse template
    tpl := string(sg.template)
    sg.parsedTemplate = sg.Tpl_Parse(tpl)
    // generate markdown
    md := markdown.New(markdown.HTML(true), markdown.Tables(true), markdown.Linkify(false))
    sg.content = []byte(md.RenderToString(sg.parsedTemplate))
    return nil
}

func (sg *MdxGenerator) GetContent() []byte {
    return sg.content
}

func (sg *MdxGenerator) Tpl_Parse(tpl string) []byte {
    parsedTpl := ""
    parsingTpl := tpl
    for {
        // calc current tpl length
        strLen := len(parsingTpl)
        // find left {{
        ldIdx := strings.Index(parsingTpl, LEFT_DELIMITER)
        // break conditions
        if strLen != 0 && ldIdx != -1 {
            substr := parsingTpl[ldIdx+2:strLen]
            err, token := sg.Tpl_FindToken(substr, ldIdx+2)
            if err != nil {
                break
            }
            if token.token == TOKEN_INCLUDE {
                fileCnt := sg.Tpl_DoInclude(token.param[0])
                parsedTpl = parsingTpl[0:ldIdx] + fileCnt + parsingTpl[token.endPos:strLen]
                parsingTpl = parsedTpl
            }
        } else {
            break
        }
    }

    return []byte(parsedTpl)
}

func (sg *MdxGenerator) Tpl_FindToken(input string, basePos int) (error, MdxToken) {
    var tokenParam = make([]string, MAX_PARAM_NUM)
    token := MdxToken{"", tokenParam, basePos, -1}
    // find space
    sdIdx := strings.Index(input, SPACE_DELIMITER)
    if sdIdx != -1 {
        substr := input[0:sdIdx]
        switch substr {
        case TOKEN_INCLUDE:
            token.token = substr
            rdIdx := strings.Index(input, RIGHT_DELIMITER)
            token.param[0] = input[sdIdx+1:rdIdx]
            token.endPos = basePos + rdIdx + 2
            return nil, token
        default:
            return NewError("No Token Found"), token
        }
    }

    return NewError("No Token Found"), token
}

func (sg *MdxGenerator) Tpl_DoInclude(path string) string {
    dat, err := ioutil.ReadFile(path)
    if err != nil {
        panic(err)
    }

    return string(dat)
}
