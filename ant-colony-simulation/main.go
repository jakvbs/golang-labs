package main

import (
	"fmt"
	"math/rand"
	"time"
)

type position struct {
	x int
	y int
}

type leaf struct {
	position position
	board    *board
	sign     string
}

type ant struct {
	position  position
	board     *board
	sign      string
	hasFood   bool
	restCount int
}

func (a *ant) getPositionsWithoutAnts() []position {
	return a.board.getPositionsWithoutAnts(a.position)
}

func (a *ant) move() {
	rand.Seed(time.Now().UnixNano())
	positionsWithoutAnts := a.getPositionsWithoutAnts()
	if len(positionsWithoutAnts) == 0 {
		return
	}

	nextPosition := positionsWithoutAnts[rand.Intn(len(positionsWithoutAnts))]

	l := a.board.getLeafFromPosition(nextPosition)

	if a.hasFood {
		if l == nil {
			neighboringLeafs := a.board.getPositionsWithLeafs(nextPosition)
			if len(neighboringLeafs) > 0 {
				a.board.addLeaf(&leaf{nextPosition, a.board, "🍀"})
				a.sign = "🐜"
				a.hasFood = false
				a.restCount = 3
			}
		}
	} else if a.restCount == 0 {
		if l != nil {
			neighboringLeafs := a.board.getPositionsWithLeafs(nextPosition)
			if len(neighboringLeafs) < 5 {
				a.board.removeLeaf(l)
				a.sign = "🐌"
				a.hasFood = true
			}
		}
	} else {
		a.restCount--
	}
	a.board.moveAnt(a, nextPosition)
}

type board struct {
	size      int
	turn      int
	positions []position
	ants      map[position]*ant
	leafs     map[position]*leaf
	separator string
}

func (b *board) positionOnBoard(p position) bool {
	return p.x >= 0 && p.x < b.size && p.y >= 0 && p.y < b.size
}

func (b *board) getPositionsWithoutAnts(p position) []position {
	return b.filterPositionsWithoutAnts(b.getNeighboringPositions(p))
}

func (b *board) getPositionsWithLeafs(p position) []position {
	return b.filterPositionsWithLeafs(b.getNeighboringPositions(p))
}

func (b *board) getNeighboringPositions(p position) []position {
	var result []position

	for y := -1; y <= 1; y++ {
		for x := -1; x <= 1; x++ {
			pomPosition := position{p.x + x, p.y + y}
			if b.positionOnBoard(pomPosition) {
				result = append(result, pomPosition)
			}
		}
	}
	return result
}

func (b *board) getAntFromPosition(p position) *ant {
	ant, ok := b.ants[p]
	if ok {
		return ant
	}
	return nil
}

func (b *board) getLeafFromPosition(p position) *leaf {
	leaf, ok := b.leafs[p]
	if ok {
		return leaf
	}
	return nil
}

func (b *board) filterFreePositions(fields []position) []position {
	var result []position

	for _, field := range fields {
		a := b.getAntFromPosition(field)
		l := b.getLeafFromPosition(field)
		if a == nil && l == nil {
			result = append(result, field)
		}
	}
	return result
}

func (b *board) filterPositionsWithoutAnts(fields []position) []position {
	var result []position

	for _, field := range fields {
		a := b.getAntFromPosition(field)
		if a == nil {
			result = append(result, field)
		}
	}
	return result
}

func (b *board) filterPositionsWithLeafs(fields []position) []position {
	var result []position
	for _, field := range fields {
		l := b.getLeafFromPosition(field)
		if l != nil {
			result = append(result, field)
		}
	}
	return result
}

func (b *board) getAllFreePositions() []position {
	return b.filterFreePositions(b.positions)
}

func (b board) String() string {
	var result string
	result += fmt.Sprintf("\nturn: %d\n", b.turn)
	for y := 0; y < b.size; y++ {
		for x := 0; x < b.size; x++ {
			a := b.getAntFromPosition(position{x, y})
			l := b.getLeafFromPosition(position{x, y})
			if a != nil {
				result += a.sign
			} else if l != nil {
				result += l.sign
			} else {
				result += b.separator
			}
		}
		result += "\n"
	}
	return result
}

func (b *board) addAnt(a *ant) {
	b.ants[a.position] = a
}

func (b *board) addLeaf(l *leaf) {
	b.leafs[l.position] = l
}

func (b *board) removeLeaf(l *leaf) {
	delete(b.leafs, l.position)
}

func (b *board) initialize(ants int, leafs int) {
	rand.Seed(time.Now().UnixNano())
	b.turn = 1
	b.positions = make([]position, b.size*b.size)
	b.ants = make(map[position]*ant)
	b.leafs = make(map[position]*leaf)

	for x := 0; x < b.size; x++ {
		for y := 0; y < b.size; y++ {
			b.positions = append(b.positions, position{x, y})
		}
	}

	for len(b.leafs) < leafs {
		freePositions := b.getAllFreePositions()
		if len(freePositions) == 0 {
			fmt.Println("no free positions")
			break
		}
		randomPosition := freePositions[rand.Intn(len(freePositions))]
		b.addLeaf(&leaf{randomPosition, b, "🍀"})
	}

	for len(b.ants) < ants {
		freePositions := b.getAllFreePositions()
		if len(freePositions) == 0 {
			fmt.Println("no free positions")
			break
		}
		randomPosition := freePositions[rand.Intn(len(freePositions))]
		b.addAnt(&ant{randomPosition, b, "🐜", false, 0})
	}
}

func (b *board) moveAnt(a *ant, p position) {
	delete(b.ants, a.position)
	a.position = p
	b.ants[p] = a
}

func (b *board) makeTurn() {
	for _, ant := range b.ants {
		ant.move()
	}
	b.turn++
}

func main() {
	board := board{size: 25, separator: "⬜"}
	board.initialize(75, 100)
	for i := 1; i <= 10000; i++ {
		if i%10 == 0 {
			fmt.Print("\033[H\033[2J")
			fmt.Println(board)
		}
		// fmt.Print("\033[H\033[2J")
		// fmt.Println(board)
		// time.Sleep(time.Millisecond * 10)
		board.makeTurn()
	}
}
