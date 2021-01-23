package main

//func main() {
//	reader := bufio.NewReader(os.Stdin)
//	N,V:=getNums(reader)
//	dp:=make([]int,V+1)
//
//	for i:=1;i<=N;i++{
//		//dp[j]表示前i个物品放入容器为j的背包中的最大价值
//		C,W:=getNums(reader)
//		for j:=V;j>=1;j--{
//			if j-C>=0{
//				dp[j]=max(dp[j],dp[j-C]+W)
//			}
//		}
//	}
//	fmt.Println(dp[V])
//}
//
//func getNums(reader *bufio.Reader)(n1,n2 int){
//	data, _, err := reader.ReadLine()
//	if err==nil{
//		str := string(data)
//		strs:=strings.Fields(str)
//		N,_:=strconv.Atoi(strs[0])
//		V,_:=strconv.Atoi(strs[1])
//		return N,V
//	}
//	return -1,-1
//}

//func main(){
//
//
//	//传的格式为application/x-www-form-urlencoded
//	//res,err := http.PostForm("http://192.168.0.17:8081/user",url.Values{"username":{"jack"},"password":{"35013501"}})
//	//
//	//if err!=nil{
//	//	println(err.Error())
//	//}
//	//defer res.Body.Close()
//
//	data := make(map[string]interface{})
//	//data["id"] = 4
//	data["username"] = "chunfuling"
//	data["login_name"] = "micheal"
//	data["password"] = "35013501"
//	data["email"] = "fulingchun@sjtu.edu.cn"
//	b, _ := json.Marshal(data)
//
//
//
//	req, _ := http.NewRequest("POST", "http://192.168.0.17:8080/user/register", bytes.NewBuffer(b))
//	req.Header.Add("Content-Type", "application/json")
//
//
//	res, _ := http.DefaultClient.Do(req)
//
//	bytes, err := ioutil.ReadAll(res.Body)
//
//	if err != nil{
//		fmt.Println("错误")
//	}
//
//	fmt.Println(string(bytes))
//
//	defer res.Body.Close()
//
//
//
//	//resp, err := http.Post("http://192.168.0.17:8081/user", "application/json", bytes.NewBuffer(b))
//	//if err != nil {
//	//	fmt.Println(err)
//	//}
//	//defer resp.Body.Close()
//
//
//
//	fmt.Println("test")
//}
