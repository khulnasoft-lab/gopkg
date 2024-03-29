# The gopkg.lock File

Where a [`gopkg.yaml`](gopkg.yaml.md) file contains the dependencies, versions (including ranges), and other configuration for the local codebase, the related `gopkg.lock` file contains the complete dependency tree and the revision (commit id) in use.

Knowing the complete dependency tree is useful for Gopkg. For example, when the complete tree is known the `gopkg install` command can install and set the proper revision for multiple dependencies concurrently. This is a fast operation to reproducibly install the dependencies.

The lock file also provides a record of the complete tree, beyond the needs of your codebase, and the revisions used. This is useful for things like audits or detecting what changed in a dependency tree when troubleshooting a problem.

The details of this file are not included here as this file should not be edited by hand. If you know how to read the [`gopkg.yaml`](gopkg.yaml.md) file you'll be able to generally understand the `gopkg.lock` file.
