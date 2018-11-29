package allocdir

import (
	"os"
	"syscall"
)

// linkDir hardlinks src to dst. The src and dst must be on the same filesystem.
func linkDir(src, dst string) error {
	return syscall.Link(src, dst)
}

// unlinkDir removes a directory link.
func unlinkDir(dir string) error {
	return syscall.Unlink(dir)
}

// createSecretDir creates the secrets dir folder at the given path
func createSecretDir(dir string) error {
	// TODO solaris has support for tmpfs so use that
	return os.MkdirAll(dir, 0777)
}

// removeSecretDir removes the secrets dir folder
func removeSecretDir(dir string) error {
	return os.RemoveAll(dir)
}

// bindMount mounts src to dst with support for read only mounting.
// Assumes that filepath.Dir(dst) exists already
func bindMount(src, dst string, readOnly bool) error {
	return syscall.Link(src, dst)
}

// unmount a bind mount.  If the target is already unmounted, no error is returned
func unmount(path string) error {
	return syscall.Unlink(path)
}
