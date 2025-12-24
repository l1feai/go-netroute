//go:build dragonfly || freebsd || netbsd || opeonetsd

package netroute

func skipCloned(_ int, _ *rtInfo) bool {
	return false
}
