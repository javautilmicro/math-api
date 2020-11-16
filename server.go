package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"sort"
)


type CalcData struct {
	NumberItems int 	`json:"numberItems"`
	Numbers    []int    `json:"numbers"`
}

type Answer struct {
	Title		string   `json:"title"`
	Result      int 	 `json:result`
}

type AnswerFloat struct {
	Title		string   `json:"title"`
	Result 		float64  `json:resultFloat`
}


func getMin(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	
	w.Header().Set("Content-Type", "application/json")
	
	
	var c CalcData
	err := json.NewDecoder(req.Body).Decode(&c)
	if err != nil {
		log.Println(">   - error: getMin(): failed at json.NewDecoder(req.Body).Decode(&c)")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	numbItems := c.NumberItems
	numbers := c.Numbers
	numbersCount := len(numbers)
	if numbItems != numbersCount {
		// must be an error!
		log.Println(">   - error: getMin(): CalcData.NumberItems-[", numbItems,"] not equal to len(numbers))--[", numbersCount,"]")
		err = fmt.Errorf("bad request data: CalcData.NumberItems=[%v] != len(Numbers)=[%v]", numbItems, numbersCount)
		// seems better than doing the fail thing
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	dex := 0
	// I know there must be quicker ways of getting this information, but
	// maybe not, I mean we've got to look at all the values, unless we
	// can find a way to not visit all values, I don't know, if we put this
	// into something that would sort things for us, it would have to do the
	// same sort of things, maybe, I'm still learning go (as you can see)
	minIntVal := 0
	for _, num := range numbers {
		if dex == 0 {
			minIntVal = num
		} else  if num < minIntVal {
			minIntVal = num
		}
		//log.Printf(">   - number[%v]:....[%v]<--<<[type=(%T)] ---->minIntVal:...(so-far):...[%v]<--<<[type=(%T)]\n", dex, num, num, minIntVal, minIntVal)
		dex = dex + 1
	}
	
	// now we package up our result: simple, right?
	ans := new(Answer)
	ans.Title = "Minimum Result"
	ans.Result = minIntVal
	// and turn into json, and return while we are at it
	err = json.NewEncoder(w).Encode(ans)
	if err != nil {
		log.Println("error: getMin(): failed at json.NewEncoder(w).Encode(ans)")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}



func getMax(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	
	var c CalcData
	err := json.NewDecoder(req.Body).Decode(&c)
	if err != nil {
		log.Println(">   - error: getMax(): failed at json.NewDecoder(req.Body).Decode(&c)")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	numbItems := c.NumberItems
	numbers := c.Numbers
	numbersCount := len(numbers)
	if numbItems != numbersCount {
		// must be an error!
		log.Println(">   - error: getMax(): CalcData.NumberItems-[", numbItems,"] not equal to len(numbers))--[", numbersCount,"]")
		err = fmt.Errorf("bad request data: CalcData.NumberItems=[%v] != len(Numbers)=[%v]", numbItems, numbersCount)
		// seems better than doing the fail thing
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	
	//log.Println(">   - now here are our number values, right?")
	dex := 0
	maxIntVal := 0
	for _, num := range numbers {
		if dex == 0 {
			maxIntVal = num
		} else  if num > maxIntVal {
			maxIntVal = num
		}
		//log.Printf(">   - number[%v]:....[%v]<--<<[type=(%T)] ---->minIntVal:...(so-far):...[%v]<--<<[type=(%T)]\n", dex, num, num, maxIntVal, maxIntVal)
		dex = dex + 1
	}
	
	ans := new(Answer)
	ans.Title = "Maximum Result"
	ans.Result = maxIntVal
	
	err = json.NewEncoder(w).Encode(ans)
	if err != nil {
		log.Println(">   - error: getMax(): at json.NewEncoder(w).Encode(ans)",err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}



func getAvg(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	
	var c CalcData
	err := json.NewDecoder(req.Body).Decode(&c)
	if err != nil {
		log.Println(">   - error: getAvg(): failed at json.NewDecoder(req.Body).Decode(&c)")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	numbItems := c.NumberItems
	numbers := c.Numbers
	numbersCount := len(numbers)
	if numbItems != numbersCount {
		// must be an error!
		log.Println(">   - error: getAvg(): CalcData.NumberItems-[", numbItems,"] not equal to len(numbers))--[", numbersCount,"]")
		err = fmt.Errorf("bad request data: CalcData.NumberItems=[%v] != len(Numbers)=[%v]", numbItems, numbersCount)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	dex := 0
	totalVals := 0
	for _, num := range numbers {
		totalVals = totalVals + num
		//log.Printf(">   - number[%v]:....[%v]<--<<[type=(%T)] ---->totalVals:...(so-far):...[%v]<--<<[type=(%T)]\n", dex, num, num, totalVals, totalVals)
		dex = dex + 1
	}
	
	// math time (but, is it in int math or it is real-number-math?
	var aveVal float64
	aveVal = (float64(totalVals)) / (float64(numbItems))
	
	
	ans := new(AnswerFloat)
	ans.Title = "Average Result"
	ans.Result = aveVal
	
	err = json.NewEncoder(w).Encode(ans)
	if err != nil {
		log.Println(">   - error: getAve(): at json.NewEncoder(w).Encode(ans)",err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}



func getMedian(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	
	var c CalcData
	err := json.NewDecoder(req.Body).Decode(&c)
	if err != nil {
		log.Println(">   - error: getMedian(): failed at json.NewDecoder(req.Body).Decode(&c)")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	numbItems := c.NumberItems
	numbers := c.Numbers
	numbersCount := len(numbers)
	if numbItems != numbersCount {
		// must be an error!
		log.Println(">   - error: getMedian(): CalcData.NumberItems-[", numbItems,"] not equal to len(numbers))--[", numbersCount,"]")
		err = fmt.Errorf("bad request data: CalcData.NumberItems=[%v] != len(Numbers)=[%v]", numbItems, numbersCount)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	
	sort.Ints(numbers)
	isOdd := numbItems % 2 != 0
	var midDayIndex1 int
	var midDayIndex2 int
	
	if isOdd == true {
		
		ans := new(Answer)
		ans.Title = "Median Result"
		midDayIndex1 = numbItems / 2
		wholeMed := numbers[midDayIndex1]
		ans.Result = wholeMed
		
		err = json.NewEncoder(w).Encode(ans)
		if err != nil {
			log.Println(">   - error: getMedian(): at json.NewEncoder(w).Encode(ans)",err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		
	} else {
		
		ans := new(AnswerFloat)
		ans.Title = "Median Result"
		midDayIndex2 = numbItems / 2
		midDayIndex1 = midDayIndex2 - 1
	
		val1 := float64(numbers[midDayIndex1])
		val2 := float64(numbers[midDayIndex2])
		realNumMed := (val1 + val2) / 2.0
		
		ans.Result = realNumMed
		
		err = json.NewEncoder(w).Encode(ans)
		if err != nil {
			log.Println(">   - error: getMedian(): at json.NewEncoder(w).Encode(ans)",err)
	    	http.Error(w, err.Error(), http.StatusInternalServerError)
	    	return
	    }
		
	}
}




func getPercentile(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	
	ans := new(Answer)
	ans.Title = "Percentile Result"
	perc := 9
	ans.Result = perc
	
	
	json.NewEncoder(w).Encode(ans)
}


func main() {
	
	mux := httprouter.New()
	//mux.GET("/", index)
	// Route handles & endpoints
	mux.POST("/min", getMin)
	mux.POST("/max", getMax)
	mux.POST("/avg", getAvg)
	mux.POST("/median", getMedian)
	mux.POST("/percentile", getPercentile)
	
	// Start server
	log.Fatal(http.ListenAndServe("localhost:8088", mux))
	
	
	
	
	

}
