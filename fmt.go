package streams

import "fmt"

const (
	DefaultFirstLen      = 5
	DefaultLastLen       = 5
	DefaultMaxSkip       = 1000
	DefaultSkippedFormat = "..."
	DefaultSeparator     = ", "
)

type formatter[T any] struct {
	stream Stream[T]
	cfg    formatterConfig
}

func (f *formatter[T]) Format(state fmt.State, verb rune) {
	if _, err := fmt.Fprint(state, f.cfg.prefix); err != nil {
		return
	}
	firstElement := true
	first, skipped, last, hasRest := f.stream.CollectFirstLast(f.cfg.firstLen, f.cfg.lastLen, f.cfg.maxSkip)
	fmtString := fmt.Sprintf("%%%c", verb)
	if f.cfg.firstLen > 0 {
		for i, v := range first {
			if !firstElement {
				if _, err := fmt.Fprint(state, f.cfg.separator); err != nil {
					return
				}
			}
			firstElement = false
			if f.cfg.showIndex {
				if _, err := fmt.Fprintf(state, "[%d]: ", i); err != nil {
					return
				}
			}
			if _, err := fmt.Fprintf(state, fmtString, v); err != nil {
				return
			}
		}
	}
	if skipped > 0 && f.cfg.firstLen > 0 && f.cfg.lastLen > 0 {
		if !firstElement {
			if _, err := fmt.Fprint(state, f.cfg.separator); err != nil {
				return
			}
		}
		if f.cfg.showSkipped && f.cfg.skippedFormat != "" {
			if _, err := fmt.Fprintf(state, f.cfg.skippedFormat, skipped); err != nil {
				return
			}
		} else {
			if _, err := fmt.Fprintf(state, f.cfg.skippedFormat); err != nil {
				return
			}
		}
		if f.cfg.skippedFormat != "" {
			firstElement = false
		}
	}
	if f.cfg.lastLen > 0 {
		if skipped < 0 {
			last = last[-skipped:]
		}
		idxOfs := len(first) + skipped
		for i, v := range last {
			if !firstElement {
				if _, err := fmt.Fprint(state, f.cfg.separator); err != nil {
					return
				}
			}
			firstElement = false
			if f.cfg.showIndex {
				if _, err := fmt.Fprintf(state, "[%d]: ", i+idxOfs); err != nil {
					return
				}
			}
			if _, err := fmt.Fprintf(state, fmtString, v); err != nil {
				return
			}
		}
		if hasRest {
			if !firstElement {
				if _, err := fmt.Fprint(state, f.cfg.separator); err != nil {
					return
				}
			}
			if _, err := fmt.Fprint(state, "..."); err != nil {
				return
			}
		}
	}
	if _, err := fmt.Fprint(state, f.cfg.suffix); err != nil {
		return
	}
}

type formatterConfig struct {
	firstLen, lastLen, maxSkip int
	showIndex, showSkipped     bool
	skippedFormat              string
	prefix, suffix, separator  string
}

func FirstLen(len int) func(cfg *formatterConfig) {
	return func(f *formatterConfig) {
		f.firstLen = len
	}
}

func LastLen(len int) func(cfg *formatterConfig) {
	return func(f *formatterConfig) {
		f.lastLen = len
	}
}

func MaxSkip(skip int) func(cfg *formatterConfig) {
	return func(f *formatterConfig) {
		f.maxSkip = skip
	}
}

func ShowIndex(show bool) func(cfg *formatterConfig) {
	return func(f *formatterConfig) {
		f.showIndex = show
	}
}

func ShowSkipped(show bool, format string) func(cfg *formatterConfig) {
	return func(f *formatterConfig) {
		f.showSkipped = show
		f.skippedFormat = format
	}
}

func Separator(s string) func(cfg *formatterConfig) {
	return func(f *formatterConfig) {
		f.separator = s
	}
}

func Prefix(s string) func(cfg *formatterConfig) {
	return func(f *formatterConfig) {
		f.prefix = s
	}
}

func Suffix(s string) func(cfg *formatterConfig) {
	return func(f *formatterConfig) {
		f.suffix = s
	}
}

func Format[T any](s Stream[T], params ...func(*formatterConfig)) fmt.Formatter {
	config := formatterConfig{
		firstLen: DefaultFirstLen, lastLen: DefaultLastLen, maxSkip: DefaultMaxSkip, separator: DefaultSeparator,
		skippedFormat: DefaultSkippedFormat,
	}
	for _, param := range params {
		param(&config)
	}
	return &formatter[T]{s, config}
}

func (s Stream[T]) Format(params ...func(*formatterConfig)) fmt.Formatter {
	return Format(s, params...)
}
