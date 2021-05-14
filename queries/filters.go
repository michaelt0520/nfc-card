package queries

func BuildQueries(data map[string]interface{}) map[string]interface{} {
	_filters := make(map[string]interface{})

	if data["name"] != nil && data["name"] != "" {
		_filters["name like ?"] = "%" + data["name"].(string) + "%"
	}

	if data["email"] != nil && data["email"] != "" {
		_filters["email = ?"] = "%" + data["email"].(string) + "%"
	}

	if data["phone_number"] != nil && data["phone_number"] != "" {
		_filters["phone_number like ?"] = "%" + data["phone_number"].(string) + "%"
	}

	if data["address"] != nil && data["address"] != "" {
		_filters["address like ?"] = "%" + data["address"].(string) + "%"
	}

	if data["code"] != nil && data["code"] != "" {
		_filters["code = ?"] = "%" + data["code"].(string) + "%"
	}

	if data["username"] != nil && data["username"] != "" {
		_filters["username = ?"] = "%" + data["username"].(string) + "%"
	}

	if data["company_id"] != nil && data["company_id"] != "" {
		_filters["company_id = ?"] = data["company_id"]
	}

	if data["user_id"] != nil {
		_filters["user_id = ?"] = data["user_id"]
	}

	if data["type"] != nil && data["type"] != "" {
		_filters["type = ?"] = data["type"]
	}

	if data["role"] != nil && data["addresroles"] != "" {
		_filters["role = ?"] = data["role"]
	}

	if data["activated"] != nil {
		_filters["activated = ?"] = data["activated"]
	}

	return _filters
}

func BuildFinds(data map[string]interface{}) map[string]interface{} {
	_filters := make(map[string]interface{})

	if data["id"] != nil {
		_filters["id = ?"] = data["id"]
	}

	if data["code"] != nil {
		_filters["code = ?"] = data["code"]
	}

	if data["name"] != nil {
		_filters["name = ?"] = data["name"]
	}

	if data["email"] != nil {
		_filters["email = ?"] = data["email"]
	}

	if data["phone_number"] != nil {
		_filters["phone_number = ?"] = data["phone_number"]
	}

	if data["username"] != nil {
		_filters["username = ?"] = data["username"]
	}

	if data["company_id"] != nil {
		_filters["company_id = ?"] = data["company_id"]
	}

	if data["user_id"] != nil {
		_filters["user_id = ?"] = data["user_id"]
	}

	if data["card_id"] != nil {
		_filters["card_id = ?"] = data["card_id"]
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
