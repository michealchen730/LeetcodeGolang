package main

import "fmt"

type MountainArray struct{}

func (this *MountainArray) get(index int) int { return 0 }
func (this *MountainArray) length() int       { return 0 }
func findInMountainArray(target int, mountainArr *MountainArray) int {
	leng := mountainArr.length()
	low, high := 0, leng-1
	if target < mountainArr.get(low) && target < mountainArr.get(high) {
		return -1
	}
	peakindex, peakheight := -1, -1
	for low < high-1 {
		mid := low + (high-low)/2
		left, right, inter := mountainArr.get(mid-1), mountainArr.get(mid+1), mountainArr.get(mid)
		if inter > left && inter < right {
			peakindex, peakheight = mid, inter
			break
		} else if inter > left {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	if peakindex == -1 {
		if mountainArr.get(low) > mountainArr.get(high) {
			peakindex, peakheight = low, mountainArr.get(low)
		} else {
			peakindex, peakheight = high, mountainArr.get(high)
		}
	}
	if peakheight < target {
		return -1
	} else if peakheight == target {
		return peakindex
	} else {
		i, j := 0, peakindex-1
		if target <= mountainArr.get(j) {
			for i <= j {
				m := i + (j-i)/2
				fmt.Println(m)
				if mountainArr.get(m) == target {
					return m
				} else if mountainArr.get(m) > target {
					j = m - 1
				} else {
					i = m + 1
				}
			}
		}
		i, j = peakindex+1, leng-1
		if target <= mountainArr.get(i) {
			for i <= j {
				m := i + (j-i)/2
				if mountainArr.get(m) == target {
					return m
				} else if mountainArr.get(m) > target {
					i = m + 1
				} else {
					j = m - 1
				}
			}
		}
	}
	return -1
}
