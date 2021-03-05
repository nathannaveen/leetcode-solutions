package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	oneCounter := 0
	twoCounter := 0
	arr := []int{}

	for oneCounter != len(nums1) || twoCounter != len(nums2) {

		if (oneCounter != len(nums1) && twoCounter != len(nums2)) && (nums1[oneCounter] < 0 || nums2[twoCounter] < 0) {
			if nums1[oneCounter] >= nums2[twoCounter] {
				arr = append(arr, nums1[oneCounter])
				oneCounter++
			} else if nums2[twoCounter] >= nums1[oneCounter] {
				arr = append(arr, nums2[twoCounter])
				twoCounter++
			}
			continue
		}
		if (oneCounter != len(nums1)) &&
			((twoCounter != len(nums2) && nums1[oneCounter] <= nums2[twoCounter]) || (twoCounter <= len(nums2))) {
			arr = append(arr, nums1[oneCounter])
			oneCounter++
		}
		if (twoCounter != len(nums2)) &&
			((oneCounter != len(nums1) && nums2[twoCounter] <= nums1[oneCounter]) || (oneCounter <= len(nums2))) {
			arr = append(arr, nums2[twoCounter])
			twoCounter++
		}
	}
	if len(arr)%2 == 0 {
		return float64(arr[len(arr)/2-1]+arr[len(arr)/2]) / 2
	} else {
		return float64(arr[len(arr)/2])
	}
}
