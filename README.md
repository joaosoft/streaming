Streaming
================

[![Build Status](https://travis-ci.org/joaosoft/streaming.svg?branch=master)](https://travis-ci.org/joaosoft/streaming) | [![codecov](https://codecov.io/gh/joaosoft/streaming/branch/master/graph/badge.svg)](https://codecov.io/gh/joaosoft/streaming) | [![Go Report Card](https://goreportcard.com/badge/github.com/joaosoft/streaming)](https://goreportcard.com/report/github.com/joaosoft/streaming) | [![GoDoc](https://godoc.org/github.com/joaosoft/streaming?status.svg)](https://godoc.org/github.com/joaosoft/streaming)

A simple streaming package with multiple solutions.

## Setup
Copy default configs
```
make setup
```

## Dependency Management
>### Dep

Project dependencies are managed using Dep. Read more about [Dep](https://github.com/golang/dep).
* Install dependencies: `dep ensure`
* Update dependencies: `dep ensure -update`


>### Go
```
go get github.com/joaosoft/streaming
```

### Code
```go
import (
	"github.com/joaosoft/streaming"
)

func main() {
    ap, err := streaming.NewStreaming()
    if err != nil {
        panic(err)
    }
    
    if err = ap.Start(); err != nil {
        panic(err)
    }
}

```

## Known issues

## Follow me at
Facebook: https://www.facebook.com/joaosoft

LinkedIn: https://www.linkedin.com/in/jo%C3%A3o-ribeiro-b2775438/

##### If you have something to add, please let me know joaosoft@gmail.com