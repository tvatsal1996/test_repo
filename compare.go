package main
import "log"
import "io/ioutil"
// import "fmt"

func read_file(filepath string) []byte {
	data, err := ioutil.ReadFile(filepath)
	if err != nil{
		log.Fatal(err)
	}
	return data
	// fmt.Printf("%s\n", data1)
}

func check (data1, data2 []byte) bool{
	var i, length1, length2 int
	var mask bool
	mask = true
	length1 = len(data1)
	length2 = len(data2)
	if (length1 != length2){
		return false
	} else {
		for i = 0; i<length1; i++ {
			if(data1[i]!=data2[i]){
				mask = false
				break
			}
		}
		return mask
	}
} 

func main() {
	file1 := "/Users/tvatsal/file1.txt"
	file2 := "/Users/tvatsal/file2.txt"
	// log.Printf("%T\n",file1)
	data1 := read_file(file1)
	data2 := read_file(file2)
	if ( check(data1, data2) ) {
		log.Printf("Both files are same\n") 
	} else {
		log.Printf("Both files are different\n")
	}
}