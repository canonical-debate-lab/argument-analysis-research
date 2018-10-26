package version

import (
	"fmt"

	"github.com/kolide/kit/version"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Print version info
func Print(b bool, name string) {
	if b {
		v := version.Version()
		fmt.Printf("-- //S/M %s --\n", name)
		fmt.Printf(" - version: %s\n", v.Version)
		fmt.Printf("   branch: \t%s\n", v.Branch)
		fmt.Printf("   revision: \t%s\n", v.Revision)
		fmt.Printf("   build user: \t%s\n", v.BuildUser)
		fmt.Printf("   go version: \t%s\n", v.GoVersion)
	}
}

// Fields returns the version keys for our logger
func Fields(debug, local bool, key string) []zapcore.Field {
	if !debug || !local {
		return []zapcore.Field{
			zap.String("app", key),
			zap.String("version", version.Version().Version),
		}
	}
	return []zapcore.Field{}
}

// Release wraps kit.version.revision
func Release() string {
	return version.Version().Revision
}
