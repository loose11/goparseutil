package goparseutil

var mimeTypes = map[string]string{
	".jpg":  "image/jpeg",
	".jpeg": "image/jpeg",
	".txt":  "text/plain",
}

func GetMimeType(extension string) string {
	value, _ := mimeTypes[extension]
	return value
}
