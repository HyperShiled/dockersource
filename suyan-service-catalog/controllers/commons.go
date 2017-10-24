package controllers

func CreateErrorData(err error) map[string]string {
	data := make(map[string]string)
	data["Error"] = err.Error()
	return data
}
