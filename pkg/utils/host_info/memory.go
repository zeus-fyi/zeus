package host_info

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/shirou/gopsutil/v3/mem"
)

func GetVirtualMemoryStats(ctx context.Context) (*mem.VirtualMemoryStat, error) {
	usage, err := mem.VirtualMemoryWithContext(ctx)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("VirtualMemoryStat")
		return nil, err
	}
	return usage, err
}
