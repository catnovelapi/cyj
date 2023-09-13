# Cyj (ciyuanji) Package

The `cyj` package is a Go library that provides a client for interacting with the Ciyuanji API. Ciyuanji is a Chinese
novel reading platform. This package allows you to perform various operations such as getting book information,
retrieving chapters, and managing user bookshelves.

## Installation

To use the `cyj` package, you need to have Go installed and set up. Then, you can install the package using the
following command:

```shell
go get github.com/catnovelapi/cyj
```
 
## Usage

To use the `cyj` package, you need to import it in your Go code:

```go
import "github.com/catnovelapi/cyj"
```

### Creating a Ciyuanji Client

To create a new Ciyuanji client, you can use the `NewCiyuanjiClient` function:

```go
client := cyj.NewCiyuanjiClient()
```

This creates a new client with default options. You can also customize the client by passing additional options:

```go
client := cyj.NewCiyuanjiClient(
cyj.WithHost("https://api.example.com"),
cyj.WithContentKey("your-content-key"),
cyj.WithParamKey("your-param-key"),
)
```

### Setting the Token

You can set the authentication token for the client using the `NewToken` method:

```go
client = client.NewToken("your-token")
```

This sets the token for subsequent API requests made by the client.

### Getting Book Information

To get information about a book, you can use the `GetBookInfoApi` method:

```go
result := client.GetBookInfoApi("book-id")
```

This returns a `gjson.Result` object containing the book information.

### Getting Account Information

To get account information, you can use the `GetAccountInfoApi` method:

```go
result := client.GetAccountInfoApi()
```

This returns a `gjson.Result` object containing the account information.

### Getting Catalog by Book ID

To get the catalog (list of chapters) for a book, you can use the `GetCatalogByBookIDApi` method:

```go
result := client.GetCatalogByBookIDApi("book-id")
```

This returns a `gjson.Result` object containing the catalog information.

### Getting Content by Book ID and Chapter ID

To get the content of a chapter in a book, you can use the `GetContentByBookIdAndChapterIdApi` method:

```go
result := client.GetContentByBookIdAndChapterIdApi("book-id", "chapter-id")
```

This returns a `gjson.Result` object containing the chapter content.

### Getting User Book Rack List

To get the list of books in the user's bookshelf, you can use the `GetUserBookRackListApi` method:

```go
result := client.GetUserBookRackListApi()
```

This returns a `gjson.Result` object containing the list of books in the bookshelf.

### Getting Phone Code by Phone Number

To get the phone code for a phone number, you can use the `GetPhoneCodeByPhoneNumberApi` method:

```go
result := client.GetPhoneCodeByPhoneNumberApi("phone-number")
```

This returns a `gjson.Result` object containing the phone code.

### Logging in with Phone Number and Phone Code

To log in using a phone number and phone code, you can use the `GetLoginByPhoneNumberAndPhoneCodeApi` method:

```go
result := client.GetLoginByPhoneNumberAndPhoneCodeApi("phone-number", "phone-code")
```

This returns a `gjson.Result` object containing the login information.

### Searching for Books by Keyword

To search for books by keyword, you can use the `GetSearchByKeywordApi` method:

```go
result := client.GetSearchByKeywordApi("keyword", "page")
```

This returns a `gjson.Result` object containing the search results.

### Getting Bookshelf

To get the user's bookshelf, you can use the `GetBookShelfApi` method:

```go
result := client.GetBookShelfApi()
```

This returns a `gjson.Result` object containing the user's bookshelf.

## Examples

Here are a few examples to demonstrate how to use the `cyj` package:

```go
package main

import (
	"fmt"
	"github.com/your-username/cyj"
)

func main() {
	client := cyj.NewCiyuanjiClient()

	// Get book information
	bookInfo := client.GetBookInfoApi("book-id")
	fmt.Println(bookInfo)

	// Get account information
	accountInfo := client.GetAccountInfoApi()
	fmt.Println(accountInfo)

	// Get catalog by book ID
	catalog := client.GetCatalogByBookIDApi("book-id")
	fmt.Println(catalog)

	// Get content by book ID and chapter ID
	content := client.GetContentByBookIdAndChapterIdApi("book-id", "chapter-id")
	fmt.Println(content)

	// Get user book rack list
	bookRackList := client.GetUserBookRackListApi()
	fmt.Println(bookRackList)

	// Get phone code by phone number
	phoneCode := client.GetPhoneCodeByPhoneNumberApi("phone-number")
	fmt.Println(phoneCode)

	// Log in with phone number and phone code
	loginResult := client.GetLoginByPhoneNumberAndPhoneCodeApi("phone-number", "phone-code")
	fmt.Println(loginResult)

	// Search for books by keyword
	searchResult := client.GetSearchByKeywordApi("keyword", "page")
	fmt.Println(searchResult)

	// Get bookshelf
	bookshelf := client.GetBookShelfApi()
	fmt.Println(bookshelf)
}
``` 