package main

import (
	"errors"
	"fmt"
)

/*
Create a function for parking system
Base prices for motorcycle is: 2000/hour
Base prices for car is: 5000/hour
The minimum fare would be 1 hour
Maximum fare for motorcycle is 20000
Maximum fare for car is 50000

Menu:
1. Input vehicle
2. Get total fare for one vehicle
3. Get all history
*/

type Vehicle struct {
	VehicleType, PlateNumber string
	TimeIn, TimeOut          float32
	Fare                     float32
}

func calculate(duration float32, vehicleType string) (float32, error) {
	if vehicleType == "motorcycle" {
		fare := duration * 2000
		if duration <= 0 {
			return 2000, nil
		} else if fare < 20000 {
			return fare, nil
		} else {
			return 20000, nil
		}
	} else if vehicleType == "car" {
		fare := duration * 5000
		if duration <= 0 {
			return 5000, nil
		} else if fare < 50000 {
			return fare, nil
		} else {
			return 50000, nil
		}
	} else {
		return -1, errors.New("vehicle type not found")
	}
}

func addVehicle(vehicleType string, plateNumber string, timeIn float32) Vehicle {
	newVehicle := Vehicle{
		VehicleType: vehicleType,
		PlateNumber: plateNumber,
		TimeIn:      timeIn,
	}
	return newVehicle
}

func findVehicle(vehicles []Vehicle, plateNumber string) (int, error) {
	for index, vehicle := range vehicles {
		if vehicle.PlateNumber == plateNumber {
			return index, nil
		}
	}
	return 0, errors.New("plate number not found")
}

func main() {
	// Declaring array of vehicles
	var vehicles []Vehicle

	// CLI MENUS
	menu := true
	for menu {
		fmt.Println("------------------------------------------------------------------")
		fmt.Println("Welcome to the golang parking system")
		fmt.Println("------------------------------------------------------------------")
		fmt.Println("Please choose one of the menu below by typing the number")
		fmt.Println("1. Add new vehicle")
		fmt.Println("2. Get the single fare")
		fmt.Println("3. Get all histories")
		fmt.Println("I want to access menu: ...")
		var chosenMenu string
		fmt.Scanln(&chosenMenu)

		if chosenMenu == "1" {
			var vehicleType string
			var plateNumber string
			var timeIn float32

			fmt.Print("Vehicle type:")
			fmt.Scanln(&vehicleType)
			fmt.Print("Plate number:")
			fmt.Scanln(&plateNumber)
			fmt.Print("Time in:")
			fmt.Scanf("%f", &timeIn)

			vehicles = append(vehicles, addVehicle(vehicleType, plateNumber, timeIn))

			fmt.Println("Vehicle successfully added!")
			fmt.Println(vehicles[len(vehicles)-1])
			fmt.Scanln()
			fmt.Println("Do you want to exit? (Y/N)")
			var yes string
			fmt.Scanln(&yes)
			if yes == "Y" {
				menu = false
			}
		} else if chosenMenu == "2" {
			var plate string
			fmt.Print("Enter plate number:")
			fmt.Scanln(&plate)
			index, err := findVehicle(vehicles, plate)
			if err == nil {
				data := vehicles[index]
				fmt.Println("Vehicle found!")
				fmt.Println(data)
				var timeOut float32
				fmt.Print("Input time out:")
				fmt.Scanf("%f", &timeOut)
				data.TimeOut = timeOut
				duration := data.TimeOut - data.TimeIn
				fmt.Println("Parking duration is", duration)
				fmt.Println("Calculating fare...")
				fare, err := calculate(duration, data.VehicleType)
				if err == nil {
					data.Fare = fare
					fmt.Println("--------------------------")
					fmt.Println("RECEIPT")
					fmt.Println("--------------------------")
					fmt.Println("Vehicle type:", data.VehicleType)
					fmt.Println("Plate number:", data.PlateNumber)
					fmt.Println("Time In:", data.TimeIn)
					fmt.Println("Time Out:", data.TimeOut)
					fmt.Println("Fare:", data.Fare)
					fmt.Println("--------------------------")
				} else {
					fmt.Println("Fail to calculate fare. Error:", err)
				}
			} else {
				fmt.Println("Error:", err)
			}
			fmt.Scanln()
			fmt.Println("Do you want to exit? (Y/N)")
			var yes string
			fmt.Scanln(&yes)
			if yes == "Y" {
				menu = false
			}
		} else if chosenMenu == "3" {
			fmt.Println("This feature still under construction, please kindly wait for this awesome menu!")
		} else {
			fmt.Println("We don't have this menu yet, but we'll consider it later!")
		}
	}
	fmt.Println("Good Bye!")
}
