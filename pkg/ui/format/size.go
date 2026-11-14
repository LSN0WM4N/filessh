package format

import "fmt"

// Size converts a size in bytes to a human-readable string using
// multiples of 1024 (B, KB, MB, GB, TB, PB, EB). Values below 1024 are
// shown in bytes without decimals; all others are shown with one decimal.
//
// Examples: 512 -> "512 B", 1024 -> "1.0 KB", 1536 -> "1.5 KB".
func Size(size int64) string {
	const unit = 1024

	if size < unit {
		return fmt.Sprintf("%d B", size)
	}

	div, exp := int64(unit), 0

	for n := size / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}

	return fmt.Sprintf(
		"%.1f %cB",
		float64(size)/float64(div),
		"KMGTPE"[exp],
	)
}
