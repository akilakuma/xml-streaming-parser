package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"strings"
)

func StreamReader(reader io.ReadCloser) {

	decoder := xml.NewDecoder(reader)

	var (
		loveNode *Root
		Depth    int
	)
	for {
		// reading from decoder
		token, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error decoding XML:", err)
			return
		}

		switch se := token.(type) {
		case xml.StartElement:
			Depth++

			if se.Name.Local == "xml" {
				break
			}

			var attrMap = make(map[string]string)
			for _, attr := range se.Attr {
				attrMap[attr.Name.Local] = attr.Value
			}

			if loveNode == nil {
				loveNode = &Root{}
				loveNode.rootPtr = &Node{
					Depth:   Depth,
					AttrMap: attrMap,
					Key:     se.Name.Local,
				}
				loveNode.currentPtr = loveNode.rootPtr
			} else {
				subNode := &Node{
					Depth:   Depth,
					AttrMap: attrMap,
					Key:     se.Name.Local,
				}

				loveNode.AppendChildren(subNode)
				loveNode.currentPtr = subNode
			}

		case xml.EndElement:
			Depth--
			if loveNode == nil {
				break
			}
			loveNode.currentPtr = loveNode.currentPtr.Parent
		case xml.CharData:
			if loveNode == nil {
				break
			}

			content := strings.TrimSpace(string(se))
			if content != "" && loveNode.currentPtr != nil {
				loveNode.currentPtr.Char = content
			}
		}
	}

	switch loveNode.rootPtr.Key {
	case "Transaction":
		var target = &TransactionStruct{}
		travelNode(loveNode.rootPtr, TransactionFill, target)
		fmt.Println("Transaction data:")
		fmt.Println(target)
		for _, room := range target.RoomData {
			fmt.Println(room)
		}

	case "OTA_HotelRateAmountNotifRQ":
		var ra = &RateAmountStruct{}
		travelNode(loveNode.rootPtr, RateAmountFill, ra)
		fmt.Println("OTA_HotelRateAmountNotifRQ data:")
		fmt.Println(ra)
		for _, rap := range ra.RatePlanSlice {
			fmt.Println(rap)
		}
	case "OTA_HotelAvailNotifRQ":
		var ava = &AvailableStruct{}
		travelNode(loveNode.rootPtr, AvailableFill, ava)
		fmt.Println("OTA_HotelAvailNotifRQ data:")
		fmt.Println(ava)
		for _, rap := range ava.AvaStatusSlice {
			fmt.Println(rap)
		}
	default:
		fmt.Println("no this router:", loveNode.rootPtr.Key)
	}

}

func travelNode(
	cPtr *Node,
	fn func(cPtr *Node, data interface{}),
	data interface{},
) {

	nowPtr := cPtr
	fn(nowPtr, data)

	if cPtr.ChildrenSlice != nil {
		for _, v := range cPtr.ChildrenSlice {
			travelNode(v, fn, data)
		}
	}
}

// verifyAncestors 確認祖譜
func verifyAncestors(n *Node, ancestors []string) bool {
	for _, ancestor := range ancestors {
		if n.Parent == nil || n.Parent.Key != ancestor {
			return false
		}
		n = n.Parent
	}
	return true
}
