package common_test

import (
		"testing"
		"github.com/zoidbergwill/example-project/lib-common/pkg/common"
)

func TestCommon(t *testing.T) {
		got := common.Common()
		want := "Common function!"
		if got != want {
				t.Errorf("Common() = %s; want %s", got, want)
		}
}
