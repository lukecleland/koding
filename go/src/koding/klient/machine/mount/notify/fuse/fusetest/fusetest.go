package fusetest

import (
	"context"

	"koding/klient/machine/index"
	"koding/klient/machine/mount/notify"
)

type BindCache struct {
	dst string
	tmp string
	idx *index.Index
}

var _ notify.Cache = (*BindCache)(nil)

func NewBindCache(dst, tmp string) (*BindCache, error) {
	idx, err := index.NewIndexFiles(dst)
	if err != nil {
		return idx, nil
	}

	return &BindCache{
		dst: dst,
		tmp: tmp,
	}, nil
}

func (bc *BindCache) Commit(idx *index.Change) context.Context {
	return nil
}

func (bc *BindCache) Index() *index.Index {
	return bc.idx
}
