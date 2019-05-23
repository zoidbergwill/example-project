package d_test

import (
		"testing"
		"github.com/zoidbergwill/example-project/lib-d/pkg/d"
)

func TestCommon(t *testing.T) {
		got := d.D()
		want := "Dependency D!"
		if got != want {
				t.Errorf("D() = %s; want %s", got, want)
		}
}
