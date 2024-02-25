package http

import (
	"context"
	"github.com/samber/do"
	"github.com/urfave/cli/v2"
	"gocourse22/cmd/flag"
	"net/http"
	"time"
)

// HTTP is the http server
type HTTP struct {
	srv             *http.Server
	shutdownTimeout time.Duration
}

// NewHTTP creates a new server
func NewHTTP(srv *http.Server, st time.Duration) *HTTP {
	return &HTTP{srv: srv, shutdownTimeout: st}
}

// Start starts the server
func (s HTTP) Start() error {
	return s.srv.ListenAndServe()
}

// Stop gracefully stops the server
// https://golang.org/pkg/net/http/#Server.Shutdown
func (s HTTP) Stop(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}

func (s HTTP) Shutdown() error {
	ctx, cancelFunc := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancelFunc()
	return s.Stop(ctx)
}

// NewServer provides a new http server
func NewServer(i *do.Injector, r *Router) *HTTP {
	c := do.MustInvoke[*cli.Context](i)

	serv := &http.Server{
		Addr:              c.String(flag.HTTPServerAddress),
		Handler:           r.Handler(),
		ReadTimeout:       c.Duration(flag.HTTPReadTimeout),
		ReadHeaderTimeout: c.Duration(flag.HTTPReadTimeout),
	}

	return NewHTTP(serv, c.Duration(flag.HTTPShutdownTimeout))
}
