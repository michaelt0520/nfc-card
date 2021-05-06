package queries

func BuildUserQueries(data map[string]interface{}) map[string]interface{} {
	_filters := make(map[string]interface{})

	if data["name"] != nil && data["name"] != "" {
		_filters["name like ?"] = "%" + data["name"].(string) + "%"
	}

	if data["email"] != nil && data["email"] != "" {
		_filters["email like ?"] = "%" + data["email"].(string) + "%"
	}

	if data["phone_number"] != nil && data["phone_number"] != "" {
		_filters["phone_number like ?"] = "%" + data["phone_number"].(string) + "%"
	}

	if data["address"] != nil && data["address"] != "" {
		_filters["address like ?"] = "%" + data["address"].(string) + "%"
	}

	if data["username"] != nil {
		_filters["username = ?"] = data["username"]
	}

	if data["company_id"] != nil {
		_filters["company_id = ?"] = data["company_id"]
	}

	if data["type"] != nil {
		_filters["type = ?"] = data["type"]
	}

	if data["role"] != nil {
		_filters["role = ?"] = data["role"]
	}

	if data["jwt"] != nil {
		_filters["jwt = ?"] = data["jwt"]
	}

	return _filters
}

func BuildCardQueries(data map[string]interface{}) map[string]interface{} {
	_filters := make(map[string]interface{})

	if data["code"] != nil && data["code"] != "" {
		_filters["code = ?"] = data["code"]
	}

	if data["user_id"] != nil {
		_filters["user_id = ?"] = data["user_id"]
	}

	if data["company_id"] != nil {
		_filters["company_id = ?"] = data["company_id"]
	}

	if data["activated"] != nil {
		_filters["activated = ?"] = data["activated"]
	}

	return _filters
}
