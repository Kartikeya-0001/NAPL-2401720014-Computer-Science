// (1) Function with at least two default parameters (shippingFee and discountPercent)
fun calculateOrderTotal(itemPrice: Double, shippingFee: Double = 50.0, discountPercent: Double = 0.0): Double {
    val discountedPrice = itemPrice * (1.0 - discountPercent / 100.0)
    return discountedPrice + shippingFee
}

// (2) Function designed to be called with named arguments in any order
fun formatStudentProfile(name: String, rollNo: Int, grade: Char) {
    println("Student: $name | Roll No: $rollNo | Grade: $grade")
}

// (3) Function with vararg accepting a variable number of Int arguments and printing them all
fun printAllNumbers(vararg numbers: Int) {
    print("Vararg Ints (${numbers.size}): ")
    for (num in numbers) {
        print("$num ")
    }
    println()
}

// (4) Higher-order function taking a lambda (Int) -> Int and applying it to a list of integers
fun transformList(numbers: List<Int>, operation: (Int) -> Int): List<Int> {
    val result = mutableListOf<Int>()
    for (n in numbers) {
        result.add(operation(n))
    }
    return result
}

fun main() {
    println(" (1) Default Parameters ---")

    println("With defaults: ${calculateOrderTotal(1000.0)}")

    println("Without defaults (custom values): ${calculateOrderTotal(1000.0, 20.0, 10.0)}")

    println("\n (2) Named Arguments Out of Order ---")
    formatStudentProfile(grade = 'A', name = "Aarav", rollNo = 104)

    println("\n (3) Vararg Function ---")
    printAllNumbers(10, 20, 30, 40, 50)

    println("\n (4) Higher-Order Function ---")
    val sampleList = listOf(1, 2, 3, 4, 5)
    val squaredList = transformList(sampleList) { x -> x * x }
    println("Original list: $sampleList")
    println("After lambda (x -> x * x): $squaredList")
}

main()