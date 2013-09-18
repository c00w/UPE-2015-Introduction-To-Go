package geometry
// OMIT
import "fmt" //OMIT

type Point struct {
	X float32
	Y float32
}
func (point Point) Add(otherPoint Point) (sumPoint Point) {
	sumPoint = Point{point.X+otherPoint.X, point.Y+otherPoint.Y}
	return
}
type Line struct {
	Slope     float32
	Intercept float32
}
func (line Line) String() (rend string) {
	return fmt.Sprintf("y = %2fx + %2f", line.Slope, line.Intercept)
}
func NewLine(pointA Point, pointB Point) (line Line){
	line = Line{}
	line.Slope = (pointA.Y - pointB.Y) / (pointA.X - pointB.X)
	line.Intercept = pointA.Y - line.Slope*pointA.X
	return
}

