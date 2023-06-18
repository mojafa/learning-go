package main

import "fmt"

func main() {
	//first string is key, second is value
	m := make(map[string][]string)
	m[`bond_james`] = []string{`Shaken, not stirred`, `Martinis`, `Women`}
	m[`moneypenny_miss`] = []string{`James Bond`, `Literature`, `Computer Science`}
	m[`no_dr`] = []string{`Being evil`, `Ice cream`, `Sunsets`}
	fmt.Printf("%#v\n", m)

	for k, v := range m {
		fmt.Println(k)
		for i, v2 := range v {
			fmt.Println(i, v2)
		}

	}
}
