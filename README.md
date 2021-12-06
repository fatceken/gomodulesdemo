# gomodulesdemo

## Usage

```
package main

import (
	"fmt"
	"github.com/fatceken/gomodulesdemo"
)

func main() {

	nums := []interface{}{"2", "3", "4"}

	parallel.ForEach(nums, func(val interface{}) {
		fmt.Println(val)
	})
}
```
