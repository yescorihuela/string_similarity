package main

import (
	"fmt"

	"github.com/yescorihuela/string_similarity/string_similarity"
)

func main() {
	a := "Don't change a tagged version of a module after publishing it. For developers using the module, Go tools authenticate a downloaded module!"
	b := "Don't change a tagged version of a module after publishing it. For developers using the module, Go tools authenticate a downloaded module"
	fmt.Printf("%f", string_similarity.Similarity(a, b))
}
