
import "fmt"

type allInterface interface {
	int | float64 | string
}

func main() {
	n, m, p := 2, 4.5, "stringToo"
	fmt.Println(printAll(n), printAll(m), printAll(p))
}

func printAll[V allInterface](allVar V) V {
	return allVar
}
