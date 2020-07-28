// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package present

import (
	"fmt"
	"strings"
)

func init() {
	Register("richmd", parseRichmd)
}

type Richmd struct {
	Cmd        string // original command from present source
	URL        string
}

func (v Richmd) PresentCmd() string   { return v.Cmd }
func (v Richmd) TemplateName() string { return "richmd" }

func parseRichmd(ctx *Context, fileName string, lineno int, text string) (Elem, error) {
	args := strings.Fields(text)
	if len(args) < 2 {
		return nil, fmt.Errorf("incorrect Richmd invocation: %q", text)
	}
	vid := Richmd{Cmd: text, URL: args[1]}
	return vid, nil
}
