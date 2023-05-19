package main

import "fmt"

func main() {

	sports := map[string]string{
		"cricket":   "sachin tendulkar",
		"baseball":  "anand",
		"runner":    "pt usha",
		"football":  "mani",
		"badminton": "serena williams",
	}
	fmt.Println(sports)
	delete(sports, "badminton")
	fmt.Printf("\n after deletion ----------------\n")
	fmt.Println(sports)
	fmt.Printf("\n Trying to print map ----------------\n")
	printColors(sports)


m := map[string]string{
   "dog": "bark",
 }
 
 changeMap(m)
 
 fmt.Println(m)
}
 
func changeMap(m map[string]string) {
	m["cat"] = "purr"
   }

func printColors(c map[string]string) {
	for col, hex := range c {
		fmt.Println("The player name for sports ", col, "is", hex)
	}
}
