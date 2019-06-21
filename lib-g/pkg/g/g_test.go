package g_test

import (
		"testing"
		"github.com/zoidbergwill/example-project/lib-g/pkg/g"
)

func TestCommon(t *testing.T) {
		got := g.G()
		want := "Dependency G!"
		if got != want {
				t.Errorf("G() = %s; want %s", got, want)
		}
}
