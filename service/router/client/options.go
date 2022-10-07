package client

import (
	"context"

	"github.com/tickoalcantara12/micro/v3/service/client"
	"github.com/tickoalcantara12/micro/v3/service/router"
)

type clientKey struct{}

// Client to call router service
func Client(c client.Client) router.Option {
	return func(o *router.Options) {
		if o.Context == nil {
			o.Context = context.WithValue(context.Background(), clientKey{}, c)
			return
		}

		o.Context = context.WithValue(o.Context, clientKey{}, c)
	}
}
