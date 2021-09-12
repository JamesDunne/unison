// Copyright ©2021 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package unison

import (
	"github.com/richardwilkes/toolbox/xmath/geom32"
)

// Separator provides a simple vertical or horizontal separator line.
type Separator struct {
	Panel
	Ink      Ink
	Vertical bool
}

// NewSeparator creates a new separator line.
func NewSeparator() *Separator {
	s := &Separator{}
	s.Self = s
	s.SetSizer(s.DefaultSizes)
	s.DrawCallback = s.DefaultDraw
	return s
}

// DefaultSizes provides the default sizing.
func (s *Separator) DefaultSizes(hint geom32.Size) (min, pref, max geom32.Size) {
	if s.Vertical {
		if hint.Height < 1 {
			pref.Height = 1
		} else {
			pref.Height = hint.Height
		}
		min.Height = 1
		max.Height = DefaultMaxSize
		min.Width = 1
		pref.Width = 1
		max.Width = 1
	} else {
		if hint.Width < 1 {
			pref.Width = 1
		} else {
			pref.Width = hint.Width
		}
		min.Width = 1
		max.Width = DefaultMaxSize
		min.Height = 1
		pref.Height = 1
		max.Height = 1
	}
	if border := s.Border(); border != nil {
		insets := border.Insets()
		min.AddInsets(insets)
		pref.AddInsets(insets)
		max.AddInsets(insets)
	}
	return min, pref, max
}

// DefaultDraw provides the default drawing.
func (s *Separator) DefaultDraw(canvas *Canvas, dirty geom32.Rect) {
	rect := s.ContentRect(false)
	if s.Vertical {
		if rect.Width > 1 {
			rect.X += (rect.Width - 1) / 2
			rect.Width = 1
		}
	} else if rect.Height > 1 {
		rect.Y += (rect.Height - 1) / 2
		rect.Height = 1
	}
	canvas.DrawRect(rect, ChooseInk(s.Ink, DividerColor).Paint(canvas, rect, Fill))
}
