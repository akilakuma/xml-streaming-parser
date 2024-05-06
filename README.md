
XML Streaming Parser
---

## 使用說明
### 用途：
這是一個用來解析http api 是xml格式的library。  
由於主流的格式為json，在處理大量資料又是xml格式的需求之下，在開源上找不太到合適的套件。  
因為xml 解析後的struct，與parser method需要高度的客製化，無法直接包裝後給大家使用。  
請複製核心的travel method之後，自行撰寫所需的parser method。
### 核心演算法:
以DFS的概念走訪xml tree，每往下走一層深度(depth)加1。  
xml decoder的同時，紀錄xml tree樹狀結構母子關係，會先從根結點開始走訪過每個節點，變成以struct pointer紀錄整個xml。  
以深度(depth)和節點(element)名稱串接為該節點的名稱, 例如3_AvailStatusMessage，意思為第3層tag叫做AvailStatusMessage的元素。  
再塞入自己想要的struct結構做，以callback的方式重新recursive走過剛剛紀錄的結構樹，將對應的資料寫入struct field。  
### 心得與注意事項：
xml 某個element是否為array無法像json可以直接判斷，最好有文件規格先參考，array資料的儲存方式請參考程式內範例。  
同樣的depth和element是可能同時出現的，因此提供verifyAncestors，藉由這個method區分上面的母節點是怎麼走下來。
### 效能測試：
GCP上N2cpu，0.25顆cpu解析5MB以下的xml，可以控制在microsecond的範圍內。


### 範例格式:
參考google hotel center的標準API作為範例，旅宿業最喜歡xml了。  
[google hotel api](https://developers.google.com/hotels/hotel-prices/xml-reference/ari-property)  

Transaction  
OTA_HotelRateAmountNotifRQ  
OTA_HotelAvailNotifRQ  