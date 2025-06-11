package dietplanperformance

func dietPlanPerformance(calories []int, k, lower, upper int) int {
	left, currSum, points := 0, 0, 0

	for right := 0; right < len(calories); right++ {
		currSum += calories[right]

		if right-left+1 == k {
			if currSum > upper {
				points++
			} else if currSum < lower {
				points--
			}
			currSum -= calories[left]
			left++
		}
	}
	return points
}

// first simple attempt, can refactor to one loop above
// func dietPlanPerformance(calories []int, k int, lower int, upper int) int {
// 	left, right, currSum, points := 0, 0, 0, 0

// 	for right < k {
// 		currSum += calories[right]
// 		right++
// 	}

// 	if currSum > upper {
// 		points++
// 	} else if currSum < lower {
// 		points--
// 	}

// 	for right < len(calories) {
// 		currSum += calories[right] - calories[left]
// 		if currSum > upper {
// 			points++
// 		} else if currSum < lower {
// 			points--
// 		}

// 		left++
// 		right++
// 	}

// 	return points
// }
