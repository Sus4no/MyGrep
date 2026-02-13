package writer

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWriter_WriteLine(t *testing.T) {
	var buf bytes.Buffer
	writer := New(&buf)

	for _, tc := range []struct {
		text []string
	}{
		{
			[]string{
				"asdads",
				"nijjsnd",
			},
		},
		{
			[]string{
				"some txt",
			},
		},
	} {
		// пишем в буфер
		for _, line := range tc.text {
			err := writer.WriteLine(line)
			require.NoError(t, err)
		}

		require.Empty(t, buf.String())

		err := writer.Flush()
		require.NoError(t, err)
		require.Equal(t, strings.Join(tc.text, "\n")+"\n", buf.String())

		buf.Reset()
	}
}
