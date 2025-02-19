package sysfs

import (
	"io/fs"
	"syscall"
)

// UnimplementedFS is an FS that returns syscall.ENOSYS for all functions,
// This should be embedded to have forward compatible implementations.
type UnimplementedFS struct{}

// String implements fmt.Stringer
func (UnimplementedFS) String() string {
	return "Unimplemented:/"
}

// OpenFile implements FS.OpenFile
func (UnimplementedFS) OpenFile(path string, flag int, perm fs.FileMode) (fs.File, error) {
	return nil, syscall.ENOSYS
}

// Mkdir implements FS.Mkdir
func (UnimplementedFS) Mkdir(path string, perm fs.FileMode) error {
	return syscall.ENOSYS
}

// Rename implements FS.Rename
func (UnimplementedFS) Rename(from, to string) error {
	return syscall.ENOSYS
}

// Rmdir implements FS.Rmdir
func (UnimplementedFS) Rmdir(path string) error {
	return syscall.ENOSYS
}

// Unlink implements FS.Unlink
func (UnimplementedFS) Unlink(path string) error {
	return syscall.ENOSYS
}

// Utimes implements FS.Utimes
func (UnimplementedFS) Utimes(path string, atimeNsec, mtimeNsec int64) error {
	return syscall.ENOSYS
}
