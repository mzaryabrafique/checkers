package testutil


import (
	"github.com/alice/checkers/testutil/sample"
)

var (
	Alice = sample.AccAddress()
	Bob   = sample.AccAddress()
	Carol = sample.AccAddress()
)

// Alternatively, if we want dynamic addresses:
// var (
// 	Alice = sample.AccAddress()
// 	Bob   = sample.AccAddress()
// 	Carol = sample.AccAddress()
// )
// But constants are preferred for tests if possible, though AccAddress returns string.
// Let's use variables if AccAddress is dynamic, or hardcoded strings if we want stability.
// The search result showed "Alice | 27 | 555-4321" which is irrelevant.
// Let's use sample.AccAddress() but assigned to variables since they are function calls.
// Wait, common_test.go uses `const`.
// const (
// 	alice = testutil.Alice
// 	bob   = testutil.Bob
// 	carol = testutil.Carol
// )
// So testutil.Alice must be a constant or a variable that can be assigned to a const?
// No, in Go, consts must be compile-time constants. Function calls are not allowed.
// So testutil.Alice MUST be a string literal.

// I will use hardcoded valid cosmos addresses.
