package main

import (
	"strconv"
)

type TransactionStruct struct {
	Partner   string
	Action    string
	HotelCode string
	RoomData  []*TSRoomDataDetail
	Package   TSPackage
}

type TSRoomDataDetail struct {
	RoomID               string
	RoomNameSlice        []TSLangText
	RoomDescriptionSlice []TSLangText
	Capacity             int
	PhotoUrl             string
	PhotoCaption         string
}

type TSLangText struct {
	language string
	text     string
}

type TSPackage struct {
	PackageId               string
	PackageNameSlice        []TSLangText
	PackageDescriptionSlice []TSLangText
	RefundableAvailable     string
	RefundableUntilDays     string
	RefundableUntilTime     string
	BreakfastIncluded       int
}

func TransactionFill(
	cPtr *Node,
	data interface{},
) {
	var (
		transData  *TransactionStruct
		ok         bool
		mappingKey = strconv.Itoa(cPtr.Depth) + "_" + cPtr.Key
	)
	if transData, ok = data.(*TransactionStruct); !ok {
		//fmt.Println("assert fail")
		//fmt.Println(transData)
		return
	}
	switch mappingKey {
	case "1_Transaction":
		transData.Partner = cPtr.AttrMap["partner"]
	case "2_PropertyDataSet":
		transData.Action = cPtr.AttrMap["action"]
	case "3_Property":
		transData.HotelCode = cPtr.Char
	case "4_RoomID":
		var roomData = &TSRoomDataDetail{
			RoomID: cPtr.Char,
		}
		transData.RoomData = append(transData.RoomData, roomData)
	case "5_Text":
		content := TSLangText{
			language: cPtr.AttrMap["language"],
			text:     cPtr.AttrMap["text"],
		}

		if verifyAncestors(cPtr, []string{"Name", "RoomData"}) {
			transData.RoomData[len(transData.RoomData)-1].RoomNameSlice = append(
				transData.RoomData[len(transData.RoomData)-1].RoomNameSlice,
				content,
			)
		} else if verifyAncestors(cPtr, []string{"Name", "PackageData"}) {
			transData.Package.PackageNameSlice = append(
				transData.Package.PackageNameSlice,
				content,
			)
		} else if verifyAncestors(cPtr, []string{"Description", "RoomData"}) {
			transData.RoomData[len(transData.RoomData)-1].RoomDescriptionSlice = append(
				transData.RoomData[len(transData.RoomData)-1].RoomDescriptionSlice,
				content,
			)
		} else if verifyAncestors(cPtr, []string{"Description", "PackageData"}) {
			transData.Package.PackageDescriptionSlice = append(
				transData.Package.PackageDescriptionSlice,
				content,
			)
		}
	case "4_Capacity":
		capacity, _ := strconv.Atoi(cPtr.Char)
		transData.RoomData[len(transData.RoomData)-1].Capacity = capacity
	case "6_URL":
		if verifyAncestors(cPtr, []string{"PhotoURL"}) {
			transData.RoomData[len(transData.RoomData)-1].PhotoUrl = cPtr.Char
		}
	case "6_Caption":
		if verifyAncestors(cPtr, []string{"PhotoURL"}) {
			transData.RoomData[len(transData.RoomData)-1].PhotoCaption = cPtr.Char
		}
	case "4_PackageID":
		transData.Package.PackageId = cPtr.Char
	case "4_Refundable":
		transData.Package.RefundableAvailable = cPtr.AttrMap["available"]
		transData.Package.RefundableUntilDays = cPtr.AttrMap["refundable_until_days"]
		transData.Package.RefundableUntilTime = cPtr.AttrMap["refundable_until_time"]
	case "4_BreakfastIncluded":
		num, _ := strconv.Atoi(cPtr.Char)
		transData.Package.BreakfastIncluded = num
	}
}
