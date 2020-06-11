package gocom

import (
	"bufio"
	"bytes"
)

const (
	wallByte = '#'
)

func loadGrid(sc *bufio.Scanner) (h, w int, grid []string) {
	h, w = nextInt(sc), nextInt(sc)
	grid = make([]string, h+2)
	wall := createWall(w)
	grid[0] = wall
	for i := 1; i <= h; i++ {
		sc.Scan()
		buf := bytes.NewBuffer([]byte{})
		buf.Write([]byte{wallByte})
		buf.WriteString(sc.Text())
		buf.Write([]byte{wallByte})
		grid[i] = buf.String()
	}
	grid[h+1] = wall
	return h, w, grid
}

func createWall(w int) string {
	buf := bytes.NewBuffer([]byte{})
	for i := 0; i < w+2; i++ {
		buf.Write([]byte{wallByte})
	}
	return buf.String()
}