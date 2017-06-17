# gen
easy random data(string/slice/map/int/etc) generator

### install

  * `go get github.com/asm-jaime/gen`

### example (how to use lib)

  * example main.go
  
```go
package main

import(
  "fmt"
  "github.com/asm-jaime/gen"
)

func main() {
  fmt.Println(gen.Str(10))
}
```
