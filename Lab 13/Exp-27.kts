import java.util.Scanner

class Student(
    var rollNo: Int = 0,
    var name: String = "",
    var course: String = "",
    percentage: Double = 0.0
) {
    // Protected/private-set style encapsulation for percentage with controlled validation
    var percentage: Double = percentage
        private set

    fun updateFromInput(scanner: Scanner) {
        print("Enter Roll No: ")
        rollNo = scanner.nextInt()
        scanner.nextLine() // Consume trailing newline

        print("Enter Name: ")
        name = scanner.nextLine().trim()

        print("Enter Course: ")
        course = scanner.nextLine().trim()

        print("Enter Percentage: ")
        val inputPercentage = scanner.nextDouble()
        scanner.nextLine() // Consume trailing newline
        percentage = if (inputPercentage in 0.0..100.0) inputPercentage else 0.0
    }

    fun displayDetails() {
        println("------------------------")
        println("Roll No    : $rollNo")
        println("Name       : $name")
        println("Course     : $course")
        println("Percentage : ${"%.2f".format(percentage)}%")
        println("------------------------")
    }
}

fun main() {
    val scanner = Scanner(System.`in`)

    val student1 = Student(101, "Riya Sharma", "B.Tech CS", 88.5)
    val student2 = Student(102, "Kabir Verma", "BCA", 79.0)

    println("--- Student 1: Update & Display ---")
    student1.updateFromInput(scanner)
    student1.displayDetails()

    println("\n--- Student 2: Update & Display ---")
    student2.updateFromInput(scanner)
    student2.displayDetails()
}

main()