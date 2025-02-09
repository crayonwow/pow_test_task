package equix

import (
	"testing"
)

func Test_equix(t *testing.T) {
	t.Run("solve", func(t *testing.T) {
		challenge := []byte("challenge")
		solution := Solve(challenge)
		t.Log(solution)
	})
}
