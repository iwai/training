package main

func main() {
	// 下記の「世界」の部分を「Go」に書き換えてください
	println("Hello, Go")
	println(2)
	println(2 + 3)
	println(4 - 7)
	println(1 + 3)
	println("1 + 3")
	println(3 * 5)
	println(6 / 2)
	println(9 % 2)
	println("Hi, " + "Gophers")
	println(3 + 5)
	println("3" + "5")

	var n int = 100
	println(n)
	n = 200
	println(n)

	// var a = 100
	// println(a)

	// b := 200
	// println(b)

	a := 100
	b := 200
	println(a, b)

	println(a + 10)
	println(a + b)

	a -= 10
	println(a)
	a++
	println(a)

	score := 100
	if score > 80 {
		println("よくできました")
	}

	x := 123
	y := 5 * 6
	if x >= 100 {
		println("よくできました")
	}
	if y < 40 {
		println("40より小さいです")
	}

	println(score > 80)
	println(score < 80)

	if score > 80 {
		println("Good")
	} else if score >= 60 {
		println("B+")
	} else {
		println("B")
	}

	if score != 77 {
		println(score != 77)
	}

	switch score {
	case 80:
		println(80)
	case 77:
		println(77)
	case 60:
		println(60)
	default:
		println("default")
	}

}
