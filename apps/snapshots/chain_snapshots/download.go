package hercules_chain_snapshots

import (
	"fmt"
	"net/http"
	"time"

	"github.com/cavaliergopher/grab/v3"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	v1_common_routes "github.com/zeus-fyi/hercules/api/v1/common"
	"github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/compression"
)

type DownloadChainSnapshotRequest struct {
	BucketRequest
}

func (t *DownloadChainSnapshotRequest) Download(c echo.Context) error {
	// download procedure
	client := grab.NewClient()
	// Downloads to your datadir
	req, err := grab.NewRequest(v1_common_routes.CommonManager.DirIn, v1_common_routes.CommonManager.BucketURL)
	if err != nil {
		log.Err(err).Msgf("DownloadChainSnapshotRequest: NewRequest, %s", v1_common_routes.CommonManager.BucketURL)
		return c.JSON(http.StatusInternalServerError, err)
	}
	// start download
	fmt.Printf("Downloading %v...\n", req.URL())
	resp := client.Do(req)
	fmt.Printf("  %v\n", resp.HTTPResponse.Status)

	// set to any desired max time
	timer := time.NewTicker(12 * time.Hour)
	defer timer.Stop()
	select {
	case <-timer.C:
		fmt.Printf("  transferred %v / %v bytes (%.2f%%)\n",
			resp.BytesComplete(),
			resp.Size(),
			100*resp.Progress())
	case <-resp.Done:
		// download is complete
		err = resp.Err()
		if err != nil {
			log.Err(err).Msg("DownloadChainSnapshotRequest")
			return c.JSON(http.StatusInternalServerError, err)
		}
	}
	cmp := compression.NewCompression()

	// choose whatever compression or file naming you want here if you're using your own source
	v1_common_routes.CommonManager.FnIn = t.ClientName + ".tar.lz4"
	v1_common_routes.CommonManager.FnOut = t.ClientName

	err = cmp.Lz4Decompress(&v1_common_routes.CommonManager.Path)
	if err != nil {
		log.Err(err).Msg("DownloadChainSnapshotRequest: Lz4Decompress")
		return c.JSON(http.StatusInternalServerError, err)
	}
	// removes compressed file
	err = v1_common_routes.CommonManager.RemoveFileInPath()
	if err != nil {
		log.Err(err).Msg("DownloadChainSnapshotRequest: RemoveFileInPath")
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, nil)
}
