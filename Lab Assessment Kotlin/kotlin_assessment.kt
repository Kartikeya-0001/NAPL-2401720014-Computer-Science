data class Product(
    val name: String,
    val category: String,
    val price: Int,
    val brand: String? = null
)

fun Product.brandName(): String {
    return brand ?: "Generic"
}

fun Product.finalPrice(): Int {
    return (price * 0.90).toInt()
}

class StoreBuilder {
    private val products = mutableListOf<Product>()

    fun product(
        name: String,
        category: String,
        price: Int,
        brand: String? = null
    ) {
        products.add(Product(name, category, price, brand))
    }

    fun build(): List<Product> = products.toList()
}

fun store(block: StoreBuilder.() -> Unit): List<Product> {
    return StoreBuilder().apply(block).build()
}

fun main() {
    val products = store {
        product("Smartphone", "Electronics", 28000, "Samsung")
        product("Desk Lamp", "Home", 1200)
        product("Jacket", "Apparel", 4200, "Adidas")
        product("Backpack", "Apparel", 1800, "Puma")
        product("Coffee Mug", "Home", 450)
    }

    println("Product Names: ${products.map { it.name }}")

    val affordable = products.filter { it.price <= 2000 }
    println("Affordable Products: ${affordable.map { it.name }}")

    val total = products.map { it.price }.reduce { a, b -> a + b }
    println("Total Price: Rs. $total")

    println("\nProducts by Category:")
    products.groupBy { it.category }.forEach { (category, items) ->
        println("$category: ${items.map { it.name }}")
    }

    println("\nProduct Details:")
    products.forEach {
        println("${it.name} | Brand: ${it.brandName()} | Final Price: Rs. ${it.finalPrice()}")
    }
}