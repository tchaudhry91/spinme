package spin_test

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/tchaudhry91/spinme/spin"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/go-redis/redis"
)

func ExamplePostgres() {
	out, err := spin.Postgres(context.Background(), nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer spin.SlashID(context.Background(), out.ID)
	// Give postgres a few seconds to boot-up, sadly there is no "ready" check yet
	time.Sleep(5 * time.Second)
	connStr, err := spin.PostgresConnString(out)
	if err != nil {
		fmt.Println(err)
		return
	}
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Connected!")
	//Output: Connected!
}

func ExampleMySQL() {
	out, err := spin.MySQL(context.Background(), nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer spin.SlashID(context.Background(), out.ID)
	// Give mysql a minute to boot-up, sadly there is no "ready" check yet
	time.Sleep(1 * time.Minute)
	connStr, err := spin.MySQLConnString(out)
	if err != nil {
		fmt.Println(err)
		return
	}
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Connected!")
	//Output: Connected!
}

func ExampleMongo() {
	out, err := spin.Mongo(context.Background(), nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer spin.SlashID(context.Background(), out.ID)
	// Give mongo a few seconds to boot-up, sadly there is no "ready" check yet
	time.Sleep(10 * time.Second)
	connStr, err := spin.MongoConnString(out)
	if err != nil {
		fmt.Println(err)
		return
	}
	client, err := mongo.NewClient(options.Client().ApplyURI(connStr))
	if err != nil {
		fmt.Println(err)
		return
	}
	err = client.Connect(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer client.Disconnect(context.Background())
	err = client.Ping(context.Background(), nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Connected!")
	// Output: Connected!
}

func ExampleRedis() {
	out, err := spin.Redis(context.Background(), nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer spin.SlashID(context.Background(), out.ID)
	// Give redis a few seconds to boot-up, sadly there is no "ready" check yet
	time.Sleep(10 * time.Second)
	connStr, err := spin.RedisConnString(out)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := redis.NewClient(&redis.Options{
		Addr: connStr,
	})
	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(pong)
	// Output: PONG
}
