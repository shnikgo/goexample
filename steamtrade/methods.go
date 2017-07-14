package steamtrade


import (
	"net/http"
	"fmt"
	"strconv"
	"github.com/gorilla/mux"
	steam "github.com/Philipp15b/go-steamapi"
)


type Responsible interface {
	MarshalJSON() ([]byte, error)
}


type Response struct {
	Status string `json:"status"`
	Description string `json:"description"`
}


type ResponseGetOfferList struct {
	Response
	Offers []*steam.CEconTradeOffer `json:"offers"`
}


type ResponseGetOffer struct {
	Response
	Offer *steam.CEconTradeOffer `json:"offer"`
}


func ResponseError(response http.ResponseWriter, err error) {
	result := Response{ Status: "error", Description: err.Error() }
	body, _ := result.MarshalJSON()
	response.WriteHeader(http.StatusBadRequest)
	fmt.Fprintf(response, string(body))	
}


func ResponseResult(response http.ResponseWriter, result Responsible) {
	body, err := result.MarshalJSON()
	if err != nil {
		ResponseError(response, err)
	} else {
		Log(string(body), "", LOG_LEVEL_DEBUG)
		response.WriteHeader(http.StatusOK)
		fmt.Fprintf(response, string(body))
	}
}


func MethodHelp(response http.ResponseWriter, request *http.Request) {
	// ToDo: Отобразить документацию
}


func MethodStop(response http.ResponseWriter, request *http.Request) {
	Log("StopServer", "", LOG_LEVEL_WARNING)
	StopServer()
}


func MethodGetOfferList(response http.ResponseWriter, request *http.Request) {
	list, err := steam.IEconGetTradeOffers(Config.AppKey, false, true, false, true, false, 0)
	if err != nil {
		Log(err.Error(), "", LOG_LEVEL_ERROR)
		ResponseError(response, err)
	} else {
		fmt.Println("ololo")
		result := ResponseGetOfferList{ Offers: list.Received }
		result.Status = "ok"
		fmt.Println(result)
		ResponseResult(response, result)
	}
}


func MethodGetOffer(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	trade_id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		Log(err.Error(), "", LOG_LEVEL_DEBUG)
		ResponseError(response, err)
	}
	offer, err := steam.IEconGetTradeOffer(Config.AppKey, trade_id)
	if err != nil {
		Log(err.Error(), "", LOG_LEVEL_DEBUG)
		ResponseError(response, err)
	} else {
		result := ResponseGetOffer{ Offer: offer }
		result.Status = "ok"
		ResponseResult(response, result)
	}
}


func MethodCancelOffer(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	trade_id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		Log(err.Error(), "", LOG_LEVEL_DEBUG)
		ResponseError(response, err)
	}
	err = steam.IEconCancelTradeOffer(Config.AppKey, trade_id)
	if err != nil {
		Log(err.Error(), "", LOG_LEVEL_DEBUG)
		ResponseError(response, err)
	} else {
		result := Response{ Status: "ok" }
		ResponseResult(response, result)
	}
}
