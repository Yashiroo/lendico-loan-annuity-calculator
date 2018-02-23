package loanhttpserver

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"github.com/yashiroo/lendico-loan-annuity-calculator/intrst_calc"
	"encoding/json"
	"github.com/yashiroo/lendico-loan-annuity-calculator/plangenerator"
)

type HTTPService struct {
	Server	*http.Server
	Router  *http.ServeMux
}

func NewHTTPService(addr string) *HTTPService{
	service := &HTTPService{
		Server:&http.Server{Addr:addr},
	}
	service.Router = http.NewServeMux()
	service.setupRouter()
	service.Server.Handler = service.Router

	return service
}

func (serv *HTTPService) setupRouter() {
	endpoints := []struct {
		name    string
		method  string
		pattern string
		handler http.Handler
	}{
		{
			"post_load_payload",
			"POST",
			"/generate-plan",
			http.HandlerFunc(serv.postGeneratePlanHandler),
		},
	}

	for _, e := range endpoints {
		serv.Router.Handle(e.pattern, e.handler)

	}
}

func (serv *HTTPService) postGeneratePlanHandler(w http.ResponseWriter, r *http.Request){
	//TODO: make sure this is a POST request
	body, err := ioutil.ReadAll(r.Body)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		//w.Write()
		fmt.Printf("Error reading body: %s",err)
		return
	}
	var loan intrstcalc.Loan
	err = json.Unmarshal(body, &loan)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		//w.Write(err)
		fmt.Printf("Error during unmarshal: %s",err)
		return
	}
	// we need to manually set the effective rate, since we're not using NewLoan
	loan.SetEffectiveRate()
	plan := loanplangenerator.GeneratePlan(&loan)

	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(plan)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		//w.Write(err)
		fmt.Printf("Error during marshal of returning data: %s",err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(data)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		//w.Write(err)
		fmt.Printf("Error replying to http request: %s",err)
		return
	}

}