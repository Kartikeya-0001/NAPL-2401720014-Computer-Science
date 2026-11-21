import java.util.Scanner
import kotlin.math.abs

fun main() {
    // (1) Print "Hello Kotlin!!!"
    println("Hello Kotlin!!!")

    val scanner = Scanner(System.`in`)

    // (2) Employee bonus based on dept_code and years of experience using `when`
    print("Enter department code (IT, HR, FINANCE, or other): ")
    val deptCode = scanner.next().trim().uppercase()

    print("Enter years of experience: ")
    val yearsOfExperience = scanner.nextInt()

    print("Enter base salary: ")
    val salary = scanner.nextDouble()

    val bonusRate = when (deptCode) {
        "IT" -> if (yearsOfExperience > 5) 0.20 else 0.10
        "HR" -> if (yearsOfExperience > 3) 0.15 else 0.08
        "FINANCE" -> if (yearsOfExperience > 4) 0.18 else 0.10
        else -> 0.10
    }

    val bonusAmount = salary * bonusRate
    println("Department: $deptCode | Bonus %: ${(bonusRate * 100).toInt()}% | Bonus Amount: $bonusAmount")

    // (3) Read an integer and compute average of its individual digits
    print("Enter an integer: ")
    val rawNumber = scanner.nextInt()

    var current = abs(rawNumber)
    var digitSum = 0
    var digitCount = 0

    if (current == 0) {
        digitSum = 0
        digitCount = 1
    } else {
        while (current > 0) {
            digitSum += current % 10
            digitCount++
            current /= 10
        }
    }

    val digitAverage = digitSum.toDouble() / digitCount
    println("Digits count: $digitCount | Sum: $digitSum | Average of digits: $digitAverage")
}

main()