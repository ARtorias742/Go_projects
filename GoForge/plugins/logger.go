package plugins

import "log"

func Log(message string) {
	log.Println("[PLUGIN]" + message)
}
