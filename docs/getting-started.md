# Getting Started With Gopkg

This is a quick start guide to using Gopkg once you have it installed.

## Initially Detecting Project Dependencies

Gopkg can detect the dependencies in use on a project and create an initial `gopkg.yaml` file for you. This detection can import the configuration from Godep, GPM, Gom, and GB. To do this change into the top level directory for the project and run:

    $ gopkg init

When this is complete you'll have a `gopkg.yaml` file populated with the projects being used. You can open up this file and even edit it to add information such as versions.

Running `gopkg init` will also ask if you would like to use a wizard to discover information about your dependencies versions and use versions or ranges. Each decision is interactive and your choice.

## Updating Dependencies

To fetch the dependencies and set them to any versions specified in the `gopkg.yaml` file use the following command:

    $ gopkg up

The `up` is short for `update`. This will fetch any dependencies specified in the `gopkg.yaml` file, walk the dependency tree to make sure any dependencies of the dependencies are fetched, and set them to the proper version. While walking the tree it will make sure versions are set and configuration from Godep, GPM, Gom, and GB is imported.

The fetched dependencies are all placed in the `vendor/` folder at the root of the project. The `go` toolchain will use the dependencies here prior to looking in the `GOPATH` or `GOROOT` if you are using Go 1.6+ or Go 1.5 with the Go 1.5 Vendor Experiment enabled.

Gopkg will then create a `gopkg.lock` file. This file contains the entire dependency tree pinned to specific commit ids. This file, as we'll see in a moment, can be used to recreate the exact dependency tree and versions used.

If you want to remove nested `vendor/` directories from within dependencies use the `--strip-vendor` or `-v` flag.

### Dependency Flattening

All of the dependencies Gopkg fetches are into the top level `vendor/` folder for a project. Go provides the ability for each package to have a `vendor/` folder. Gopkg only uses a top level folder for two reasons:

1. Each import location will be compiled into the binary. If the same dependency is imported into three `vendor/` folders it will be in the compiled binary three times. This can quickly lead to binary bloat.
2. Instances of types created in a dependency in one `vendor/` folder are not compatible with the same dependency in other locations. Even if they are the same version. Go sees them as different packages because they are in different locations. This is a problem for database drivers, loggers, and many other things. If you [try to pass an instance created from one location of a package to another you'll encounter errors](https://github.com/mattfarina/golang-broken-vendor).

If a dependency has a `vendor/` directory of its own Gopkg does not remove it by default. The resolution in the `go` toolchain will use these nested versions if they are present. To remove them use the `--strip-vendor` or `-v` flag on the `up` or `install` commands.

## Installing Dependencies

If you want to install the dependencies needed by a project use the `install` command like so:

    $ gopkg install

This command does one of two things:

* If a `gopkg.lock` file is present it retrieves, if missing from the `vendor/` folder, the dependency and sets it to the exact version in the `gopkg.lock` file. The dependencies are fetched and versions set concurrently so this operation is fairly quick.
* If there is no `gopkg.lock` file then an `update` will be performed.

If you're not managing the dependency versions for a project but need to install the dependencies you should use the `install` command.

## Adding More Dependencies

Gopkg can help you add more dependencies to the `gopkg.yaml` file with the `get` command.

    $ gopkg get github.com/Khulnasoft-lab/semver

The `get` command is similar to `go get` but instead fetches dependencies into the `vendor/` folder and adds them to the `gopkg.yaml` file. This command can take one or more dependencies to fetch.

The `get` command can also work with versions.

    $ gopkg get github.com/Khulnasoft-lab/semver#~1.2.0

The `#` is used as a separator between the dependency name and a version to use. The version can be a semantic version, version range, branch, tag, or commit id.

If no version or range is specified and the dependency uses Semantic Versions Gopkg will prompt you to ask if you want to use them.
