package loanhttpserver

import "testing"

func TestNewHTTPService(t *testing.T) {
	service := NewHTTPService(":8888")
	err := service.Server.ListenAndServe()
	if err != nil{
		t.Fail()
	}
}
