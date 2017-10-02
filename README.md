# shellwords

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![Build Status](https://travis-ci.org/Wing924/shellwords.svg?branch=master)](https://travis-ci.org/Wing924/shellwords)
[![Go Report Card](https://goreportcard.com/badge/github.com/Wing924/shellwords)](https://goreportcard.com/report/github.com/Wing924/shellwords)
[![codecov](https://codecov.io/gh/Wing924/shellwords/branch/master/graph/badge.svg)](https://codecov.io/gh/Wing924/shellwords)
[![GoDoc](https://godoc.org/github.com/Wing924/shellwords?status.svg)](https://godoc.org/github.com/Wing924/shellwords)

A Golang library to manipulate strings according to the word parsing rules of the UNIX Bourne shell.


## Installation

```sh
go get github.com/Wing924/shellwords
```

## Usage

```go
import "github.com/Wing924/shellwords"
```

```go
args, err := shellwords.Split("./foo --bar=baz")
// args should be ["./foo", "--bar=baz"]

args, err := shellwords.Split("./foo 'a b c'")
// args should be ["./foo", "a b c"]
```

```go
line := shellwords.Join([]string{"abc", "d e f"})
// line should be `abc d\ e\ f`
```

```go
line := shellwords.Escape("./foo 'a b c'")
// line should be `./foo\ \'a\ b\ c\'\"`
```

## Thanks

This is based on [go-shellwords](https://github.com/mattn/go-shellwords) and ruby module: [rubysl-shellwords](https://github.com/rubysl/rubysl-shellwords)
