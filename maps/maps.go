package maps

import "fmt"

/*
For maps, the funciton make accepts only two args, not three as seen at slices
*/

func main() {
	// m := make(map[string]string, 100)
	// m := map[string]string{
	// 	"Pedro": "Pessoa",
	// 	"Joao": "Legal",
	// }

	// m := map[string][]int{
	// 	"Pedro": []int{1,2,3},
	// }

	// m := make(map[string]string)

	// m["Pedro"] = "Pessoa"
	// val, ok := m["Pedro"]

	// delete(m, "Pedro")

	// fmt.Println(val)
	// fmt.Println(ok)
	// fmt.Println(m)
	
	// clear(m) /* empty a map */

	m := map[string]string{
		"Pedro": "Pessoa",
		"Joao": "Legal",
	}

	for k, v := range m {
		fmt.Println(k)
		fmt.Println(v)
	}
}