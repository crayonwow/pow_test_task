package hashchash

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_equix(t *testing.T) {
	t.Run("solve", func(t *testing.T) {
		// defer func(start time.Time) {
		// 	t.Log(time.Since(start))
		// }(time.Now())
		p := New(newInMemmoryStorage(), Confing{
			Expiration: 1000,
			Difficulty: 20,
		})
		challenge := []byte("challenge_string")
		solution, err := p.Solve(context.Background(), challenge)
		require.NoError(t, err)
		_ = solution

		verified, err := p.Verify(solution)
		require.NoError(t, err)
		require.True(t, verified)

		verified, err = p.Verify(solution)
		require.Error(t, err)

		t.Log(string(solution))
	})
}
