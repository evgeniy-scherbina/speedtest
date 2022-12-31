package speedtest

import (
	"github.com/ddo/go-fast"
)

type netflix struct{}

func newNetflix() *netflix {
	return &netflix{}
}

func (n *netflix) GetResult() (*Result, error) {
	fastCom := fast.New()

	err := fastCom.Init()
	if err != nil {
		return nil, err
	}

	urls, err := fastCom.GetUrls()
	if err != nil {
		return nil, err
	}

	var (
		kbpsChan             = make(chan float64)
		totalKbps            float64
		numberOfMeasurements int
	)
	go func() {
		for kbps := range kbpsChan {
			totalKbps += kbps
			numberOfMeasurements++
		}
	}()

	err = fastCom.Measure(urls, kbpsChan)
	if err != nil {
		return nil, err
	}

	avgKbps := totalKbps / float64(numberOfMeasurements)
	avgMbps := avgKbps / 1000

	uploadSpeed, err := measureUploadWithFast()
	if err != nil {
		return nil, err
	}

	return &Result{
		DLSpeed: avgMbps,
		ULSpeed: uploadSpeed,
	}, nil
}
