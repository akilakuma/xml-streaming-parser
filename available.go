package main

import "strconv"

type AvailableStruct struct {
	Partner        string
	HotelCode      string
	AvaStatusSlice []AvaStatus
}

type AvaStatus struct {
	BookingLimit                 string
	Start                        string
	End                          string
	InvType                      string
	RatePlanCode                 string
	LengthOfStaySlice            []AvaLengthOfStay
	RestrictionStatusStatus      string
	RestrictionStatusRestriction string
}

type AvaLengthOfStay struct {
	Time  string
	MType string
}

func AvailableFill(
	cPtr *Node,
	data interface{},
) {
	var (
		availableData *AvailableStruct
		ok            bool
		mappingKey    = strconv.Itoa(cPtr.Depth) + "_" + cPtr.Key
	)
	if availableData, ok = data.(*AvailableStruct); !ok {
		//fmt.Println("assert fail")
		//fmt.Println(transData)
		return
	}
	switch mappingKey {
	case "4_RequestorID":
		availableData.Partner = cPtr.AttrMap["ID"]
	case "2_AvailStatusMessages":
		availableData.HotelCode = cPtr.AttrMap["HotelCode"]
	case "3_AvailStatusMessage":
		availableData.AvaStatusSlice = append(
			availableData.AvaStatusSlice,
			AvaStatus{
				BookingLimit: cPtr.AttrMap["BookingLimit"],
			},
		)
	case "4_StatusApplicationControl":
		availableData.AvaStatusSlice[len(availableData.AvaStatusSlice)-1].Start = cPtr.AttrMap["Start"]
		availableData.AvaStatusSlice[len(availableData.AvaStatusSlice)-1].End = cPtr.AttrMap["End"]
		availableData.AvaStatusSlice[len(availableData.AvaStatusSlice)-1].InvType = cPtr.AttrMap["InvType"]
		availableData.AvaStatusSlice[len(availableData.AvaStatusSlice)-1].RatePlanCode = cPtr.AttrMap["RatePlanCode"]

	case "5_LengthOfStay":
		if verifyAncestors(cPtr, []string{"LengthsOfStay", "AvailStatusMessage"}) {
			availableData.AvaStatusSlice[len(availableData.AvaStatusSlice)-1].LengthOfStaySlice = append(
				availableData.AvaStatusSlice[len(availableData.AvaStatusSlice)-1].LengthOfStaySlice,
				AvaLengthOfStay{
					Time:  cPtr.AttrMap["Time"],
					MType: cPtr.AttrMap["MinMaxMessageType"],
				},
			)
		}
	case "4_RestrictionStatus":
		availableData.AvaStatusSlice[len(availableData.AvaStatusSlice)-1].RestrictionStatusStatus = cPtr.AttrMap["Status"]
		availableData.AvaStatusSlice[len(availableData.AvaStatusSlice)-1].RestrictionStatusRestriction = cPtr.AttrMap["Restriction"]
	}
}
