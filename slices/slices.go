package slices

import "fmt"

var moviesInsideDb = []string{
	"movie1",
	"movie2",
	"movie3",
	"movie4",
	"movie5",
	"movie6",
	"movie7",
	"movie8",
	"movie9",
	"movie10",
}

func main() {
	resultsFromApi := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	matrix := [][]int{}     // 2D
	matrix3D := [][][]int{} // 3D

	movies := make([]string, 0, 10)

	for _, id := range resultsFromApi {
		movie := moviesInsideDb[id]
		movies = append(movies, movie)
	}

	fmt.Println(movies)
}
