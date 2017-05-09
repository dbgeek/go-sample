package main

import "fmt"

func main() {
    slice := []int{1, 2, 3, 4, 5, 6, 6, 10, 15, 15, 17, 18, 19,30, 33, 39, 66, 70, 79, 83, 90, 95, 99, 111, 112, 115, 119, 125, 129}
    find := binarySearch(slice, 17, 0, len(slice)-1)
    fmt.Printf("Found the elemt by recursive binary search at index %d \n", find)
    find = iterBinarySearch(slice, 17, 0, len(slice)-1)
    fmt.Printf("Found the elemt by iterBinarySearch binary search at index %d \n", find)
}

func binarySearch(array []int, target int, lowIndex int, highIndex int) int {
    if highIndex < lowIndex {
        return -1
    }
    mid := int((lowIndex + highIndex) / 2)
    
    if array[mid] > target {
        return binarySearch(array, target, lowIndex, mid)
    } else if array[mid] < target {
        return binarySearch(array, target, mid + 1, highIndex)
    } else {
        return mid
    }
}

func iterBinarySearch(array []int, target int, lowIndex int, highIndex int) int {
    startIndex := lowIndex
    endIndex := highIndex
    var mid int
    for startIndex <= endIndex {
        mid = int((startIndex + endIndex) / 2)
        if array[mid] > target {
            endIndex = mid
        } else if array[mid] < target {
            startIndex = mid + 1
        } else {
            return mid
        }
    }
    return - 1
}