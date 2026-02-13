package matcher

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMatcher_MatchLine(t *testing.T) {
	for _, tc := range []struct {
		line          string
		pattern       string
		caseSensetive bool
		result        bool
	}{
		{
			"asdsadas",
			"asd",
			false,
			true,
		},
		{
			"wordword1",
			"d",
			false,
			true,
		},
	} {
		matcher, err := New(tc.pattern, tc.caseSensetive)
		require.NoError(t, err)

		res := matcher.MatchLine(tc.line)
		require.Equal(t, tc.result, res)
	}
}
