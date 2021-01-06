package util

func SplitStringIntoChunks(s string, chunkSize int) []string {
	chunks := make([]string, 0, len(s)/chunkSize+1)

	for i := 0; i < len(s); i += chunkSize {
		chunks = append(chunks, s[i:Min(i+chunkSize, len(s))])
	}

	return chunks
}
