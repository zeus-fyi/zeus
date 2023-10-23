package poseidon

import (
	"context"
	"fmt"
	"time"

	"github.com/cavaliergopher/grab/v3"
	"github.com/rs/zerolog/log"
)

func DownloadFile(ctx context.Context, dataDir, url string) error {
	// download procedure
	client := grab.NewClient()
	// Downloads to your datadir
	req, err := grab.NewRequest(dataDir, url)
	if err != nil {
		log.Err(err).Msgf("DownloadFile: NewRequest, %s", url)
		return err
	}
	// start download
	fmt.Printf("Downloading %v...\n", req.URL())
	resp := client.Do(req)
	fmt.Printf("  %v\n", resp.HTTPResponse.Status)

	// set to any desired time ticker increment for update prints
	timer := time.NewTicker(10 * time.Second)
	defer timer.Stop()

	for {
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
				log.Err(err).Msg("DownloadFile")
				return nil
			}
			fmt.Printf("Downloading Complete")
			return nil
		}
	}
}
