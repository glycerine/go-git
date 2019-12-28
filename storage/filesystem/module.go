package filesystem

import (
	"github.com/glycerine/go-git/plumbing/cache"
	"github.com/glycerine/go-git/storage"
	"github.com/glycerine/go-git/storage/filesystem/dotgit"
)

type ModuleStorage struct {
	dir *dotgit.DotGit
}

func (s *ModuleStorage) Module(name string) (storage.Storer, error) {
	fs, err := s.dir.Module(name)
	if err != nil {
		return nil, err
	}

	return NewStorage(fs, cache.NewObjectLRUDefault()), nil
}
