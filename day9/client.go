package day9

import (
	"fmt"
	"log"

	"github.com/vpilkauskas/adventofcode/solver"
)

type Client struct {
	currentMarble   *Marble
	scoreBoard      map[int]int
	players         int
	lastMarbleValue int
}

var _ solver.Client = (*Client)(nil)

func New(input string) *Client {
	currentMarble := &Marble{
		value: 0,
	}
	currentMarble.next = currentMarble
	currentMarble.previous = currentMarble

	players, lastMarbleValue, err := parseInput(input)
	if err != nil {
		log.Fatal("failed to parse input", err)
	}

	return &Client{
		currentMarble:   currentMarble,
		scoreBoard:      createScoreBoard(446),
		players:         players,
		lastMarbleValue: lastMarbleValue,
	}
}

func (d *Client) Solve() string {
	for value := 1; value <= d.lastMarbleValue; value++ {
		if value%23 == 0 {
			seventhMarble := d.getSeventhMarble()
			addScore(d.scoreBoard, d.players, value, seventhMarble.value)
			d.removeMarble(seventhMarble)
			continue
		}
		d.placeMarble(value)
	}

	player, score := getWinner(d.scoreBoard)

	return fmt.Sprintf("Player nr [%d] won with score [%d]", player, score)
}

func (d *Client) placeMarble(marbleValue int) {
	newMarble := &Marble{
		value:    marbleValue,
		next:     d.currentMarble.next.next,
		previous: d.currentMarble.next,
	}

	d.currentMarble.next.next.previous = newMarble
	d.currentMarble.next.next = newMarble

	d.currentMarble = newMarble
}

func (d *Client) getSeventhMarble() *Marble {
	seventhMarble := d.currentMarble.previous

	for i := 0; i < 6; i++ {
		seventhMarble = seventhMarble.previous
	}

	return seventhMarble
}

func (d *Client) removeMarble(marble *Marble) {
	marble.previous.next = marble.next
	marble.next.previous = marble.previous
	d.currentMarble = marble.next
}
