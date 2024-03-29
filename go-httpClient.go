package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {

	resp, err := http.Get("https://api.sampleapis.com/codingresources/codingResources")

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	fmt.Println("Status respose is", resp.Status)

	scanner := bufio.NewScanner(resp.Body)

	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

}
