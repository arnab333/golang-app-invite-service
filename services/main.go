package services

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/arnab333/golang-app-invite-service/helpers"
	"github.com/arnab333/golang-app-invite-service/models"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var mysqlOnce sync.Once

var redisOnce sync.Once

type dbConnection struct {
	DB *gorm.DB
}

var DBConn dbConnection

type redisConnection struct {
	redisClient *redis.Client
}

var redisConn redisConnection

func dsn() string {
	userName := os.Getenv(helpers.EnvKeys.MYSQL_USERNAME)
	password := os.Getenv(helpers.EnvKeys.MYSQL_PASSWORD)
	dbname := os.Getenv(helpers.EnvKeys.DBNAME)
	hostname := os.Getenv(helpers.EnvKeys.MYSQL_HOST_ADDR)
	port := os.Getenv(helpers.EnvKeys.MYSQL_HOST_PORT)

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", userName, password, hostname, port, dbname)
}

func InitDBConnection() func() {
	mysqlOnce.Do(func() {
		conn, err := gorm.Open(mysql.Open(dsn()), &gorm.Config{})

		if err != nil {
			panic("could not connect to the database")
		}

		fmt.Println("DB Connection Success")
		conn.AutoMigrate(&models.User{})
		DBConn.DB = conn
	})

	return func() {
		sqlDB, err := DBConn.DB.DB()
		sqlDB.Close()

		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
}

func InitRedisConnection() func() {
	redisOnce.Do(func() {
		dsn := os.Getenv(helpers.EnvKeys.REDIS_DSN)

		redisConn.redisClient = redis.NewClient(&redis.Options{
			Addr:     dsn,                                       //redis port
			Password: os.Getenv(helpers.EnvKeys.REDIS_PASSWORD), // no password set
			DB:       0,                                         // use default DB
		})
		result, err := redisConn.redisClient.Ping(context.Background()).Result()
		if err != nil {
			panic(err)
		}
		log.Println("redis ==>", result)
	})

	return func() {
		closeRedisConnection(redisConn.redisClient)
	}
}

func closeRedisConnection(client *redis.Client) {
	defer func() {
		if err := client.Close(); err != nil {
			log.Fatal("Close Error ==>", err)
		}
	}()
}
