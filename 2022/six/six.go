package six

func findMarker(in string, markerSize int) int {
	marker := 0

	// Do a sliding window starting at markerSize
	for i := markerSize; i <= len(in); i++ {
		window := in[i-markerSize : i]
		if seen(window) {
			return i
		}
	}

	return marker
}

func seen(w string) bool {
	m := make(map[rune]bool)
	for _, v := range w {
		if _, exists := m[v]; exists {
			return false
		}
		m[v] = true
	}
	return true
}
