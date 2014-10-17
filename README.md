Go - Merge maps
---------------
This [Go](http://golang.org/) package provides functions to merge two maps.

[![GoDoc](https://godoc.org/github.com/VojtechVitek/go-merge-maps?status.png)](https://godoc.org/github.com/VojtechVitek/go-merge-maps)
[![Travis](https://travis-ci.org/VojtechVitek/go-merge-maps.svg?branch=master)](https://travis-ci.org/VojtechVitek/go-merge-maps)

Example
-------

```
package main

import (
    "github.com/VojtechVitek/go-merge-maps"
)

func main() {
    m := map[string]int{}
    m1 := map[string]int{"foo": 0}
    m2 := map[string]int{"bar": 1, "baz": 2}

    Merge(m, m1, 0)
    Merge(m, m2, 0)
}
```

License
-------
Go - Merge maps is licensed under the [Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0).
