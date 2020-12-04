package cmd

// EnvironMap returns a lookup map for a given set of env vars.
func EnvironMap(vars []string) (env map[string]int) {
	env = make(map[string]int)
	for i, s := range vars {
		for j := 0; j < len(s); j++ {
			if s[j] == '=' {
				key := s[:j]
				if _, ok := env[key]; !ok {
					env[key] = i // first mention of key
				} else {
					vars[i] = ""
				}
				break
			}
		}
	}
	return
}
