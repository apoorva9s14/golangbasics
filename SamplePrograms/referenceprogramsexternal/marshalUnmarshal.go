package referenceprogramsexternal

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	Name    string
	Age     int
	Address string
}
type Response struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

func main() {
	emp := Employee{Name: "George Smith", Age: 30, Address: "Newyork, USA"}
	empData, err := json.Marshal(emp)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(empData))
	var resp Response
	newErr := json.Unmarshal(empData, &resp)
	fmt.Println(newErr, resp)

}
