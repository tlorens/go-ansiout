/**
 *
 *	ANSI Output library.
 *	For creating interesting CLI apps.
 *
 * 	Author: Timothy Lorens (tlorens@cyberdyne.org)
 * 	http://www.cyberdyne.org/~icebrkr
 *
 */
package ansiout

import (
	"fmt"
	"bufio"
	"os"
	"math"
	"strings"
	"strconv"
	"time"
	"github.com/tlorens/ibkey"
)

const ESC = "\033["


// Storage variables for preventing echoing of color esc
// sequences multiple times if the colors haven't changed.
var curFg int = 7  // Default forground color
var curBg int = 0  // Default background color


/**
 *
 *	Reset text color to terminal default.
 *
 */
func ColorReset() {
	fmt.Print(ESC + "0m")
}


/**
 *
 * 	Private method to clear the terminal
 *
 * 	integer i
 * 		0 Clear from cursor to end of screen;
 * 		1 Clear from cursor to beginning of screen
 * 		2 Clear entire screen and home cursor (1,1) (ANSI.SYS)
 */
func clear(i int) {
	fmt.Printf(ESC + "%dJ", i)
}


func clearline(i int) {
	fmt.Printf(ESC + "%dK", i)
}

/**
 *
 *	Clear entire screen and home cursor (1,1)
 *
 * 	Since ^[[2J doesn't work in bash/linux.
 *
 */
func ClearScr() {
	clear(2)
	GotoXY(1,1)
}


/**
 *
 *	Clears entire line.
 *
 */
func ClearLine() {
	clearline(2)
}


/**
 *
 *	Put the cursor at X, Y coordinates
 *
 */
func GotoXY(x int, y int) {
	fmt.Printf(ESC + "%d;%dH", x, y)
}


/**
 *
 *	Get / Report cursors current postition.
 *
 * 	@FIXME: Refactor this garbage.
 *
 */
func CursorXY() (int, int) {
	var ch int
	var line string
	var ret [] int

	fmt.Print(ESC + "6n")

	for ch != 'R' {
		ch = ibkey.ReadKey(true)
		if ch != 27 && ch != '[' {
			ret = append(ret, ch)
		}
	}

	for _, value := range ret  {
		line += string(int(value))
	}

	line = strings.TrimRight(line, "R")
	parts := strings.Split(line, ";")

	X, err := strconv.Atoi(parts[0])
	if err == nil {

	}

	Y, err := strconv.Atoi(parts[1])
	if err == nil {

	}


	return X, Y
}

/**
 *
 *	Move cursor up i number of lins
 *
 */
func CursorUp(i int) {
	fmt.Printf(ESC + "%dA", i)
}


/**
 *
 *	Move cursor down i number of lines.
 *
 */
func CursorDn(i int) {
	fmt.Printf(ESC + "%dB", i)
}


/**
 *
 *	Move cursor right i number of columns
 *
 */
func CursorRt(i int) {
	fmt.Printf(ESC + "%dC", i)
}


/**
 *
 *	Move cursor left i number of columns
 *
 */
func CursorLf(i int) {
	fmt.Printf(ESC + "%dD", i)
}


/**
 *
 *	Save cursors X,Y coordinates
 *
 */
func CursorSave() {
	fmt.Print(ESC + "s")
}


/**
 *
 *	Restore / Goto saved cursor coordinates.
 *
 */
func CursorRestore() {
	fmt.Print(ESC + "u")
}


/**
 *
 *	Short-cut to print a string is specified forecolor/backcolor.
 *
 *	integer f Forground color value 0-15
 *	integer b background color value 0-7
 *	string s String to print
 *
 */
func Print(f int, b int, s string) {
	Color(f, b)
	fmt.Printf("%s", s)
}


/**
 *
 *	Set forground and background color.
 *
 *	integer f Forground color value 0-15
 *	integer b background color value 0-7
 *
 */
func Color(f int, b int) {
	var tmp string
	var colors = [8]int {0,4,2,6,1,5,3,7}

	// Prevent sending color codes if they haven't changed.
	if curFg != f || curBg != b {
		curFg = f
		curBg = b

		fg := colors[int(math.Mod(float64(f), 8))] + 30
		bg := colors[int(math.Mod(float64(b), 8))] + 40

		if b > 7 {
			tmp += "5;"
		} else {
			tmp += "0;"
		}

		if f > 7 {
			tmp += "1;"
		}

		fmt.Printf(ESC + "%s%d;%dm", tmp, fg, bg)
	}
}

/**
 *
 *	Print a string at X, Y coordinates
 *
 * 	string s Message to print
 * 	integer x X coordinate
 * 	integer y Y coordinate
 *
 */
func PrintXY(s string, x int, y int) {
	GotoXY(x,y)
	fmt.Printf("%s",s)
}


/**
 *
 *	Silly whirly-cursor progress indicator.
 *
 */
func Wait(i int) {
	for x:=1; x < i*5; x++ {
		fmt.Print("|")
		time.Sleep(50 * time.Millisecond)
		fmt.Print("\010/")
		time.Sleep(50 * time.Millisecond)
		fmt.Print("\010-")
		time.Sleep(50 * time.Millisecond)
		fmt.Print("\010\\")
		time.Sleep(50 * time.Millisecond)
		fmt.Print("\010")
	}
}


/**
 *
 *	Private method to read strings from files.
 *
 */
func readln(r *bufio.Reader) (string, error) {
	var (isPrefix bool = true
		err error = nil
		line, ln []byte
	)

  for isPrefix && err == nil {
	  line, isPrefix, err = r.ReadLine()
	  ln = append(ln, line...)
  }

  return string(ln),err
}


/**
 *
 *	Print a file to the screen.
 *
 */
func PrintFile(filename string) int {
	f, err := os.Open(filename)

	if err != nil {
		fmt.Printf("error opening file: %v\n",err)
		return -1
	}

	r := bufio.NewReader(f)

	s, e := readln(r)

	for e == nil {
		fmt.Println(s)
		s,e = readln(r)
	}

	fmt.Print("\n")

	return 1
}
