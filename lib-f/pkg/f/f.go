package f

import (
		"github.com/zoidbergwill/example-project/lib-g/pkg/g"
)

func F() string {
		return "Dependency F! " + g.G()
}
