package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

//https://codeburst.io/unit-testing-for-rest-apis-in-go-86c70dada52d

/*func TestGetWord(t *testing.T) {
	// Set up a mock request and response
	req, err := http.NewRequest("GET", "/api/words/en/single/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	body, err := io.ReadAll(req.Body)
	// b, err := ioutil.ReadAll(resp.Body)  Go.1.15 and earlier
	if err != nil {
		log.Fatalln(err)
	}

	rr := httptest.NewRecorder()

	// Create a test router and call the handler
	router := mux.NewRouter()
	router.HandleFunc("/api/words/{languagecode}/single/{id}", getWord)
	//router.HandleFunc("/api/words/{languagecode}/single/{id}", getWord).Methods("GET")
	router.ServeHTTP(rr, req)
	//handler := http.HandlerFunc(getTenWordsByID)

	//handler.ServeHTTP(rr, req)
	// Verify the response
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{
		"id": "22",
		"english": "surface",
		"foreignword": "surface",
		"examplesentence_english": "On the surface, the spy looked like a typical businessman.",
		"examplesentence_foreign": "On the surface, the spy looked like a typical businessman.",
		"english_definition": "The overside or up-side of a flat object such as a table, or of a liquid.",
		"foreign_definition": "The overside or up-side of a flat object such as a table, or of a liquid.",
		"audiofilelink": "https://api.dictionaryapi.dev/media/pronunciations/en/surface-us.mp3"
	}`

	if string(body) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
		fmt.Println("body of response: ", req.Body)
		//fmt.Println("body: ", rr.Body.String())
	}
	//log.Fatal(http.ListenAndServe(":8000", router))
} */

/*func TestGetTenWordsByID(t *testing.T) {
	// Set up a mock request and response
	req, err := http.NewRequest("GET", "/api/words/zh/package/11", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	// Create a test router and call the handler
	router := mux.NewRouter()
	//api/words/{languagecode}/package/{id}
	//need to be /api/words..., not api/words
	router.HandleFunc("/api/words/{languagecode}/package/{id}", getTenWordsByID)
	//router.HandleFunc("/api/words/{languagecode}/single/{id}", getWord).Methods("GET")
	router.ServeHTTP(rr, req)
	//handler := http.HandlerFunc(getTenWordsByID)
	//handler.ServeHTTP(rr, req)
	// Verify the response
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{
		"tenwords": [
			{
				"id": "11",
				"english": "reveal",
				"foreignword": "揭示",
				"examplesentence_english": "The comedian had been telling us about his sleep being disturbed by noise. Then came the reveal: he was sleeping on a bed in a department store.",
				"examplesentence_foreign": "喜剧演员一直告诉我们他的睡眠被噪音打扰了。然后揭露：他睡在一家百货公司的床上。",
				"english_definition": "The outer side of a window or door frame; the jamb.",
				"foreign_definition": "窗户或门框的外侧；门框。",
				"audiofilelink": "https://api.dictionaryapi.dev/media/pronunciations/en/reveal-au.mp3"
			},
			{
				"id": "12",
				"english": "human",
				"foreignword": "人类",
				"examplesentence_english": "Humans share common ancestors with other apes.",
				"examplesentence_foreign": "人类与其他类人猿有共同的祖先。",
				"english_definition": "A human being, whether man, woman or child.",
				"foreign_definition": "一个人，无论是男人、女人还是孩子。",
				"audiofilelink": "https://api.dictionaryapi.dev/media/pronunciations/en/human-us.mp3"
			},
			{
				"id": "13",
				"english": "brain",
				"foreignword": "脑",
				"examplesentence_english": "She was a total brain.",
				"examplesentence_foreign": "她是一个完整的大脑。",
				"english_definition": "The control center of the central nervous system of an animal located in the skull which is responsible for perception, cognition, attention, memory, emotion, and action.",
				"foreign_definition": "动物中枢神经系统的控制中心，位于颅骨内，负责知觉、认知、注意力、记忆、情感和行动。",
				"audiofilelink": "https://api.dictionaryapi.dev/media/pronunciations/en/brain-uk.mp3"
			},
			{
				"id": "14",
				"english": "south",
				"foreignword": "南",
				"examplesentence_english": "The moon souths at nine.",
				"examplesentence_foreign": "月亮九点南下。",
				"english_definition": "One of the four major compass points, specifically 180°, directed toward the South Pole, and conventionally downwards on a map, abbreviated as S.",
				"foreign_definition": "四个主要罗盘点之一，具体为 180°，指向南极，在地图上通常向下，缩写为 S。",
				"audiofilelink": "https://api.dictionaryapi.dev/media/pronunciations/en/south-uk.mp3"
			},
			{
				"id": "15",
				"english": "historical",
				"foreignword": "历史的",
				"examplesentence_english": "July 4, 1776, is a historic date. A great deal of historical research has been done on the events leading up to that day.",
				"examplesentence_foreign": "1776 年 7 月 4 日是一个具有历史意义的日子。对导致那一天发生的事件进行了大量的历史研究。",
				"english_definition": "A historical romance.",
				"foreign_definition": "一段历史浪漫。",
				"audiofilelink": "https://api.dictionaryapi.dev/media/pronunciations/en/historical-us.mp3"
			},
			{
				"id": "16",
				"english": "campaign",
				"foreignword": "活动",
				"examplesentence_english": "The company is targeting children in their latest advertising campaign.",
				"examplesentence_foreign": "该公司在最新的广告活动中以儿童为目标。",
				"english_definition": "A series of operations undertaken to achieve a set goal.",
				"foreign_definition": "为实现既定目标而进行的一系列操作。",
				"audiofilelink": "https://api.dictionaryapi.dev/media/pronunciations/en/campaign-us.mp3"
			},
			{
				"id": "17",
				"english": "year",
				"foreignword": "年",
				"examplesentence_english": "we moved to this town a year ago;  I quit smoking exactly one year ago",
				"examplesentence_foreign": "我们一年前搬到了这个小镇；我整整一年前就戒烟了",
				"english_definition": "A solar year, the time it takes the Earth to complete one revolution of the Sun (between 365.24 and 365.26 days depending on the point of reference).",
				"foreign_definition": "一个太阳年，即地球完成太阳自转一圈所需的时间（在 365.24 到 365.26 天之间，具体取决于参考点）。",
				"audiofilelink": "https://api.dictionaryapi.dev/media/pronunciations/en/year-1-uk.mp3"
			},
			{
				"id": "18",
				"english": "ah",
				"foreignword": "啊",
				"examplesentence_english": "Mom drove my sister and I to school.",
				"examplesentence_foreign": "妈妈开车送我和妹妹上学。",
				"english_definition": "The speaker or writer, referred to as the grammatical subject, of a sentence.",
				"foreign_definition": "说话者或作者，被称为句子的语法主语。",
				"audiofilelink": "https://api.dictionaryapi.dev/media/pronunciations/en/ah-1-us.mp3"
			},
			{
				"id": "19",
				"english": "point",
				"foreignword": "观点",
				"examplesentence_english": "point de Venise; Brussels point",
				"examplesentence_foreign": "威尼斯点;布鲁塞尔点",
				"english_definition": "A discrete division of something.",
				"foreign_definition": "某物的离散划分。",
				"audiofilelink": "https://api.dictionaryapi.dev/media/pronunciations/en/point-au.mp3"
			},
			{
				"id": "20",
				"english": "tell",
				"foreignword": "告诉",
				"examplesentence_english": "All told, there were over a dozen.  Can you tell time on a clock?  He had untold wealth.",
				"examplesentence_foreign": "总而言之，有十几个。你能告诉时钟上的时间吗？他拥有数不清的财富。",
				"english_definition": "A reflexive, often habitual behavior, especially one occurring in a context that often features attempts at deception by persons under psychological stress (such as a poker game or police interrogation), that reveals information that the person exhibiting the behavior is attempting to withhold.",
				"foreign_definition": "一种反射性的、通常是习惯性的行为，尤其是发生在通常以心理压力下的人试图欺骗为特征的环境中（例如扑克游戏或警察审讯），它揭示了表现出该行为的人试图隐瞒的信息。",
				"audiofilelink": "https://api.dictionaryapi.dev/media/pronunciations/en/tell-us.mp3"
			}
		],
		"date": "03-25-2023"
	}`

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
		fmt.Println("body: ", rr.Body.String())
	}
	//log.Fatal(http.ListenAndServe(":8000", router))
} */

