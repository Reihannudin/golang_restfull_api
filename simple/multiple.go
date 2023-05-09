package simple

type Database struct {
	Name string
}

type DatabaseMySQL Database
type DatabaseMongoDB Database

type DatabaseRepository struct {
	*DatabaseMySQL
	*DatabaseMongoDB
}

func NewDatabaseMySQL() *DatabaseMySQL {
	return (*DatabaseMySQL)(&Database{
		Name: "MySQL",
	})
}

func NewDatabaseMongoDB() *DatabaseMongoDB {
	return (*DatabaseMongoDB)(&Database{
		Name: "MongoDB",
	})
}

func NewDatabaseRepository(
	databaseMySQL *DatabaseMySQL,
	databaseMongoDB *DatabaseMongoDB,
) *DatabaseRepository {
	return &DatabaseRepository{
		DatabaseMySQL:   databaseMySQL,
		DatabaseMongoDB: databaseMongoDB,
	}
}
