package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(shortestPathAllKeys([]string{"@.a.#", "###.#", "b.A.B"}))
	fmt.Println(shortestPathAllKeys([]string{"@..aA", "..B#.", "....b"}))
	fmt.Println(shortestPathAllKeys([]string{"@Aa"}))
	fmt.Println(shortestPathAllKeys([]string{"@...a", ".###A", "b.BCc"}))
}

var distence [][]int = [][]int{
	{0, -1},
	{0, 1},
	{-1, 0},
	{1, 0},
}

func shortestPathAllKeys(grid []string) int {
	m, n := len(grid), len(grid[0])
	startM, startN := 0, 0
	keyNum := map[byte]int{}
	quenen := [][]int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '@' {
				startM, startN = i, j
				continue
			}
			if grid[i][j] >= 'a' && grid[i][j] <= 'z' {
				keyNum[grid[i][j]] = len(keyNum)
			}
		}
	}
	dist := make([][][]int, m)
	for i := range dist {
		dist[i] = make([][]int, n)
		for j := range dist[i] {
			dist[i][j] = make([]int, 1<<len(keyNum))
			for k := range dist[i][j] {
				dist[i][j][k] = -1
			}
		}
	}
	dist[startM][startN][0] = 0
	step := 0
	quenen = append(quenen, []int{startM, startN, step})
	for len(quenen) > 0 {
		point := quenen[0]
		quenen = quenen[1:]
		x, y, st := point[0], point[1], point[2]
		for _, move := range distence {
			mX, mY := x+move[0], y+move[1]
			if mX < 0 || mX >= m || mY < 0 || mY >= n || isWall(grid[mX][mY]) {
				continue
			}
			char := grid[mX][mY]
			if isKey(char) {
				t := st | 1<<keyNum[char]
				if dist[mX][mY][t] == -1 {
					dist[mX][mY][t] = dist[x][y][st] + 1
					if t == 1<<len(keyNum)-1 {
						return dist[mX][mY][t]
					}
					quenen = append(quenen, []int{mX, mY, t})
				}
				continue
			}
			if isLock(char) {
				idx := keyNum[KeyToLock(char)]
				if st>>idx&1 > 0 && dist[mX][mY][st] == -1 {
					dist[mX][mY][st] = dist[x][y][st] + 1
					quenen = append(quenen, []int{mX, mY, st})
				}
				continue
			}
			if dist[mX][mY][st] == -1 {
				dist[mX][mY][st] = dist[x][y][st] + 1
				quenen = append(quenen, []int{mX, mY, st})
			}

		}
	}
	return -1
}

func isWall(b byte) bool {
	return b == '#'
}
func isKey(b byte) bool {
	return b >= 'a' && b <= 'z'
}
func isLock(b byte) bool {
	return b >= 'A' && b <= 'Z'
}
func KeyToLock(b byte) byte {
	return b - 'A' + 'a'
}

var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func shortestPathAllKeys1(grid []string) int {
	m, n := len(grid), len(grid[0])
	sx, sy := 0, 0
	keyToIdx := map[rune]int{}
	for i, row := range grid {
		for j, c := range row {
			if c == '@' {
				sx, sy = i, j
			} else if unicode.IsLower(c) {
				if _, ok := keyToIdx[c]; !ok {
					keyToIdx[c] = len(keyToIdx)
				}
			}
		}
	}

	dist := make([][][]int, m)
	for i := range dist {
		dist[i] = make([][]int, n)
		for j := range dist[i] {
			dist[i][j] = make([]int, 1<<len(keyToIdx))
			for k := range dist[i][j] {
				dist[i][j][k] = -1
			}
		}
	}
	dist[sx][sy][0] = 0
	q := [][3]int{{sx, sy, 0}}
	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		x, y, mask := p[0], p[1], p[2]
		for _, d := range dirs {
			nx, ny := x+d.x, y+d.y
			if 0 <= nx && nx < m && 0 <= ny && ny < n && grid[nx][ny] != '#' {
				c := rune(grid[nx][ny])
				if c == '.' || c == '@' {
					if dist[nx][ny][mask] == -1 {
						dist[nx][ny][mask] = dist[x][y][mask] + 1
						q = append(q, [3]int{nx, ny, mask})
					}
				} else if unicode.IsLower(c) {
					t := mask | 1<<keyToIdx[c]
					if dist[nx][ny][t] == -1 {
						dist[nx][ny][t] = dist[x][y][mask] + 1
						if t == 1<<len(keyToIdx)-1 {
							return dist[nx][ny][t]
						}
						q = append(q, [3]int{nx, ny, t})
					}
				} else {
					idx := keyToIdx[unicode.ToLower(c)]
					if mask>>idx&1 > 0 && dist[nx][ny][mask] == -1 {
						dist[nx][ny][mask] = dist[x][y][mask] + 1
						q = append(q, [3]int{nx, ny, mask})
					}
				}
			}
		}
	}
	return -1
}

//作者：力扣官方题解
//链接：https://leetcode.cn/problems/shortest-path-to-get-all-keys/solutions/1959739/huo-qu-suo-you-yao-chi-de-zui-duan-lu-ji-uu5m/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
