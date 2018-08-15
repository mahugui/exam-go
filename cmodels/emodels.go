package cmodels

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

var connKey = "student"
var examConnKey = "exam"
var examKey = "exam:%s"
var examInfoKey = "info:%s"
var productKey = "P:%d"
var studentProductKey = "%d:%d"
var examStudentKey = "%s:student"

func GetProductAllExams(productId int) (map[string]string, error){
	conn := RedisConn[examConnKey].Get()
	defer conn.Close()

	reply, err := redis.StringMap(conn.Do("hgetall", fmt.Sprintf(productKey, productId)))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return reply, nil
}

func GetUserAllExam(userId int, productId int) (map[string]string, error) {
	conn := RedisConn[connKey].Get()
	defer conn.Close()

	result, err := redis.StringMap(conn.Do("hgetall", fmt.Sprintf(studentProductKey, userId, productId)))
	if err != nil{
		fmt.Println(err)
		return nil, err
	}
	return result, nil
}

func GetUserAttendExam(userId int, examId string) (bool, error) {
	conn := RedisConn[examConnKey].Get()
	defer conn.Close()

	result, err := redis.Bool(conn.Do("sismember", fmt.Sprintf(examStudentKey, examId), userId))
	if err != nil{
		fmt.Println(err)
		return false, err
	}
	return result, nil
}

func GetExamRandom(examId string, num int) ([][]byte, error) {
	conn := RedisConn[examConnKey].Get()
	defer conn.Close()

	result, err := redis.ByteSlices(conn.Do("srandmember", fmt.Sprintf(examKey, examId), num))
	if err != nil{
		return nil, err
	}
	return result, err
}