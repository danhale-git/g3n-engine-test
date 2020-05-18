package main

import (
	"fmt"

	"github.com/g3n/engine/window"
)

const (
	cmdForward = iota
	cmdBackward
	cmdLeft
	cmdRight
)

//TODO: Store key events in maps and provide functions to get ket state

// keyState := window.NewKeyState(<my app window>)

// MovePlayer is subscribed to winow.OnKeyDown and moves the player object based on key presses
func movePlayer(evname string, ev interface{}) {

	/*var state bool
	if evname == window.OnKeyDown {
		state = true
	} else {
		state = false
	}*/

	kev := ev.(*window.KeyEvent)
	switch kev.Key {
	case window.KeyW:
		fmt.Println("W")
	case window.KeyS:
		fmt.Println("S")
	case window.KeyA:
		fmt.Println("A")
	case window.KeyD:
		fmt.Println("D")
	}
}
