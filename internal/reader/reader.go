package reader

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func printLine(arr []byte, bytesRead int, offset int, cols int, groupSize int) {
	components := make([]string, 0)
	for i := 0; i < cols; i++ {
		elem := "  "
		if i < bytesRead {
			elem = fmt.Sprintf("%02x", arr[i])

			// Replace all characters that can't be renderded in terminal with `.`
			if !(arr[i] >= 32 && arr[i] <= 126) {
				arr[i] = '.'
			}

		}

		components = append(components, elem)

		if i%groupSize == (groupSize-1) || i == cols-1 {
			components = append(components, " ")
		}
	}

	printedGroup := strings.Join(components, "")
	fmt.Printf("%08x: %s %s\n", offset, printedGroup, arr[:bytesRead])
}

func Read(file os.File, cols int, groupSize int, maxBytesToRead int) {
	offset := 0
	maxBytes := maxBytesToRead
	for {
		if maxBytes == 0 {
			break
		}

		arr := make([]byte, cols)
		bytesRead, err := file.Read(arr)
		if err != nil && bytesRead != 0 {
			log.Fatalln("Failed to read file")
		}

		if bytesRead == 0 {
			break
		}

		bytesToPrint := bytesRead
		if maxBytes > 0 && maxBytes < bytesRead {
			bytesToPrint = maxBytes
		}

		printLine(arr, bytesToPrint, offset, cols, groupSize)
		offset += cols
		maxBytes -= bytesToPrint
	}
}
