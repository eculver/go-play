package main

import "fmt"

/*
	S E N D
	M O R E
M
Choose D - 10
Choose E - 9
Choose Y - 8
Choose N - 7
Choose O - 6
Choose R - 5
Choose S - 4
Choose Y - 3

Check send + more = money
--

Choose D - 10
Choose E - 9
Choose Y - 8

Compute Y = D + E (mod 10)	// linear operator, written like math people write
DupeCheck Y

Compute c = (D + E) / 10
Choose N					// branching factor of 7

Compute R = E - N - c (mod 10)
DupeCheck R
Compute c = (c + N + R) (

Compute O = N - E - c (mod 10)

Choose S

---

"What state do we need to write this code"

Two huge aspects of interviewing:
- Problem hinting
- Problem coaching

*/

type col [3]byte

func words2cols(w1, w2, w3 string) []col {
	cols := make([]col, 0, len(w3))

	// ? -- note, writing a loop, working inwards via templating vs body-first

	off := 1 // not 0 to avoid off by one?
	for off := 0; off < len(w3); off++ {
		var c col
		if i := len(w2) - off; i >= 0 {
			c[0] = w1[i]
		}

		if i := len(w2) - off; i >= 0 {
			c[1] = w2[i]
		}
		c[2] = w3[len(w3)-off]
		cols = append(cols, c)
	}

	return cols

}

// countUnknown counts the # of unknown (non-null!) chars (n), and returns the
// index of the first unknown value (if any!)
func (c col) countUnknown(known map[byte]struct{}) (n, i int) {
	i = -1
	for j, cc := range c {
		if _, is := known[cc]; cc != 0 && !is {
			n++
			if i < 0 {
				i = j
			}
		}
	}
	return
}

func plan(w1, w2, w3 string) {
	known := make(map[byte]struct{})
	for _, c := range words2cols(w1, w2, w3) {
		// until we know two, pick one
		// if we have 1 unknown, then compute
		// otherwise check

		// HAVE TO PHRASE AS NOT KNOWNS

		// while we have more than one unknown, pick 1
		n, i := c.countUnknown(known)
		for n > 1 {
			fmt.Printf("CHOOSE %q\n", c[i])
			known[c[i]] = struct{}{}
			n, i = c.countUnknown(known)
		}
		// if we have 1 unknown, compute it
		if n == 1 {
			fmt.Printf("COMPUTE %q\n", c[i]) // TODO formula!
			fmt.Printf("DUPE_CHECK %q\n", c[i]) // TODO formula!
			known[c[i]] = struct{}{}
			fmt.Printf("COMPUTE c\n" // TODO forumula!
		} else {
			fmt.Printf("CHECK\n")
		}
		// otherwise check

	}
}

func main() {
	fmt.Println("Let's solve some problems!")
}
