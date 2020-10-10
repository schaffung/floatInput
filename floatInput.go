package floatInput

import (
	"os"
	"log"
	"bufio"
	"strings"
	"strconv"
)

func GetFloatNumUser() ( float64, error ){

	// Creating a bufio reader.
	reader := bufio.NewReader(os.Stdin)

	// Reading the user input.
	valStr, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	// Trim the spaces.
	valStr = strings.TrimSpace(valStr)

	// Convert the string to float value
	val, err := strconv.ParseFloat(valStr)
	if err != nil {
		return 0, err
	}

	// Return the value.
	return val, nil
}
