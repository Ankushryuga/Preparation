package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Dice struct {
	sides int
}

func (d *Dice) Roll() int {
	return rand.Intn(d.sides) + 1
}

type Player struct {
	ID       int
	Name     string
	Position int
}

type Board struct {
	Size    int
	Snakes  map[int]int
	Ladders map[int]int
}

type Game struct {
	Board      *Board
	Dice       *Dice
	Players    []*Player
	Winners    []*Player
	TurnCh     chan int
	ResultCh   chan bool
	DoneCh     chan struct{}
	Active     map[int]bool
	MaxPos     int
	WinnerLock sync.Mutex
	WG         sync.WaitGroup
}

func NewGame(size, numSnakes, numPlayers int) *Game {
	rand.Seed(time.Now().UnixNano())

	board := &Board{
		Size:    size,
		Snakes:  map[int]int{},
		Ladders: map[int]int{},
	}

	// Random snakes
	for len(board.Snakes) < numSnakes {
		head := rand.Intn(size*size-1) + 2
		tail := rand.Intn(head-1) + 1
		if board.Snakes[head] == 0 {
			board.Snakes[head] = tail
		}
	}

	// Random ladders
	for len(board.Ladders) < numSnakes {
		start := rand.Intn(size*size-1) + 1
		end := rand.Intn(size*size-start) + start + 1
		if board.Snakes[start] == 0 && board.Ladders[start] == 0 {
			board.Ladders[start] = end
		}
	}

	// Players
	players := []*Player{}
	for i := 0; i < numPlayers; i++ {
		players = append(players, &Player{ID: i, Name: fmt.Sprintf("Player-%d", i+1), Position: 0})
	}

	active := make(map[int]bool)
	for i := 0; i < numPlayers; i++ {
		active[i] = true
	}

	return &Game{
		Board:    board,
		Dice:     &Dice{6},
		Players:  players,
		TurnCh:   make(chan int),
		ResultCh: make(chan bool),
		DoneCh:   make(chan struct{}),
		Active:   active,
		MaxPos:   size * size,
	}
}

func (g *Game) Start() {
	for _, player := range g.Players {
		g.WG.Add(1)
		go g.playerLoop(player)
	}

	go g.gameMaster()

	g.WG.Wait()
// 	close(g.DoneCh)
}

func (g *Game) playerLoop(p *Player) {
	defer g.WG.Done()
	for {
		select {
		case <-g.DoneCh:
			return
		case id, ok := <-g.TurnCh:
		    if !ok{
		        return
		    }
			if id != p.ID {
				// Not this player's turn, put it back
				g.TurnCh <- id
				continue
			}

			roll := g.Dice.Roll()
			fmt.Printf("🎲 %s rolled a %d\n", p.Name, roll)

			newPos := p.Position + roll
			if newPos > g.MaxPos {
				fmt.Printf("%s can't move, stays at %d\n", p.Name, p.Position)
				g.nextTurn(p.ID)
				continue
			}

			if dest, ok := g.Board.Snakes[newPos]; ok {
				fmt.Printf("🐍 %s bitten by snake! %d → %d\n", p.Name, newPos, dest)
				newPos = dest
			} else if dest, ok := g.Board.Ladders[newPos]; ok {
				fmt.Printf("🪜 %s climbed a ladder! %d → %d\n", p.Name, newPos, dest)
				newPos = dest
			}

			fmt.Printf("%s moved to %d\n", p.Name, newPos)
			p.Position = newPos

			if newPos == g.MaxPos {
				fmt.Printf("🏁 %s WINS!\n", p.Name)
				g.WinnerLock.Lock()
				g.Winners = append(g.Winners, p)
				g.Active[p.ID] = false
				g.WinnerLock.Unlock()
			}

			// Check if game should end
			activePlayers := 0
			for _, v := range g.Active {
				if v {
					activePlayers++
				}
			}
			if activePlayers <= 1 {
				g.ResultCh <- true
				return
			}

			g.nextTurn(p.ID)
		}
	}
}

func (g *Game) nextTurn(currentID int) {
	next := (currentID + 1) % len(g.Players)
	for {
		if g.Active[next] {
			g.TurnCh <- next
			return
		}
		next = (next + 1) % len(g.Players)
	}
}

func (g *Game) gameMaster() {
	// Start the game with player 0
	g.TurnCh <- 0

	<-g.ResultCh
	fmt.Println("Game finished. Closing all players.")
	close(g.DoneCh) //notify all players to exit
}

func main() {
	var size, snakes, players int
	fmt.Print("Enter board size (n for n x n): ")
	fmt.Scan(&size)
	fmt.Print("Enter number of snakes/ladders: ")
	fmt.Scan(&snakes)
	fmt.Print("Enter number of players: ")
	fmt.Scan(&players)

	if size < 2 || snakes < 1 || players < 2 {
		fmt.Println("❌ Invalid input. Minimum size: 2, snakes: 1, players: 2")
		return
	}

	game := NewGame(size, snakes, players)
	fmt.Println("🕹️ Starting Snake and Ladder Game...")
	game.Start()
}
