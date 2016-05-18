package porto

import (
	"fmt"

	"golang.org/x/net/context"

	"github.com/noxiouz/stout/isolate"
)

// Box implements isolate.Box interface using Porto
type Box struct {
}

// NewBox creates Box with Porto inside
func NewBox(ctx context.Context, cfg isolate.BoxConfig) (isolate.Box, error) {
	b := &Box{}
	return b, nil
}

func (b *Box) Spool(ctx context.Context, name string, opts isolate.Profile) (err error) {
	return fmt.Errorf("Not implemented")
}

func (b *Box) Spawn(ctx context.Context, opts isolate.Profile, name, executable string, args, env map[string]string) (isolate.Process, error) {

	return nil, fmt.Errorf("Not implemented")
}

// Close releases all resources
func (b *Box) Close() error {
	return nil
}
