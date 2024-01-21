# Commands

The following are the Glide commands, most of which are to help you manage your workspace.

## gopkg create (aliased to init)

Initialize a new workspace. Among other things, this creates a `gopkg.yaml` file
while attempting to guess the packages and versions to put in it. For example,
if your project is using Godep it will use the versions specified there. Glide
is smart enough to scan your codebase and detect the imports being used whether
they are specified with another package manager or not.

    $ gopkg create
    [INFO]	Generating a YAML configuration file and guessing the dependencies
    [INFO]	Attempting to import from other package managers (use --skip-import to skip)
    [INFO]	Scanning code to look for dependencies
    [INFO]	--> Found reference to github.com/Khulnasoft-lab/semver
    [INFO]	--> Found reference to github.com/Khulnasoft-lab/vcs
    [INFO]	--> Found reference to github.com/codegangsta/cli
    [INFO]	--> Found reference to gopkg.in/yaml.v2
    [INFO]	Writing configuration file (gopkg.yaml)
    [INFO]	Would you like Glide to help you find ways to improve your gopkg.yaml configuration?
    [INFO]	If you want to revisit this step you can use the config-wizard command at any time.
    [INFO]	Yes (Y) or No (N)?
    n
    [INFO]	You can now edit the gopkg.yaml file. Consider:
    [INFO]	--> Using versions and ranges. See https://gopkg.sh/docs/versions/
    [INFO]	--> Adding additional metadata. See https://gopkg.sh/docs/gopkg.yaml/
    [INFO]	--> Running the config-wizard command to improve the versions in your configuration

The `config-wizard`, noted here, can be run here or manually run at a later time.
This wizard helps you figure out versions and ranges you can use for your
dependencies.

### gopkg config-wizard

This runs a wizard that scans your dependencies and retrieves information on them
to offer up suggestions that you can interactively choose. For example, it can
discover if a dependency uses semantic versions and help you choose the version
ranges to use.

## gopkg get [package name]

You can download one or more packages to your `vendor` directory and have it added to your
`gopkg.yaml` file with `gopkg get`.

    $ gopkg get github.com/Khulnasoft-lab/cookoo

When `gopkg get` is used it will introspect the listed package to resolve its dependencies including using Godep, GPM, Gom, and GB config files.

The `gopkg get` command can have a [version or range](versions.md) passed in with the package name. For example,

    $ gopkg get github.com/Khulnasoft-lab/cookoo#^1.2.3

The version is separated from the package name by an anchor (`#`). If no version or range is specified and the dependency uses Semantic Versions Glide will prompt you to ask if you want to use them.

## gopkg update (aliased to up)

Download or update all of the libraries listed in the `gopkg.yaml` file and put
them in the `vendor` directory. It will also recursively walk through the
dependency packages to fetch anything that's needed and read in any configuration.

    $ gopkg up

This will recurse over the packages looking for other projects managed by Glide,
Godep, gb, gom, and GPM. When one is found those packages will be installed as needed.

A `gopkg.lock` file will be created or updated with the dependencies pinned to
specific versions. For example, if in the `gopkg.yaml` file a version was
specified as a range (e.g., `^1.2.3`) it will be set to a specific commit id in
the `gopkg.lock` file. That allows for reproducible installs (see `gopkg install`).

To remove any nested `vendor/` directories from fetched packages see the `-v` flag.

## gopkg install

When you want to install the specific versions from the `gopkg.lock` file use `gopkg install`.

    $ gopkg install

This will read the `gopkg.lock` file, warning you if it's not tied to the `gopkg.yaml` file, and install the commit id specific versions there.

When the `gopkg.lock` file doesn't tie to the `gopkg.yaml` file, such as there being a change, it will provide an warning. Running `gopkg up` will recreate the `gopkg.lock` file when updating the dependency tree.

If no `gopkg.lock` file is present `gopkg install` will perform an `update` and generates a lock file.

To remove any nested `vendor/` directories from fetched packages see the `-v` flag.

## gopkg novendor (aliased to nv)

When you run commands like `go test ./...` it will iterate over all the subdirectories including the `vendor` directory. When you are testing your application you may want to test your application files without running all the tests of your dependencies and their dependencies. This is where the `novendor` command comes in. It lists all of the directories except `vendor`.

    $ go test $(gopkg novendor)

This will run `go test` over all directories of your project except the `vendor` directory.

## gopkg name

When you're scripting with Glide there are occasions where you need to know the name of the package you're working on. `gopkg name` returns the name of the package listed in the `gopkg.yaml` file.

## gopkg list

Glide's `list` command shows an alphabetized list of all the packages that a project imports.

    $ gopkg list
    INSTALLED packages:
    	vendor/github.com/Khulnasoft-lab/cookoo
    	vendor/github.com/Khulnasoft-lab/cookoo/fmt
    	vendor/github.com/Khulnasoft-lab/cookoo/io
    	vendor/github.com/Khulnasoft-lab/cookoo/web
    	vendor/github.com/Khulnasoft-lab/semver
    	vendor/github.com/Khulnasoft-lab/vcs
    	vendor/github.com/codegangsta/cli
    	vendor/gopkg.in/yaml.v2

## gopkg help

Print the gopkg help.

    $ gopkg help

## gopkg --version

Print the version and exit.

    $ gopkg --version
    gopkg version 0.12.0

## gopkg mirror

Mirrors provide the ability to replace a repo location with
another location that's a mirror of the original. This is useful when you want
to have a cache for your continuous integration (CI) system or if you want to
work on a dependency in a local location.

The mirrors are stored in an `mirrors.yaml` file in your `GLIDE_HOME`.

The three commands to manage mirrors are `list`, `set`, and `remove`.

Use `set` in the form:

    gopkg mirror set [original] [replacement]

or

    gopkg mirror set [original] [replacement] --vcs [type]

for example,

    gopkg mirror set https://github.com/example/foo https://git.example.com/example/foo.git

or

    gopkg mirror set https://github.com/example/foo file:///path/to/local/repo --vcs git

Use `remove` in the form:

    gopkg mirror remove [original]

for example,

    gopkg mirror remove https://github.com/example/foo
