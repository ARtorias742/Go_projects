package framework

import (
	"log"
	"os/exec"
)

func BuildWASM(srcFile, output string) error {
	cmd := exec.Command("go", "build", "-o", output, srcFile)
	cmd.Env = append(cmd.Env, "GOOS=js", "GOARCH=wasm")
	outputBytes, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("WASM Build Error: %s", outputBytes)
		return err
	}
	return nil
}
