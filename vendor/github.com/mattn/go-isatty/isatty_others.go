// +build !windows
// +build !appengine

package isatty

<<<<<<< HEAD
// IsCygwinTerminal return true if the file descriptor is a cygwin or msys2
=======
// IsCygwinTerminal() return true if the file descriptor is a cygwin or msys2
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
// terminal. This is also always false on this environment.
func IsCygwinTerminal(fd uintptr) bool {
	return false
}
