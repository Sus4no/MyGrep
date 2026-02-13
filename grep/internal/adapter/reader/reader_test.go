package reader

import (
	"bytes"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReader_ReadLine(t *testing.T) {
	var buf bytes.Buffer
	reader := New(&buf)

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
		for _, line := range tc.text {
			_, err := buf.WriteString(line + "\n")
			require.NoError(t, err)
		}

		builder := strings.Builder{}
		for {
			line, err := reader.ReadLine()
			if err == io.EOF {
				break
			}
			require.NoError(t, err)

			builder.WriteString(line)
			builder.WriteByte('\n')
		}
		require.Equal(t, strings.Join(tc.text, "\n")+"\n", builder.String())

		buf.Reset()
	}
}
