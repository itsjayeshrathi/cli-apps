package main

import (
	"strings"
	"testing"
)

func TestCount(t *testing.T) {
	testCases := []struct {
		name       string
		input      string
		countLines bool
		countBytes bool
		want1      int
		want2      int
	}{
		{
			name:       "Count lines only",
			input:      "line1\nline2\nline3\n",
			countLines: true,
			countBytes: false,
			want1:      3,
			want2:      0,
		},
		{
			name:       "Count words only",
			input:      "hello world\nfoo bar",
			countLines: false,
			countBytes: false,
			want1:      4,
			want2:      0,
		},
		{
			name:       "Count lines and bytes",
			input:      "hi\nthere\n",
			countLines: true,
			countBytes: true,
			want1:      2,
			want2:      9,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res1, res2 := count(strings.NewReader(tc.input), tc.countLines, tc.countBytes)

			if res1 != tc.want1 || res2 != tc.want2 {
				t.Errorf("got (%d, %d), want (%d, %d)", res1, res2, tc.want1, tc.want2)
			}
		})
	}
}
