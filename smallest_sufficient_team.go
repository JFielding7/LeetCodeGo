package main

import "fmt"

func getSkillNums(skills []string) map[string]int {
	skillNums := make(map[string]int)
	for i, skill := range skills {
		skillNums[skill] = i
	}
	return skillNums
}

func getSkillMasks(skillNums map[string]int, people [][]string) []int {
	skillMasks := make([]int, len(people))
	for i, person := range people {
		skillMask := 0
		for _, skill := range person {
			skillMask |= 1 << skillNums[skill]
		}
		skillMasks[i] = skillMask
	}
	return skillMasks
}

func getTeam(size int, smallestTeams [][]int, skillMasks []int, people int, remaining int) []int {
	team := make([]int, size)
	i := size - 1
	fmt.Println(remaining)
	for remaining > 0 {
		if smallestTeams[people][remaining] != smallestTeams[people-1][remaining] {
			remaining &= ^skillMasks[people-1]
			team[i] = people - 1
			i--
			fmt.Println(team, remaining)
		}
		people--
	}

	return team
}

func smallestSufficientTeam(reqSkills []string, people [][]string) []int {
	skillNums := getSkillNums(reqSkills)
	skillMasks := getSkillMasks(skillNums, people)
	skillsCombos := 1 << len(reqSkills)
	numPeople := len(people)

	smallestTeams := make([][]int, numPeople+1)
	smallestTeams[0] = make([]int, skillsCombos)
	for i := 1; i < skillsCombos; i++ {
		smallestTeams[0][i] = -1
	}

	for i, skillMask := range skillMasks {
		smallestTeams[i+1] = make([]int, skillsCombos)

		for mask := range skillsCombos {
			remaining := (mask & skillMask) ^ mask
			if smallestTeams[i][remaining] == -1 {
				smallestTeams[i+1][mask] = -1
			} else if smallestTeams[i][mask] == -1 {
				smallestTeams[i+1][mask] = smallestTeams[i][remaining] + 1
			} else {
				smallestTeams[i+1][mask] = min(smallestTeams[i][remaining]+1, smallestTeams[i][mask])
			}
		}
	}

	fmt.Println(smallestTeams)
	return getTeam(smallestTeams[numPeople][skillsCombos-1], smallestTeams, skillMasks, numPeople, skillsCombos-1)
}
