## About

The control package provides Go bindings for the libhbsdcontrol
library for the [HardenedBSD](https://git.hardenedbsd.org/hardenedbsd/hardenedbsd)
operating system. The library provides an interface that can enable, disable,
restore and query feature states for a given file. See
[hbcontrol](https://git.hardenedbsd.org/0x1eef/hbcontrol)
as an example of a command line application that uses this library.

## Examples

#### Features

The following example demonstrates how to create an instance of
`control.Context` and then how to query all feature names:

```go
package main

import (
	"fmt"

	"git.hardenedbsd.org/0x1eef/control"
)

func main() {
	ns := control.Namespace("system")
	ctx, err := control.NewContext(ns)
	if err != nil {
		panic(err)
	}
	defer ctx.Free()
	features, err := ctx.FeatureNames()
	if err != nil {
		panic(err)
	}
	for _, name := range features {
		fmt.Printf("feature: %s\n", name)
	}
}
```

#### Settings

The next example shows how to enable, disable, and restore the system default
settings for a given file and feature:

```go
package main

import (
	"fmt"

	"git.hardenedbsd.org/0x1eef/control"
)

func main() {
	ns := control.Namespace("system")
	ctx, err := control.NewContext(ns)
	if err != nil {
		panic(err)
	}
	defer ctx.Free()
	feature, target := "mprotect", "/usr/bin/mdo"
	if err := ctx.Enable(feature, target); err != nil {
		panic(err)
	}
	if err := ctx.Disable(feature, target); err != nil {
		panic(err)
	}
	if err := ctx.Sysdef(feature, target); err != nil {
		panic(err)
	}
}
```

#### Status

The last example demonstrates how to query the status of a feature
for a given file:

```go
package main

import (
	"fmt"

	"git.hardenedbsd.org/0x1eef/control"
)

func main() {
	ns := control.Namespace("system")
	ctx, err := control.NewContext(ns)
	if err != nil {
		panic(err)
	}
	defer ctx.Free()
	feature, target := "mprotect", "/usr/bin/mdo"
	status, err := ctx.Status(feature, target)
	if err != nil {
		panic(err)
	}
	fmt.Printf("The %s feature is %s\n", feature, status)
}
```

#### Concurrency

The control package expects that each instance of `control.Context`
is not shared across goroutines, otherwise the behavior is undefined
and it could lead to program crashes. In other words, create one context
per goroutine. The following example spawns three goroutines that
correctly create one context per goroutine, and most important,
the context is not shared between goroutines:

```go
package main

import (
	"fmt"

	"git.hardenedbsd.org/0x1eef/control"
)

func worker() {
	ns := control.Namespace("system")
	ctx, err := control.NewContext(ns)
	if err != nil {
		panic(err)
	}
	defer ctx.Free()
}

func main() {
	for i := 0; i < 3; i++ {
		go worker()
	}
}
```

## Install

    go get git.hardenedbsd.org/0x1eef/control

## Sources

* [github.com/@0x1eef](https://git.hardenedbsd.org/0x1eef/control#readme)
* [gitlab.com/@0x1eef](https://gitlab.com/0x1eef/control#about)
* [hardenedbsd.org/@0x1eef](https://git.HardenedBSD.org/0x1eef/control#about)

## License

[BSD Zero Clause](https://choosealicense.com/licenses/0bsd/)
<br>
See [LICENSE](./LICENSE)
