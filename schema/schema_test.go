package schema

import (
	"fmt"
	"mrsydar/geve/schema/kind"
)

func ExampleSchema() {

	items := Schema{
		"_id": kind.String{
			Common: kind.Common{
				Required: true,
			},
		},
	}

	fmt.Println(items)

}
