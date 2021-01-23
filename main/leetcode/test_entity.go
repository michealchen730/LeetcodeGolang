package main

//func main(){
//	//data := make(map[string]interface{})
//	//data["entity_id"] = 12
//	//data["entity_name"] = "lok"
//	//b, _ := json.Marshal(data)
//	//
//	//req, _ := http.NewRequest("PUT", "http://192.168.0.17:8080/entity", bytes.NewBuffer(b))
//	//req.Header.Add("Content-Type", "application/json")
//
//	res,_:=http.Get("http://192.168.0.17:8080/entity?entity_name=phone")
//
//
//	//res, _ := http.DefaultClient.Do(req)
//
//	bytes, err := ioutil.ReadAll(res.Body)
//
//	if err != nil{
//		fmt.Println("错误")
//	}
//	fmt.Println(string(bytes))
//	defer res.Body.Close()
//	fmt.Println("test")
//}

//func main(){
//	data := make(map[string]interface{})
//	data["attribute_name"] = "attr"
//	data["entity_name"] = "test"
//	data["entity_id"] = 7
//	b, _ := json.Marshal(data)
//	//
//	req, _ := http.NewRequest("POST", "http://192.168.0.17:8080/attribute", bytes.NewBuffer(b))
//	req.Header.Add("Content-Type", "application/json")
//
//	//res,_:=http.Get("http://192.168.0.17:8080/entity?entity_name=test")
//
//
//	res, _ := http.DefaultClient.Do(req)
//
//	bytes, err := ioutil.ReadAll(res.Body)
//
//	if err != nil{
//		fmt.Println("错误")
//	}
//	fmt.Println(string(bytes))
//	defer res.Body.Close()
//	fmt.Println("test")
//}
