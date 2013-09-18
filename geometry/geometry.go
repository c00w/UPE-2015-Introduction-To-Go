package geometry

import "fmt" //OMIT
// OMIT
type Point struct {
	X float32
	Y float32
}

func (point Point) Add(otherPoint Point) (sumPoint Point) {
	sumPoint = Point{point.X+otherPoint.X, point.Y+otherPoint.Y}
	return
}

type Line struct { //OMIT
	Slope     float32 //OMIT
	Intercept float32 //OMIT
} //OMIT
 //OMIT
func (line Line) String() (rend string) { //OMIT
	return fmt.Sprintf("y = %2fx + %2f", line.Slope, line.Intercept) //OMIT
} //OMIT
 //OMIT
func NewLine(pointA Point, pointB Point) (line Line){ //OMIT
	line = Line{} //OMIT
	line.Slope = (pointA.Y - pointB.Y) / (pointA.X - pointB.X) //OMIT
	line.Intercept = pointA.Y - line.Slope*pointA.X //OMIT
	return //OMIT
} //OMIT
