//using strings.Join is faster and more efficient than string concat with +

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))

	fmt.Println(os.Args[1:])
}
