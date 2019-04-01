package main

import (
	"context"

	proto "./genproto"
)

const (
	avoidNoopCurrencyConversionRPC = false
)

func (fe *frontendServer) sayHello(ctx context.Context) ([]string, error) {
	resp, err := proto.NewAccountServiceClient(fe.accountSvcConn).
		SayHello(ctx, &proto.Empty{})
	return resp, err
}
