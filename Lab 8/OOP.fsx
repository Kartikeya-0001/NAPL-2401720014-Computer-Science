type Vehicle() =
    abstract member Mileage : unit -> int
    default this.Mileage() = 0

type Car(km: int) =
    inherit Vehicle()

    override this.Mileage() =
        km

type Bike(km: int) =
    inherit Vehicle()

    override this.Mileage() =
        km

let car1 = Car(18)
let bike1 = Bike(45)

printfn "Car mileage: %d km/l" (car1.Mileage())
printfn "Bike mileage: %d km/l" (bike1.Mileage())

let vehicles: Vehicle list = [car1; bike1]

printfn "\nVehicle details:"
vehicles
|> List.iter (fun v -> printfn "Mileage: %d km/l" (v.Mileage()))
