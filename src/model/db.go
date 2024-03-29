package model

import (
	"log"
	"os"
	"time"

	"github.com/asterich/CV-analyzer-backend/src/utils"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var newLogger = logger.New(
	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	logger.Config{
		SlowThreshold:             time.Second, // Slow SQL threshold
		LogLevel:                  logger.Info, // Log level
		IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
		ParameterizedQueries:      true,        // Don't include params in the SQL log
		Colorful:                  true,        // Disable color
	},
)

var Db, err = gorm.Open(sqlite.Open(utils.DbPath), &gorm.Config{
	NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
	},
	Logger: newLogger,
})

func InitDb() {
	if err != nil {
		log.Fatalln("Failed to connect database! err: ", err.Error())
	}

	var sqlDB, err1 = Db.DB()
	if err1 != nil {
		log.Fatalln("Failed to get *sql.DB, err: ", err1.Error())
	}

	Db.AutoMigrate(
		&CV{},
		&Education{},
		&WorkExperience{},
		&SchoolExperience{},
		&InternshipExperience{},
		&ProjectExperience{},
		&Award{},
		&Skill{},

		&Position{},
	)
	if err != nil {
		log.Fatalln("Failed to build many2many associations, err: ", err.Error())
	}

	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second)

	//	sqlDB.Close()

}
