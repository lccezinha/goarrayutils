A bunch of methods to make array verifications in golang easier, inspired in Ruby array methods.

### Methods

  - [Include](#include)
  - [Any](#any)
  - [None](#none)
  - [Collect](#collect)
  - [Compact](#compact)

### Examples of Use

#### Include

```go
package main

import "fmt"
import "github.com/lccezinha/goarrayutils"

func main() {
  var arr []int
  arr = append(arr, 1, 2, 3)
  expected := 1

  result := arrayutils.Include(arr, expected)

  fmt.Println(result)
}

// true
```

#### Any

```go
package main

import "fmt"
import "github.com/lccezinha/goarrayutils"

func main() {
  var arr []int
  arr = append(arr, 1, 2, 3)

  condition := func(item interface{}) bool {
    return item.(int) >= 2
  }

  result := arrayutils.Any(arr, condition)

  fmt.Println(result)
}

// true
```

#### None

```go
package main

import "fmt"
import "github.com/lccezinha/goarrayutils"

func main() {
  var arr []int
  arr = append(arr, 1, 2, 3)

  condition := func(item interface{}) bool {
    return item.(int) >= 5
  }

  result := arrayutils.None(arr, condition)

  fmt.Println(result)
}

// true
```

#### Collect

```go
package main

import "fmt"
import "github.com/lccezinha/goarrayutils"

func main() {
  var arr []interface{}
  arr = append(arr, 1, 2, 3)

  condition := func(item interface{}) bool {
    return item.(int) >= 2
  }

  result := arrayutils.Collect(arr, condition)

  fmt.Println(result)
}

// [2, 3]
```

#### Compact

```go
package main

import "fmt"
import "github.com/lccezinha/goarrayutils"

func main() {
  var arr []interface{}
  arr = append(arr, 1, 2, 3, nil, 4, nil, 5)

  result := arrayutils.Compact(arr)

  fmt.Println(result)
}

// [1, 2, 3, 4, 5]
```