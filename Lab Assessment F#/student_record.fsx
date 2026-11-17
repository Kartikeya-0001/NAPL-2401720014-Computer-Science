
#r "nuget: Microsoft.Data.Sqlite"

open Microsoft.Data.Sqlite

type Student = {
    Id: int
    Name: string
    Course: string
    Marks: int
}

let db = "Data Source=students.db"

let createTable() = async {
    use con = new SqliteConnection(db)
    do! con.OpenAsync() |> Async.AwaitTask

    use cmd = con.CreateCommand()
    cmd.CommandText <-
        "CREATE TABLE IF NOT EXISTS Students
         (Id INTEGER PRIMARY KEY, Name TEXT, Course TEXT, Marks INTEGER)"

    let! _ = cmd.ExecuteNonQueryAsync() |> Async.AwaitTask
    return ()
}

let result marks =
    match marks with
    | m when m >= 75 -> "Distinction"
    | m when m >= 40 -> "Pass"
    | _ -> "Fail"


let addStudent (s: Student) = async {
    use con = new SqliteConnection(db)
    do! con.OpenAsync() |> Async.AwaitTask

    use cmd = con.CreateCommand()
    cmd.CommandText <-
        "INSERT OR REPLACE INTO Students VALUES ($id,$name,$course,$marks)"

    cmd.Parameters.AddWithValue("$id", s.Id) |> ignore
    cmd.Parameters.AddWithValue("$name", s.Name) |> ignore
    cmd.Parameters.AddWithValue("$course", s.Course) |> ignore
    cmd.Parameters.AddWithValue("$marks", s.Marks) |> ignore

    let! _ = cmd.ExecuteNonQueryAsync() |> Async.AwaitTask
    return ()
}

let getStudents() = async {
    use con = new SqliteConnection(db)
    do! con.OpenAsync() |> Async.AwaitTask

    use cmd = con.CreateCommand()
    cmd.CommandText <- "SELECT * FROM Students"

    use! r = cmd.ExecuteReaderAsync() |> Async.AwaitTask
    let students = ResizeArray<Student>()

    while r.Read() do
        students.Add({
            Id = r.GetInt32(0)
            Name = r.GetString(1)
            Course = r.GetString(2)
            Marks = r.GetInt32(3)
        })

    return students |> Seq.toList
}

let updateStudent id marks = async {
    use con = new SqliteConnection(db)
    do! con.OpenAsync() |> Async.AwaitTask

    use cmd = con.CreateCommand()
    cmd.CommandText <- "UPDATE Students SET Marks=$marks WHERE Id=$id"
    cmd.Parameters.AddWithValue("$marks", marks) |> ignore
    cmd.Parameters.AddWithValue("$id", id) |> ignore

    let! _ = cmd.ExecuteNonQueryAsync() |> Async.AwaitTask
    return ()
}

let deleteStudent id = async {
    use con = new SqliteConnection(db)
    do! con.OpenAsync() |> Async.AwaitTask

    use cmd = con.CreateCommand()
    cmd.CommandText <- "DELETE FROM Students WHERE Id=$id"
    cmd.Parameters.AddWithValue("$id", id) |> ignore

    let! _ = cmd.ExecuteNonQueryAsync() |> Async.AwaitTask
    return ()
}

createTable() |> Async.RunSynchronously

let student = {
    Id = 1
    Name = "Kartikeya"
    Course = "Computer Science"
    Marks = 85
}

addStudent student |> Async.RunSynchronously

let students = getStudents() |> Async.RunSynchronously

let names = students |> List.map (fun s -> s.Name)
let passed = students |> List.filter (fun s -> s.Marks >= 40)

printfn "Student Names: %A" names
printfn "Passed Students: %A" passed

for s in students do
    printfn "%s: %s" s.Name (result s.Marks)

updateStudent 1 90 |> Async.RunSynchronously
deleteStudent 1 |> Async.RunSynchronously

printfn "Student operations completed."
