package sway

import (
	"context"
	"fmt"
)

func (s Sway) Run(ctx context.Context, cmdfmt string, arg ...any) error {
	cmd := fmt.Sprintf(cmdfmt, arg...)
	_, err := s.cli.RunCommand(ctx, cmd)
	if err != nil {
		return err
	}
	return nil
}
