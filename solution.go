package solution

import (
	"fmt"

	"github.com/kyokomi/emoji"
)

func GetMessage() string {
	return fmt.Sprint("Hello ", emoji.CodeMap()[":world_map:"], "!")
}
