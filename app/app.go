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

func inputForm() Form {
	scanner := bufio.NewScanner(os.Stdin)
	form := Form{}

	fields := reflect.TypeOf(form)
	values := reflect.ValueOf(&form).Elem()

	for i := 0; i < fields.NumField(); i++ {
		field := fields.Field(i)
		fmt.Printf("insert your %s : ", field.Name)
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
