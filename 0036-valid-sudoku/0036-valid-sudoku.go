func isValidSudoku(board [][]byte) bool {
	uniqueRows := make(map[int]map[string]bool)
	boxes := make(map[string]map[string]bool)
	for i := 0; i < len(board); {
		for j := 0; j < len(board[i]); {
			if board[i][j] != '.' {
				if _, exists := uniqueRows[i][string(board[i][j])]; exists {
					return false
				}
				if uniqueRows[i] == nil {
					uniqueRows[i] = make(map[string]bool)
				}
				uniqueRows[i][string(board[i][j])] = true
			}

			j++
		}
		columns := make(map[string]bool)
		for j := 0; j < len(board[i]); {
			if board[j][i] != '.' {
				str := string(board[j][i])

				if _, ok := columns[str]; ok {
					return false
				}
				columns[str] = true

			}
			j++
		}

		if i%3 == 0 {
			for boxCol := 0; boxCol < 3; boxCol++ {
				startRow := i
				startCol := boxCol * 3

				for row := startRow; row < startRow+3; row++ {
					for col := startCol; col < startCol+3; col++ {
						if board[row][col] == '.' {
							continue
						}

						boxIndex := fmt.Sprintf("%d%d", row/3, col/3)
						value := string(board[row][col])
						if boxes[boxIndex] == nil {
							boxes[boxIndex] = make(map[string]bool)
						}
						if _, ok := boxes[boxIndex][value]; ok {
							log.Println("no it is me")
							return false
						}

						boxes[boxIndex][value] = true

					}
				}
			}
		}

		i++
	}

	return true
}