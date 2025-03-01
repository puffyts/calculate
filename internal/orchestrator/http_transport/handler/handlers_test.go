package handler

//import (
//	//"bufio"
//	//"bytes"
//	//"net/http_transport"
//	//"net/http_transport/httptest"
//	"testing"
//)
//
//func TestCalculateHandler(t *testing.T) {
//	testCases := []struct {
//		models  string
//		expected string
//	}{
//		{`{"expression": "(48-67+220)/5"}`, `{"result":40.2}`},
//		{`{"expression": "5*5"}`, `{"result":25}`},
//		{`{"expression": "2+3"}`, `{"result":5}`},
//		{`{"expression": "5/2"}`, `{"result":2.5}`},
//		{`{"expression": "555-444"}`, `{"result":111}`},
//		{`{"expression": "1+2+3-4+5*8"}`, `{"result":42}`},
//		{`{"expression": "(5-2*2.5)*111"}`, `{"result":0}`},
//		{`{"expression": "(48+677)/5"}`, `{"result":145}`},
//	}
//	return
