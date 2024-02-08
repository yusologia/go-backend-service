package Parser

func IDNameParser(id uint, name string) map[string]interface{} {
	if id == 0 {
		return nil
	}

	return map[string]interface{}{
		"id":   id,
		"name": name,
	}
}
