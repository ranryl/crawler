package utils

import (
	"crawler/conf"
	"fmt"
	"os"
	"testing"
)

var redisClient *CacherClient

type TestData struct {
	ID    int
	Name  string
	Woman bool
}

func init() {
	basePath := string(os.PathSeparator) +
		"../config" + string(os.PathSeparator)
	cacheFile, err := GetFilePath(basePath + "cache.yaml")
	PanicError(err)
	cacheConf := conf.Cache{}
	err = BindYamlConf(&cacheConf, cacheFile)
	PanicError(err)
	redisClient = NewRedisPool(cacheConf.Master)
}
func TestSet(t *testing.T) {
	redisClient.SetPrefix("hello_")
	if err := redisClient.Set("a", 100); err != nil {
		t.Error(err)
	}

	if err := redisClient.Set("a", 100, "PX", 100000); err != nil {
		t.Error(err)
	}

	if err := redisClient.Set("a", 100, "PX", 100000, "NX"); err != nil {
		t.Error(err)
	}

	if err := redisClient.Set("a", 100, "PX", 100000, "XX"); err != nil {
		t.Error(err)
	}

	if err := redisClient.Set("a", 100, "NX"); err != nil {
		t.Error(err)
	}

	if err := redisClient.Set("a", 100, "XX"); err != nil {
		t.Error(err)
	}
}

func TestObject(t *testing.T) {
	redisClient.SetPrefix("hahah_")
	data := TestData{
		ID:    100,
		Name:  "crala",
		Woman: true,
	}
	if err := redisClient.Set("crala", data); err != nil {
		t.Error(err)
	}

	var d TestData
	if err := redisClient.GetObject("crala", &d); err != nil {
		t.Error(err)
	}
}
func TestDo(t *testing.T) {
	_, err := redisClient.Do("set", "a", "100")
	if err != nil {
		t.Error(err)
	}
}

func TestDel(t *testing.T) {
	redisClient.Set("b", 100)
	err := redisClient.Del("b")
	if err != nil {
		t.Error(err)
	}
}

func TestH(t *testing.T) {
	err1 := redisClient.HSet("search", "google", "www.google.com")
	_, err2 := redisClient.HGet("search", "google")
	if err1 != nil || err2 != nil {
		t.Error(err1, err2)
	}

	err1 = redisClient.HSet("search", "baidu", "www.baidu.com")
	_, err2 = redisClient.HGet("search", "baidu")
	if err1 != nil || err2 != nil {
		t.Error(err1, err2)
	}
}

func TestHM(t *testing.T) {
	err1 := redisClient.HMSet("phone", "huawei", "mate 30", "xiaomi", "xiaomi9")
	if err1 != nil {
		t.Error(err1)
	}
	// data, err2 := redisClient.HMGet("phone", "huawei", "xiaomi")
	// fmt.Printf("%v", data)
	// if err1 != nil || err2 != nil {
	// 	t.Error(err1, err2)
	// }
}
func TestHObject(t *testing.T) {
	data1 := TestData{
		ID:    100,
		Name:  "crala1",
		Woman: true,
	}
	data2 := TestData{
		ID:    200,
		Name:  "alice",
		Woman: true,
	}
	data3 := TestData{
		ID:    300,
		Name:  "bob",
		Woman: false,
	}
	if err := redisClient.HSet("human", "crala", data1); err != nil {
		t.Error(err)
	}
	if err := redisClient.HMSet("human", "crala", data1, "alice", data2, "bob", data3); err != nil {
		t.Error(err)
	}
	var d TestData
	if err := redisClient.HGetObject("human", "bob", &d); err != nil {
		t.Error(err)
	}
	fmt.Println(d)
}

func TestPush(t *testing.T) {
	if err := redisClient.LPush("list", 100); err != nil {
		t.Error(err)
	}
	if err := redisClient.RPush("list", 200); err != nil {
		t.Error(err)
	}
	if err := redisClient.LPushX("list", 400, 500); err != nil {
		t.Error(err)
	}
	if err := redisClient.RPushX("list", 600, 700); err != nil {
		t.Error(err)
	}
	// 500 400 100 200 600 700
}
