package main

import (
		"fmt"
		"github.com/zoidbergwill/example-project/lib-common/pkg/common"
		"github.com/zoidbergwill/example-project/lib-d/pkg/d"
		"github.com/zoidbergwill/example-project/lib-e/pkg/e"
		"github.com/zoidbergwill/example-project/lib-f/pkg/f"
)

func Run() {
		fmt.Println("App A running!")
		common.Common()
		d.D();
		e.E();
		f.F();

}

func main() {
		Run()
}

