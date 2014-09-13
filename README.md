## Example

```
package main

import (
  "fmt"
  "github.com/tlorens/ansiout"
)


func main() {
	ansiout.ClearScr()
	ansiout.Wait(5)
	ansiout.PrintFile("header.ans")

	ansiout.Color(7,1)
	fmt.Print("Grey on Blue Text")
}
```

Unfortunately, this is a bit dependent on my [ibkey](https://github.com/tlorens/ibgetkey) library.
