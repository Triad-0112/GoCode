import (
	"flag"
	"fmt"
	//"os"
	//"strings"
)

func main() {
	name := flag.String("n", "", "Nama")
	flag.Parse()
	fmt.Printf("Hello, %s!\n", *name)
}
