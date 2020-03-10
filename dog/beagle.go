package dog

import (
	"fmt"
	"github.com/jfrog-solutiontest/food"
)

func BeaglePackageName () {
	fmt.Println ("Package Name: beagle")
	ColliePackageName()
	food.PurinaPackageName()
}
