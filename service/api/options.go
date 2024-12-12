package api

type Options interface {
	apply(*options)
}

type options struct {
	apiKey string
}

// funcOption wraps a function that modifies options into an implementation of the Option interface.
type funcOption struct {
	f func(*options)
}

func (fdo *funcOption) apply(do *options) {
	fdo.f(do)
}

func newFuncDialOption(f func(*options)) *funcOption {
	return &funcOption{
		f: f,
	}
}

// WithAPIKey returns an Option which configures the API key of the Scrapeless.
func WithAPIKey(apiKey string) Options {
	return newFuncDialOption(func(o *options) {
		o.apiKey = apiKey
	})
}
