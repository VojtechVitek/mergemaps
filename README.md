Go MergeMaps - Merge two maps
-----------------------------
This [Go (Golang)](http://golang.org/) package provides functions to merge two maps using [reflection](http://golang.org/pkg/reflect/).

[![GoDoc](https://godoc.org/github.com/VojtechVitek/mergemaps?status.png)](https://godoc.org/github.com/VojtechVitek/mergemaps)
[![Travis](https://travis-ci.org/VojtechVitek/mergemaps.svg?branch=master)](https://travis-ci.org/VojtechVitek/mergemaps)

This pkg was merged into [OpenShift 3](https://github.com/openshift/origin) [util pkg](https://github.com/openshift/origin/blob/master/pkg/util/mergemap.go).

Example
-------

See [more examples](examples/examples.go).

```
package main

import (
    map "github.com/VojtechVitek/mergemaps"
)

func main() {
    dst := map[string]string{"foo": "bar"}
    src := map[string]string{"baz": ""}

    // Merge src map into dst map
    err := mergemaps.MergeInto(dst, src, 0 /*flags*/)
    if err != nil {
        fmt.Errorf("Can't merge src into dst: %v", err)
    }

    // Prints map[string]string{"baz":"", "foo":"bar"}
    fmt.Printf("%#v\n", dst)
}
```

License
-------
Go MergeMaps is licensed under the [Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0).
