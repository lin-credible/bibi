package main

import "fmt"

/**
 * 1. https://stackoverflow.com/questions/68042972/why-append-slice-in-go-change-the-original-slice
 * 2. https://go.dev/blog/slices
 *
 * A slice is a data structure describing a contiguous section of an array stored separately from the slice variable
 * itself. A slice is not an array. A slice describes a piece of an array.
 */

func main() {
	c1()
	c2()
}

func c1() {
	chunks := []uint32{5, 4, 4, 4, 4}
	fmt.Println("groups: ", chunks)
	ids := []uint32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21}
	fmt.Println("init ids len: ", len(ids))
	fmt.Println("init ids: ", ids)
	startIdx := 0
	endIdx := 0
	pp := make([][]uint32, 0)
	for i := 0; i < len(chunks); i++ {
		endIdx += int(chunks[i])
		if len(ids) < endIdx {
			endIdx = len(ids)
		}
		pp = append(pp, ids[startIdx:endIdx])
		startIdx = endIdx
	}
	fmt.Println("\n---------- GROUP BEFORE --------------------------------------")
	fmt.Println(pp)
	fmt.Println("---------- GROUP BEFORE --------------------------------------")
	ppp := make([][]uint32, 0)
	for _, item := range pp {
		eee(item)
		ppp = append(ppp, item)
	}
	fmt.Println("\n---------- GROUP AFTER ---------------------------------------")
	fmt.Println(ppp)
	fmt.Println("after ids: ", ids)
	fmt.Println("---------- GROUP AFTER ---------------------------------------")
}

func eee(ids []uint32) {
	ids = append(ids, 0)
	fmt.Println(ids)
}

func c2() {
	chunks := []uint32{5, 4, 4, 4, 4}
	fmt.Println(chunks)
	ids := []uint32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21}
	startIdx := 0
	endIdx := 0
	pp := make([][]uint32, 0)
	for i := 0; i < len(chunks); i++ {
		endIdx += int(chunks[i])
		if len(ids) < endIdx {
			endIdx = len(ids)
		}
		list := make([]uint32, endIdx-startIdx)
		copy(list, ids[startIdx:endIdx])
		pp = append(pp, list)
		// pp = append(pp, ids[startIdx:endIdx])
		startIdx = endIdx
	}
	fmt.Println("\n++++++++++ GROUP BEFORE +++++++++++++++++++++++++++++++++++++++")
	fmt.Println(pp)
	fmt.Println("++++++++++ GROUP BEFORE +++++++++++++++++++++++++++++++++++++++")
	ppp := make([][]uint32, 0)
	for _, item := range pp {
		fmt.Println("___ item: ", item)
		// ...... others ......
		eee(item)
		ppp = append(ppp, item)
	}
	fmt.Println("\n++++++++++ GROUP AFTER +++++++++++++++++++++++++++++++++++++++")
	fmt.Println(ppp)
	fmt.Println("++++++++++ GROUP AFTER +++++++++++++++++++++++++++++++++++++++")
}
