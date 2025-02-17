// Copyright 2020 Asim Aslam
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Original source: github.com/micro/go-micro/v3/client/mucp/mucp_test.go

package mucp

import (
	"context"
	"fmt"
	"testing"

	"github.com/tickoalcantara12/micro/v3/service/client"
	"github.com/tickoalcantara12/micro/v3/service/errors"
	"github.com/tickoalcantara12/micro/v3/service/registry"
	"github.com/tickoalcantara12/micro/v3/service/registry/memory"
	"github.com/tickoalcantara12/micro/v3/service/router"
	regRouter "github.com/tickoalcantara12/micro/v3/service/router/registry"
)

func newTestRouter() router.Router {
	reg := memory.NewRegistry(memory.Services(testData))
	return regRouter.NewRouter(router.Registry(reg))
}

func TestCallAddress(t *testing.T) {
	var called bool
	service := "test.service"
	endpoint := "Test.Endpoint"
	address := "10.1.10.1:8080"

	wrap := func(cf client.CallFunc) client.CallFunc {
		return func(ctx context.Context, node string, req client.Request, rsp interface{}, opts client.CallOptions) error {
			called = true

			if req.Service() != service {
				return fmt.Errorf("expected service: %s got %s", service, req.Service())
			}

			if req.Endpoint() != endpoint {
				return fmt.Errorf("expected service: %s got %s", endpoint, req.Endpoint())
			}

			if node != address {
				return fmt.Errorf("expected address: %s got %s", address, node)
			}

			// don't do the call
			return nil
		}
	}

	r := newTestRouter()

	c := NewClient(
		client.Router(r),
		client.WrapCall(wrap),
	)

	req := c.NewRequest(service, endpoint, nil)

	// test calling remote address
	if err := c.Call(context.Background(), req, nil, client.WithAddress(address)); err != nil {
		t.Fatal("call with address error", err)
	}

	if !called {
		t.Fatal("wrapper not called")
	}

}

func TestCallRetry(t *testing.T) {
	service := "test.service"
	endpoint := "Test.Endpoint"
	address := "10.1.10.1"

	var called int

	wrap := func(cf client.CallFunc) client.CallFunc {
		return func(ctx context.Context, node string, req client.Request, rsp interface{}, opts client.CallOptions) error {
			called++
			if called == 1 {
				return errors.InternalServerError("test.error", "retry request")
			}

			// don't do the call
			return nil
		}
	}

	r := newTestRouter()
	c := NewClient(
		client.Router(r),
		client.WrapCall(wrap),
		client.Retry(client.RetryOnError),
	)

	req := c.NewRequest(service, endpoint, nil)

	// test calling remote address
	if err := c.Call(context.Background(), req, nil, client.WithAddress(address)); err != nil {
		t.Fatal("call with address error", err)
	}

	// num calls
	if called < c.Options().CallOptions.Retries+1 {
		t.Fatal("request not retried")
	}
}

func TestCallWrapper(t *testing.T) {
	var called bool
	id := "test.1"
	service := "test.service"
	endpoint := "Test.Endpoint"
	address := "10.1.10.1:8080"

	wrap := func(cf client.CallFunc) client.CallFunc {
		return func(ctx context.Context, node string, req client.Request, rsp interface{}, opts client.CallOptions) error {
			called = true

			if req.Service() != service {
				return fmt.Errorf("expected service: %s got %s", service, req.Service())
			}

			if req.Endpoint() != endpoint {
				return fmt.Errorf("expected service: %s got %s", endpoint, req.Endpoint())
			}

			if node != address {
				return fmt.Errorf("expected address: %s got %s", address, node)
			}

			// don't do the call
			return nil
		}
	}

	r := newTestRouter()
	c := NewClient(
		client.Router(r),
		client.WrapCall(wrap),
	)

	r.Options().Registry.Register(&registry.Service{
		Name:    service,
		Version: "latest",
		Nodes: []*registry.Node{
			{
				Id:      id,
				Address: address,
			},
		},
	})

	req := c.NewRequest(service, endpoint, nil)
	if err := c.Call(context.Background(), req, nil); err != nil {
		t.Fatal("call wrapper error", err)
	}

	if !called {
		t.Fatal("wrapper not called")
	}
}
