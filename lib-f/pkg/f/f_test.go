package f_test

import (
		"testing"
		"github.com/zoidbergwill/example-project/lib-f/pkg/f"
)

func TestCommon(t *testing.T) {
		got := f.F()
		want := "Dependency F!"
		if got != want {
				t.Errorf("F() = %s; want %s", got, want)
		}
}
