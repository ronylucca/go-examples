package main

import fmt "fmt"


func main() {
		salarios := map[string]int{ "user1":1000, "user2":2000, "user3":3000}
		fmt.Println(salarios["user3"])
		delete(salarios, "user3")
		salarios["user3.1"] = 4000

		fmt.Println(salarios["user3.1"])


		sal := make(map[string]int) 
		sal["user1_sal"] = 1000
		fmt.Println(sal["user1_sal"])

		delete(sal, "user1_sal")
		fmt.Println(sal["user1_sal"])

		for k, v := range salarios{
			fmt.Printf("O salario de %s é %d\n", k, v)
		}
		
	}



