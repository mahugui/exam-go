package cmodels

import (
	"fmt"
	"encoding/json"

	"github.com/gomodule/redigo/redis"
)

var StudentQuestionKey = "%d:%d:%s"

func GetQuestionListFromPaper(key string, start int, end int) ([][]byte, error) {
	conn := RedisConn[examConnKey].Get()
	defer conn.Close()

	result, err := redis.ByteSlices(conn.Do("lrange", key, start, end))
	if err != nil{
		return nil, err
	}

	return result, nil
}

func GetStudentAnswerQuestion(userId int, examId string, paperKey string) ([][]byte, error) {
	conn := RedisConn[connKey].Get()
	defer conn.Close()

	result, err := redis.ByteSlices(conn.Do("hkeys", fmt.Sprintf(StudentQuestionKey,
		userId, examId, paperKey)))
	if err != nil{
		return nil, err
	}
	return result, nil
}

func SetStudentInfo(userId int, examId string, productId int, studentInfo interface{}) bool {
	conn := RedisConn[connKey].Get()
	defer conn.Close()

	studentInfo, err := json.Marshal(studentInfo)
	if err != nil{
		return false
	}

	result, err := redis.Bool(conn.Do("hset", fmt.Sprintf(studentProductKey, userId, productId),
		examId, studentInfo))
	if err != nil{
		return false
	}
	return result
}