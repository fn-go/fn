/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package term

import (
	"fmt"
	"io"
	"os"

	wordwrap "github.com/mitchellh/go-wordwrap"
	"github.com/moby/term"
)

type wordWrapWriter struct {
	limit  uint
	writer io.Writer
}

// GetWidth returns the current width of the terminal associated with fd.
func GetWidth(fd uintptr) (*uint16, error) {
	winsize, err := term.GetWinsize(fd)
	if err != nil {
		return nil, fmt.Errorf("unable to get terminal size: %v", err)
	}

	size := winsize.Width

	return &size, nil
}

// NewResponsiveWriter creates a Writer that detects the column width of the
// terminal we are in, and adjusts every line width to fit and use recommended
// terminal sizes for better readability. Does proper word wrapping automatically.
//    if terminal width >= 120 columns		use 120 columns
//    if terminal width >= 100 columns		use 100 columns
//    if terminal width >=  80 columns		use  80 columns
// In case we're not in a terminal or if it's smaller than 80 columns width,
// doesn't do any wrapping.
func NewResponsiveWriter(w io.Writer) io.Writer {
	// TODO capture errors here so that we can be helpful when we think we aren't in a terminal
	file, ok := w.(*os.File)
	if !ok {
		return w
	}
	fd := file.Fd()
	if !term.IsTerminal(fd) {
		return w
	}

	var terminalWidth uint16
	{
		width, _ := GetWidth(fd)
		if width == nil {
			return w
		}

		terminalWidth = *width
	}

	var limit uint
	switch {
	case terminalWidth >= 120:
		limit = 120
	case terminalWidth >= 100:
		limit = 100
	case terminalWidth >= 80:
		limit = 80
	}

	return NewWordWrapWriter(w, limit)
}

// NewWordWrapWriter is a Writer that supports a limit of characters on every line
// and does auto word wrapping that respects that limit.
func NewWordWrapWriter(w io.Writer, limit uint) io.Writer {
	return &wordWrapWriter{
		limit:  limit,
		writer: w,
	}
}

func (w wordWrapWriter) Write(p []byte) (nn int, err error) {
	if w.limit == 0 {
		return w.writer.Write(p)
	}
	original := string(p)
	wrapped := wordwrap.WrapString(original, w.limit)
	return w.writer.Write([]byte(wrapped))
}
