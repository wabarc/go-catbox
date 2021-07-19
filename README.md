# go-catbox

`go-catbox` is a toolkit to help upload files to [Catbox](https://catbox.moe/).

## Installation

The simplest, cross-platform way is to download from [GitHub Releases](https://github.com/wabarc/go-catbox/releases) and place the executable file in your PATH.

Via Golang package get command

```sh
go get -u github.com/wabarc/go-catbox/cmd/catbox
```

From [gobinaries.com](https://gobinaries.com):

```sh
$ curl -sf https://gobinaries.com/wabarc/go-catbox | sh
```

## Usage

Command-line:

```sh
$ catbox
A CLI tool help upload files to Catbox.

Usage:

  catbox [options] [file1] ... [fileN]
```

Go package:
```go
import (
        "fmt"

        "github.com/wabarc/go-catbox"
)

func main() {
        if url, err := catbox.New(nil).Upload(path); err != nil {
            fmt.Fprintf(os.Stderr, "catbox: %v\n", err)
        } else {
            fmt.Fprintf(os.Stdout, "%s  %s\n", url, path)
        }
}
```

## License

This software is released under the terms of the MIT. See the [LICENSE](https://github.com/wabarc/go-catbox/blob/main/LICENSE) file for details.
