package Parser

func IDNameParser(id any, name string) map[string]interface{} {
	if id == 0 || id == "" {
		return nil
	}

	return map[string]interface{}{
		"id":   id,
		"name": name,
	}
}
