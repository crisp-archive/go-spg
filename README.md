# Go-spg

## Introduction

SPG(single page generator) is a single static page generator with Go lang, with the power of Markdown and Go templates.

## Installation

```
go get -u github.com/russross/blackfriday
go get github.com/crispgm/go-spg
```

## Getting Started

```
// generate mode
./spg -generate -src=/path/to/src -dst=/path/to/dst -layout=/path/to/index.tpl
   
// live mode
./spg -live=8082 -src=/path/to/src -dst=/path/to/dst -layout=/path/to/index.tpl
```

## License

[MIT LICENSE](https://github.com/crispgm/go-spg/blob/master/LICENSE)

