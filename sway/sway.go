package sway

import (
	"github.com/joshuarubin/go-sway"
)

type Sway struct {
	cli sway.Client
}

func New(c sway.Client) Sway {
	return Sway{c}
}
