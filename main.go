package main

import "fmt"

func solve(n int) []string{
	ans := []string{}
	
	for count := 1; count <= n; count++ {
		if count%3==0 && count%5==0 {
			ans = append(ans,"FizzBuzz");
		} else if count%3==0 {
			ans = append(ans,"Fizz");
		} else if count%5==0 {
			ans = append(ans,"Buzz");
		}else {
			ans = append(ans,"0");
		}
   
	}
	return ans
}
func main(){
	
	output:= []string{}
	output=solve(100)
	fmt.Println(output)
}