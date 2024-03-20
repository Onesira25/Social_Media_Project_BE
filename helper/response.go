package helper

// var message = map[int]string{
// 	200: "success",
// 	201: "the data sent was successfully registered",
// 	400: "the data sent is incorrect",
// 	401: "invalid or expired jwt",
// 	500: "an error occurred in the server process",
// }

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
