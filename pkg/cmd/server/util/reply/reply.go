package reply

type Payload = map[string]any

func JSON(success bool, message string, data map[string]any) map[string]any {
	return map[string]any{
		"success": success,
		"message": message,
		"data":    data,
	}
}
