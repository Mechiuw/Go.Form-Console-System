package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

func main() {
	showForm()
	form := inputForm()
	fmt.Println(form)
}

func showForm() {
	fmt.Println("------------------FORM-----------------")
	fmt.Println("1. Insert your First Name")
	fmt.Println("2. Insert your Last Name")
	fmt.Println("3. Insert your Age")
	fmt.Println("4. Insert your Address")
	fmt.Println("5. Insert your Email")
	fmt.Println("6. Insert your Phone Number")
	fmt.Println("7. Insert your Proficiency/Profession")
	fmt.Println("8. Insert your Company's Name")
	fmt.Println("9. Insert your Salary")
	fmt.Println("10. Insert your Colleague")
	fmt.Println("---------------------------------------")
}

func inputForm() Form {
	scanner := bufio.NewScanner(os.Stdin)
	form := Form{}

	fields := reflect.TypeOf(form)
	values := reflect.ValueOf(&form).Elem()

	for i := 0; i < fields.NumField(); i++ {
		field := fields.Field(i)
		fmt.Printf("Insert your %s: ", field.Name)
		scanner.Scan()
		input := scanner.Text()

		valueField := values.Field(i)
		switch valueField.Kind() {
		case reflect.Int:
			val, _ := strconv.Atoi(input)
			valueField.SetInt(int64(val))
		case reflect.String:
			valueField.SetString(input)
		}
	}
	return form
}

type Form struct {
	FirstName   string
	LastName    string
	Age         int
	Address     string
	Email       string
	PhoneNumber string
	Profession  string
	CompName    string
	Salary      int
	Colleague   string
}
