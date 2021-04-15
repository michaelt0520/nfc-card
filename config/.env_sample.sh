# App
export app_host=http://localhost:$app_port

# Postgres
export db_host=localhost
export db_driver=postgres
export db_user="$(whoami)"
export db_password=
export db_name="nfc_card_${app_env}"
export db_port=5432

# JWT
export jwt_key=samplekey
