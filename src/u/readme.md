run
===

[github.com/dwitp/golang/src/u](https://github.com/dwitp/golang/tree/master/src/u)

Debug code library

## functions

#### **D** ~ Print Type and Value of variable with separator

```go
package main

import "github.com/dwitp/golang/src/u"

func main() {
    u.D("Hello World!")
}
```

output

```bash
user@desktop:~/workspace/src/github.com/dwitp/go-start$ run

-----------------------------------------------------
{string} Hello World!
-----------------------------------------------------
```

#### **S** ~ Print Slice Information

```go
package main

import "github.com/dwitp/golang/src/u"

func main() {
    s := make([]int, 5)
    u.D("s")
    u.S(s)
}
```

output

```bash
user@desktop:~/workspace/src/github.com/dwitp/go-start$ run

-----------------------------------------------------
{string} s
-----------------------------------------------------

len=5 cap=5 [0 0 0 0 0]
```
