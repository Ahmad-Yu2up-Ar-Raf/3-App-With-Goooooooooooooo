package List



import (
	"bufio" 

	. "fmt"
	"os"
	"slices"
)

type students struct {
	Id int
	Name string
	Grade int
}

func List()  {
	name := string("")


 SepuluhSatu := []students{
  
 }

 
  NumberEnd := int(0)



menu := string("")
var grade int





 for {

	Println()
	Println()
	Println("1. Show List")
	Println("2. Add List")
	Println("3. Remove Data")
	Println("4. Clear Data")
	Println("5. Exit")
    
	Println()
		
		Print("Select Menu: ")
		Scanner := bufio.NewScanner(os.Stdin)
		if Scanner.Scan() {
			menu = Scanner.Text()
		}



	switch menu {
	case "1":
		Println()
		if len(SepuluhSatu) == 0 {
			Print("Your Have 0 Student, Pls Add")
		}else{
	
			 for _, v := range SepuluhSatu {
				Printf("%d. %s %d \n" , v.Id, v.Name, v.Grade)
			 }
}
	case "2":	

	NumberEnd++
	Println()
	Print("Enter User : ")
	Scanner := bufio.NewScanner(os.Stdin)
	if Scanner.Scan() {
		name = Scanner.Text()
	}


	Print("Enter Student Grade : ")
	Scanln(&grade)

SepuluhSatu = append(SepuluhSatu, students{Id: len(SepuluhSatu) + 1, Name: name, Grade: grade })

case "3":
	Println()
	if len(SepuluhSatu) == 0 {
		Print("Your Have 0 Student, Pls Add")
	}else{

		

		 for _, v := range SepuluhSatu {
			Printf("%d. %s %d \n" , v.Id, v.Name, v.Grade)
		 }
var ge int
Println()
   Printf("Select Student ID: ")
   Scanln(&ge)




   


  for _, v := range SepuluhSatu {
	if v.Id > ge {
	   
	
		SepuluhSatu = append(SepuluhSatu, students{Id: v.Id - 1, Name: v.Name, Grade: v.Grade })

	
	
		
	}

}

SepuluhSatu = slices.Delete(SepuluhSatu, ge-1, NumberEnd )
NumberEnd--
}
case "4":	
if len(SepuluhSatu) == 0 {
	Print("Your Have 0 Student, Pls Add")
}else{
SepuluhSatu =  nil
}

	default:
		Println("Bruh Wrong Input")
		os.Exit(0)
 }

 }
}