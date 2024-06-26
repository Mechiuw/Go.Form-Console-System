package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	showForm()
	inputForm()
}

func showForm() {
	fmt.Println("------------------FORM-----------------")
	fmt.Println("1. insert your first name 			")
	fmt.Println("2. insert your last name 				")
	fmt.Println("3. insert your age 					")
	fmt.Println("4. insert your address 				")
	fmt.Println("5. insert your email 					")
	fmt.Println("6. insert your phone number 			")
	fmt.Println("7. insert your proficiency/profession ")
	fmt.Println("8. insert your company's name 		")
	fmt.Println("9. insert your salary 				")
	fmt.Println("10. insert your collegue 				")
	fmt.Println("---------------------------------------")
}

func inputForm() {
	scanner := bufio.NewScanner(os.Stdin)
	var form1 Form
	scanner.Scan()
	form1.firstName = scanner.Text()
	scanner.Scan()
	form1.lastName = scanner.Text()
	scanner.Scan()
	form1.age, _ = strconv.Atoi(scanner.Text())
	scanner.Scan()
	form1.address = scanner.Text()
	scanner.Scan()
	form1.email = scanner.Text()
	scanner.Scan()
	form1.phoneNumber = scanner.Text()
	scanner.Scan()
	form1.profession = scanner.Text()
	scanner.Scan()
	form1.compName = scanner.Text()
	scanner.Scan()
	form1.salary, _ = strconv.Atoi(scanner.Text())
	scanner.Scan()
	form1.collegue = scanner.Text()

	fmt.Println(form1)
}

type Form struct {
	firstName   string
	lastName    string
	age         int
	address     string
	email       string
	phoneNumber string
	profession  string
	compName    string
	salary      int
	collegue    string
}
