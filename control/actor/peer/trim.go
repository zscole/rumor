package peer

import (
	"context"
	"github.com/protolambda/rumor/control/actor/base"
	"time"
)

type PeerTrimCmd struct {
	*base.Base
	Timeout  time.Duration `ask:"[which]" help:"Timeout for trimming."`
}

func (c *PeerTrimCmd) Help() string {
	return "Trim peers, with timeout."
}

func (c *PeerTrimCmd) Default() {
	c.Timeout = time.Second * 2
}

func (c *PeerTrimCmd) Run(ctx context.Context, args ...string) error {
	h, err := c.Host()
	if err != nil {
		return err
	}
	trimCtx, _ := context.WithTimeout(ctx, c.Timeout)
	h.ConnManager().TrimOpenConns(trimCtx)
	return nil
}
