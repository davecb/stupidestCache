package main

// Exercise -- a program to exercise the simplest/stupidest possible cache for a
//	particular task.

// main -- get options and commands
func main() {

}

func exercise() {
	var key, value string
	var c = cache.New()

	for {
		line, eof := xxx.Read()
		if eof {
			break
		}
		switch operation {
		case "r":
			x, present = c.Get(key)
			if x != operand {
				error
			}
		case "w":
			err = c.Put(key, value)
			if err != nil {
				error
			}
		default:
			//illformed 			line
		}

	}
	// all done
}
