package TaskManager

import (
	"TaskManager/bootstrapper"
)

func main() {
	bootstrapper := bootstrapper.New("TaskManager", "xzw")
	bootstrapper.Start()
}
