package jobs

import (
	"net/http"
	"strings"
	"io/ioutil"
	"fmt"
	//"encoding/json"
	"encoding/json"
	//"strconv"
   ss "github.com/as/structslice"
)

func Canyoninlotinfopark()  {
	var b = `"ID<1000"`

	client := &http.Client{}

	req, err := http.NewRequest("POST", "http://123.132.229.198:60009/api/InVehicle/GetByFunc", strings.NewReader(b))
	if err != nil {
		// handle error
	}

	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	resp, err := client.Do(req)

	defer resp.Body.Close()

	//body, err := ioutil.ReadAll(resp.Body)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	//fmt.Println(body)
	type InfoStruct struct {
		State struct {
			IsSucess       bool        `json:"IsSucess"`
			IsError        bool        `json:"IsError"`
			Code           int         `json:"Code"`
			RecordAffected int         `json:"RecordAffected"`
			Describe       interface{} `json:"Describe"`
		} `json:"State"`
		Models []struct {
			ID              int    `json:"ID"`
			ParkingID       string `json:"ParkingId"`
			TokenNo         string `json:"TokenNo"`
			TokenType       int    `json:"TokenType"`
			ParkID          string `json:"ParkId"`
			TcmID           string `json:"TcmId"`
			TcmName         string `json:"TcmName"`
			StaffNo         string `json:"StaffNo"`
			StaffName       string `json:"StaffName"`
			RegPlate        string `json:"RegPlate"`
			InAutoPlate     string `json:"InAutoPlate"`
			VehicleBand     string `json:"VehicleBand"`
			VehicleColor    string `json:"VehicleColor"`
			InLaneName      string `json:"InLaneName"`
			InTime          string `json:"InTime"`
			InPicture       string `json:"InPicture"`
			InPicture2      string `json:"InPicture2"`
			InPictureStaff  string `json:"InPictureStaff"`
			InOperatorID    string `json:"InOperatorId"`
			InOperatorName  string `json:"InOperatorName"`
			InType          int    `json:"InType"`
			InFlag          int    `json:"InFlag"`
			LotFullRemark   string `json:"LotFullRemark"`
			InLaneID        string `json:"InLaneId"`
			State           int    `json:"State"`
			LastChargeTime  string `json:"LastChargeTime"`
			GroupLotState   int    `json:"GroupLotState"`
			ReservationNo   string `json:"ReservationNo"`
			InRemark        string `json:"InRemark"`
			TerminalID      string `json:"TerminalId"`
			TerminalName    string `json:"TerminalName"`
			Gid             string `json:"Gid"`
			Rid             string `json:"Rid"`
			VehicleCategory string `json:"VehicleCategory"`
		} `json:"Models"`
	}

	var Jsoninlotinfo InfoStruct

	//fmt.Println(string(body))
	json.Unmarshal([]byte(body), &Jsoninlotinfo)

    //fmt.Println(Jsoninlotinfo.Models)

	//声明一个map暂存数据
	if Jsoninlotinfo.State.IsSucess == true {
		JsonData := Jsoninlotinfo.Models
		data := make(map[string]string)
		ss.SortByName(JsonData, "ID")
		ss.SortStableByName(JsonData, "ID")
		//JsonData.Print()
		//fmt.Println(JsonData)

		for k,v := range JsonData{
		//	//fmt.Printf("%+d\n",k, v.ID)
		//	for w,n :=range v{
		//		//switch sa:= n.(type) {
		//		//	case int:
		//		//		if n != 0 {
		//		//			v = v
		//		//		}
		//		//	case string:
		//		//		if n != "" {
		//		//			v = v
		//		//		}
		//		//}
		//		fmt.Println(n)
		//	}
		//
		if v.ID != 0{
			data[k]=v.ID
		}

		}
		fmt.Println(data)
	}

	//data := make(map[string]string)
	//
	//if JsonPark.Model.Name != "" {
	//	data["name"] = JsonPark.Model.Name
	//}
	//
	//if JsonPark.Model.LotCount != 0 {
	//	data["lot_count"] = strconv.Itoa(JsonPark.Model.LotCount)
	//}
	//
	//if JsonPark.Model.LotFree != 0 {
	//	data["lot_free"] = strconv.Itoa(JsonPark.Model.LotFree)
	//}
	//
	//if JsonPark.Model.ParentNo != "" {
	//	data["parent_id"] = JsonPark.Model.ParentNo
	//}
	//
	//if JsonPark.Model.Name != "" {
	//	data["remark"] = JsonPark.Model.Remark
	//}else{
	//	data["remark"] = ""
	//}
	//
	//if JsonPark.Model.Rid != "" {
	//	data["rid"] = JsonPark.Model.Rid
	//}
	//if JsonPark.Model.Gid != "" {
	//	data["gid"] = JsonPark.Model.Gid
	//}
	//
	//var lot_count_int , lot_count_error = strconv.Atoi(data["lot_count"])
	//if lot_count_error != nil {
	//	panic(lot_count_error.Error())
	//}
	//
	//
	//var lot_free_int , lot_free_err = strconv.Atoi(data["lot_free"])
	//if lot_free_err != nil {
	//	panic(lot_free_err.Error())
	//}
	//fmt.Println(Jsoninlotinfo)
}