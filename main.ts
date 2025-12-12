/**
 * @author Ethan White
 * @version 1.0.0
 * @date 2025-12-11
 * @fileoverview  This program will take the make,model,colour,odometer,oil change and distance driven and update them based on new wrap, and how long you drove
 */

// Car Stats Program

let carMake = "Honda";
let carModel = "Civic";
let carColor = "Red";
let odometer = 65000;
let oilChangeKM = 65000;
let gasCost: number[] = new Array(10).fill(0);
let fillUpIndex = 0;

// Function to return car stats as a string
function carStats(): string {
  return `Car Make: ${carMake}\nCar Model: ${carModel}\nCar Color: ${carColor}\nOdometer: ${odometer} km\nLast Oil Change KM: ${oilChangeKM}\nGas Costs: ${gasCost.slice(0, fillUpIndex).join(", ")}`;
}

// Function to update car color
function wrapCar(): string {
  let newColor = prompt("Enter a new color for your car: ")!;
  return newColor;
}

// Function to simulate driving a random distance
function drive(): number {
  let distance = Math.floor(Math.random() * (1000 - 100 + 1)) + 100;
  odometer += distance;
  return distance;
}

// Function to record a fill-up
function fillUp(): void {
  let cost = parseFloat(prompt("Enter the cost of gas for this fill-up: ")!);
  gasCost[fillUpIndex] = cost;
  fillUpIndex++;
}

// Function to display gas costs and average
function displayCostToFillUp(): number {
  let total = 0;
  for (let i = 0; i < fillUpIndex; i++) {
    console.log(`Fill-up ${i + 1}: $${gasCost[i]}`);
    total += gasCost[i];
  }
  let average = total / fillUpIndex;
  console.log(`Average cost per fill-up: $${average.toFixed(2)}`);
  return average;
}

// Function to check if oil change is needed
function oilChange(mileage: number, oilChangeKM: number): boolean {
  if (mileage - oilChangeKM >= 5000) {
    console.log("An oil change was done.");
    oilChangeKM = mileage;
    return true;
  } else {
    console.log("Your car does not need an oil change.");
    return false;
  }
}

// Main program
carColor = wrapCar();
console.log(`You drove ${drive()} km.`);
fillUp();
fillUp();
displayCostToFillUp();
oilChange(odometer, oilChangeKM);
console.log("\nCar Stats:");
console.log(carStats());

console.log("\nDone.")