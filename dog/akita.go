package dog

import (
	"fmt"
        "github.com/jfrog-solutiontest/food"
)

func AkitaPackageName () {
	fmt.Println ("Package Name: akita")
	BeaglePackageName()
	ColliePackageName()
	food.AlpoPackageName()
}
