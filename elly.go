package main
import ("os"
		"fmt"
		"flag"
		"path/filepath"
)

var version = "0.0.1"
var outputpath = "../out"

func main() {
	inputpath := flag.String("i", ".", "source folder")
	flag.StringVar(&outputpath, "o", "../out", "destination folder")
	flagversion := flag.Bool("v", false, "display version")
	flag.Parse()

	if *flagversion {
		fmt.Printf("\nElly version %s\nOld school PHP site constructor\n", version)
		fmt.Printf("\nCommand-line options")
		fmt.Printf("\n--------------------\n")
		flag.PrintDefaults()
		fmt.Printf("\n")
		os.Exit(0)
	}

	filepath.Walk(*inputpath, walkpath)

/*	file, err := os.Open("file.go") // For read access.

	if err != nil {
		log.Fatal(err)
	}
	data := make([]byte, 100)
	count, err := file.Read(data)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("read %d bytes: %q\n", count, data[:count])
*/
	fmt.Printf("Hello World!\n%s\n%s\n%b\n", *inputpath, outputpath, *flagversion)
}

func walkpath(path string, f os.FileInfo, err error) error {
	fmt.Printf("%s with %d bytes\n", path,f.Size())
	return nil
}
