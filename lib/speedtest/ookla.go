package speedtest

import (
	"github.com/pkg/errors"
	"github.com/showwin/speedtest-go/speedtest"
)

type ookla struct {}

func newOokla() *ookla {
	return &ookla{}
}

func (o *ookla) GetResult() (*Result, error) {
	user, err := speedtest.FetchUserInfo()
	if err != nil {
		return nil, errors.Wrap(err, "cannot fetch user information. http://www.speedtest.net/speedtest-config.php is temporarily unavailable.")
	}

	servers, err := speedtest.FetchServers(user)
	if err != nil {
		return nil, errors.Wrap(err, "can't fetch servers")
	}
	if len(servers) == 0 {
		return nil, errors.Errorf("no available servers")
	}
	server := servers[0]

	if err := o.startTest(server); err != nil {
		return nil, err
	}

	return &Result{
		DLSpeed: server.DLSpeed,
		ULSpeed: server.ULSpeed,
	}, nil
}

func (o *ookla) startTest(server *speedtest.Server) error {
	err := server.PingTest()
	if err != nil {
		return errors.Wrap(err, "can't ping server")
	}

	err = server.DownloadTest(false)
	if err != nil {
		return errors.Wrap(err, "can't test download")
	}
	err = server.UploadTest(false)
	if err != nil {
		return errors.Wrap(err, "can't test upload")
	}

	valid := server.CheckResultValid()
	if !valid {
		return errors.Errorf("result seems to be wrong. Please speedtest again.")
	}

	return nil
}
