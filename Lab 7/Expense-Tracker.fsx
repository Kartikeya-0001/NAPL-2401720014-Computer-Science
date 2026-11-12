
printfn "-- PERSONAL EXPENSE TRACKER --"

// Daily expenses
let expenses = [120; 450; 80; 300; 650]

printfn "Daily Expenses: %A" expenses

// List.map - Add a 10% extra charge to each expense
let updatedExpenses =
    List.map (fun amount -> amount + amount / 10) expenses

printfn "Updated Expenses: %A" updatedExpenses


// Shopping expenses
let shopping = [150; 700; 250; 900; 400]

printfn "\nShopping Amounts: %A" shopping

// List.filter - Find expenses above 400
let costlyItems =
    List.filter (fun amount -> amount > 400) shopping

printfn "Expenses Above 400: %A" costlyItems


// Weekly expense calculation
let weeklyExpenses = [120; 350; 200; 500; 180; 450]

printfn "\nWeekly Expenses: %A" weeklyExpenses

// List.fold - Calculate total spending
let total =
    List.fold (fun sum amount -> sum + amount) 0 weeklyExpenses

printfn "Total Weekly Spending: %d" total

// Calculate highest expense
let highest =
    List.fold (fun current amount ->
        if amount > current then amount else current) 0 weeklyExpenses

printfn "Highest Single Expense: %d" highest

printfn "\nExpense calculation completed."
