Go MergeMaps - Merge two maps
-----------------------------
This [Go (Golang)](http://golang.org/) package provides functions to merge two maps using reflection.

[![GoDoc](https://godoc.org/github.com/VojtechVitek/mergemaps?status.png)](https://godoc.org/github.com/VojtechVitek/mergemaps)
[![Travis](https://travis-ci.org/VojtechVitek/mergemaps.svg?branch=master)](https://travis-ci.org/VojtechVitek/mergemaps)

Example
-------

```
package main

import (
    "github.com/VojtechVitek/mergemaps"
)

func main() {
    var m map[string]int

    m1 := map[string]int{"foo": 0}
    m2 := map[string]int{"bar": 1, "baz": 2}

    mergemaps.Merge(m, m1, 0)
    mergemaps.Merge(m, m2, 0)
}
```

License
-------
Go MergeMaps is licensed under the [Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0).
