// Copyright (c) 2020 Proton Technologies AG
//
// This file is part of ProtonMail Bridge.
//
// ProtonMail Bridge is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// ProtonMail Bridge is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with ProtonMail Bridge.  If not, see <https://www.gnu.org/licenses/>.

package address

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParseDateTime(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: `Fri, 21 Nov 1997 09:55:06`,
			want:  `1997-11-21T09:55:06Z`,
		},
		{
			input: `Fri, 21 Nov 1997 09:55:06 -0600`,
			want:  `1997-11-21T09:55:06-06:00`,
		},
		{
			input: `Tue, 1 Jul 2003 10:52:37 +0200`,
			want:  `2003-07-01T10:52:37+02:00`,
		},
		{
			input: `Thu, 13 Feb 1969 23:32:54 -0330`,
			want:  `1969-02-13T23:32:54-03:30`,
		},
		{
			input: "Thu, 13 Feb 1969 23:32 -0330 (Newfoundland Time)",
			want:  `1969-02-13T23:32:00-03:30`,
		},
	}
	for _, test := range tests {
		test := test

		t.Run(test.input, func(t *testing.T) {
			got, err := ParseDateTime(test.input)
			assert.NoError(t, err)
			assert.Equal(t, test.want, got.Format(time.RFC3339))
		})
	}
}

func TestParseDateTimeObsolete(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: `21 Nov 97 09:55:06 GMT`,
			want:  `1997-11-21T09:55:06Z`,
		},
	}
	for _, test := range tests {
		test := test

		t.Run(test.input, func(t *testing.T) {
			got, err := ParseDateTime(test.input)
			assert.NoError(t, err)
			assert.Equal(t, test.want, got.Format(time.RFC3339))
		})
	}
}