# Gopkg: Vendor Package Management for Go

[Gopkg](https://gopkg.sh) is a package manager for [Go](https://golang.org) that is conceptually similar to package managers for other languages such as Cargo for Rust, NPM for Node.js, Pip for Python, Bundler for Ruby, and so forth.

Gopkg provides the following functionality:

* Records dependency information in a `gopkg.yaml` file. This includes a name, version or version range, version control information for private repos or when the type cannot be detected, and more.
* Tracks the specific revision each package is locked to in a `gopkg.lock` file. This enables reproducibly fetching the dependency tree.
* Works with Semantic Versions and Semantic Version ranges.
* Supports Git, Bzr, HG, and SVN. These are the same version control systems supported by `go get`.
* Utilizes `vendor/` directories, known as the Vendor Experiment, so that different projects can have differing versions of the same dependencies.
* Allows for aliasing packages which is useful for working with forks.
* Import configuration from Godep, GPM, Gom, and GB.

## Installing Gopkg

There are a few ways to install Gopkg.

1. Use the shell script to try an automatically install it. `curl https://gopkg.sh/get | sh`
2. Download a [versioned release](https://github.com/Khulnasoft-lab/gopkg/releases). Gopkg releases are semantically versioned.
3. Use a system package manager to install Gopkg. For example, using `brew install gopkg` can be used if you're using [Homebrew](http://brew.sh) on Mac.
4. The latest development snapshot can be installed with `go get`. For example, `go get -u github.com/Khulnasoft-lab/gopkg`. This is not a release version.
