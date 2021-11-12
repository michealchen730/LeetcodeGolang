package main

//按照教程来一遍
//func shoppingOffers(price []int, special [][]int, needs []int) int {
//	minRes:=0
//	for k,v:=range needs{
//		minRes+=v*price[k]
//	}
//	//过滤无用的大礼包
//	var special2 [][]int
//	for _,v:=range special{
//		canUse:=true
//		oPrice:=0
//		for k,need:=range needs{
//			if v[k]>need{
//				canUse=false
//				break
//			}
//			oPrice+=v[k]*price[k]
//		}
//		if canUse&&v[len(v)-1]<oPrice{
//			special2=append(special2,v)
//		}
//	}
//	//对有用的大礼包进行遍历
//	//可以选择要或者不要（完全递归）
//	for _,v:=range special2{
//		//不要
//		tRes:=shoppingOffers(price,special2[1:],needs)
//		if tRes<minRes{
//			minRes=tRes
//		}
//		//要
//		priceOfPack:=v[len(v)-1]
//		//oPrice:=0
//		needsNew:=make([]int,len(needs))
//		for i:=0;i<len(v)-1;i++{
//			//oPrice+=v[i]*price[i]
//			needsNew[i]=needs[i]-v[i]
//		}
//		//同一礼包可以一直购买
//		tRes=shoppingOffers(price,special2,needsNew)+priceOfPack
//		if tRes<minRes{
//			minRes=tRes
//		}
//	}
//	return minRes
//}

// 上一轮代码有误
// 第26行不应该使用for循环
//func shoppingOffers(price []int, special [][]int, needs []int) int {
//	minRes:=0
//	for k,v:=range needs{
//		minRes+=v*price[k]
//	}
//	//过滤无用的大礼包
//	var special2 [][]int
//	for _,v:=range special{
//		canUse:=true
//		oPrice:=0
//		for k,need:=range needs{
//			if v[k]>need{
//				canUse=false
//				break
//			}
//			oPrice+=v[k]*price[k]
//		}
//		if canUse&&v[len(v)-1]<oPrice{
//			special2=append(special2,v)
//		}
//	}
//	//对有用的大礼包进行遍历
//	//可以选择要或者不要（完全递归）
//	if len(special2)>0{
//		v:=special2[0]
//		//不要
//		tRes := shoppingOffers(price, special2[1:], needs)
//		if tRes < minRes {
//			minRes = tRes
//		}
//		//要
//		priceOfPack := v[len(v)-1]
//		//oPrice:=0
//		needsNew := make([]int, len(needs))
//		for i := 0; i < len(v)-1; i++ {
//			//oPrice+=v[i]*price[i]
//			needsNew[i] = needs[i] - v[i]
//		}
//		//同一礼包可以一直购买
//		tRes = shoppingOffers(price, special2, needsNew) + priceOfPack
//		if tRes < minRes {
//			minRes = tRes
//		}
//	}
//	return minRes
//}
