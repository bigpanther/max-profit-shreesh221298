package main

// maxProfit returns the maximum profit possible for the stock
func maxProfit(stockPricesByDay []int) int {
	//get first price
	//find the largest number
	//find a number smaller than first
	//check if it is before the largest
	//largest - smallest (if before largest)

	//set max profit as 0
	//loop through each element
	//make second loop which goes from 1st loop element onwards
	//everytime the difference is greater than max profit, change max profit to new difference
	//return max profit

	//consider first element as smallest
	//second element is largest
	//if next element is smaller than smallest, then make this element smallest
	//if this element is larger than largest, make it as largest
	//
	maxProfit := 0
	arrayLen := len(stockPricesByDay)

	if arrayLen != 0 {
		max := stockPricesByDay[arrayLen-1]
		for current := arrayLen - 1; current > 0; current-- {
			if stockPricesByDay[current] < 0 {
				maxProfit = 0
				break
			}
			if stockPricesByDay[current] >= max {
				max = stockPricesByDay[current]
				for next := current - 1; next >= 0; next-- {
					profit := stockPricesByDay[current] - stockPricesByDay[next]
					if profit > maxProfit {
						maxProfit = profit
					}

				}
			}

		}

	}

	return maxProfit

}

// func maxShreesh(stockPricesByDay []int) int {
// 	//set max profit as 0
// 	//loop through each element
// 	//make second loop which goes from 1st loop element onwards
// 	//everytime the difference is greater than max profit, change max profit to new difference
// 	//return max profit

// 	maxProfit := 0
// 	for current := 0; current < len(stockPricesByDay); current++ {
// 		if stockPricesByDay[current] < 0 {
// 			maxProfit = 0
// 			break
// 		}
// 		for next := current + 1; next < len(stockPricesByDay); next++ {
// 			profit := stockPricesByDay[next] - stockPricesByDay[current]
// 			if profit > maxProfit {
// 				maxProfit = profit
// 			}

// 		}
// 	}
// 	return maxProfit

// }
// func maxShreeshOpt(stockPricesByDay []int) int {
// 	//set max profit as 0
// 	//loop through each element
// 	//make second loop which goes from 1st loop element onwards
// 	//everytime the difference is greater than max profit, change max profit to new difference
// 	//return max profit

// 	maxProfit := 0
// 	// for current := 0; current < len(stockPricesByDay); current++ {
// 	// 	if stockPricesByDay[current] < 0 {
// 	// 		maxProfit = 0
// 	// 		break
// 	// 	}
// 	// 	for next := current + 1; next < len(stockPricesByDay); next++ {
// 	// 		profit := stockPricesByDay[next] - stockPricesByDay[current]
// 	// 		if profit > maxProfit {
// 	// 			maxProfit = profit
// 	// 		}

// 	// 	}
// 	// }
// 	arrayLen := len(stockPricesByDay)
// 	max := stockPricesByDay[arrayLen-1]
// 	for current := arrayLen - 1; current >= 0; current-- {
// 		if stockPricesByDay[current] < 0 {
// 			maxProfit = 0
// 			break
// 		}
// 		if stockPricesByDay[current] >= max {
// 			max = stockPricesByDay[current]
// 			for next := current - 1; next >= 0; next-- {
// 				profit := stockPricesByDay[current] - stockPricesByDay[next]
// 				if profit > maxProfit {
// 					maxProfit = profit
// 				}

// 			}
// 		}

// 	}
// 	return maxProfit

// }