/*func add(x int, y int) int {
	return x + y
}

func TestAdd(t *testing.T) {
	t.Run("add 2 + 2", func(t *testing.T) {
		want := 4

		// Call the function you want to test.
		got := add(2, 2)

		// Assert that you got your expected response
		if got != want {
			t.Fail()
		}
	})
}

func TestGetWord(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/words/es/package/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getTenWordsByID)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"id":"","english":"","foreignword":"","examplesentence_english":"","examplesentence_foreign":"","english_definition":"","foreign_definition":"","audiofilelink":""}`

	resp := rr.Result()
	body, _ := io.ReadAll(resp.Body)

	t.Log(string(body))
	t.Log("HI")

	if string(body) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

//Maybe this will fix the problem..? -->
//https://stackoverflow.com/questions/65108369/why-is-the-response-body-empty-when-running-a-test-of-mux-api
*/

func TestGetWord(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/words/zh/single/11", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getWord)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{
		"id": "11",
		"english": "reveal",
		"foreignword": "揭示",
		"examplesentence_english": "The comedian had been telling us about his sleep being disturbed by noise. Then came the reveal: he was sleeping on a bed in a department store.",
		"examplesentence_foreign": "喜剧演员一直告诉我们他的睡眠被噪音打扰了。然后揭露：他睡在一家百货公司的床上。",
		"english_definition": "The outer side of a window or door frame; the jamb.",
		"foreign_definition": "窗户或门框的外侧；门框。",
		"audiofilelink": "https://api.dictionaryapi.dev/media/pronunciations/en/reveal-au.mp3"
	}`

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

//Might solve: https://stackoverflow.com/questions/71263812/what-do-i-wrong-while-coding-gorm-api-unit-test
