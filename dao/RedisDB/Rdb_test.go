package RedisDB

import (
	"fmt"
	"testing"
)

func TestRDB(t *testing.T) {
	InitRdb()
	conn := GetRedisConn()
	fmt.Println(conn)
	do, err := conn.Do("ping")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(do)
}
