package medium

import (
	"math"
	"slices"
)

// 74. Search a 2D Matrix
func SearchMatrix(matrix [][]int, target int) bool {
    sr, er := 0, len(matrix)-1
    tr := -1
    for sr <= er {
        mr := (sr + er) / 2
        // fmt.Printf("sr=%v, mr=%v, er=%v\n", sr, mr, er)
        mrow := matrix[mr]
        if mrow[0] <= target && target <= mrow[len(mrow)-1] {
            tr = mr
            break
        }
        if target < mrow[0] {
            er = mr - 1
        } else {
            sr = mr + 1
        }
    }

    if tr >= 0 && binarySearch(matrix[tr], target) >= 0 {
        return true
    }
    return false
}

func binarySearch(arr []int, t int) int {
    s, e := 0, len(arr)-1
    for s <= e {
        m := (s + e) / 2
        if arr[m] == t {
            return m
        }
        if e - s <= 0 {
            return -1
        }
        if arr[m] < t {
            s = m + 1
        } else {
            e = m - 1
        }
    }
    return -1
}

// 875. Koko Eating Bananas
func MinEatingSpeed(piles []int, h int) int {
    canEatUp := func(k int) bool {
        th := 0
        for _, p := range piles {
            th += int(math.Ceil(float64(p) / float64(k)))
        }
        return th <= h
    }

    // binary search
    s, e := 1, slices.Max(piles)
    var k int
    for s <= e {
        m := (e + s) / 2
        if canEatUp(m) {
            // if can eat up in time update k
            k = m
            // continue into left subarray (decrease eating speed)
            e = m-1
        } else {
            // otherwise to right subarray (increase eating speed)
            s = m+1
        }
    }

    return k
}

// 153. Find Minimum in Rotated Sorted Array
func FindMin(nums []int) int {
    l, r := 0, len(nums)-1
    mn := nums[0]
    for l <= r {
        if nums[l] < nums[r] {
            return min(mn, nums[l])
        }
        m := (l + r) / 2
        mn = min(mn, nums[m])
        if nums[l] <= nums[m] {
            // m in rotated subarray, search right
            l = m + 1
        } else {
            // m in original subarray, search left
            r = m - 1
        }
    }
    return mn
}

func FindMinFast(nums []int) int {
    l, r, m, result := 0, len(nums) - 1, 0, nums[0]

	for l <= r {
        m = (l + r) / 2
		if result > nums[m] {
			result = nums[m]
		}

        if nums[l] < nums[m] && nums[l] < nums[r] || nums[l] > nums[m] && nums[l] > nums[r] { 
            // ascending subarr, or mid on not rotated side
            r = m - 1
        } else { 
            // mid on rotated side
            l = m + 1
        }
	}
    return result
}

// 33. Search in Rotated Sorted Array
func Search(nums []int, target int) int {
    l, r, m := 0, len(nums) - 1, 0

    for l <= r {
        m = (l + r) / 2
        if nums[m] == target {
            return m
        }

        if nums[l] <= nums[m] {
            // left portion
            if target > nums[m] {
                // left portion, t > m: right
                l = m + 1
            } else {
                // left portion, t < m:
                if target >= nums[l] {
                    // t >= l: left
                    r = m - 1
                } else {
                    // else: right
                    l = m + 1
                }
            }
        } else {
            // not rotated
            if target < nums[m] {
                // right portion, t < m: left
                r = m - 1
            } else {
                // right portion, t > m: 
                if target <= nums[r] {
                    // t <= r: right
                    l = m + 1
                } else {
                    // else: left
                    r = m - 1
                }
            }
        }
    }

    return -1
}

// 981. Time Based Key-Value Store
/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
type TimeItem struct {
    ts int
    data string
}
type TimeMap struct {
    m map[string][]TimeItem
}

func TMapConstructor() TimeMap {
    return TimeMap{
        m: make(map[string][]TimeItem),
    }
}

func (this *TimeMap) Set(key string, value string, timestamp int)  {
    vals := this.m[key]
    if vals == nil {
        vals = []TimeItem{}
    } 
    this.m[key] = append(vals, TimeItem{
        ts: timestamp,
        data: value,
    })
}

func (this *TimeMap) Get(key string, timestamp int) string {
    vals := this.m[key]
    if vals == nil {
        return ""
    }

    l, r, m := 0, len(vals)-1, 0
    for l <= r {
        m = (l + r) / 2
        vm := vals[m]
        if vm.ts == timestamp {
            return vm.data
        }
        if l == r {
            if vm.ts < timestamp {
                return vm.data
            } else {
                if m-1 >= 0 {
                    return vals[m-1].data
                } else {
                    return ""
                }
            }
        }
        if timestamp < vm.ts {
            r = m - 1
        } else {
            l = m + 1
        }
    }

    return ""
}

// 4. Median of Two Sorted Arrays
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    m := int(math.Ceil(float64(len(nums1)) + float64(len(nums2))) / 2)
}
