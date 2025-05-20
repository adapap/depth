package depth

import (
	"go/build"
	"sync"
)

type CachingImporter struct {
	mu    sync.Mutex
	cache map[string]*build.Package
}

func NewCachingImporter() *CachingImporter {
	return &CachingImporter{
		cache: make(map[string]*build.Package),
	}
}

func (c *CachingImporter) Import(path, srcDir string, mode build.ImportMode) (*build.Package, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if pkg, ok := c.cache[path]; ok {
		return pkg, nil
	}
	pkg, err := build.Default.Import(path, srcDir, mode)
	if err == nil {
		if existingPkg, ok := c.cache[path]; ok {
			return existingPkg, nil
		}
		c.cache[path] = pkg
	}
	return pkg, err
}
