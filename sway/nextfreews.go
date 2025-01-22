package sway

import (
	"context"
	"sort"
)

func (s Sway) NextFreeWS(ctx context.Context) (int, error) {
	ws, err := s.cli.GetWorkspaces(ctx)
	if err != nil {
		return 0, err
	}

	// we need it in order of the ws numbers to sanely find a "gap"
	sort.Slice(ws, func(i, j int) bool {
		return ws[i].Num < ws[j].Num
	})

	for i, v := range ws {
		// it's a gap!
		if v.Num > int64(i+1) {
			return i + 1, nil
		}
	}

	// no gaps? it'll just be after.
	return len(ws) + 1, nil
}
