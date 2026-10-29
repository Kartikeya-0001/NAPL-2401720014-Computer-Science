let isPrime n =
    let rec check i =
        if i * i > n then true
        elif n % i = 0 then false
        else check (i + 1)

    if n < 2 then false
    else check 2

printfn "%b" (isPrime 7)
printfn "%b" (isPrime 9)
printfn "%b" (isPrime 13)