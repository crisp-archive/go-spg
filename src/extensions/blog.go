package main

import (
    "time"
)

type (
    Blog struct {
        title string
        publishTime int32
        tags []string
        isRepost bool
        abstract string
        content string
        fileName string
    }
)

func (blog *Blog) getTitle() string {
    return blog.title
}

func (blog *Blog) getAbstract() string {
    return blog.abstract
}

func (blog *Blog) getContent() string {
    return blog.content
}

func (blog *Blog) render(tpl Template) string {
}