package main

import (
	"strconv"
)

type RateAmountStruct struct {
	Partner       string
	HotelCode     string
	RatePlanSlice []RatePlan
}
type RatePlan struct {
	Start        string
	End          string
	Mon          string
	Tue          string
	Weds         string
	Thur         string
	Fri          string
	Sat          string
	Sun          string
	InvTypeCode  string
	RatePlanCode string
	RateSlice    []Rate
}
type Rate struct {
	AmountAfterTax string
	CurrencyCode   string
	NumberOfGuests string
}

func RateAmountFill(
	cPtr *Node,
	data interface{},
) {
	var (
		rateAmountData *RateAmountStruct
		ok             bool
		mappingKey     = strconv.Itoa(cPtr.Depth) + "_" + cPtr.Key
	)
	if rateAmountData, ok = data.(*RateAmountStruct); !ok {
		//fmt.Println("assert fail")
		//fmt.Println(transData)
		return
	}
	switch mappingKey {
	case "4_RequestorID":
		rateAmountData.Partner = cPtr.AttrMap["ID"]
	case "2_RateAmountMessages":
		rateAmountData.HotelCode = cPtr.AttrMap["HotelCode"]
	case "4_StatusApplicationControl":
		rateAmountData.RatePlanSlice = append(
			rateAmountData.RatePlanSlice,
			RatePlan{
				Start:        cPtr.AttrMap["Start"],
				End:          cPtr.AttrMap["End"],
				Mon:          cPtr.AttrMap["Mon"],
				Tue:          cPtr.AttrMap["Tue"],
				Weds:         cPtr.AttrMap["Weds"],
				Thur:         cPtr.AttrMap["Thur"],
				Fri:          cPtr.AttrMap["Fri"],
				Sat:          cPtr.AttrMap["Sat"],
				Sun:          cPtr.AttrMap["Sun"],
				InvTypeCode:  cPtr.AttrMap["InvTypeCode"],
				RatePlanCode: cPtr.AttrMap["RatePlanCode"],
			},
		)
	case "7_BaseByGuestAmt":
		if verifyAncestors(cPtr, []string{"BaseByGuestAmts", "Rate", "Rates", "RateAmountMessage"}) {
			rateAmountData.RatePlanSlice[len(rateAmountData.RatePlanSlice)-1].RateSlice = append(rateAmountData.
				RatePlanSlice[len(rateAmountData.RatePlanSlice)-1].RateSlice,
				Rate{
					AmountAfterTax: cPtr.AttrMap["AmountAfterTax"],
					CurrencyCode:   cPtr.AttrMap["CurrencyCode"],
					NumberOfGuests: cPtr.AttrMap["NumberOfGuests"],
				},
			)
		}
	}
}
