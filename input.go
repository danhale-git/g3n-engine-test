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

// MovePlayer is subscribed to winow.OnKeyDown and moves the player object based on key presses
func MovePlayer(evname string, ev interface{}) {

	/*var state bool
	if evname == window.OnKeyDown {
		state = true
	} else {
		state = false
	}*/

	// TODO: Apply torque and acceleration.
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
