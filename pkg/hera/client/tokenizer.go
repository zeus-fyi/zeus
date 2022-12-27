package hera_client

import (
	"os/exec"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/zeus-fyi/zeus/pkg/python"
)

func (h *HeraClient) GetTokenApproximate(prompt string) int {
	python.ForceDirToPythonDir()
	cmd := exec.Command("python", "tokenizer.py", prompt)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Err(err)
		return 0
	}
	// Convert the output to a string and split it by newline
	outputStr := string(output)
	outputLines := strings.Split(outputStr, "\n")

	// Convert the first line of the output (which should be the result) to an integer
	result, err := strconv.Atoi(outputLines[0])
	if err != nil {
		log.Err(err)
		return 0
	}
	return result
}
