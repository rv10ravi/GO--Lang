package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    type menuItem struct {
        name   string
        prices map[string]float64
    }

    menu := []menuItem{
        {name: "Coffee", prices: map[string]float64{"small": 165, "medium": 250, "large": 450}},
        {name: "Espresso", prices: map[string]float64{"single": 120, "double": 220, "triple": 330}},
    }
    //fmt.Println(menu)
loop:
    for {
        fmt.Println("Please Enter your Choice")
        fmt.Println("1) Print Menu")
        fmt.Println("2) Add Item")
        fmt.Println("q)quit")
        in := bufio.NewReader(os.Stdin)
        choice, _ := in.ReadString('\n')

        switch strings.TrimSpace(choice) {
        case "1":
            for _, item := range menu {
                fmt.Println(item.name)
                fmt.Println(strings.Repeat("-", 10))
                for size, price := range item.prices {
                    fmt.Printf("\t%10s%10.2f\n", size, price)
                }
            }
        case "2":
            fmt.Println("Enter an new Menu Item")
            item, _ := in.ReadString('\n')
            menu = append(menu, menuItem{name: item, prices: make(map[string]float64)})

        case "q":
            break loop
        default:
            fmt.Println("Please Enter the Correct Option")
        }
    }
}