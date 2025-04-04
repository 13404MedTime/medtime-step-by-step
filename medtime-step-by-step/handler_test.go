package function

import (
	"encoding/json"
	"fmt"
	"testing"
)

type TestHandlerI interface {
	GetAsserts() []Asserts
	GetBenchmarkRequest() Asserts
}

func NewAssert(f FunctionAssert) TestHandlerI {
	return f
}

func TestHandler(t *testing.T) {
	
	var Request = NewRequestBody{
		Data: Data{
			AppId: "test",
			ObjectData: struct {
				UserID   string  "json:\"user_id\""
				Steps    int     "json:\"steps\""
				Km       float64 "json:\"km\""
				MoveTime struct {
					Hour   float64 "json:\"hour\""
					Minute int     "json:\"minute\""
				} "json:\"move_time\""
				Date string "json:\"date\""
			}{
				UserID: "218d41e3-4475-4795-b4e2-54c843b7c18f",
				Steps:  204,
				Km:     151,
				MoveTime: struct {
					Hour   float64 "json:\"hour\""
					Minute int     "json:\"minute\""
				}{Hour: 1, Minute: 1},
				Date: "2023-12-11",
			},
		},
	}


	req, err := json.Marshal(Request)
	if err != nil {
		t.Errorf("Error on marshal request::: %s", err.Error())
	}
	fmt.Println(string(req))

	got := Handle(req)
	var resp Response
	err = json.Unmarshal([]byte(got), &resp)

	if err != nil {
		t.Errorf("Error on unmarshal response::: %s", err.Error())
	}
}

// func BenchmarkHandler(b *testing.B) {
// 	if !IsHTTP {
// 		return
// 	}
// 	a := NewAssert(FunctionAssert{})
// 	var start time.Time

// 	for i := 0; i < b.N; i++ {
// 		reqByte, err := json.Marshal(a.GetBenchmarkRequest().Request)
// 		assert.Nil(b, err)

// 		start = time.Now()

// 		response := Handle(reqByte)

// 		resStatus, err := ConvertResponse([]byte(response))
// 		assert.Nil(b, err)
// 		assert.Equal(b, "done", resStatus.Status)

// 		if time.Since(start) > time.Millisecond*5000 {
// 			assert.Nil(b, fmt.Errorf("took more time than %d ms: %v", 500, time.Since(start)))
// 		}
// 	}
// }
