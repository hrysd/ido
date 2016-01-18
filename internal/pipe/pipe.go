package pipe

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func Read() string {
	var lines bytes.Buffer
	scanner := bufio.NewScanner(os.Stdin)

	lines.Write([]byte("```\n"))

	for scanner.Scan() {
		fmt.Println(scanner.Text())
		lines.Write(scanner.Bytes())
		lines.Write([]byte("\n"))
	}

	lines.Write([]byte("\n```"))

	return lines.String()
}
