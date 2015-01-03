Mersenne Twister 64bit version in Go
================

An implementation of Takuji Nishimura's and Makoto Matsumoto's Mersenne Twister pseudo random number generator in Go. ([http://www.math.sci.hiroshima-u.ac.jp/~m-mat/MT/emt64.html](http://www.math.sci.hiroshima-u.ac.jp/~m-mat/MT/emt64.html))

## Sample

```go
package main

import (
	"fmt"
	"github.com/disktnk/mt19937"
	"math/rand"
)

func main() {
	rand := rand.New(mt19937.New())
	rand.Seed(20150103)
	fmt.Println(rand.Float64()) // 0.5814099686652303
}
```
