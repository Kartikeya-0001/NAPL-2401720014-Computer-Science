
package main

import "fmt"

func main() {

    team := [3]string{"Ronaldo", "Messi", "Neymar"}
    fmt.Println("Team Players:", team)

    players := []string{"Ronaldo", "Messi", "Neymar"}
    fmt.Println("\nPlayer List:", players)

    players = append(players, "Mbappe")
    fmt.Println("After Adding:", players)

    players = append(players[:1], players[2:]...)
    fmt.Println("After Removing:", players)

    players[1] = "Haaland"
    fmt.Println("After Updating:", players)

    jerseys := map[string]int{
        "Ronaldo": 7,
        "Messi":   10,
        "Neymar":  11,
    }

    fmt.Println("\nJersey Numbers:", jerseys)

    jerseys["Mbappe"] = 9
    fmt.Println("After Inserting:", jerseys)

    delete(jerseys, "Neymar")
    fmt.Println("After Deleting:", jerseys)

    number, found := jerseys["Messi"]

    if found {
        fmt.Println("Messi's Jersey Number:", number)
    } else {
        fmt.Println("Player not found")
    }
}
