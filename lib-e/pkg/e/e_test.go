package e_test

import (
		"testing"
		"github.com/zoidbergwill/example-project/lib-e/pkg/e"
)

func TestCommon(t *testing.T) {
		got := e.E()
		want := "Dependency E! Version 2!"
		if got != want {
				t.Errorf("E() = %s; want %s", got, want)
		}
}
