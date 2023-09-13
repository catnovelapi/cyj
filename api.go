package cyj

import (
	"github.com/tidwall/gjson"
	"log"
)

func (book *Client) GetBookInfoApi(bookId string) gjson.Result {
	return book.get("/book/getBookDetail", map[string]any{"bookId": bookId})

}
func (book *Client) GetAccountInfoApi() gjson.Result {
	return book.get("/account/getAccountByUser", nil)
}

func (book *Client) GetCatalogByBookIDApi(bookId string) gjson.Result {
	return book.get("/chapter/getChapterListByBookId", map[string]any{"sortType": "1", "pageNo": "1", "pageSize": "9999", "bookId": bookId})
}

func (book *Client) NewGetCatalogByBookIDApi(bookId string) []gjson.Result {
	return book.GetCatalogByBookIDApi(bookId).Get("data.bookChapter.chapterList").Array()
}
func (book *Client) GetContentByBookIdAndChapterIdApi(bookId, chapterId string) gjson.Result {
	return book.get("/chapter/getChapterContent", map[string]any{"chapterId": chapterId, "bookId": bookId})
}
func (book *Client) NewGetContentByBookIdAndChapterIdApi(bookId, chapterId string) string {
	for i := 0; i < 5; i++ {
		response := book.GetContentByBookIdAndChapterIdApi(chapterId, bookId)
		if response.Get("code").String() == "200" {
			content := book.decryptDESECB([]byte(response.Get("data.chapter.content").String()), []byte("ZUreQN0E"))
			if content != "" {
				return content
			}
		} else {
			log.Printf("获取章节内容失败,chapterId:%v\t%s\n", chapterId, response.Get("msg").String())
		}
	}
	return ""
}

func (book *Client) GetUserBookRackListApi() gjson.Result {
	return book.get("/bookrack/getUserBookRackList", map[string]any{"RankType": "1", "PageNo": "1", "PageSize": "9999"})
}

func (book *Client) GetPhoneCodeByPhoneNumberApi(phoneNumber string) gjson.Result {
	return book.post("/login/getPhoneCode", map[string]any{"phone": phoneNumber, "smsType": "1"})
}

func (book *Client) GetLoginByPhoneNumberAndPhoneCodeApi(phoneNumber, phoneCode string) gjson.Result {
	return book.post("/login/phone", map[string]any{"phone": phoneNumber, "phoneCode": phoneCode})
}

func (book *Client) GetSearchByKeywordApi(keyword, page string) gjson.Result {
	return book.get("/book/searchBookList", map[string]any{"keyword": keyword, "pageNo": page, "pageSize": "15", "rankType": "0"})
}

func (book *Client) GetBookShelfApi() gjson.Result {
	return book.get("/bookrack/getUserBookRackList", map[string]any{"pageNo": 1, "pageSize": 100, "rankType": 1})
}
