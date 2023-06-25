package head_row

type HeadRow struct {
	Index int
	Value string
}

type HeadRowManager struct {
	headRows   []HeadRow
	headRowMap map[string]int
}

func InitHeadRowManager(headRows []HeadRow) *HeadRowManager {
	headRowManager := HeadRowManager{}
	headRowManager.headRows = headRows
	headRowManager.headRowMap = map[string]int{}
	for i, headRow := range headRows {
		headRowManager.headRowMap[headRow.Value] = i
	}
	return &headRowManager
}

func (headRowManager HeadRowManager) GetHeadRowByValue(value string) (HeadRow, bool) {
	i, exists := headRowManager.headRowMap[value]
	var headRow HeadRow
	if exists {
		headRow = headRowManager.headRows[i]
	}
	return headRow, exists
}

func (headRowManager HeadRowManager) GetNextHeadRow(headRow HeadRow) (HeadRow, bool) {
	i, _ := headRowManager.headRowMap[headRow.Value]
	nextI := i + 1
	var nextHeadRow HeadRow
	var exists bool
	if nextI < len(headRowManager.headRows) {
		nextHeadRow = headRowManager.headRows[nextI]
		exists = true
	}
	return nextHeadRow, exists
}
