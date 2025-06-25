package mergesortedarray

func merge(nums1 []int, m int, nums2 []int, n int) {
	// set two pointers one for each list, set to the last
	// element of the respective list
	p1, p2 := m-1, n-1

	// loop backwards through nums1, set a pointer to the full end of nums1
	// p will be the pointer where we insert the final values into nums1
	for p := m + n - 1; p >= 0; p-- {
		// if out of nums2 values we need to break out of the loop
		if p2 < 0 {
			break
		}

		// if pointer on nums1 not out of values, and the value of nums1 is
		// greater than the value of nums2 we need to add nums1 value at the
		// p index and move p1 towards the first element of nums1
		if p1 >= 0 && nums1[p1] > nums2[p2] {
			nums1[p] = nums1[p1]
			p1 -= 1
		} else {
			nums1[p] = nums2[p2]
			p2 -= 1
		}
	}
}
