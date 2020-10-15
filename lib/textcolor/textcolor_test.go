package textcolor

import (
	"fmt"
	"testing"
)

func TestBlack(t *testing.T) {
	fmt.Println(Black("Black"))
	fmt.Println(Red("Red"))
	fmt.Println(Green("Green"))
	fmt.Println(Yellow("Yellow"))
	fmt.Println(Blue("Blue"))
	fmt.Println(Magenta("Magenta"))
	fmt.Println(Cyan("Cyan"))
	fmt.Println(White("White"))
	fmt.Println(String("Red", 0, 0, TextRed))
}

func BenchmarkString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		String("Blue", 0, 0, TextBlue)
	}
}
