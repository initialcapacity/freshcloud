package cmds

import (
	"fmt"
	"os"
)

func must(variable string) string {
	if f := os.Getenv(variable); f == "" {
		panic(fmt.Sprintf("please set the %v environemnt variable.", variable))
	} else {
		return f
	}
}
