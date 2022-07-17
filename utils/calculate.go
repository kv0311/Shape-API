package utils

func CalculatePerimeterRectangle(width int, length int) (perimeter int) {
	return (width + length) * 2
}

func CalculateAreaRectangle(width int, length int) (area int) {
	return width * length
}

func CalculatePerimeterSquare(side int) (perimeter int) {
	return side * 4
}

func CalculateAreaSquare(side int) (area int) {
	return side * side
}

func CalculatePerimeterDiamond(side int) (perimeter int) {
	return side * 4
}

func CalculateAreaDiamond(diagonal1 int, diagonal2 int) (area int) {
	return diagonal1 * diagonal2
}

func CalculatePerimeterTriangle(side1 int, side2 int, base int) (perimeter int) {
	return side1 + side2 + base
}

func CalculateAreaTriangle(base int, height int) (area int) {
	return base * height
}
