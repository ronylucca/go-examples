package main

import fmt "fmt"


func main() {

		s:=[]int{10,20,30,40,50,60,70,80,90,100}
		fmt.Printf("Len=%d, Capacidade=%d %v\n", len(s) , cap(s) , s)

		s = append(s, 110)
		fmt.Printf("Len=%d, Capacidade=%d %v\n",  len(s[2:]) , cap(s[2:]) , s[2:])
	}



