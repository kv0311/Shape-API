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
