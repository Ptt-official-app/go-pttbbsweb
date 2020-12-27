package schema

func getBlock(theLen int, maxBlock int) int {
	if theLen < maxBlock {
		return theLen
	}

	return maxBlock
}
