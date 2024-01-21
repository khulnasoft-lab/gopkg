# Gopkg Plugins

Gopkg supports a simple plugin system similar to Git.

## Existing Plugins

Some plugins exist today for Gopkg including:

* [gopkg-vc](https://github.com/sgotti/gopkg-vc) - The vendor cleaner allows you to strip files not needed for building your application from the `vendor/` directory.
* [gopkg-brew](https://github.com/heewa/gopkg-brew) - Convert Go deps managed by gopkg to Homebrew resources to help you make brew formulas for you Go programs.
* [gopkg-hash](https://github.com/mattfarina/gopkg-hash) - Generates a hash of the `gopkg.yaml` file compatible with Gopkgs internal hash.
* [gopkg-cleanup](https://github.com/ngdinhtoan/gopkg-cleanup) - Removing unused packages from the `gopkg.yaml` file.
* [gopkg-pin](https://github.com/multiplay/gopkg-pin) - Take all dependencies from the `gopkg.lock` and pin them explicitly in the `gopkg.yaml` file.

_Note, to add plugins to this list please create a pull request._

## How Plugins Work

When Gopkg encounters a subcommand that it does not know, it will try to delegate it to another executable according to the following rules.

Example:

```
$ gopkg install # We know this command, so we execute it
$ gopkg foo     # We don't know this command, so we look for a suitable
                # plugin.
```

In the example above, when gopkg receives the command `foo`, which it does not know, it will do the following:

1. Transform the name from `foo` to `gopkg-foo`
2. Look on the system `$PATH` for `gopkg-foo`. If it finds a program by that name, execute it...
3. Or else, look at the current project's root for `gopkg-foo`. (That is, look in the same directory as `gopkg.yaml`). If found, execute it.
4. If no suitable command is found, exit with an error.

## Writing a Gopkg Plugin

A Gopkg plugin can be written in any language you wish, provided that it can be executed from the command line as a subprocess of Gopkg. The example included with Gopkg is a simple Bash script. We could just as easily write Go, Python, Perl, or even Java code (with a wrapper) to
execute.

A Gopkg plugin must be in one of two locations:

1. Somewhere on the PATH
2. In the same directory as `gopkg.yaml`

It is recommended that system-wide Gopkg plugins go in `/usr/local/bin` or `$GOPATH/bin` while project-specific plugins go in the same directory as `gopkg.yaml`.

### Arguments and Flags

Say Gopkg is executed like this:

```
$ gopkg foo -name=Matt myfile.txt
```

Gopkg will interpret this as a request to execute `gopkg-foo` with the arguments `-name=Matt myfile.txt`. It will not attempt to interpret those arguments or modify them in any way.

Hypothetically, if Gopkg had a `-x` flag of its own, you could call this:

```
$ gopkg -x foo -name=Matt myfile.txt
```

In this case, gopkg would interpret and swollow the -x and pass the rest on to `gopkg-foo` as in the example above.

## Example Plugin

File: gopkg-foo

```bash
#!/bin/bash

echo "Hello"
```
