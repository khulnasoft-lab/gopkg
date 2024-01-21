// Package repo provides tools for working with VCS repositories.
//
// Gopkg manages repositories in the vendor directory by using the native VCS
// systems of each repository upon which the code relies.
package repo

// concurrentWorkers is the number of workers to be used in concurrent operations.
var concurrentWorkers = 20

// UpdatingVendored indicates whether this run of Gopkg is updating a vendored vendor/ path.
//
// It is related to the --update-vendor flag for update and install.
//
// TODO: This is legacy, and maybe we should handle it differently. It should
// be set either 0 or 1 times, and only at startup.
//var UpdatingVendored bool = false
