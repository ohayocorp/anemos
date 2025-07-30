package js

import (
	_ "embed"
	"log/slog"
	"os"
	"time"
)

func RunTsc(tsconfigPath string) error {
	// Running tsc with Goja requires many more NodeJS module definitions to be implemented. A working implementation
	// is available in a private branch, but the performance is not acceptable for production use. So, it is not included
	// in the main branch.

	// Running tsc with Bun is about 10x faster than running it with Goja. Still, it can take seconds to compile
	// large projects.
	// When https://github.com/microsoft/typescript-go is released, we can switch to it.
	return runTscWithBun(tsconfigPath)
}

func runTscWithBun(directory string) error {
	slog.Info(
		"Running tsc with Bun to compile ${directory}",
		slog.String("directory", directory))

	startTime := time.Now()

	err := RunBunCommand(BunCommand{
		Description: "TypeScript compilation",
		Args:        []string{"run", "tsc"},
		Cwd:         &directory,
		Stdout:      os.Stdout,
		Stderr:      os.Stderr,
		Stdin:       os.Stdin,
	})

	endTime := time.Now()
	slog.Debug("tsc execution time: ${duration}", slog.String("duration", endTime.Sub(startTime).String()))

	return err
}
