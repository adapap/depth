package main

import (
	"fmt"
	"testing"

	"github.com/adapap/depth"
	"github.com/stretchr/testify/assert"
)

func Test_parse(t *testing.T) {
	tests := []struct {
		internal bool
		test     bool
		depth    int
		json     bool
		explain  string
	}{
		{true, true, 0, true, ""},
		{false, false, 10, false, ""},
		{true, false, 10, false, ""},
		{false, true, 5, true, ""},
		{false, true, 5, true, "github.com/adapap/depth"},
		{false, true, 5, true, ""},
	}

	for _, tc := range tests {
		tr, options := parse([]string{
			fmt.Sprintf("-internal=%v", tc.internal),
			fmt.Sprintf("-test=%v", tc.test),
			fmt.Sprintf("-max=%v", tc.depth),
			fmt.Sprintf("-json=%v", tc.json),
			fmt.Sprintf("-explain=%v", tc.explain),
		})

		assert.Equal(t, tc.internal, tr.ResolveInternal)
		assert.Equal(t, tc.test, tr.ResolveTest)
		assert.Equal(t, tc.depth, tr.MaxDepth)
		assert.Equal(t, tc.json, options.OutputJSON)
		assert.Equal(t, tc.explain, options.ExplainPkg)
	}
}

func Example_handlePkgsStrings() {
	var tree depth.Tree

	_ = handlePkgs(&tree, &depth.Options{PackageNames: []string{"strings"}})
	// Output:
	// strings
	//   ├ errors
	//   ├ internal/abi
	//   ├ internal/bytealg
	//   ├ internal/stringslite
	//   ├ io
	//   ├ sync
	//   ├ unicode
	//   ├ unicode/utf8
	//   └ unsafe
	// 9 dependencies (9 internal, 0 external, 0 testing).
}

func Example_handlePkgsTestStrings() {
	var tree depth.Tree
	tree.ResolveTest = true

	_ = handlePkgs(&tree, &depth.Options{PackageNames: []string{"strings"}})
	// Output:
	// strings
	//   ├ bytes
	//   ├ errors
	//   ├ fmt
	//   ├ internal/abi
	//   ├ internal/bytealg
	//   ├ internal/stringslite
	//   ├ internal/testenv
	//   ├ io
	//   ├ math
	//   ├ math/rand
	//   ├ reflect
	//   ├ strconv
	//   ├ sync
	//   ├ testing
	//   ├ unicode
	//   ├ unicode/utf8
	//   └ unsafe
	// 17 dependencies (17 internal, 0 external, 8 testing).
}

func Example_handlePkgsDepth() {
	var tree depth.Tree

	_ = handlePkgs(&tree, &depth.Options{PackageNames: []string{"github.com/adapap/depth/cmd/depth"}})
	// Output:
	// github.com/adapap/depth/cmd/depth
	//   ├ encoding/json
	//   ├ flag
	//   ├ fmt
	//   ├ io
	//   ├ os
	//   ├ strings
	//   └ github.com/adapap/depth
	//     ├ bytes
	//     ├ errors
	//     ├ go/build
	//     ├ os
	//     ├ path
	//     ├ sort
	//     └ strings
	// 12 dependencies (11 internal, 1 external, 0 testing).
}

func Example_handlePkgsUnknown() {
	var tree depth.Tree

	_ = handlePkgs(&tree, &depth.Options{PackageNames: []string{"notreal"}})
	// Output:
	// 'notreal': FATAL: unable to resolve root package
}

func Example_handlePkgsJson() {
	var tree depth.Tree
	_ = handlePkgs(&tree, &depth.Options{PackageNames: []string{"strings"}, OutputJSON: true})

	// Output:
	// {
	//   "name": "strings",
	//   "internal": true,
	//   "resolved": true,
	//   "deps": [
	//     {
	//       "name": "errors",
	//       "internal": true,
	//       "resolved": true,
	//       "deps": null
	//     },
	//     {
	//       "name": "internal/abi",
	//       "internal": true,
	//       "resolved": true,
	//       "deps": null
	//     },
	//     {
	//       "name": "internal/bytealg",
	//       "internal": true,
	//       "resolved": true,
	//       "deps": null
	//     },
	//     {
	//       "name": "internal/stringslite",
	//       "internal": true,
	//       "resolved": true,
	//       "deps": null
	//     },
	//     {
	//       "name": "io",
	//       "internal": true,
	//       "resolved": true,
	//       "deps": null
	//     },
	//     {
	//       "name": "sync",
	//       "internal": true,
	//       "resolved": true,
	//       "deps": null
	//     },
	//     {
	//       "name": "unicode",
	//       "internal": true,
	//       "resolved": true,
	//       "deps": null
	//     },
	//     {
	//       "name": "unicode/utf8",
	//       "internal": true,
	//       "resolved": true,
	//       "deps": null
	//     },
	//     {
	//       "name": "unsafe",
	//       "internal": true,
	//       "resolved": true,
	//       "deps": null
	//     }
	//   ]
	// }

}

func Example_handlePkgsExplain() {
	var tree depth.Tree

	_ = handlePkgs(&tree, &depth.Options{PackageNames: []string{"github.com/adapap/depth/cmd/depth"}, ExplainPkg: "strings"})
	// Output:
	// github.com/adapap/depth/cmd/depth -> strings
	// github.com/adapap/depth/cmd/depth -> github.com/adapap/depth -> strings
}
