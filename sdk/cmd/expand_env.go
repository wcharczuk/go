package cmd

import (
	"context"
	"os"
)

// ExpandEnv expands a given statement within a given context.
func ExpandEnv(ctx context.Context, statement string) string {
	envVars := GetEnviron(ctx, os.Environ())
	envMap := EnvironMap(envVars)
	return os.Expand(statement, func(name string) string {
		if index, ok := envMap[name]; ok {
			s := envVars[index]
			for i := 0; i < len(s); i++ {
				if s[i] == '=' {
					return s[i+1:]
				}
			}
			return ""
		}
		return ""
	})
}
