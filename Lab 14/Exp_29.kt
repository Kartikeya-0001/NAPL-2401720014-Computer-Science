fun main() {

    val list = listOf(10, 20, 30)
    val mutableList = mutableListOf(10, 20, 30)
    mutableList.add(40)
    mutableList.remove(10)
    println("List: $list")
    println("Mutable List: $mutableList")

    val set = setOf(1, 2, 3)
    val mutableSet = mutableSetOf(1, 2, 3)
    mutableSet.add(4)
    mutableSet.remove(2)
    println("Set: $set")
    println("Mutable Set: $mutableSet")

    val map = mapOf("A" to 1, "B" to 2)
    val mutableMap = mutableMapOf("A" to 1, "B" to 2)
    mutableMap["C"] = 3
    mutableMap.remove("A")
    println("Map: $map")
    println("Mutable Map: $mutableMap")

    val numbers = listOf(1, 2, 3, 4, 5)
    val doubled = numbers.map { it * 2 }
    val even = numbers.filter { it % 2 == 0 }
    val sum = numbers.reduce { a, b -> a + b }

    println("Doubled: $doubled")
    println("Even numbers: $even")
    println("Sum: $sum")

    data class Student(val name: String, val percentage: Int)

    val students = listOf(
        Student("Aman", 75),
        Student("Riya", 90),
        Student("Karan", 65),
        Student("Neha", 85)
    )

    val sortedStudents = students.sortedByDescending { it.percentage }
    println("Sorted Students: $sortedStudents")

    val grouped = students.groupBy {
        if (it.percentage >= 80) "Above 80"
        else "Below 80"
    }

    println("Grouped Students: $grouped")
}