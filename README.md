## About

The hbcontrol package provides Go bindings for the libhbsdcontrol
library for the [HardenedBSD](https://git.hardenedbsd.org/hardenedbsd/hardenedbsd)
operating system. The library provides an interface that can enable, disable,
restore and query feature states for a given file. This repository also
includes a commmand line utility by the name of
[hbcontrol](https://git.hardenedbsd.org/0x1eef/hbcontrol),
that provides an alternative implementation of hbsdcontrol(8).

## Examples

### CLI

    usage: hbcontrol [-hHnv] <command> [feature] <file>

    COMMANDS
        enable           Enable a feature
        disable          Disable a feature
        sysdef           Restore the system default
        [query|status]   Query feature states

    OPTIONS
        -h          Show help
        -H          Follow symlinks
        -n          Set the namespace (either 'user' or 'system')
        -v          Print current version

    EXAMPLES
        hbcontrol enable mprotect /bin/ls
        hbcontrol disable pageexec /bin/ls
        hbcontrol sysdef segvguard /bin/ls
        hbcontrol query /bin/ls
        hbcontrol status /bin/ls

    FEATURES
    shlibrandom              segvguard                prohibit_ptrace_capsicum
    pageexec                 mprotect                 insecure_kmod
    harden_shm               disallow_map32bit        aslr

### Package

#### Features

The following example queries feature names available in the "system"
namespace:

```go
package main

import (
	"fmt"

	"git.hardenedbsd.org/0x1eef/control"
)

func main() {
	ns := hbcontrol.Namespace("system")
	ctx, err := hbcontrol.NewContext(ns)
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

The following example enables, disables, and restores the system default
settings for a given file and feature:

```go
package main

import (
	"fmt"

	"git.hardenedbsd.org/0x1eef/control"
)

func main() {
	ns := hbcontrol.Namespace("system")
	ctx, err := hbcontrol.NewContext(ns)
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
	ns := hbcontrol.Namespace("system")
	ctx, err := hbcontrol.NewContext(ns)
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

The control package expects that each instance of `hbcontrol.Context`
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
	ns := hbcontrol.Namespace("system")
	ctx, err := hbcontrol.NewContext(ns)
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
