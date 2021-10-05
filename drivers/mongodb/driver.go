package mongodb

// import (
// 	"log"

// 	"gopkg.in/mgo.v2"
// )

// type ConfigMongo struct {
// 	// DB_Username string
// 	// DB_Password string
// 	DB_Host string
// 	// DB_Port     string
// 	// DB_Database string
// }

// // func (config *ConfigDB) InitialDB() *gorm.DB {
// // 	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
// // 		config.DB_Username,
// // 		config.DB_Password,
// // 		config.DB_Host,
// // 		config.DB_Port,
// // 		config.DB_Database)

// // 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// // 	if err != nil {
// // 		log.Fatal(err)
// // 	}

// // 	return db
// // }

// func (config *ConfigMongo) InitMongo() *mgo.Session {
// 	session, err := mgo.Dial(config.DB_Host)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// err = session.DB(database).C(collection).Insert(user)
// 	return session
// };
