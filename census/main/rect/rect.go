package rect

type rect struct {
	width, height int
}

func (r *rect) squareIt() {
	r.height = r.width
}

// r := rect{width: 10, height: 20}
// fmt.Printf("Width: %d, Height: %d\n", r.width, r.height)
// // Output: Width: 10, Height: 20

// r.squareIt()
// fmt.Printf("Width: %d, Height: %d\n", r.width, r.height)
// // Output: Width: 10, Height: 10
