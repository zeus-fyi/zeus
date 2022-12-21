package hercules_chain_snapshots

import "github.com/labstack/echo/v4"

func DownloadChainSnapshotHandler(c echo.Context) error {
	request := new(DownloadChainSnapshotRequest)
	if err := c.Bind(request); err != nil {
		return err
	}
	return request.Download(c)
}
