open System
open System.Text.RegularExpressions

let reverse (s: string) =
    String(Array.rev (s.ToCharArray()))

let transform (s: string) =
    if s <> "" && Seq.forall Char.IsLetter s then
        s.ToUpper()
    elif s <> "" && Seq.forall Char.IsDigit s then
        s + s
    else
        s

let checkEmail (s: string) =
    Regex.IsMatch(s, @"^[^@\s]+@[^@\s]+\.[^@\s]+$")

let strings = ["madam"; "hello"; "level"]

strings |> List.iter (fun s ->
    printfn "%s -> Reverse: %s, Palindrome: %b" s (reverse s) (s = reverse s))

["hello"; "123"; "abc123"]
|> List.iter (fun s -> printfn "%s -> Transform: %s" s (transform s))

["abc@gmail.com"; "hello"; "user@yahoo.com"]
|> List.iter (fun s -> printfn "%s -> Valid Email: %b" s (checkEmail s))