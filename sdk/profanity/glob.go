package profanity

import "strings"

// GlobAnyMatch tests if a file matches a (potentially) csv of glob filters.
func GlobAnyMatch(filters []string, file string) bool {
	for _, part := range filters {
		if matches := Glob(strings.TrimSpace(part), file); matches {
			return true
		}
	}
	return false
}

// Glob returns if a given pattern matches a given subject.
func Glob(pattern, subj string) bool {
	// Empty pattern can only match empty subject
	if pattern == "" {
		return subj == pattern
	}

	// If the pattern _is_ a glob, it matches everything
	if pattern == Star {
		return true
	}

	parts := strings.Split(pattern, Star)

	if len(parts) == 1 {
		// No globs in pattern, so test for equality
		return subj == pattern
	}

	leadingGlob := strings.HasPrefix(pattern, Star)
	trailingGlob := strings.HasSuffix(pattern, Star)
	end := len(parts) - 1

	// Go over the leading parts and ensure they match.
	for i := 0; i < end; i++ {
		idx := strings.Index(subj, parts[i])

		switch i {
		case 0:
			// Check the first section. Requires special handling.
			if !leadingGlob && idx != 0 {
				return false
			}
		default:
			// Check that the middle parts match.
			if idx < 0 {
				return false
			}
		}

		// Trim evaluated text from subj as we loop over the pattern.
		subj = subj[idx+len(parts[i]):]
	}

	// Reached the last section. Requires special handling.
	return trailingGlob || strings.HasSuffix(subj, parts[end])
}
