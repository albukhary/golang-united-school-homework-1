package solution

import (
	"fmt"

	"github.com/kyokomi/emoji"
)

func GetMessage() string {
	return fmt.Sprintf("Hello %v!", emoji.CodeMap()[":world_map:"])
}
