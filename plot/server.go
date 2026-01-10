package plot

import (
	"fmt"
	"net/http"
)

// HTTPServer defines the interface for an HTTP server that Chart will use
type HTTPServer interface {
	// RegisterHandler registers a handler for a specific route
	RegisterHandler(path string, handler http.HandlerFunc)

	// RegisterFileServer registers a handler to serve static files
	RegisterFileServer(path string, fs http.FileSystem)

	// Start starts the HTTP server on the specified port
	Start(port int) error
}

// StandardHTTPServer implements the HTTPServer interface using the standard http package
type StandardHTTPServer struct{}

// NewStandardHTTPServer creates a new instance of StandardHTTPServer
func NewStandardHTTPServer() *StandardHTTPServer {
	return &StandardHTTPServer{}
}

// RegisterHandler registers a handler for a specific route
func (s *StandardHTTPServer) RegisterHandler(path string, handler http.HandlerFunc) {
	http.HandleFunc(path, handler)
}

// RegisterFileServer registers a handler to serve static files
func (s *StandardHTTPServer) RegisterFileServer(path string, fs http.FileSystem) {
	http.Handle(path, http.FileServer(fs))
}

// Start starts the HTTP server on the specified port
func (s *StandardHTTPServer) Start(port int) error {
	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

// MuxHTTPServer implements the HTTPServer interface using the standard http package
type MuxHTTPServer struct {
	mux    *http.ServeMux
	server *http.Server
}

// NewMuxHTTPServer creates a new instance of MuxHTTPServer
func NewMuxHTTPServer() *MuxHTTPServer {
	mux := http.NewServeMux()
	return &MuxHTTPServer{
		mux: mux,
	}
}

// RegisterHandler registers a handler for a specific route
func (s *MuxHTTPServer) RegisterHandler(path string, handler http.HandlerFunc) {
	s.mux.HandleFunc(path, handler)
}

// RegisterFileServer registers a handler to serve static files
func (s *MuxHTTPServer) RegisterFileServer(path string, fs http.FileSystem) {
	s.mux.Handle(path, http.FileServer(fs))
}

// Start starts the HTTP server on the specified port
func (s *MuxHTTPServer) Start(port int) error {
	s.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: s.mux,
	}
	return s.server.ListenAndServe()
}

func (s *MuxHTTPServer) Close() error {
	if s.server == nil {
		return nil
	}
	return s.server.Close()
}
