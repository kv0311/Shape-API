package utils

func BuildShapeKey(shapeType string, shapeInfo string, name string) (key string) {
	return (shapeType + "-" + shapeInfo + "-" + name)
}
