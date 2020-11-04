package printer

import (
	"fmt"
	"io"
)

// DefaultHeader is a header which is used for all generators.
// TODO: use a function to return a header according to the version/time etc.
var DefaultHeader = Printer{
	"Code generated by Semaphore. DO NOT EDIT.",
	"Semaphore version: v2.0.0",
	"Timestamp: Thu, 01 Oct 2020 14:02:49 GMT",
}

// Options contains printer settings.
type Options struct {
	StreamStart string
	LineStart   string
	LineEnd     string
	StreamEnd   string
}

// Printer contains lines to be printed.
type Printer []string

// Print lines to the provided writer.
func (printer Printer) Print(dst io.Writer, options Options) error {
	if options.StreamStart != "" {
		if _, err := fmt.Fprint(dst, options.StreamStart); err != nil {
			return err
		}
	}

	for _, line := range printer {
		if _, err := fmt.Fprintf(dst, "%s%s%s", options.LineStart, line, options.LineEnd); err != nil {
			return err
		}
	}

	if options.StreamEnd != "" {
		_, err := fmt.Fprint(dst, options.StreamEnd)

		return err
	}

	return nil
}
