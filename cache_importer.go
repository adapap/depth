package depth

import (
	"go/build"
)

type CachingImporter struct {
	cache map[string]*build.Package
}

func (c *CachingImporter) Import(path, srcDir string, mode build.ImportMode) (*build.Package, error) {
	if pkg, ok := c.cache[path]; ok {
		return pkg, nil
	}
	pkg, err := build.Default.Import(path, srcDir, mode)
	if err == nil {
		c.cache[path] = pkg
	}
	return pkg, err
}
