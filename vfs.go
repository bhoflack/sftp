package sftp

import "os"

// Vfs implementation to allow to override the filesystem.
type Vfs interface {
	Stat(path string) (os.FileInfo, error)
	Lstat(path string) (os.FileInfo, error)
	Mkdir(path string, perm os.FileMode) error
	Remove(path string) error
	Rename(oldpath string, newpath string) error
	Symlink(oldname string, newname string) error
	Readlink(path string) (string, error)
}

type osFs struct {
}

func (*osFs) Stat(path string) (os.FileInfo, error) {
	return os.Stat(path)
}

func (*osFs) Lstat(path string) (os.FileInfo, error) {
	return os.Lstat(path)
}

func (*osFs) Mkdir(path string, perm os.FileMode) error {
	return os.Mkdir(path, perm)
}

func (*osFs) Remove(path string) error {
	return os.Remove(path)
}

func (*osFs) Rename(oldpath string, newpath string) error {
	return os.Rename(oldpath, newpath)
}

func (*osFs) Symlink(oldname string, newname string) error {
	return os.Symlink(oldname, newname)
}

func (*osFs) Readlink(path string) (string, error) {
	return os.Readlink(path)
}

// NewOsFs for a Vfs implementation mapping to os.
func NewOsFs() Vfs {
	osFs := &osFs{}
	return osFs
}
