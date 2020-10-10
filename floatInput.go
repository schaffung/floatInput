// package for getting a single float64 input from user.
package floatInput

import (
	"os"
	"bufio"
	"strings"
	"strconv"
)

// GetFloatNumUser reads a floating-point number from keyboard
// It returns the number read and any error encountered
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
	val, err := strconv.ParseFloat(valStr, 64)
	if err != nil {
		return 0, err
	}

	// Return the value.
	return val, nil
}
