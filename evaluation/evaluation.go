package evaluation

import (
	//"fmt"
	"math/rand"
)


func Shuffle(from []string) {
	rows := len(from)
	for i:=0; i < rows; i++ {
		j := rand.Intn(i + 1)
		from[i], from[j] = from[j], from[i]
	}
	return from
}

/*
func InstancesTrainTestSplit(src FixedDataGrid, prop float64) (FixedDataGrid, FixedDataGrid) {
	trainingRows := make([]int, 0)
	testingRows := make([]int, 0)
	src = Shuffle(src)

	// Create the return structure
	_, rows := src.Size()
	for i := 0; i < rows; i++ {
		trainOrTest := rand.Intn(101)
		if trainOrTest > int(100*prop) {
			trainingRows = append(trainingRows, i)
		} else {
			testingRows = append(testingRows, i)
		}
	}

	allAttrs := src.AllAttributes()

	return NewInstancesViewFromVisible(src, trainingRows, allAttrs), NewInstancesViewFromVisible(src, testingRows, allAttrs)

}
*/