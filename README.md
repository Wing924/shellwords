# shellwords
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
line should be `./foo\ \'a\ b\ c\'\"`
```

## Thanks

This is based on [go-shellwords](https://github.com/mattn/go-shellwords) and ruby module: (rubysl-shellwords)[https://github.com/rubysl/rubysl-shellwords]
