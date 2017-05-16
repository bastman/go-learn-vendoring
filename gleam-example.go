package main
import (
	"fmt"
	//"github.com/yanatan16/itertools"
	"github.com/chrislusf/glow/flow"
	//"github.com/chrislusf/gleam/gio"
	//"os"
	//"github.com/chrislusf/gleam/distributed"
	"flag"
	"strings"
)


func main() {
	fmt.Println(1)

	flag.Parse()

	flow.New().TextFile(
		"/etc/passwd", 3,
	).Filter(func(line string) bool {
		return !strings.HasPrefix(line, "#")
	}).Map(func(line string, ch chan string) {
		for _, token := range strings.Split(line, ":") {
			ch <- token
		}
	}).Map(func(key string) int {
		return 1
	}).Reduce(func(x int, y int) int {
		return x + y
	}).Map(func(x int) {
		println("count:", x)
	}).Run()

}
