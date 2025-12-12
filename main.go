//Author: Ethan White
//Version: 1.0.0
//Date: 2025-12-11
//Fileoverview: This program will take the make,model,colour,odometer,oil change and distance driven and update them based on new wrap, and how long you drove

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func carStats(carMake, carModel, carColor string, odometer, oilChangeKM int, gasCost []float64, fillUpIndex int) string {
	return fmt.Sprintf("Car Make: %s\nCar Model: %s\nCar Color: %s\nOdometer: %d km\nLast Oil Change KM: %d\nGas Costs: %v",
		carMake, carModel, carColor, odometer, oilChangeKM, gasCost[:fillUpIndex])
}
func wrapCar(reader *bufio.Reader) string {
	fmt.Print("Enter a new color for your car: ")
	color, _ := reader.ReadString('\n')
	return strings.TrimSpace(color)
}

func drive(odometer *int) int {
	distance := rand.Intn(1000-100+1) + 100
	*odometer += distance
	return distance
}

func fillUp(reader *bufio.Reader, gasCost []float64, fillUpIndex *int) {
	fmt.Print("Enter the cost of gas for this fill-up: ")
	input, _ := reader.ReadString('\n')
	cost, _ := strconv.ParseFloat(strings.TrimSpace(input), 64)
	gasCost[*fillUpIndex] = cost
	*fillUpIndex++
}

func displayCostToFillUp(gasCost []float64, fillUpIndex int) float64 {
	total := 0.0
	for i := 0; i < fillUpIndex; i++ {
		fmt.Printf("Fill-up %d: $%.2f\n", i+1, gasCost[i])
		total += gasCost[i]
	}
	average := total / float64(fillUpIndex)
	fmt.Printf("Average cost per fill-up: $%.2f\n", average)
	return average
}

func oilChange(odometer, oilChangeKM *int) bool {
	if *odometer-*oilChangeKM >= 5000 {
		fmt.Println("An oil change was done.")
		*oilChangeKM = *odometer
		return true
	} else {
		fmt.Println("Your car does not need an oil change.")
		return false
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	reader := bufio.NewReader(os.Stdin)

	carMake := "Honda"
	carModel := "Civic"
	carColor := "Red"
	odometer := 65000
	oilChangeKM := 65000
	gasCost := make([]float64, 10)
	fillUpIndex := 0

	// Update car color
	carColor = wrapCar(reader)

	// Simulate driving
	fmt.Printf("You drove %d km.\n", drive(&odometer))

	// Fill up twice
	fillUp(reader, gasCost, &fillUpIndex)
	fillUp(reader, gasCost, &fillUpIndex)

	// Display gas costs
	displayCostToFillUp(gasCost, fillUpIndex)

	// Check oil change
	oilChange(&odometer, &oilChangeKM)

	// Display car stats
	fmt.Println("\nCar Stats:")
	fmt.Println(carStats(carMake, carModel, carColor, odometer, oilChangeKM, gasCost, fillUpIndex))
  fmt.Println("\nDone.")
}

