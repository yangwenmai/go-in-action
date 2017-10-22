package example

import "testing"
import . "github.com/smartystreets/goconvey/convey"
import "fmt"

func TestAdd(t *testing.T) {
	Convey("test add", t, func() {
		a := 3
		b := 4
		c := Add(a, b)
		So(c, ShouldEqual, 7)
	})
}

func Example() {
	sum := Add(3, 4)
	fmt.Println("3+4=", sum)

	//Output:
	//3+4= 7
}
