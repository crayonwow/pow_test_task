package hashchash

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_equix(t *testing.T) {
	t.Run("solve", func(t *testing.T) {
		p := New(newInMemmoryStorage(), Confing{
			Expiration: 1000,
		})
		challenge := []byte("challenge_string")
		solution, err := p.Solve(context.Background(), challenge, 20)
		require.NoError(t, err)
		_ = solution

		verified, err := p.Verify(solution, 20)
		require.NoError(t, err)
		require.True(t, verified)

		verified, err = p.Verify(solution, 20)
		require.Error(t, err)

		t.Log(string(solution))
	})
}
