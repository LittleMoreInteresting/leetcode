package main

import (
	"fmt"
	"os"
)

func readMaze(file string) [][]int {
	open, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	var row, col int
	_, _ = fmt.Fscanf(open, "%d %d", &row, &col)
	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			_, _ = fmt.Fscanf(open, "%d", &maze[i][j])
		}
	}
	return maze
}

type point struct {
	x, y int
}

func (p point) add(r point) point {
	return point{p.x + r.x, p.y + r.y}
}

func (p point) at(grid [][]int) (int, bool) {
	if p.x < 0 || p.x >= len(grid) {
		return 0, false
	}
	if p.y < 0 || p.y >= len(grid[0]) {
		return 0, false
	}
	return grid[p.x][p.y], true
}

var dirs = [4]point{
	{-1, 0}, {0, -1}, {1, 0}, {0, 1},
}

func walk(maze [][]int, start, end point) [][]int {
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}
	Q := []point{start}
	for len(Q) > 0 {
		current := Q[0]

		Q = Q[1:]

		if current == end {
			break
		}
		for _, dir := range dirs {
			next := current.add(dir)
			val, ok := next.at(maze)
			if !ok || val == 1 {
				continue
			}
			val, ok = next.at(steps)
			if !ok || val != 0 {
				continue
			}

			if next == start {
				continue
			}
			curStep, _ := current.at(steps)
			steps[next.x][next.y] = curStep + 1
			Q = append(Q, next)
		}
	}
	return steps
}

func main() {
	maze := readMaze("algorithm/maze/maze.in")
	fmt.Printf("%v", maze)
	steps := walk(maze, point{0, 0},
		point{len(maze) - 1, len(maze[0]) - 1})

	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%3d", val)
		}
		fmt.Println()
	}
}
