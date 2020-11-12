package main

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/go-kit/kit/log"
	kitHttp "github.com/go-kit/kit/transport/http"
	goaHttp "goa.design/goa/v3/http"
	httpMdlwr "goa.design/goa/v3/http/middleware"
	"goa.design/goa/v3/middleware"

	"github.com/selmison/dojo-1/gen/decolar"
	decolarKitSvr "github.com/selmison/dojo-1/gen/http/decolar/kitserver"
	decolarSvr "github.com/selmison/dojo-1/gen/http/decolar/server"
)

// handleHTTPServer starts configures and starts a HTTP server on the given
// URL. It shuts down the server if any error is received in the error channel.
func handleHTTPServer(ctx context.Context, u *url.URL, decolarEndpoints *decolar.Endpoints, wg *sync.WaitGroup, errc chan error, logger log.Logger, debug bool) {

	// Provide the transport specific request decoder and response encoder.
	// The goa http package has built-in support for JSON, XML and gob.
	// Other encodings can be used by providing the corresponding functions,
	// see goa.design/implement/encoding.
	var (
		dec = goaHttp.RequestDecoder
		enc = goaHttp.ResponseEncoder
	)

	// Build the service HTTP request multiplexer and configure it to serve
	// HTTP requests to the service endpoints.
	var mux goaHttp.Muxer
	{
		mux = goaHttp.NewMuxer()
	}

	// Wrap the endpoints with the transport specific layers. The generated
	// server packages contains code generated from the design which maps
	// the service input and output data structures to HTTP requests and
	// responses.
	var (
		decolarCreatePaisHandler     *kitHttp.Server
		decolarCreateCompaniaHandler *kitHttp.Server
		decolarServer                *decolarSvr.Server
	)
	{
		eh := errorHandler(logger)
		decolarCreatePaisHandler = kitHttp.NewServer(
			decolarEndpoints.CreatePais,
			decolarKitSvr.DecodeCreatePaisRequest(mux, dec),
			decolarKitSvr.EncodeCreatePaisResponse(enc),
			kitHttp.ServerErrorEncoder(decolarKitSvr.EncodeCreatePaisError(enc, nil)),
		)
		decolarCreateCompaniaHandler = kitHttp.NewServer(
			decolarEndpoints.CreateCompania,
			decolarKitSvr.DecodeCreateCompaniaRequest(mux, dec),
			decolarKitSvr.EncodeCreateCompaniaResponse(enc),
			kitHttp.ServerErrorEncoder(decolarKitSvr.EncodeCreateCompaniaError(enc, nil)),
		)
		decolarServer = decolarSvr.New(decolarEndpoints, mux, dec, enc, eh, nil)
	}

	// Configure the mux.
	decolarKitSvr.MountCreatePaisHandler(mux, decolarCreatePaisHandler)
	decolarKitSvr.MountCreateCompaniaHandler(mux, decolarCreateCompaniaHandler)

	// Wrap the multiplexer with additional middlewares. Middlewares mounted
	// here apply to all the service endpoints.
	var handler http.Handler = mux
	{
		handler = httpMdlwr.Log(logger)(handler)
		handler = httpMdlwr.RequestID()(handler)
	}

	// Start HTTP server using default configuration, change the code to
	// configure the server as required by your service.
	srv := &http.Server{Addr: u.Host, Handler: handler}
	for _, m := range decolarServer.Mounts {
		logger.Log("info", fmt.Sprintf("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern))
	}

	(*wg).Add(1)
	go func() {
		defer (*wg).Done()

		// Start HTTP server in a separate goroutine.
		go func() {
			logger.Log("info", fmt.Sprintf("HTTP server listening on %q", u.Host))
			errc <- srv.ListenAndServe()
		}()

		<-ctx.Done()
		logger.Log("info", fmt.Sprintf("shutting down HTTP server at %q", u.Host))

		// Shutdown gracefully with a 30s timeout.
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		_ = srv.Shutdown(ctx)
	}()
}

// errorHandler returns a function that writes and logs the given error.
// The function also writes and logs the error unique ID so that it's possible
// to correlate.
func errorHandler(logger log.Logger) func(context.Context, http.ResponseWriter, error) {
	return func(ctx context.Context, w http.ResponseWriter, err error) {
		id := ctx.Value(middleware.RequestIDKey).(string)
		_, _ = w.Write([]byte("[" + id + "] encoding: " + err.Error()))
		logger.Log("info", fmt.Sprintf("[%s] ERROR: %s", id, err.Error()))
	}
}
