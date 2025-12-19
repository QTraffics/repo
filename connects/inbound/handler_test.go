package inbound

import (
	"context"
	"testing"

	"github.com/qtraffics/qtfra/ex"

	"github.com/stretchr/testify/assert"
)

func TestHandlerChain(t *testing.T) {
	var chain HandlerChain[*ConnHandlerContext] = &ConnHandlerChain{}
	var (
		afterNext   bool
		requireTrue bool

		shouldNotNil error
	)
	chain.Add(func(ctx *ConnHandlerContext) error {
		assert.False(t, afterNext)
		if err := ctx.Next(); err != nil {
			return err
		}
		assert.True(t, afterNext)
		return nil
	})

	chain.Add(func(ctx *ConnHandlerContext) error {
		requireTrue = true
		afterNext = true
		return ctx.Next()
	})

	chain.Add(func(ctx *ConnHandlerContext) error {
		shouldNotNil = ex.New("not nil")
		return ctx.Next()
	})

	chain.Add(func(ctx *ConnHandlerContext) error {
		assert.NotNil(t, shouldNotNil)
		return ctx.Next()
	})

	err := chain.Entrypoint(GetConnContext(context.Background()))

	assert.Nil(t, err)
	assert.True(t, requireTrue)
}
