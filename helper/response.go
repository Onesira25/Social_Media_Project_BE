package helper

// var message = map[int]string{
// 	200: "success",
// 	201: "the data sent was successfully registered",
// 	400: "the data sent is incorrect",
// 	401: "invalid or expired jwt",
// 	500: "an error occurred in the server process",
// }

func ResponseFormat(status int, message any, data ...map[string]any) (int, map[string]any) {
	result := map[string]any{
		"code":    status,
		"message": message,
	}

	for _, part := range data {
		for key, value := range part {
			result[key] = value
		}
	}
	return status, result
}
