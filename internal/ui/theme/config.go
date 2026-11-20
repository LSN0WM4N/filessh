package theme

import (
	"os"
	"strings"
)

type Config struct {
	// utf8 enable/disable the UTF-8 Output
	utf8 bool
}

type Options struct {
	// forceBackward disables UTF-8 output regardless of the environment.
	forceBackward bool

	// getenv looks up environment variables. If nil, os.Getenv is used.
	getenv func(string) string
}

// New creates a new Config using the provided Options.
//
// The behavior is determined as follows:
//   - If opts.getenv is nil, it defaults to os.Getenv.
//   - If opts.forceBackward is true, UTF-8 support is disabled
//     (Config.utf8 is false) and no environment detection is performed.
//   - Otherwise, UTF-8 support is detected from the environment
//     variables returned by opts.getenv.
//
// Injecting a custom getenv function is mainly useful for testing,
// as it allows simulating different environments without modifying
// the real process environment.
func New(opts Options) Config {
	if opts.getenv == nil {
		opts.getenv = os.Getenv
	}

	if opts.forceBackward {
		return Config{utf8: false}
	}
	return Config{utf8: detectUTF8(opts.getenv)}
}

// UTF8 say if the UTF-8  Output is enabled in the program
func (c Config) UFT8() bool {
	return c.utf8
}

func detectUTF8(getenv func(string) string) bool {
	for _, key := range []string{"LC_ALL", "LC_CTYPE", "LANG"} {
		v := strings.ToLower(getenv(key))
		if strings.Contains(v, "utf-8") || strings.Contains(v, "utf8") {
			return true
		}
	}
	return false
}
