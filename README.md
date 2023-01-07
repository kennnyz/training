# training
This repository contain homework from udemy course

func main() {
	//Setting a new cache where we can contain everything of type key(string)-value(any)
	cache := cache.New()

	// Here I set to my cache in key "UserId" value - true
	cache.Set("UserId", true)

	// Here is another one example but with int value
	cache.Set("Muhammed", 14)
	cache.Set("KOla", 13)

	fmt.Println(cache)
	fmt.Println(cache.Get("Muhammed"))

	//Deleting record by key
	cache.Delete("Muhammed")
	fmt.Println(cache)
}
