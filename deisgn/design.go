package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// http://qiita.com/ikawaha/items/6638ee8b6978aef50d65
// APIサーバーの定義
var _ = API("deisgn", func() { // API defines the microservice endpoint and
	// APIのタイトル
	Title("The virtual wine deisgn") // other global properties. There should be one
	// 詳しい説明
	Description("A simple goa service") // and exactly one API definition appearing in
	// URLのスキーム
	Scheme("http") // the design.
	// ホスト名とポート
	Host("localhost:8080")
	// Version,Licence,Docsなどもある
})

// APIが管理するデータへのアクセス方法
// エンドポイントなどを定義する

// リソースの名前をつける
var _ = Resource("bottle", func() {
	// エンドポイントのprefix部分の定義
	BasePath("/bottles")
	// メディア・タイプの定義
	DefaultMedia(BottleMedia)

	// リソースに対する操作名 list show add createなど
	Action("show", func() {
		Description("Get bottle by id")
		// エンドポイントの定義
		// GET POST PUT PATCH DELETEなど
		// /bottles/{bottleID}
		Routing(GET("/:bottleID")) // Routing(POST("/:id"), PATHC("/:id"))　こういうことも出来る
		// 受け付けるパラメーター定義
		//Params(func() {
		//	// bottleIDという名前でIntegerと説明
		//	// curl -XGET localhost:8080/bottles/1?category=red
		//	Param("bottleID", Integer, "Bottle ID", func() {
		//		// バリデーション定義
		//		Minimum(0)
		//		Maximum(127)
		//	})
		//	Param("category", String, "Category", func() {
		//		Enum("red", "white", "rose")
		//	})
		//})

		// ペイロードで設定するとこうなる
		Params(func() {
			Param("bottleID", Integer, "Bottle ID")
		})
		Payload(BottlePayload)
		Response(OK)
		Response(NotFound)
	})
})

// レスポンスデータの定義
var BottleMedia = MediaType("application/vnd.goa.example.bottle+json", func() {
	Description("A bottle of wine")
	Attributes(func() {
		// idは整数型
		Attribute("id", Integer, "Unique bottle ID")
		// hrefは文字型
		Attribute("href", String, "API href for making requests on the bottle")
		// nameは文字型
		Attribute("name", String, "Name of wine")
		// 上記のうちで必須なもの
		Required("id", "href", "name")
	})
	// defaultは必須
	View("default", func() {
		Attribute("id")
		Attribute("href")
		Attribute("name")
	})
})

// Payload
// endpoint で受け付けるデータの形式を定義します．
var BottlePayload = Type("BottlePayload", func() {
	Member("bottleID", Integer, "Bottle ID", func() {
		Minimum(0)
		Maximum(127)
	})
	Member("category", String, "Category", func() {
		Enum("red", "whilte", "rose")
		// 何も指定がなかったらこれになる
		Default("red")
	})
	Member("comment", String, "Comment", func() {
		MaxLength(256)
	})
	Required("bottleID", "category")
})
