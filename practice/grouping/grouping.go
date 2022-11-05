package main

import (
	"fmt"
	"math/rand"
)

func main() {
	groupList := [][]string{
		[]string{"小名", "小红", "小马", "小丽", "小强"},
		[]string{"大壮", "大力", "大1", "大2", "大3"},
		[]string{"阿花", "阿朵", "阿蓝", "阿紫", "阿红"},
		[]string{"A", "B", "C", "D", "E"},
		[]string{"一", "二", "三", "四", "五"},
		[]string{"建国", "建军", "建民", "建超", "建跃"},
		[]string{"爱民", "爱军", "爱国", "爱辉", "爱月"},
	}
	result := grouping(groupList)
	for _, team := range result {
		fmt.Println(team)
	}
}

func grouping(groutList [][]string) [][]string {
	teams := [][]string{}
	team := []string{}
	length := len(groutList)
	for length > 0 {
		for i := range groutList {
			group := groutList[i]
			if len(group) == 0 {
				continue
			}
			idx := randIdx(len(group) - 1)
			if len(team) == 2 {
				teams = append(teams, team)
				team = []string{}
			}
			team = append(team, group[idx])
			groutList[i] = append(group[:idx], group[idx+1:]...)
			if len(groutList[i]) == 0 {
				length--
			}
		}
	}
	if len(team) == 1 {
		teams[0] = append(teams[0], team[0])
	}
	return teams
}

func randIdx(max int) int {
	if max == 0 {
		return 0
	}
	return rand.Intn(max)
}
