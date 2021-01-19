// +build prod

package cmd

import "context"

func createContext() context.Context {
	ctx := context.Background()
	return ctx
}

func checkOptions() {
	options.Debug = false
}
