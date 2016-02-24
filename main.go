package gemuboi

import (
	"log"
	"os"

	_ "github.com/hajimehoshi/ebiten"
)

func parseArgs() string {
	if len(os.Args) != 2 {
		log.Fatal("File name necessary for ROM load")
	} else {
		if _, err := os.Stat(os.Args[1]); os.IsNotExist(err) {
			log.Fatal("File provided does not exist")
		}
	}

	return os.Args[1]
}

func main() {
	var Memory MemoryIO
	var CPU LR35902

	if Memory.loadROM(parseArgs()) != nil {
		log.Fatal("ROM data failed to load")
	}

	CPU.Init(&Memory)
	CPU.Run()
}
