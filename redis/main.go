package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func main() {
	err := initClient()
	if err != nil {
		return
	}
	redisExample()
	redisExample2()
}

var rdb *redis.Client

func initClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, err = rdb.Ping().Result()
	if err != nil {
		fmt.Println("连接失败", err)
		return err
	}
	fmt.Println("连接成功")
	return nil
}
func redisExample() {
	err := rdb.Set("name", "zhangsan", 0).Err()
	if err != nil {
		fmt.Println("设置失败", err)
		return
	}
	val, err := rdb.Get("name").Result()
	if err != nil {
		fmt.Println("获取失败", err)
		return
	}
	fmt.Println("获取成功", val)
	val2, err := rdb.Get("name2").Result()
	if err == redis.Nil {
		fmt.Println("该记录不存在", err)
		return
	} else if err != nil {
		fmt.Println("获取失败", err)
		return
	} else {
		fmt.Println("获取成功", val2)
	}
}
func redisExample2() {
	fmt.Println("--------zset示例---------")
	zsetKey := "language_rank"
	languages := []redis.Z{
		redis.Z{Score: 90.0, Member: "golang"},
		redis.Z{Score: 80.0, Member: "python"},
		redis.Z{Score: 70.0, Member: "java"},
		redis.Z{Score: 60.0, Member: "c"},
	}
	num, err := rdb.ZAdd(zsetKey, languages...).Result()
	if err != nil {
		fmt.Println("添加失败", err)
		return
	}
	fmt.Println("添加成功", num)
	//	把golang分数加10
	newScore, err := rdb.ZIncrBy(zsetKey, 10.0, "golang").Result()
	if err != nil {
		fmt.Println("更新失败", err)
		return
	}
	fmt.Println("更新成功,golang新的分数:", newScore)
	//	取分数最高的3个
	res, err := rdb.ZRevRangeWithScores(zsetKey, 0, 2).Result()
	if err != nil {
		fmt.Println("取值失败", err)
		return
	}
	fmt.Println("取前3名成功", res)
	for _, z := range res {
		fmt.Println("每一项---", z.Member, z.Score)
	}
	//	取95-100
	op := redis.ZRangeBy{
		Min: "95",
		Max: "100",
	}
	res, err = rdb.ZRangeByScoreWithScores(zsetKey, op).Result()
	if err != nil {
		fmt.Println("取95-100值失败", err)
		return
	}
	fmt.Println("取95-100值成功", res)
	for _, z := range res {
		fmt.Println("每一项---", z.Member, z.Score)
	}
}
