package schema

func config() {
	MONGO_HOST = setStringConfig("MONGO_HOST", MONGO_HOST)
	MONGO_PORT = setIntConfig("MONGO_PORT", MONGO_PORT)
	MONGO_DBNAME = setStringConfig("MONGO_DBNAME", MONGO_DBNAME)
	MONGO_PROTOCOL = setStringConfig("MONGO_PROTOCOL", MONGO_PROTOCOL)
}
