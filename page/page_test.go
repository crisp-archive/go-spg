package page

import (
    "testing"
)

func TestNewPage(t *testing.T) {
    err, page := New(
        "/home/users/zhangwanlong/project/go-spg/testing/",
        "/home/users/zhangwanlong/project/go-spg/testing/favorite-things.html",
        "favorite-things.md",
        []byte(``), []byte(``), []byte(``))

    err = page.LoadTemplate()
    if err != nil {
        t.Error("Test LoadTemplate failed")
    }

    err = page.Compile()
    if err != nil {
        t.Error("Test Compile failed")
    }

    err = page.Build()
    if err != nil {
        t.Error("Test Build failed")
    }

    page.Output()
}
