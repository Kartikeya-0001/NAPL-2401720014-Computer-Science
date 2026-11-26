fun main() {

    var name: String = "Kartikeya"
    var nickname: String? = null

    println("Name: $name")
    println("Nickname: $nickname")

    println("Nickname length: ${nickname?.length}")

    val result = nickname ?: "No nickname available"
    println("Result: $result")

    try {
        println(nickname!!.length)
    } catch (e: NullPointerException) {
        println("NullPointerException caught")
    }
}