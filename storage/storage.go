package storage


var urlStore = make(map[string]string)
func Save(shortCode, originalURL string) {
	urlStore[shortCode] = originalURL
}
func Save(shortCode, originalURL string) {
	urlStore[shortCode] = originalURL
}