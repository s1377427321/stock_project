package main

import "common"

type SearchMatrixS struct {
	saveMatrix []common.IntSlice
}

func (this *SearchMatrixS) searchMatrix(target int) (isExist bool) {
	rowSize := len(this.saveMatrix)
	if rowSize < 1 {
		return false
	}

	colSize := len(this.saveMatrix[rowSize-1])

	if target < this.saveMatrix[0][0] || target > this.saveMatrix[rowSize-1][colSize-1] {
		return false
	}

	var rowIndex int

	rowHigh := rowSize - 1
	rowLow := 0
	rowMid := int((rowHigh + rowLow) / 2)

	for ; rowLow <= rowHigh; {
		colSize := len(this.saveMatrix[rowMid]) - 1
		if this.saveMatrix[rowMid][0] == target || this.saveMatrix[rowMid][colSize] == target {
			return true
		} else if this.saveMatrix[rowMid][0] < target && this.saveMatrix[rowMid][colSize] > target {
			rowIndex = rowMid
			break
		} else if this.saveMatrix[rowMid][0] > target {
			rowHigh = rowMid - 1
			rowMid = (rowLow + rowHigh) / 2
		}else if this.saveMatrix[rowMid][0] < target{
			rowLow = rowMid +1
			rowMid = (rowLow+rowHigh) /2
		}
	}

	colHigh := len(this.saveMatrix[rowIndex]) -1
	colLow := 0
	colMid := int((colHigh + colLow) / 2)

	for ;colLow<=colHigh; {
		if this.saveMatrix[rowIndex][colMid] == target {
			return  true
		}else if this.saveMatrix[rowIndex][colMid] < target {
			colLow = colMid+1
			colMid  = (colLow + colHigh) /2
		}else if this.saveMatrix[rowIndex][colMid] > target {
			colHigh = colMid-1
			colMid  = (colLow + colHigh) /2
		}
	}

	return  false
}
