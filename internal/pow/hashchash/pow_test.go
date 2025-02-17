package hashchash

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_hashcash(t *testing.T) {
	t.Run("solve", func(t *testing.T) {
		p := New(NewInMemmoryStorage(), Confing{
			Expiration: 1000,
		})
		difficulty := int64(10)
		challenge := []byte("challenge_string")
		solution, err := p.Solve(context.Background(), challenge, difficulty)
		require.NoError(t, err)
		_ = solution

		verified, err := p.Verify(solution, difficulty)
		require.NoError(t, err)
		require.True(t, verified)

		verified, err = p.Verify(solution, difficulty)
		require.Error(t, err)
	})
}
