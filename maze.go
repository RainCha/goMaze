package main

import (
	"fmt"
	"os"
)

// 读取迷宫文件
func readMaze(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	// 读取文件的行和列
	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)

	// 构造迷宫 slice 结构
	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			// 读取迷宫的每一项
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}

	return maze
}

// 点结构
type point struct {
	i, j int
}

// 方向结构
var dirs = [4]point{
	{-1, 0}, // 下
	{0, -1}, // 左
	{1, 0},  // 上
	{0, 1},  // 右
}

// 点的移动操作
func (p point) move(r point) point {
	return point{p.i + r.i, p.j + r.j}
}

// 返回点在集合中的值，用来判断是否是有效的
func (p point) at(grid [][]int) (int, bool) {
	// 纵坐标没有超出集合边界
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}
	// 横坐标没有超出集合边界
	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}

	return grid[p.i][p.j], true
}

// 走迷宫算法
func walk(maze [][]int, start, end point) [][]int {
	// 构造步数集合，在对应迷宫的坐标位置，存储走到该位置的步数
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}

	// 构造探测队列
	Q := []point{start}

	for len(Q) > 0 {

		// fmt.Printf("%v", Q)
		// fmt.Println()

		cur := Q[0]
		Q = Q[1:]

		if cur == end {
			break
		}

		// 探测四个方向
		for _, dir := range dirs {
			next := cur.move(dir)

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

			// 调整当前步数
			curSteps, _ := cur.at(steps)
			steps[next.i][next.j] = curSteps + 1

			Q = append(Q, next)
		}
	}

	return steps
}

func main() {
	maze := readMaze("maze.in")

	start := point{0, 0}
	end := point{len(maze) - 1, len(maze[0]) - 1}

	steps := walk(maze, start, end)

	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%3d", val)
		}
		fmt.Println()
	}

}
