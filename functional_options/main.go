package main

import "github.com/IvanSharovarov/golang-patterns/functional_options/file"

func main() {
	err := file.New("/tmp/empty.txt")
	if err != nil {
		panic(err)
	}

	err = file.New("/tmp/file.txt", file.UID(1000), file.Contents("My text for file"))
	if err != nil {
		panic(err)
	}
}
