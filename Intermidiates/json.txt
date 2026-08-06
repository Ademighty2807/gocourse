package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID       int     `json:"id"`
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Address  Address `json:"address"`
}

type Address struct {
	City  string `json:"city"`
	State string `json:"state"`
}

func main() {
	user := User{
		ID:       1,
		Username: "John Doe",
		Email:    "john.doe@example.com",
		Address:  Address{City: "New York", State: "NY"},
	}

	jsonData, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	fmt.Println("JSON data:", string(jsonData))

	user2 := User{ID: 2, Username: "Jane Doe", Email: "jane.doe@example.com", Address: Address{City: "Los Angeles", State: "CA"}}
	jsonData2, err := json.Marshal(user2)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}
	fmt.Println("JSON data 2:", string(jsonData2))

	jsonData1 := `{"full_name": "Jane Doe", "emp_id": "6345", "age": 30, "address": {"city": "Los Angeles", "state": "CA"}}`

	var employeeFromJson Employee
	err = json.Unmarshal([]byte(jsonData1), &employeeFromJson)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}
	fmt.Println("Employee from JSON:", employeeFromJson)
	fmt.Println("Employee Full Name:", employeeFromJson.FullName)
	fmt.Println("Employee ID:", employeeFromJson.EmpID)
	fmt.Println("Employee Age:", employeeFromJson.Age)

	listofCityState := []Address{
		{City: "New York", State: "NY"},
		{City: "Los Angeles", State: "CA"},
		{City: "Chicago", State: "IL"},
		{City: "Houston", State: "TX"},
		{City: "Phoenix", State: "AZ"},
	}

	fmt.Println(listofCityState)

	jsonList, err := json.Marshal(listofCityState)
	if err != nil {
		fmt.Println("Error marshaling JSON list:", err)
		return
	}

	fmt.Println("JSON list of city and state:", string(jsonList))

	// Handling unknown json structure
	jsonDataUnknown := `{"name": "Alice", "age": 25, "address": {"city": "San Francisco", "state": "CA"}}`

	var data map[string]interface{}

	err = json.Unmarshal([]byte(jsonDataUnknown), &data)

	if err != nil {
		fmt.Println("Error unmarshaling unknown JSON:", err)
		return
	}
	fmt.Println("Unknown JSON data:", data)
	fmt.Println("Name:", data["name"])
	fmt.Println("Age:", data["age"])
	fmt.Println("Address:", data["address"])
}

type Employee struct {
	FullName string  `json:"full_name"`
	EmpID    string  `json:"emp_id"`
	Age      int     `json:"age"`
	Address  Address `json:"address"`
}
