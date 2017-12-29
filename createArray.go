package main
import (
"log"
"os"
"math/rand"
	"strconv"
	"fmt"
)

var (
newFile *os.File
err     error
)

func main() {
newFile, err = os.Create("array.txt")
if err != nil {
log.Fatal(err)
}
log.Println(newFile)
	for i := 0; i < 10000000; i++ {
		_, err=newFile.WriteString(strconv.Itoa(rand.Intn(10000000))+"\n")
	}
newFile.Close()
fmt.Println("==> done creating file")
fmt.Println("You can use it as your mergesort input")
}