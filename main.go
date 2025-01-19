package main

import "fmt"

func main() {
	req_skills := []string{"java", "nodejs", "reactjs"}
	people := [][]string{{"java"}, {"nodejs"}, {"nodejs", "reactjs"}}
	fmt.Println(smallestSufficientTeam(req_skills, people))
}
