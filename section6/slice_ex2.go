package section6

import (
	"fmt"
	"sort"
)

func SliceEx2() {

	fmt.Println("======================= 슬라이스 일부 추출")

	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("slice1[:] : ", slice1[:])
	fmt.Println("slice1[0:] : ", slice1[0:])
	fmt.Println("slice1[:5] : ", slice1[:5])
	fmt.Println("slice1[0:len(slice1)] : ", slice1[0:len(slice1)])
	fmt.Println("slice1[3:] : ", slice1[3:])
	fmt.Println("slice1[:3] : ", slice1[:3])
	fmt.Println("slice1[1:3] : ", slice1[1:3])

	fmt.Println("======================= 슬라이스 정렬")

	slice2 := []int{3, 100, 20, 5, 7, 1, 4, 31}
	slice3 := []string{"b", "Bc", "Ba", "x", "100", "Y", "10"}

	fmt.Println("slice2 - sort.IntsAreSorted() : ", sort.IntsAreSorted(slice2))
	sort.Ints(slice2) // = sort.Sort(sort.IntSlice(slice2))
	fmt.Println("Asc Sorted slice2 : ", slice2)
	sort.Sort(sort.Reverse(sort.IntSlice(slice2)))
	fmt.Println("Desc Sorted slice2 : ", slice2)

	fmt.Println("slice3 - sort", sort.StringsAreSorted(slice3))
	sort.Strings(slice3)
	fmt.Println("Asc Sorted slice3 : ", slice3)
	sort.Sort(sort.Reverse(sort.StringSlice(slice3)))
	fmt.Println("Desc Sorted slice3 : ", slice3)
}
