package helper

func ResponseFormat(status int, message any, data ...any) (int, map[string]any) {
	result := map[string]any{
		"code":    status,
		"message": message,
	}

	if len(data) > 0 {
		result["data"] = data[0]
	}
	return status, result
}
