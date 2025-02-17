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
// Original source: github.com/micro/go-micro/v3/network/transport/http/http_proxy.go

package http

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
)

const (
	proxyAuthHeader = "Proxy-Authorization"
)

func getURL(addr string) (*url.URL, error) {
	r := &http.Request{
		URL: &url.URL{
			Scheme: "https",
			Host:   addr,
		},
	}
	return http.ProxyFromEnvironment(r)
}

type pbuffer struct {
	net.Conn
	r io.Reader
}

func (p *pbuffer) Read(b []byte) (int, error) {
	return p.r.Read(b)
}

func proxyDial(conn net.Conn, addr string, proxyURL *url.URL) (_ net.Conn, err error) {
	defer func() {
		if err != nil {
			conn.Close()
		}
	}()

	r := &http.Request{
		Method: http.MethodConnect,
		URL:    &url.URL{Host: addr},
		Header: map[string][]string{"User-Agent": {"micro/latest"}},
	}

	if user := proxyURL.User; user != nil {
		u := user.Username()
		p, _ := user.Password()
		auth := []byte(u + ":" + p)
		basicAuth := base64.StdEncoding.EncodeToString(auth)
		r.Header.Add(proxyAuthHeader, "Basic "+basicAuth)
	}

	if err := r.Write(conn); err != nil {
		return nil, fmt.Errorf("failed to write the HTTP request: %v", err)
	}

	br := bufio.NewReader(conn)
	rsp, err := http.ReadResponse(br, r)
	if err != nil {
		return nil, fmt.Errorf("reading server HTTP response: %v", err)
	}
	defer rsp.Body.Close()
	if rsp.StatusCode != http.StatusOK {
		dump, err := httputil.DumpResponse(rsp, true)
		if err != nil {
			return nil, fmt.Errorf("failed to do connect handshake, status code: %s", rsp.Status)
		}
		return nil, fmt.Errorf("failed to do connect handshake, response: %q", dump)
	}

	return &pbuffer{Conn: conn, r: br}, nil
}

// Creates a new connection
func newConn(dial func(string) (net.Conn, error)) func(string) (net.Conn, error) {
	return func(addr string) (net.Conn, error) {
		// get the proxy url
		proxyURL, err := getURL(addr)
		if err != nil {
			return nil, err
		}

		// set to addr
		callAddr := addr

		// got proxy
		if proxyURL != nil {
			callAddr = proxyURL.Host
		}

		// dial the addr
		c, err := dial(callAddr)
		if err != nil {
			return nil, err
		}

		// do proxy connect if we have proxy url
		if proxyURL != nil {
			c, err = proxyDial(c, addr, proxyURL)
		}

		return c, err
	}
}
