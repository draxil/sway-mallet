package sway

import (
	"context"
	"fmt"
)

func (s Sway) FocusedWin(ctx context.Context) (string, error) {
	t, err := s.cli.GetTree(ctx)
	if err != nil {
		return "", err
	}

	fn := t.FocusedNode()

	if fn == nil {
		return "", nil
	}

	id := fmt.Sprintf("%d", fn.ID)

	return id, nil
}
