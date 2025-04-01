package worldcount

import (
	"bufio"
	. "fmt"
	"os"


	
	"strings"
)




func World()  {
	r := string("y")

for r == "y" {
	var y string = ""




	for y == "" {
		Print("Enter The Name : ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		y = scanner.Text()
	}
	}
	
	// y = strings.Trim(y, " ")
	
	 s := make(map[string][]int) 
	
	
	
	
	for i, v := range y {
	
	t := string(v)
	

	l := strings.Count(y,t)
	
	r := l > 1
	
	
	if r {
		s[t] = append(s[t], i)
	}
	
	
	if !(t == " ") {
		if !(r && i > s[t][0]) {
			Println(t ,":" ,l,)  
		}
	}
	
	}
	
	Print("do u want to contionue ? (y/n) "   )
	Scanln(&r)
}

}