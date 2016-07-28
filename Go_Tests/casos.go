package main

import "fmt"
import "time"

func main() {

    //Switch.
    i := 2
    fmt.Print("Caso ", i, " es ")
    switch i {
    case 1:
        fmt.Println("UNOO")
    case 2:
        fmt.Println("DOSS")
    case 3:
        fmt.Println("TRES")
    }

    // You can use commas to separate multiple expressions
    // in the same `case` statement. We use the optional
    // `default` case in this example as well. Interesante
	
    switch time.Now().Weekday() 
	{
    case time.Saturday, time.Sunday:
        fmt.Println("Hoy es Finde")
    default:
        fmt.Println("Hoy es una semana")
    }

  
    t := time.Now()
    switch 
	{
		case t.Hour() < 12:
			fmt.Println("Son antes de las doce")
		default:
			fmt.Println("Es el mismo default que java si no son las 12  o antes")
    }
}
