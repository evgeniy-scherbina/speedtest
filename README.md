# speedtest

## Install dependencies
`go mod tidy`

## Run tests
`go test -v ./lib/...`

## Dependencies
`Ookla speedtest library - github.com/showwin/speedtest-go`

`Netflix speedtest library - github.com/ddo/go-fast`

## Note:
It seems that most libraries for Netflix(https://fast.com) provides only API for getting download speed. So I found this lib `https://github.com/calin014/speedfast` and took code for getting upload speed from there.

## Usage API
```go
package main

import (
	"fmt"
	"log"

	"github.com/evgeniy-scherbina/sandbox/speedtest/lib/speedtest"
)

func main() {
	result, err := speedtest.GetResult(speedtest.OoklaProvider)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Result: %v\n", result)

	result, err := speedtest.GetResult(speedtest.NetflixProvider)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Result: %v\n", result)
}

```
