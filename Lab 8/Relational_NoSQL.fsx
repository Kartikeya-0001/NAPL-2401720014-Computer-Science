
#r "nuget: Microsoft.Data.Sqlite, 8.0.0"
#r "nuget: MongoDB.Driver, 2.30.0"

open Microsoft.Data.Sqlite
open MongoDB.Driver
open MongoDB.Bson.Serialization.Attributes

printfn "       CRUD OPERATIONS"

// SQLITE DATABASE

printfn "\nSQLITE DATABASE"

let con = new SqliteConnection("Data Source=gym.db")
con.Open()

let execute sql =
    use cmd = con.CreateCommand()
    cmd.CommandText <- sql
    cmd.ExecuteNonQuery() |> ignore

let showSqlite() =
    use cmd = con.CreateCommand()
    cmd.CommandText <- "SELECT * FROM Members"
    use reader = cmd.ExecuteReader()

    if not reader.HasRows then
        printfn "No records found."

    while reader.Read() do
        printfn "%d  %s  %s"
            (reader.GetInt32(0))
            (reader.GetString(1))
            (reader.GetString(2))

execute """CREATE TABLE IF NOT EXISTS Members
           (Id INTEGER PRIMARY KEY, Name TEXT, Plan TEXT)"""

execute "DELETE FROM Members WHERE Id = 201"

// Create
printfn "\n1. Adding member..."
execute "INSERT INTO Members VALUES (201, 'Kartikeya', 'Monthly')"
showSqlite()

// Read
printfn "\n2. Reading member..."
showSqlite()

// Update
printfn "\n3. Updating membership..."
execute "UPDATE Members SET Plan = 'Yearly' WHERE Id = 201"
showSqlite()

// Delete
printfn "\n4. Deleting member..."
execute "DELETE FROM Members WHERE Id = 201"
showSqlite()

con.Close()

// MONGODB DATABASE

printfn "\nMONGODB DATABASE"

type Member() =
    [<BsonId>]
    member val Id = 0 with get, set

    member val Name = "" with get, set
    member val Plan = "" with get, set

let client = MongoClient("mongodb://localhost:27017")
let db = client.GetDatabase("FitnessCenter")
let members = db.GetCollection<Member>("Members")

let filter = Builders<Member>.Filter.Eq("Id", 201)

let showMongo() =
    let result = members.Find(filter).ToList()

    if result.Count = 0 then
        printfn "No records found."

    for m in result do
        printfn "%d  %s  %s" m.Id m.Name m.Plan

// Remove old test record
members.DeleteOne(filter) |> ignore

// Create
let m1 = Member()
m1.Id <- 201
m1.Name <- "Kartikeya"
m1.Plan <- "Monthly"

printfn "\n1. Adding member..."
members.InsertOne(m1)
showMongo()

// Read
printfn "\n2. Reading member..."
showMongo()

// Update
printfn "\n3. Updating membership..."
let change =
    Builders<Member>.Update.Set("Plan", "Yearly")

members.UpdateOne(filter, change) |> ignore
showMongo()

// Delete
printfn "\n4. Deleting member..."
members.DeleteOne(filter) |> ignore
showMongo()

printfn "       ALL OPERATIONS COMPLETED"