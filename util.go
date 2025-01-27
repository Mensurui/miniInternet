package miniinternet

func SsplitIntoChunks(data string, chunkSize int) []string {
	var chunks []string
	for i := 0; i < len(data); i += chunkSize {
		end := i + chunkSize
		if end > len(data) {
			end = len(data)
		}
		chunks = append(chunks, data[i:end])
	}
	return chunks
}
