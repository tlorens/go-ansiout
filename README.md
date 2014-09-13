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

Unfortunately, this is a bit dependent on my [go-ibgetkey](https://github.com/tlorens/go-ibgetkey) library.


Screenshot of what my intent was.

![BBS Menu](https://dl.dropboxusercontent.com/u/2934311/stupid-bbs.png)