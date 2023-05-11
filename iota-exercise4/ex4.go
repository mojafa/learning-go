// create a program that uses iota to calculate the size of each measurement of bytes
// ○ KB
// ○ MB
// ○ GB
// ○ TB
// ○ PB
// ○ EB
// 1024 bytes 1024 KB 1024 MB 1024 GB 1024 TB 1024 EB
// ● hint: use int and not float64 as shown in effective go (we aren't going up to the larger values of ZB
// and YB so we don't need to capacity of float64)

package main

import "fmt"

const (
	// _ = 1 << (10 * iota)
	_           = iota // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota)

	MB
	GB
	TB
	PB
	EB
)

type ByteSize int

func main() {
	// fmt.Printf("1 KB = %d bytes\n", KB)
	// fmt.Printf("1 MB = %d bytes\n", MB)
	// fmt.Printf("1 GB = %d bytes\n", GB)
	// fmt.Printf("1 TB = %d bytes\n", TB)
	// fmt.Printf("1 PB = %d bytes\n", PB)
	// fmt.Printf("1 EB = %d bytes\n", EB)

	fmt.Printf("1 KB = %d bytes \t\t %b\n", KB, KB)
	fmt.Printf("1 MB = %d bytes \t\t %b\n", MB, MB)
	fmt.Printf("1 GB = %d bytes \t %b\n", GB, GB)
	fmt.Printf("1 TB = %d bytes \t %b\n", TB, TB)
	fmt.Printf("1 PB = %d bytes \t %b\n", PB, PB)
	fmt.Printf("1 EB = %d bytes %b\n", EB, EB)

}
