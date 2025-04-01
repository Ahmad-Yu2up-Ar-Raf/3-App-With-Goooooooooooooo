package pyramid

import (
	"fmt"
	"time"
)

func pyramid(l string,y int,k string)  {
    
	for i := 1; i < y; i++ {
 
		time.Sleep(1 * time.Second)
	 
	  fmt.Println()
 switch k {
 case "Full":
	 for s := y; s > i ; s-- {
				 
		 fmt.Print(" ")	
		 }
		 
		 for s := 0; s < i; s++ {
		 fmt.Print(l + " ")	
		 
		 
		 }
	 case "Inverted":
		 for s := 0; s < i  ; s++ {
			 fmt.Print(" ")	
		   }
		   for s := y; s > i  ; s-- {
			 fmt.Print(l + " ")	
		   }
	 
		 default:
			 for s := 0; s < i && i <= y/2 ; s++ {
				 fmt.Print(" ")	
				 }
				 for s := y/2; s > i  ; s-- {
				 fmt.Print(l + " ")	
				 }
				 for s := y; s > i && i > y/2; s-- {
				 
				 fmt.Print(" ")	
				 }
				 
				 for s := y/2; s < i; s++ {
				 fmt.Print(l + " ")	
				 
				 
				 }
 }
 
	 
 }
	  
	 
 
	 
		
	 
	}


func Input()  {

		 
	
	 symbol := string("")
	 pyr := string("")
	
	 var length int
	
	 fmt.Println(length)
	 fmt.Println(symbol)
	
	r := string("y")
	

for r == "y"{

for symbol == "" {
	
	fmt.Print("Enter the pyramid symbol : ")
 fmt.Scanln(&symbol)

}
		for length == 0  {
		
			fmt.Print("Enter the length : ")
			fmt.Scanln(&length)
		}


			
			fmt.Print("Full / Inverted Pyramid ? or both ? : ")
		 fmt.Scanln(&pyr)
	

	pyramid(symbol,length ,pyr)
	
 fmt.Println()


	fmt.Print("do u want to contionue ? (y/n) "   )
	fmt.Scanln(&r)


symbol = ""
length = 0

}
	
	
	
	}
	