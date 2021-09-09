package office

import (
	"fmt"
)

type Office struct {
	OfficeName string
	MaxPackWt  int
	MinPackWt  int
	// TODO: list? of packages(uid, Wt)
}

func (e Employee) LeavesRemaining() {
	fmt.Printf("%s %s has %d leaves remaining\n", e.FirstName, e.LastName, (e.TotalLeaves - e.LeavesTaken))
}
