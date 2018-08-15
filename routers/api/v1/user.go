package v1

import (
	"fmt"
	"time"
	"strings"
	"strconv"
	"net/http"
	"encoding/json"

	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	beeUtils "github.com/astaxie/beego/utils"
	"github.com/gin-gonic/gin"

	"github.com/goexam/models"
	"github.com/goexam/cmodels"
	"github.com/goexam/utils"
	"github.com/goexam/utils/httpbase"
)
//18100086
type ExamInfo struct {
	MaxSubmits int `json:"max_submits"`
	PublishTime string `json:"publish_time"`
	Descriptions string `json:"descriptions"`
	Quantity int `json:"quantity"`
	OpenTime string `json:"open_time"`
	EndTime string `json:"end_time"`
	Score int `json:"score"`
	UpdateTime int `json:"update_time"`
	Name string `json:"name"`
	Duration int `json:"duration"`
	Interval int `json:"interval"`
	IsCertification bool `json:"is_certification"`
 }

type StudentInfo struct {
	EndTime int `json:"end_time"`
	IsAnswer int `json:"is_answer"`
	IsMakeup int `json:"is_makeup"`
	LastPaperId string `json:"last_paper_id"`
	LastSubmit int `json:"last_submit"`
	MakeupId string `json:"makeup_id"`
	MaxSubmits int `json:"max_submits"`
	OpenTime int `json:"open_time"`
	PaperId string `json:"paper_id"`
	PaperNum int `json:"paper_num"`
	PublishTime int `json:"publish_time"`
	Score int `json:"score"`
	StartTime int64 `json:"start_time"`
	Status int `json:"status"`
	StatusName string `json:"status_name"`
	Submit int `json:"submit"`
	ToScore int `json:"to_score"`
}

type ExamData struct{
	ExamId string `json:"exam_id"`
	IsAttend bool `json:"is_attend"`
	ExamInfo ExamInfo `json:"exam_info"`
	StudentInfo StudentInfo `json:"student_info"`
	UnAnswerQuestionList []string `json:"un_answer_question_list"`
}

type ResultData struct {
	Product map[string]interface{} `json:"product"`
	ExamList []ExamData `json:"exam_list"`
	Time int64 `json:"time"`
}

func GetUserProducts(c *gin.Context)  {
	appG := httpbase.Gin{c}
	userId := com.StrTo(c.Param("user_id")).MustInt()
	valid := validation.Validation{}
	valid.Min(userId, 1, "user_id").Message("用户id必须大于0的数字")

	if valid.HasErrors(){
		appG.Response(http.StatusOK, valid.Errors, nil)
		return
	}

	results, err := models.GetProductsByUserID(userId)
	if err != nil {
		appG.Response(http.StatusOK, httpbase.ERROR, err)
		return
	}

	appG.Response(http.StatusOK, httpbase.SUCCESS, results)
}

func ExamDataView(c *gin.Context) {
	appG := httpbase.Gin{c}
	userId := com.StrTo(c.Param("user_id")).MustInt()
	productId := com.StrTo(c.Param("product_id")).MustInt()
	valid := validation.Validation{}
	valid.Min(userId, 1, "user_id").Message("用户id必须大于0")
	valid.Min(productId, 1, "product_id").Message("课程id必须大于0")

	if valid.HasErrors(){
		appG.Response(http.StatusOK, httpbase.INVALID_PARAMS, nil)
		return
	}

	var returnData ResultData
	returnData.Time = time.Now().Unix()
	returnData.Product = make(map[string]interface{})
	returnData.Product["id"]=productId

	productExam, err := cmodels.GetProductAllExams(productId)
	if err != nil{
		appG.Response(http.StatusOK, httpbase.ERROR, err)
		return
	}

	name, ok := productExam["name"]
	if ok{
		delete(productExam, "name")
		returnData.Product["name"] = name
	}

	exams, err := cmodels.GetUserAllExam(userId, productId)
	if err != nil{
		appG.Response(http.StatusOK, httpbase.ERROR, err)
		return
	}

	for key, value := range productExam{
		var examData ExamData
		examData.ExamId = key

		err := json.Unmarshal([]byte(value), &examData.ExamInfo)
		if err != nil{
			continue
		}

		stuExam, ok := exams[key]

		if ok{
			err := json.Unmarshal([]byte(stuExam), &examData.StudentInfo)
			if err != nil{
				continue
			}

			if examData.StudentInfo.IsAnswer == 0{
				examData.StudentInfo.StartTime = returnData.Time
			}

			attend, err := cmodels.GetUserAttendExam(userId, key)
			if err != nil{
				fmt.Println(err)
				attend = false
			}

			examData.IsAttend = attend
		}else{
			paperId, err:= cmodels.GetExamRandom(key, 1)
			if (err != nil) || (len(paperId) == 0) {
				fmt.Println(err)
				continue
			}

			IdNum := strings.Split(string(paperId[0][:]), "#")
			if len(IdNum) != 2 {
				fmt.Println("split paperIdNum err")
				continue
			}

			examData.StudentInfo.PaperId = IdNum[0]
			examData.StudentInfo.PaperNum, err = strconv.Atoi(IdNum[1])
			if err != nil{
				examData.StudentInfo.PaperNum = 0
			}

			examData.StudentInfo.StartTime = returnData.Time
			examData.StudentInfo.MaxSubmits = examData.ExamInfo.MaxSubmits
			if !cmodels.SetStudentInfo(userId, key, productId,  examData.StudentInfo) {
				fmt.Println("set student info into redis error")
			}
		}

		questionList, _ := cmodels.GetQuestionListFromPaper(examData.StudentInfo.PaperId, 0 ,-1)
		answerQuestionList, _ :=cmodels.GetStudentAnswerQuestion(userId, key, examData.StudentInfo.PaperId)
		unQuestionList := beeUtils.SliceDiff(utils.ArrBytesToInterface(questionList),
			utils.ArrBytesToInterface(answerQuestionList))
		examData.UnAnswerQuestionList = utils.ArrInterfaceByteToArrString(unQuestionList)

		returnData.ExamList = append(returnData.ExamList, examData)
	}
	appG.Response(http.StatusOK, httpbase.SUCCESS, returnData)
	return
}
