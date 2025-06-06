package structpages

import (
	"fmt"
	"testing"
)

type (
	TopPage struct {
		Level1 `route:"/level1 Level 1"`
		Level2 `route:"GET /level2 Level 2"`
	}
	Level1 struct {
		Level2 `route:"POST /level2 true level 2"`
	}
	Level2 struct{}
)

func (l *Level1) Page() component    { return testComponent{} }
func (l *Level1) Partial() component { return testComponent{} }
func (l *Level2) Page() component    { return testComponent{} }
func (l *Level2) Partial() component { return testComponent{} }

func TestPrint(t *testing.T) {
	s := PrintRoutes(&TopPage{}, "/", "Top Page")
	fmt.Println(s)
}
