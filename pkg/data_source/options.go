package data_source

import "time"

// Option -.
type Option func(*DataSource)

// MaxPoolSize -.
func MaxPoolSize(size int) Option {
	return func(c *DataSource) {
		c.maxPoolSize = size
	}
}

// ConnAttempts -.
func ConnAttempts(attempts int) Option {
	return func(c *DataSource) {
		c.connAttempts = attempts
	}
}

// ConnTimeout -.
func ConnTimeout(timeout time.Duration) Option {
	return func(c *DataSource) {
		c.connTimeout = timeout
	}
}
