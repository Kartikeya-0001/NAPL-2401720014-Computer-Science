let checkDivision marks =
    let percentage = float (List.sum marks) / float marks.Length

    printfn "Percentage = %f" percentage

    match percentage with
    | p when p >= 60.0 -> printfn "First Division"
    | p when p >= 50.0 -> printfn "Second Division"
    | p when p >= 40.0 -> printfn "Third Division"
    | _ -> printfn "Fail"

checkDivision [80; 70; 60; 50]
checkDivision [55; 50; 55; 40]
checkDivision [45; 40; 42; 41]
checkDivision [20; 30; 25; 35]