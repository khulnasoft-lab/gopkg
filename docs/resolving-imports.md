# Resolving Imports

Gopkg scans an applications codebase to discover the projects to manage in the `vendor/` directory. This happens in a few different ways. Knowing how this works can help you understand what Gopkg is doing.

## At Initialization

When you run `gopkg create` or `gopkg init` to create a `gopkg.yaml` file for a codebase Gopkg will scan your codebase to identify the imports. It does this by walking the filesystem to identify packages. In each package it reads the imports within the Go files.

From this it will attempt to figure out the external packages. External packages are grouped by the root version control system repo with their sub-packages listed underneath. Figuring out the root version control repo compared with the packages underneath it follows the same rules for the `go` tool.

1. GitHub, Bitbucket, Launchpad, IBM Jazz, and go.googlesource.com are evaluated with special rules. We know or can talk to an API to learn about these repos.
2. If the package associated with the repo ends in `.git`, `.hg`, `.bzr`, or `.svn` this is used to determine the root and the type of version control system.
3. If the rules don't provide an answer a `go get` request occurs to try and lookup the information.

Again, this is the same way `go` tries to determine an external location when you use `go get`.

If the project has dependency configuration stored in a Godep, GPM, Gom, or GB file that information will be used to populate the version within the `gopkg.yaml` file.

## At Update

When `gopkg update`, `gopkg up`, `gopkg get`, or `gopkg install` (when no `gopkg.lock` is present) Gopkg will attempt to discover the complete dependency tree. That is all dependencies including dependencies of dependencies of dependencies.

### The Default Option

The default method is to walk the referenced import tree. The resolver starts by scanning the local application to get a list of imports. Then it looks at the specific package imports, scans the imported package for imports, and repeats the lookup cycle until the complete tree has been fetched.

That means that only imports referenced in the source are fetched.

When a version control repo is fetched it does fetch the complete repo. But, it doesn't scan all the packages in the repo for dependencies. Instead, only the packages referenced in the tree are scanned with the imports being followed.

Along the way configuration stored in Gopkg, Godep, GPM, Gom, and GB files are used to work out the version to set and fetched repos to. The first version found while walking the import tree wins.

### All Possible Dependencies

Using the `--all-dependencies` flag on `gopkg update` will change the behavior of the scan. Instead of walking the import tree it walks the filesystem and fetches all possible packages referenced everywhere. This downloads all packages in the tree. Even those not referenced in an applications source or in support of the applications imports.

As in other cases, Gopkg, Godep, GPM, Gom, and GB files are used to set the version of the fetched repo.
