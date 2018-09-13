package main

import (
	"github.com/liuzwei/npcs"
	"fmt"
)

func main() {

	mob := npcs.NonPlayerCharacter{Name:"Sper1"}

	fmt.Println(mob)

	hortense := npcs.NonPlayerCharacter{Name:"Hortense",Loc:npcs.Location{X:1,Y:1,Z:1}}

	fmt.Println(hortense)

	fmt.Printf("Mob is %f units from Hortense", mob.DistanceTo(hortense))
}

