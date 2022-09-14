# ALP - Accept-Language Parser

![License](https://img.shields.io/badge/license-MIT-green?style=for-the-badge)
![Go](https://img.shields.io/badge/Go-1.19-blue?style=for-the-badge)

This is a parser for the http header Accept-Language.
\
It will split parts into a Language struct {name, quality}.

## Installation

Include this module into your existing go project

```zsh
go get github.com/ThomasBoom89/accept-language-parser
```

## Update

```zsh
go get -u github.com/ThomasBoom89/accept-language-parser
```

## Usage

It is simple to use, just throw your header into the parse function and you will return an array of Languages

```go
languages, err := alp.Parse(header)
if err != nil {
    panic(err)
}
fmt.Println(languages)
```

## License

ALP- Accept-Header Parser
\
Copyright (C) 2022 ThomasBoom89. MIT license.