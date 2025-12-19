package inbound

import (
	"context"
	"net"

	"github.com/qtraffics/qnetwork/addrs"
	"github.com/qtraffics/qnetwork/netio"
	"github.com/qtraffics/qtfra/enhancements/iolib"
)

type HandlerFunc[C HandlerContext] func(C) error

type HandlerContext interface {
	Next() error
}

type HandlerChain[C HandlerContext] interface {
	Len() int
	Add(hh ...HandlerFunc[C])
	Purge()
	Entrypoint(c C) error
}

type baseHandlerChain[C HandlerContext, HS ~[]HandlerFunc[C]] struct {
	handlers HS
}

func (b *baseHandlerChain[C, HS]) Len() int {
	return len(b.handlers)
}

func (b *baseHandlerChain[C, HS]) Add(hh ...HandlerFunc[C]) {
	b.handlers = append(b.handlers, hh...)
}

func (b *baseHandlerChain[C, HS]) Add0(name string, h HandlerFunc[C]) {
}

func (b *baseHandlerChain[C, HS]) Purge() {
	b.handlers = nil
}

func (b *baseHandlerChain[C, HS]) Entrypoint(c C) error {
	panic("not implement")
}

type (
	ConnHandler   = HandlerFunc[*ConnHandlerContext]
	PacketHandler = HandlerFunc[*PacketHandlerContext]

	ConnHandlerChain struct {
		baseHandlerChain[*ConnHandlerContext, []ConnHandler]
	}

	PacketHandlerChain struct {
		baseHandlerChain[*PacketHandlerContext, []PacketHandler]
	}
)

func (c *PacketHandlerChain) Entrypoint(chc *PacketHandlerContext) error {
	chc.reset()
	chc._chain = c.handlers
	err := chc.Next()
	if err == nil && chc._index < len(chc._chain) {
		panic("context quit too early")
	}
	return err
}

func (c *ConnHandlerChain) Entrypoint(chc *ConnHandlerContext) error {
	chc.reset()
	chc._chain = c.handlers
	err := chc.Next()
	if err == nil && chc._index < len(chc._chain) {
		panic("context quit too early")
	}
	return err
}

type ConnHandlerContext struct {
	Context context.Context

	Conn net.Conn

	_index int
	_chain []HandlerFunc[*ConnHandlerContext]
}

func (c *ConnHandlerContext) Next() error {
	c._index++
	if c._index < len(c._chain) {
		return c._chain[c._index](c)
	}
	return nil
}

func (c *ConnHandlerContext) reset() {
	c.Context = context.Background()
	c.Conn = nil

	c._chain = nil
	c._index = -1
}

type PacketHandlerContext struct {
	Context     context.Context
	Source      addrs.Socksaddr
	Destination addrs.Socksaddr

	Writer netio.PacketWriter

	_index int
	_chain []HandlerFunc[*PacketHandlerContext]
}

func (c *PacketHandlerContext) Next() error {
	c._index++
	if c._index < len(c._chain) {
		return c._chain[c._index](c)
	}
	return nil
}

func (c *PacketHandlerContext) reset() {
	c.Context = context.Background()
	c.Source = addrs.Socksaddr{}
	c.Destination = addrs.Socksaddr{}
	if c.Writer != nil {
		_ = iolib.Close(c.Writer)
	}
	c._index = -1
	c._chain = nil
}
