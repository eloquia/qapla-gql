package postgres

import (
	"qaplagql/graph/model"
)

type UserServiceImpl struct {
	Message    string
	DbUserName string
	DbPassword string
	DbURL      string
	DbName     string
}

func (p *UserServiceImpl) CreateUser(firstName string, lastName string, goesBy string, middleName string, email string, gender string, ethnicity string, position string, institution string, isActive bool) (*model.User, error) {
	// fmt.Printf("I have a message: %s\n", p.)
	panic("oops!")
}

func (p *UserServiceImpl) Initialize() error {
	// dbConnectionString := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", p.DbUserName, p.DbPassword, p.DbURL, p.DbName)
	// db, err := sql.Open("pgx", dbConnectionString)
	// if err != nil {
	// 	return err
	// }
	// driver, err := postgres.WithInstance(db, &postgres.Config{})
	// if err != nil {
	// 	return err
	// }
	// m, err := migrate.NewWithDatabaseInstance("file://../../postgres/migrations", p.DbName, driver)
	// if err != nil {
	// 	return err
	// }
	// err = m.Steps(2)
	// if err != nil {
	// 	return err
	// }
	return nil
}
