// Package httpserver implements HTTP server.
package httpserver

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v3"
)

const (
	_defaultAddr            = ":80"
	_defaultReadTimeout     = 5 * time.Second
	_defaultWriteTimeout    = 5 * time.Second
	_defaultShutdownTimeout = 3 * time.Second
)

// Server -.
type Server struct {
	App    *fiber.App
	notify chan error

	address         string
	prefork         bool
	readTimeout     time.Duration
	writeTimeout    time.Duration
	shutdownTimeout time.Duration
}

type structValidator struct {
	validate *validator.Validate
}

// Validate Validator needs to implement the Validate method
func (v *structValidator) Validate(out any) error {
	return v.validate.Struct(out)
}

// New -.
func New(opts ...Option) *Server {
	s := &Server{
		App:             nil,
		notify:          make(chan error, 1),
		address:         _defaultAddr,
		readTimeout:     _defaultReadTimeout,
		writeTimeout:    _defaultWriteTimeout,
		shutdownTimeout: _defaultShutdownTimeout,
	}

	// Custom options
	for _, opt := range opts {
		opt(s)
	}

	app := fiber.New(fiber.Config{
		ReadTimeout:     s.readTimeout,
		WriteTimeout:    s.writeTimeout,
		JSONDecoder:     json.Unmarshal,
		JSONEncoder:     json.Marshal,
		StructValidator: &structValidator{validate: validator.New()},
	})

	s.App = app

	return s
}

// Start -.
func (s *Server) Start() {
	go func() {
		s.notify <- s.App.Listen(s.address, fiber.ListenConfig{
			EnablePrefork:     s.prefork,
			EnablePrintRoutes: true,
		})

		close(s.notify)
	}()
}

// Notify -.
func (s *Server) Notify() <-chan error {
	return s.notify
}

// Shutdown -.
func (s *Server) Shutdown() error {
	return s.App.ShutdownWithTimeout(s.shutdownTimeout)
}
