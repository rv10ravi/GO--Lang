// Q3 You're developing a weather monitoring system in GoLang for a research institute. The system needs to store data about temperature, humidity, and wind speed. Define variables to hold these measurements for a specific location at a given point in time. Ensure you choose suitable data types for representing numerical measurements accurately.

package main

import "fmt"

//using struct to store the weather details
type weather struct {
	temperature float64
	humidity float64
	windSpeed float64
}

func main(){
	//Declaring arrray of structures
	var w[100] weather
	var n int
	fmt.Println("Enter the number of locations you want to enter:")
	fmt.Scanln(&n)
	for i:=1;i<=n;i++ {
		fmt.Println("Enter Location ",i," Temperature (in Celcius):")
		fmt.Scanln(&w[i].temperature)
		fmt.Println("Enter Location ",i," Humidity(in %):")
		fmt.Scanln(&w[i].humidity)
		fmt.Println("Enter Location ",i," Wind Speed(in km/hr):")
		fmt.Scanln(&w[i].windSpeed)
		fmt.Println("Location details entered successfully")
		fmt.Println("====================================")
	}
	for i:=1;i<=n;i++ {
		//Print the product details
		fmt.Println("Location ",i," Temperature:",w[i].temperature,"Celcius")
		fmt.Println("Location ",i," Humidity:",w[i].humidity,"%")
		fmt.Println("Loc45ation ",i," Wind Speed:",w[i].windSpeed,"km/hr")
		fmt.Println("====================================")
	}

}