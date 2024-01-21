// Package cfg handles working with the Gopkg configuration files.
//
// The cfg package contains the ability to parse (unmarshal) and write (marshal)
// gopkg.yaml and gopkg.lock files. These files contains the details about
// projects managed by Gopkg.
//
// To convert yaml into a cfg.Config instance use the cfg.ConfigFromYaml function.
// The yaml, typically in a gopkg.yaml file, has the following structure.
//
//	package: github.com/Khulnasoft-lab/gopkg
//	homepage: https://khulnasoft-lab.github.io/gopkg
//	license: MIT
//	owners:
//	- name: Matt Butcher
//	  email: technosophos@gmail.com
//	  homepage: http://technosophos.com
//	- name: Matt Farina
//	  email: matt@mattfarina.com
//	  homepage: https://www.mattfarina.com
//	ignore:
//	- appengine
//	excludeDirs:
//	- node_modules
//	import:
//	- package: gopkg.in/yaml.v2
//	- package: github.com/Khulnasoft-lab/vcs
//	  version: ^1.2.0
//	  repo:    git@github.com:Khulnasoft-lab/vcs
//	  vcs:     git
//	- package: github.com/codegangsta/cli
//	- package: github.com/Khulnasoft-lab/semver
//	  version: ^1.0.0
//
// These elements are:
//
//   - package: The top level package is the location in the GOPATH. This is used
//     for things such as making sure an import isn't also importing the top level
//     package.
//   - homepage: To find the place where you can find details about the package or
//     applications. For example, http://k8s.io
//   - license: The license is either an SPDX license string or the filepath to the
//     license. This allows automation and consumers to easily identify the license.
//   - owners: The owners is a list of one or more owners for the project. This
//     can be a person or organization and is useful for things like notifying the
//     owners of a security issue without filing a public bug.
//   - ignore: A list of packages for Gopkg to ignore importing. These are package
//     names to ignore rather than directories.
//   - excludeDirs: A list of directories in the local codebase to exclude from
//     scanning for dependencies.
//   - import: A list of packages to import. Each package can include:
//   - package: The name of the package to import and the only non-optional item.
//   - version: A semantic version, semantic version range, branch, tag, or
//     commit id to use.
//   - repo: If the package name isn't the repo location or this is a private
//     repository it can go here. The package will be checked out from the
//     repo and put where the package name specifies. This allows using forks.
//   - vcs: A VCS to use such as git, hg, bzr, or svn. This is only needed
//     when the type cannot be detected from the name. For example, a repo
//     ending in .git or on GitHub can be detected to be Git. For a repo on
//     Bitbucket we can contact the API to discover the type.
//   - testImport: A list of development packages not already listed under import.
//     Each package has the same details as those listed under import.
package cfg
