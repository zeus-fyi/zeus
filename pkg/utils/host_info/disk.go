package host_info

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/shirou/gopsutil/v3/disk"
)

func GetDiskUsageStats(ctx context.Context, path string) (*disk.UsageStat, error) {
	usage, err := disk.UsageWithContext(ctx, path)
	if err != nil {
		log.Ctx(ctx).Err(err).Msgf("GetDiskUsageStats %s", path)
		return nil, err
	}
	return usage, err
}
