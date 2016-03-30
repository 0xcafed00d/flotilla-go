package flotilla

import (
	"fmt"
	"os"
)

func ExitOnError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
