package ciyuanjiAPI

import (
	"github.com/tidwall/gjson"
)

func (book *CiyuanjiClient) GetBookInfoApi(bookId string) gjson.Result {
	return book.get("/book/getBookDetail", map[string]any{"bookId": bookId})

}
func (book *CiyuanjiClient) GetAccountInfoApi() gjson.Result {
	return book.get("/account/getAccountByUser", nil)
}

func (book *CiyuanjiClient) GetCatalogByBookIDApi(bookId string) gjson.Result {
	return book.get("/chapter/getChapterListByBookId", map[string]any{"sortType": "1", "pageNo": "1", "pageSize": "9999", "bookId": bookId})
}

func (book *CiyuanjiClient) NewGetCatalogByBookIDApi(bookId string) []gjson.Result {
	return book.GetCatalogByBookIDApi(bookId).Get("data.bookChapter.chapterList").Array()
}
func (book *CiyuanjiClient) GetContentByBookIdAndChapterIdApi(bookId, chapterId string) string {
	response := book.get("/chapter/getChapterContent", map[string]any{"chapterId": chapterId, "bookId": bookId})
	if response.Get("code").String() == "200" {
		return book.decryptDESECB([]byte(response.Get("data.chapter.content").String()), []byte("ZUreQN0E"))
	}
	return ""
}

func (book *CiyuanjiClient) GetUserBookRackListApi() gjson.Result {
	return book.get("/bookrack/getUserBookRackList", map[string]any{"RankType": "1", "PageNo": "1", "PageSize": "9999"})
}

func (book *CiyuanjiClient) GetPhoneCodeByPhoneNumberApi(phoneNumber string) gjson.Result {
	return book.post("/login/getPhoneCode", map[string]any{"phone": phoneNumber, "smsType": "1"})
}

func (book *CiyuanjiClient) GetLoginByPhoneNumberAndPhoneCodeApi(phoneNumber, phoneCode string) gjson.Result {
	return book.post("/login/phone", map[string]any{"phone": phoneNumber, "phoneCode": phoneCode})
}

func (book *CiyuanjiClient) GetSearchByKeywordApi(keyword, page string) gjson.Result {
	return book.get("/book/searchBookList", map[string]any{"keyword": keyword, "pageNo": page, "pageSize": "15", "rankType": "0"})
}

func (book *CiyuanjiClient) GetBookShelfApi() gjson.Result {
	return book.get("/bookrack/getUserBookRackList", map[string]any{"pageNo": 1, "pageSize": 100, "rankType": 1})
}
