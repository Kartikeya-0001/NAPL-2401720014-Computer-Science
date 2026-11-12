
open System

let temperatureCheck = async {
    printfn "Temperature sensor: Collecting data..."
    do! Async.Sleep 1500
    printfn "Temperature sensor: Data received."
    return "Temperature: 29 C"
}

let humidityCheck = async {
    printfn "Humidity sensor: Checking air moisture..."
    do! Async.Sleep 600
    printfn "Humidity sensor: Check completed."
    return "Humidity: 65 percent"
}

let reportGeneration = async {
    printfn "Weather system: Preparing daily report..."
    do! Async.Sleep 1000
    printfn "Weather system: Report generated."
    return "Daily weather report is ready"
}

printfn "-- SMART WEATHER MONITOR --"
printfn "\nStarting sensor activities...\n"

let reports =
    [| temperatureCheck
       humidityCheck
       reportGeneration |]
    |> Async.Parallel
    |> Async.RunSynchronously

printfn "\n-- WEATHER SUMMARY --"

for report in reports do
    printfn "%s" report

printfn "\nAll monitoring tasks completed."
