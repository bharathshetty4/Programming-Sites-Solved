
import "fmt"

type Number interface {
	int | float64
}

func main() {
	n := 2
	m := 4.5
	fmt.Println(printNum(n))
	fmt.Println(printNum(m))
}

func printNum[V Number](num V) V {
	return num
}
