<strong>How to use?</strong><br/>
```go
package main

import "github.com/rafaelsouzaribeiro/check-required-fields/pkg/utils"

func main() {
	var m = map[string]interface{}{"campo": "", "testar": 0}
	message := utils.CheckRequired(m, "The field", "is required")

	println(message)
}
```

