package routines

import (
	"context"
	"fmt"
	"syscall"

	"github.com/rs/zerolog/log"
	"github.com/shirou/gopsutil/v3/process"
)

func KillProcessWithCtx(ctx context.Context, name string) error {
	processes, err := process.Processes()
	if err != nil {
		log.Ctx(ctx).Err(err).Msgf("killProcessWithCtx %s", name)
		return err
	}
	for _, p := range processes {
		n, perr := p.Name()
		if perr != nil {
			log.Ctx(ctx).Err(err).Msgf("killProcessWithCtx %s", name)
			return err
		}
		if n == name {
			log.Ctx(ctx).Info().Msgf("killing process with syscall.SIGINT %s", name)
			return p.SendSignalWithContext(ctx, syscall.SIGINT)
		}
	}
	return fmt.Errorf("process not found")
}

func SuspendProcessWithCtx(ctx context.Context, name string) error {
	processes, err := process.Processes()
	if err != nil {
		log.Ctx(ctx).Err(err).Msgf("SuspendProcessWithCtx %s", name)
		return err
	}
	for _, p := range processes {
		n, perr := p.Name()
		if perr != nil {
			log.Ctx(ctx).Err(err).Msgf("suspendProcessWithCtx %s", name)
			return err
		}
		if n == name {
			log.Ctx(ctx).Info().Msgf("suspending process %s", name)
			return p.SuspendWithContext(ctx)
		}
	}
	return fmt.Errorf("process not found")
}

func ResumeProcessWithCtx(ctx context.Context, name string) error {
	processes, err := process.Processes()
	if err != nil {
		log.Ctx(ctx).Err(err).Msgf("ResumeProcessWithCtx %s", name)
		return err
	}
	for _, p := range processes {
		n, perr := p.Name()
		if perr != nil {
			log.Ctx(ctx).Err(err).Msgf("ResumeProcessWithCtx %s", name)
			return err
		}
		if n == name {
			log.Ctx(ctx).Info().Msgf("resuming process %s", name)
			return p.ResumeWithContext(ctx)
		}
	}
	return fmt.Errorf("process not found")
}
