package main

import "firestore_clean/drivers"

func main() {
	drivers.ServeUsers(":8000")
}
