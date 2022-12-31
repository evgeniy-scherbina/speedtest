package speedtest

import "github.com/pkg/errors"

type Provider uint8

const (
	OoklaProvider Provider = iota
	NetflixProvider
)

type Result struct {
	DLSpeed float64
	ULSpeed float64
}

func GetResult(provider Provider) (*Result, error) {
	switch provider {
	case OoklaProvider:
		ookla := newOokla()
		return ookla.GetResult()
	case NetflixProvider:
		netflix := newNetflix()
		return netflix.GetResult()
		return &Result{}, nil
	default:
		return nil, errors.Errorf("unknown provider")
	}
}
