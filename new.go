package main
import "fmt"

func main() {
	filepath := "/Users/tvatsal/file1.txt"
	data := main.read_file(filepath)
	fmt.Printf("%s\n", data) 
}