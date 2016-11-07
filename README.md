[![Build Status](https://travis-ci.org/lccezinha/gosliceutils.svg?branch=master)](https://travis-ci.org/lccezinha/gosliceutils)

Some methods to make slice verifications in Golang easier.

### Methods

  - [Include](#include)
  - [Any](#any)
  - [None](#none)
  - [Collect](#collect)
  - [Compact](#compact)
  - [All](#all)
  - [Reject](#reject)

### Examples of Use

#### Include

```go
package main

import "fmt"
import "github.com/lccezinha/sliceutils"

func main() {
  var slice []int
  slice = append(slice, 1, 2, 3)
  expected := 1

  result := sliceutils.Include(slice, expected)

  fmt.Println(result)
}

// true
```

#### Any

```go
package main

import "fmt"
import "github.com/lccezinha/gosliceutils"

func main() {
  var slice []int
  slice = append(slice, 1, 2, 3)

  condition := func(item interface{}) bool {
    return item.(int) >= 2
  }

  result := sliceutils.Any(slice, condition)

  fmt.Println(result)
}

// true
```

#### None

```go
package main

import "fmt"
import "github.com/lccezinha/gosliceutils"

func main() {
  var slice []int
  slice = append(slice, 1, 2, 3)

  condition := func(item interface{}) bool {
    return item.(int) >= 5
  }

  result := sliceutils.None(slice, condition)

  fmt.Println(result)
}

// true
```

#### Collect

```go
package main

import "fmt"
import "github.com/lccezinha/gosliceutils"

func main() {
  var slice []interface{}
  slice = append(slice, 1, 2, 3)

  condition := func(item interface{}) bool {
    return item.(int) >= 2
  }

  result := sliceutils.Collect(slice, condition)

  fmt.Println(result)
}

// [2, 3]
```

#### Compact

```go
package main

import "fmt"
import "github.com/lccezinha/gosliceutils"

func main() {
  var slice []interface{}
  slice = append(slice, 1, 2, 3, nil, 4, nil, 5)

  result := sliceutils.Compact(slice)

  fmt.Println(result)
}

// [1, 2, 3, 4, 5]
```

#### All

```go
package main

import "fmt"
import "github.com/lccezinha/gosliceutils"

func main() {
  var slice []interface{}
  slice = append(slice, 1, 2, 3)

  condition := func(item interface{}) bool {
    return item.(int) > 0
  }

  result := sliceutils.All(slice, condition)

  fmt.Println(result)
}

// true
```

#### Reject

```go
package main

import "fmt"
import "github.com/lccezinha/gosliceutils"

func main() {
  var slice []interface{}
  slice = append(slice, 1, 2, 3)

  condition := func(item interface{}) bool {
    return item.(int) > 1
  }

  result := sliceutils.Reject(slice, condition)

  fmt.Println(result)
}

// [1]
```
