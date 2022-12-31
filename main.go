package main

import (
	"fmt"
	"log"

	"github.com/evgeniy-scherbina/sandbox/speedtest/lib/speedtest"

)

func main() {
	if false {
		result, err := speedtest.GetResult(speedtest.OoklaProvider)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("result: %v\n", result)
	}

	if true {
		result, err := speedtest.GetResult(speedtest.NetflixProvider)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("result: %v\n", result)
	}
}
