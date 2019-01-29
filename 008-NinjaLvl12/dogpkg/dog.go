// Package dogpkg will perform various tasks that a real dogpkg may do.
package dogpkg

// Years will take an int that represents human years and convert it to go years.  The returning type is an int.
func Years(hy int) int {
	return hy * 7
}
