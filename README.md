## About

The control(8) utility provides an alternative implementation of
the official hbsdcontrol(8) utility. The alternative is implemented
in Go.

## Examples

#### Usage

    usage: control [-hn] <command> <feature> <file>

    COMMANDS
        enable      Enable a feature
        disable     Disable a feature
        restore     Restore a feature to the system default

    OPTIONS
        -h  Show help
        -n  Set the namespace (either user or system)

    EXAMPLES
        control enable mprotect /usr/bin/ls
        control disable pageexec /usr/bin/ls
        control restore segvguard /usr/bin/ls

    FEATURES
    shlibrandom                     segvguard                       prohibit_ptrace_capsicum
    pageexec                        mprotect                        insecure_kmod
    harden_shm                      disallow_map32bit               aslr

## Sources

* [github.com/@0x1eef](https://github.com/0x1eef/control#readme)
* [gitlab.com/@0x1eef](https://gitlab.com/0x1eef/control#about)
* [hardenedbsd.org/@0x1eef](https://git.HardenedBSD.org/0x1eef/controlm#about)

## License

[BSD Zero Clause](https://choosealicense.com/licenses/0bsd/)
<br>
See [LICENSE](./LICENSE)
