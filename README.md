## About

The control(8) utility is implemented in Go for the [HardenedBSD](https://git.hardenedbsd.org)
operating system, and it provides an alternative implementation
of the official hbsdcontrol(8) utility that is implemented in C.

## Motivation

#### Why?

The primary motivation behind this project was to have a testbed
for the [majortom](https://github.com/0x1eef/majortom#readme) library,
which provides Go bindings for libhbsdcontrol. But it is also a place
to experiment with a different command-line interface that requires
less typing and is hopefully more inuitive to use.

## Examples

#### Usage

    usage: control [-hnv] <command> [feature] <file>

    COMMANDS
        enable           Enable a feature
        disable          Disable a feature
        sysdef           Restore the system default
        [query|status]   Query feature states

    OPTIONS
        -h          Show help
        -n          Set the namespace (either 'user' or 'system')
        -v          Print current version

    EXAMPLES
        control enable mprotect /bin/ls
        control disable pageexec /bin/ls
        control sysdef segvguard /bin/ls
        control query /bin/ls
        control status /bin/ls

    FEATURES
    shlibrandom              segvguard                prohibit_ptrace_capsicum
    pageexec                 mprotect                 insecure_kmod
    harden_shm               disallow_map32bit        aslr

## Sources

* [github.com/@0x1eef](https://github.com/0x1eef/control#readme)
* [gitlab.com/@0x1eef](https://gitlab.com/0x1eef/control#about)
* [hardenedbsd.org/@0x1eef](https://git.HardenedBSD.org/0x1eef/controlm#about)

## License

[BSD Zero Clause](https://choosealicense.com/licenses/0bsd/)
<br>
See [LICENSE](./LICENSE)
