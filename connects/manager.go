package connects

import (
	"context"

	"github.com/qtraffics/qtfra/ex"
	"github.com/qtraffics/qtfra/log"
	"github.com/qtraffics/qtfra/services"
	"github.com/qtraffics/repo/connects/inbound"
	"github.com/qtraffics/repo/connects/outbound"
	"github.com/qtraffics/repo/connects/tracker"
)

type NetworkManager interface {
	services.LifeCycle
}

var _ NetworkManager = (*DefaultNetworkManager)(nil)

type DefaultNetworkManager struct {
	logger log.Logger

	inbounds  []inbound.Inbound
	outbounds []outbound.Outbound

	trackers []tracker.Tracker
}

func (nm *DefaultNetworkManager) Start(ctx context.Context) error {
	if len(nm.inbounds) != len(nm.outbounds) {
		return ex.New("len(inbounds) should equal with len(outbounds)")
	}
	return nil
}

func (nm *DefaultNetworkManager) Close() error {
	// TODO implement me
	panic("implement me")
}

func (nm *DefaultNetworkManager) AddInboundOutbound(in inbound.Inbound, out outbound.Outbound) {
	nm.inbounds = append(nm.inbounds, in)
	nm.outbounds = append(nm.outbounds, out)
}

func (nm *DefaultNetworkManager) AddTracker(t ...tracker.Tracker) {
	nm.trackers = append(nm.trackers, t...)
}

func (nm *DefaultNetworkManager) bind(ctx context.Context, in inbound.Inbound, out outbound.Outbound) error {
	if err := in.Start(ctx); err != nil {
		return ex.Cause(err, "start inbound")
	}
	if err := out.Start(ctx); err != nil {
		return ex.Cause(err, "start outbound")
	}
	go func() {
		defer in.Close()
		defer out.Close()
	}()

	return nil
}
