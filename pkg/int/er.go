package er

import (
	"fmt"
)

func Log(m string, e error) error {
	if e == nil { return nil }
	fmt.Print(e)
	return fmt.Errorf("%s : %w",m, e)
}