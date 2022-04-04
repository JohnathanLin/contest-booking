package datasource

import (
	"backend/config"
	"backend/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
)

type ContestResult struct {
	BMatchId     string `json:"bMatchId"`
	BMatchName   string `json:"bMatchName"`
	GameId       string `json:"GameId"`
	GameName     string `json:"GameName"`
	GameTypeId   string `json:"GameTypeId"`
	GameMode     string `json:"GameMode"`
	GameTypeName string `json:"GameTypeName"`
	GameProcId   string `json:"GameProcId"`
	GameProcName string `json:"GameProcName"`
	TeamA        string `json:"TeamA"`
	ScoreA       string `json:"ScoreA"`
	TeamB        string `json:"TeamB"`
	ScoreB       string `json:"ScoreB"`
	MatchDate    string `json:"MatchDate"`
	MatchStatus  string `json:"MatchStatus"`
	MatchWin     string `json:"MatchWin"`
	IQTMatchId   string `json:"iQTMatchId"`
	BGameId      string `json:"bGameId"`
	NewsId       string `json:"NewsId"`
	HighlightsId string `json:"HighlightsId"`
}

type ContestMsg struct {
	Total     string          `json:"total"`
	Totalpage string          `json:"totalpage"`
	Page      string          `json:"page"`
	Result    []ContestResult `json:"result"`
}
type ContestAPIResponse struct {
	Status string     `json:"status"`
	Msg    ContestMsg `json:"msg"`
}

type Contest struct {
	ContestId        int    `json:"contestId"`
	Status           int    `json:"status"`
	ContestName      string `json:"contestName"`
	ContestStage     string `json:"contestStage"`
	ContestMarchName string `json:"contestMarchName"`
	Time             string `json:"time"`
	Result           string `json:"result"`
}

type ContestDetail struct {
	ContestId int    `json:"matchId"`
	TeamA     string `json:"teamAName"`
	TeamB     string `json:"teamBName"`
	ScoreA    int    `json:"teamAScore"`
	ScoreB    int    `json:"teamBScore"`
}

type GetContestDetailResponse struct {
	Success bool          `json:"success"`
	Data    ContestDetail `json:"data"`
}

// GetLatestContestInLPL 根据lpl赛程官网的状态获取比赛，获取状态1（未开始）和状态2（进行中）的比赛
func GetLatestContestInLPL() []Contest {
	var contestList []Contest
	contestList = append(contestList, getContestInLPL("2")...)
	contestList = append(contestList, getContestInLPL("1")...)
	return contestList
}

//根据lpl赛程官网的状态获取比赛
func getContestInLPL(status string) []Contest {
	var contestRes []Contest
	var header = make(map[string]string)
	var res = utils.Get(getUrlByContestStatus(status), header)
	res = res[11 : len(res)-1]
	str := []byte(res)
	fmt.Println(res)
	resObj := ContestAPIResponse{}
	err := json.Unmarshal(str, &resObj)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resObj)
	for _, contestLPL := range resObj.Msg.Result {
		contestId, _ := strconv.Atoi(contestLPL.BMatchId)
		status, _ := strconv.Atoi(contestLPL.MatchStatus)
		var contest Contest = Contest{
			ContestId:        contestId,
			Status:           status,
			ContestName:      contestLPL.GameName,
			ContestStage:     contestLPL.GameTypeName,
			ContestMarchName: contestLPL.BMatchName,
			Time:             contestLPL.MatchDate,
			Result:           contestLPL.ScoreA + " : " + contestLPL.ScoreB,
		}
		contestRes = append(contestRes, contest)
	}
	return contestRes
}

//拼接请求字符串
func getUrlByContestStatus(status string) string {
	return "https://apps.game.qq.com/lol/match/apis/searchBMatchInfo_bak.php?p8=5&p1=167&p4=" + status + "&p2=%C8%AB%B2%BF&p9=&p10=&p6=3&p11=&p12=&page=1&pagesize=10&r1=retObj&_=1648892356929"
}

// GetContestDetail 获取某一场比赛的详情（仅支持获取比赛开始的比赛）
func GetContestDetail(contestId int) GetContestDetailResponse {
	contestIdStr := strconv.Itoa(contestId)
	var res = GetContestDetailStr(contestIdStr)
	str := []byte(res)
	//fmt.Println(res)
	resObj := GetContestDetailResponse{}
	err := json.Unmarshal(str, &resObj)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(resObj)
	return resObj
}

// GetContestDetailStr 拼接请求字符串
func GetContestDetailStr(contestId string) string {
	if config.SystemConfig.TestMode == 0 {
		var header = make(map[string]string)
		header["authorization"] = "7935be4c41d8760a28c05581a7b1f570"
		return utils.Get("https://open.tjstats.com/match-auth-app/open/v1/compound/matchDetail?matchId="+contestId, header)
	} else {
		filePtr, err := ioutil.ReadFile("test/contestDetail.txt") //config的文件目录
		if err != nil {

		}
		//fmt.Println(string(filePtr))
		return string(filePtr)
	}

}
