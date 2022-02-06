import "fmt"

// strings are immutable and you cannot do str[i]

func main() {
	str := "ABCD"
	runes := []rune(str)
	fmt.Println("Hello", str)
	for i := 0; i < len(runes)/2; i++ {
		runes[i], runes[len(runes)-1-i] = runes[len(runes)-1-i], runes[i]
	}
	str = string(runes)
	fmt.Println("Hello1", str)
}
