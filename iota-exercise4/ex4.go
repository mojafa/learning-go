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

const (
	_           = iota // c0 == 0
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	
	b
	c
	d
	e
	f
)

func main() {

}
