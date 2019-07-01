# KD PBL1

食材の廃棄料を減らすことを目的としたショッピングサイトのAPIサーバー

## APIドキュメント

### [POST] ログイン

#### エンドポイント
```
POST http://54.238.92.95:8080/api/v1/auth/login
```

#### Request Parameters
場所|パラメータ名|指定する値|説明|
|:-|:-|:-|:-|
|Body|user_id|文字列|ログインしたいユーザーのID|
|Body|password|文字列|ログインしたいユーザーのパスワード|

#### Example Response

+ Response 200 (application/json)

    <details>
    <summary>Body</summary>
    <pre>
    <code>
    {
        "user_id": "goya",
        "user_name": "destinyZero",
        "is_admin": true
    }
    </code>
    </pre>
    </details>



#### Example Request

```
curl -F "user_id=goya" -F "password=syogo" http://54.238.92.95:8080/api/v1/auth/login
```
---

### [POST] ログアウト

#### エンドポイント
```
POST http://54.238.92.95:8080/api/v1/auth/logout
```

#### Request Parameters
場所|パラメータ名|指定する値|説明|
|:-|:-|:-|:-|
|Body|user_id|文字列|ログインしたいユーザーのID|

#### Example Request

```
curl -F "user_id=goya" http://54.238.92.95:8080/api/v1/auth/logout
```
---

### [GET] 渡したレシピIDを元にそのレシピのデータ(材料)を取得する

#### エンドポイント
```
GET http://54.238.92.95:8080/api/v1/ingredient/:recipe_id
```

#### Request Parameters
場所|パラメータ名|指定する値|説明|
|:-|:-|:-|:-|
|Path|recipe_id|数値|データを取得したいレシピID|

#### Example Response

+ Response 200 (application/json)

    <details>
    <summary>Body</summary>
    <pre>
    <code>
    [
        {
            "food_name": "ライス",
            "quantity": "250g"
        },
        {
            "food_name": "牛肉",
            "quantity": "85g"
        },
        {
            "food_name": "牛丼のタレ",
            "quantity": "適量"
        }
    ]
    </code>
    </pre>
    </details>

#### Example Request

```
curl http://54.238.92.95:8080/api/v1/ingredient/1
```
---

### [POST] 受け取ったレシピIDを元にそのレシピ画面を表示するのに必要なデータ（レシピ名・材料・画像URL）を取得する

#### エンドポイント
```
GET http://54.238.92.95:8080/api/v2/recipes/:recipe_id
```

#### Request Parameters
場所|パラメータ名|指定する値|説明|
|:-|:-|:-|:-|
|Path|recipe_id|数値|データを取得したいレシピID|

#### Example Request

```
curl http://54.238.92.95:8080/api/v2/recipes/1
```
---

### [POST] 材料IDに結びつく材料データをカートに追加する(ログイン必須)

#### エンドポイント
```
POST http://54.238.92.95:8080/api/v1/carts
```

#### Request Parameters
場所|パラメータ名|指定する値|説明|
|:-|:-|:-|:-|
|Body|user_id|文字列|材料をカートに追加したいユーザID|
|Body|recipe_id|数値|カートに追加したい材料のあるレシピID|

#### Example Request

```
curl -F "user_id=goya" -F "recipe_id=1" http://54.238.92.95:8080/api/v1/carts
```
---

### [GET] 全てのカテゴリの名前を取得する

#### エンドポイント
```
GET  http://54.238.92.95:8080/api/v1/categories
```

#### Example Response

+ Response 200 (application/json)

    <details>
    <summary>Body</summary>
    <pre>
    <code>
    [
        {
            "category_name": "全て"
        },
        {
            "category_name": "和食"
        },
        {
            "category_name": "洋食"
        },
        {
            "category_name": "中華"
        },
        {
            "category_name": "フレンチ"
        },
        {
            "category_name": "イタリアン"
        },
        {
            "category_name": "その他"
        }
    ]

    </code>
    </pre>
    </details>


#### Example Request

```
curl http://54.238.92.95:8080/api/v1/categories
```
---

### [GET] 与えられたキーとカテゴリーIDを元に検索した結果のレシピデータを取得する

#### エンドポイント
```
GET http://54.238.92.95:8080/api/v1/search-recipes/categories/:category_id/keys/:search_key
```

#### Request Parameters
場所|パラメータ名|指定する値|説明|
|:-|:-|:-|:-|
|Path|category_id|数値|絞り込みたいカテゴリーID|
|Path|search_key|数値|検索したい文字列|


#### Example Response

+ Response 200 (application/json)

    <details>
    <summary>Body</summary>
    <pre>
    <code>
    [
        {
            "recipe_id": 1,
            "recipe_name": "牛丼",
            "recipe_Description": "お肉がのったご飯だよ！",
            "recipe_image_url": "https://www.pakutaso.com/shared/img/thumb/KAZUHIRO171013022_TP_V.jpg",
            "recipe_page_url": "hogehoge.com",
            "price": "￥150",
            "point": 15
        },
        {
            "recipe_id": 2,
            "recipe_name": "肉だけカレーライス",
            "recipe_Description": "男は黙って肉食っとけカレーライス",
            "recipe_image_url": "https://d2l930y2yx77uc.cloudfront.net/production/uploads/images/7120502/picture_pc_bd3805fab5e332c67b1862c988179471.jpg",
            "recipe_page_url": "fugafuga.com",
            "price": "￥200",
            "point": 20
        }
    ]
    </code>
    </pre>
    </details>


#### Example Request

```
curl http://54.238.92.95:8080/api/v1/search-recipes/categories/1/keys/肉
```
---

### [GET] 与えられたカテゴリーIDだけを元に検索した結果のレシピデータを取得する

#### エンドポイント
```
GET http://54.238.92.95:8080/api/v1/search-recipes/categories/:category_id
```

#### Request Parameters
場所|パラメータ名|指定する値|説明|
|:-|:-|:-|:-|
|Path|category_id|数値|絞り込みたいカテゴリーID|


#### Example Response

+ Response 200 (application/json)

    <details>
    <summary>Body</summary>
    <pre>
    <code>
    [
        {
            "recipe_id": 2,
            "recipe_name": "肉だけカレーライス",
            "recipe_Description": "男は黙って肉食っとけカレーライス",
            "recipe_image_url": "https://d2l930y2yx77uc.cloudfront.net/production/uploads/images/7120502/picture_pc_bd3805fab5e332c67b1862c988179471.jpg",
            "recipe_page_url": "fugafuga.com",
            "price": "¥200",
            "point": 20
        },
        {
            "recipe_id": 1,
            "recipe_name": "牛丼",
            "recipe_Description": "お肉がのったご飯だよ！",
            "recipe_image_url": "https://www.pakutaso.com/shared/img/thumb/KAZUHIRO171013022_TP_V.jpg",
            "recipe_page_url": "hogehoge.com",
            "price": "¥150",
            "point": 15
        },
        {
            "recipe_id": 4,
            "recipe_name": "米食っとけ、そう米だけさ",
            "recipe_Description": "米オンリー",
            "recipe_image_url": "https://d2dcan0armyq93.cloudfront.net/photo/odai/600/222569875b57db9b87ae55845b35315d_600.jpg",
            "recipe_page_url": "komekome.com",
            "price": "¥50",
            "point": 5
        }
    ]
    </code>
    </pre>
    </details>


#### Example Request

```
curl http://54.238.92.95:8080/api/v1/search-recipes/categories/1
```
---

### [GET] 対象ユーザのカートの中身を取得する(ログイン必須)

#### エンドポイント
```
GET http://54.238.92.95:8080/api/v1/carts/users/:user_id
```

#### Request Parameters
場所|パラメータ名|指定する値|説明|
|:-|:-|:-|:-|
|Path|user_id|文字列|対象ユーザーのID|


#### Example Response

+ Response 200 (application/json)

    <details>
    <summary>Body</summary>
    <pre>
    <code>
    [
        {
            "recipe_name": "牛丼",
            "recipe_count": 1,
            "recipe_image_url": "https://www.pakutaso.com/shared/img/thumb/KAZUHIRO171013022_TP_V.jpg",
            "food_names": [
                {
                    "food_name": "ライス",
                    "food_count": 1,
                    "food_quantity": "250g"
                },
                {
                    "food_name": "牛肉",
                    "food_count": 1,
                    "food_quantity": "85g"
                },
                {
                    "food_name": "牛丼のタレ",
                    "food_count": 1,
                    "food_quantity": "適量"
                }
            ]
        },
        {
            "recipe_name": "肉だけカレーライス",
            "recipe_count": 1,
            "recipe_image_url": "https://d2l930y2yx77uc.cloudfront.net/production/uploads/images/7120502/picture_pc_bd3805fab5e332c67b1862c988179471.jpg",
            "food_names": [
                {
                    "food_name": "ライス",
                    "food_count": 1,
                    "food_quantity": "250g"
                },
                {
                    "food_name": "牛肉",
                    "food_count": 1,
                    "food_quantity": "100g"
                },
                {
                    "food_name": "カレールー",
                    "food_count": 1,
                    "food_quantity": "欠片4分の1ほど"
                }
            ]
        }
    ]
    </code>
    </pre>
    </details>


#### Example Request

```
curl http://54.238.92.95:8080/api/v1/carts/users/goya
```
---

### [PUT] 対象ユーザのカートの中身(レシピの数量)を更新する(ログイン必須)

#### エンドポイント
```
PUT http://54.238.92.95:8080/api/v1/carts/recipe-counts
```

#### Request Parameters
場所|パラメータ名|指定する値|説明|
|:-|:-|:-|:-|
|Body|user_id|文字列|対象ユーザーのID|
|Body|recipe_id|数値|対象のレシピID|
|Body|recipe_count|数値|購入したいレシピのセット数|

#### Example Request

```
curl -X PUT -d "user_id=goya&recipe_id=1&recipe_count=2" http://54.238.92.95:8080/api/v1/carts/recipe-counts
```
---

### [PUT] 対象ユーザのカートの中身(食料の数量)を更新する(ログイン必須)

#### エンドポイント
```
PUT http://54.238.92.95:8080/api/v1/carts/food-counts
```

#### Request Parameters
場所|パラメータ名|指定する値|説明|
|:-|:-|:-|:-|
|Body|user_id|文字列|対象のユーザーID|
|Body|recipe_id|数値|対象のレシピID|
|Body|food_id|数値|対象の食料ID|
|Body|food_count|数値|購入したい食料のセット数|

#### Example Request

```
curl -X PUT -d "user_id=goya&recipe_id=1&food_id=1&food_count=5" http://54.238.92.95:8080/api/v1/carts/food-counts
```
---

### [GET] 対象ユーザの購入履歴を取得する(ログイン必須)

#### エンドポイント
```
GET http://54.238.92.95:8080/api/v1/purchase-histories/users/:user_id
```

#### Request Parameters
場所|パラメータ名|指定する値|説明|
|:-|:-|:-|:-|
|Path|user_id|文字列|対象ユーザーのID|


#### Example Response

+ Response 200 (application/json)

    <details>
    <summary>Body</summary>
    <pre>
    <code>
    {
        "date": "2019-06-10T19:51:23+09:00",
        "total_price": 50,
        "total_point": 5,
        "PurchaseHistoryCards": [
            {
                "recipe_name": "米食っとけ、そう米だけさ",
                "recipe_count": 1,
                "recipe_image_url": "https://d2dcan0armyq93.cloudfront.net/photo/odai/600/222569875b57db9b87ae55845b35315d_600.jpg",
                "price": 50,
                "point": 5,
                "foods": [
                    {
                        "food_name": "ライス",
                        "food_quantity": "500g",
                        "food_count": 1
                    }
                ]
            }
        ]
    }
    </code>
    </pre>
    </details>


#### Example Request

```
curl http://54.238.92.95:8080/api/v1/purchase-histories/users/goya
```
---

### [POST] 対象ユーザのマイページ情報を取得する(ログイン必須)

#### エンドポイント
```
POST http://54.238.92.95:8080/api/v1/mypage
```

#### Request Parameters
場所|パラメータ名|指定する値|説明|
|:-|:-|:-|:-|
|Body|user_id|文字列|対象ユーザーのID|


#### Example Response

+ Response 200 (application/json)

    <details>
    <summary>Body</summary>
    <pre>
    <code>
    {
        "user_id": "goya",
        "user_name": "goyaShogo",
        "user_email": "goyagoya.com",
        "user_password": "syogo",
        "user_street_address": "兵庫県呉屋市呉屋村585-5",
        "cumulative_points": 1000
    }
    </code>
    </pre>
    </details>


#### Example Request

```
curl -F "user_id=goya"  http://54.238.92.95:8080/api/v1/mypage
```
---

### [PUT] 対象ユーザのマイページ情報を更新する(ログイン必須)

#### エンドポイント
```
GET http://54.238.92.95:8080/api/v1/mypage
```

#### Request Parameters
場所|パラメータ名|指定する値|説明|
|:-|:-|:-|:-|
|Body|user_id|文字列|対象ユーザーのID|
|Body|user_name|文字列|対象ユーザーの名前|
|Body|user_email|文字列|対象ユーザーのメールアドレス|
|Body|user_street_address|文字列|対象ユーザーの住所|
|Body|old_password|文字列|対象ユーザーの変更前のパスワード|
|Body|new_password|文字列|対象ユーザーの変更後のパスワード|

#### Example Request

```
curl -X PUT -d "user_id=goya&user_name=goyaShogo&user_email=goyagoya.com&user_street_address=兵庫県呉屋市呉屋村585-5&old_password=syogo&new_password=syogo" localhost:8080/api/v1/mypage
```
---

### [GET] 対象ユーザのブックマーク情報を取得する(ログイン必須)

#### エンドポイント
```
GET http://54.238.92.95:8080/api/v1/bookmark/users/:user_id
```

#### Request Parameters
場所|パラメータ名|指定する値|説明|
|:-|:-|:-|:-|
|Path|user_id|文字列|対象ユーザーのID|


#### Example Response

+ Response 200 (application/json)

    <details>
    <summary>Body</summary>
    <pre>
    <code>
    [
        {
            "recipe_id": 1,
            "recipe_name": "牛丼",
            "recipe_Description": "お肉がのったご飯だよ！",
            "recipe_image_url": "https://www.pakutaso.com/shared/img/thumb/KAZUHIRO171013022_TP_V.jpg",
            "price": "¥150",
            "point": 15
        },
        {
            "recipe_id": 4,
            "recipe_name": "米食っとけ、そう米だけさ",
            "recipe_Description": "米オンリー",
            "recipe_image_url": "https://d2dcan0armyq93.cloudfront.net/photo/odai/600/222569875b57db9b87ae55845b35315d_600.jpg",
            "price": "¥50",
            "point": 5
        }
    ]
    </code>
    </pre>
    </details>


#### Example Request

```
curl http://54.238.92.95:8080/api/v1/bookmark/users/goya
```
---

### [POST] 対象ユーザのブックマーク情報を追加する(ログイン必須)

#### エンドポイント
```
POST http://54.238.92.95:8080/api/v1/bookmark
```


#### Request Parameters
場所|パラメータ名|指定する値|説明|
|:-|:-|:-|:-|
|Body|user_id|文字列|対象ユーザーのID|
|Body|recipe_id|数値|対象レシピのID|


#### Example Request

```
curl -F "user_id=goya" -F "recipe_id=1" http://54.238.92.95:8080/api/v1/bookmark
```
---

### [DELETE] 対象ユーザのブックマーク情報を削除する(ログイン必須)

#### エンドポイント
```
DELETE http://54.238.92.95:8080/api/v1/bookmark
```


#### Request Parameters
場所|パラメータ名|指定する値|説明|
|:-|:-|:-|:-|
|Path|user_id|文字列|対象ユーザーのID|
|Path|recipe_id|数値|対象レシピのID|


#### Example Request

```
curl -X DELETE localhost:8080/api/v1/bookmark/users/goya/recipes/2
```
---

### [GET] 飢餓貢献ポイントと交換できる商品データを取得する

#### エンドポイント
```
GET http://54.238.92.95:8080/api/v1/trade-items
```

#### Request Parameters
場所|パラメータ名|指定する値|説明|
|:-|:-|:-|:-|


#### Example Response

+ Response 200 (application/json)

    <details>
    <summary>Body</summary>
    <pre>
    <code>
    [
        {
            "ID": 1,
            "Name": "スコッティ フラワーパック",
            "ImageURL": "http://ur0.work/W8jw",
            "Point": 300
        },
        {
            "ID": 2,
            "Name": "クリネックス ティシュー",
            "ImageURL": "http://ur0.work/UR5I",
            "Point": 400
        },
        {
            "ID": 3,
            "Name": "ウタマロ石鹸",
            "ImageURL": "http://ur0.work/02OC",
            "Point": 450
        },
        {
            "ID": 4,
            "Name": "ハイディスク モバイルバッテリー5000mAh",
            "ImageURL": "http://ur0.work/Xaxx",
            "Point": 2000
        },
        {
            "ID": 5,
            "Name": "プラズマクラスター搭載 卓上加湿器",
            "ImageURL": "http://ur0.work/WrVy",
            "Point": 15000
        },
        {
            "ID": 6,
            "Name": "オーガニック オリーブオイル",
            "ImageURL": "http://ur0.work/HwUk",
            "Point": 400
        },
        {
            "ID": 7,
            "Name": "キャノーラ油",
            "ImageURL": "http://ur0.work/3Vcm",
            "Point": 350
        },
        {
            "ID": 8,
            "Name": "おーいお茶2L×6本",
            "ImageURL": "http://urx.blue/SmGu",
            "Point": 350
        },
        {
            "ID": 9,
            "Name": "鼻セレブ3箱パック",
            "ImageURL": "http://urx.blue/P7lw",
            "Point": 450
        },
        {
            "ID": 10,
            "Name": "電気タコ焼き機",
            "ImageURL": "http://urx.blue/VpeM",
            "Point": 1500
        },
        {
            "ID": 11,
            "Name": "Dyson Cool",
            "ImageURL": "http://urx.blue/fIvJ",
            "Point": 25000
        },
        {
            "ID": 12,
            "Name": "人をダメにするソファ",
            "ImageURL": "http://urx.blue/V5dk",
            "Point": 2000
        },
        {
            "ID": 13,
            "Name": "SanDisc USB 128GB",
            "ImageURL": "http://urx.blue/Y0YM",
            "Point": 1600
        }
    ]
    </code>
    </pre>
    </details>


#### Example Request

```
curl http://54.238.92.95:8080/api/v1/purchase-histories/users/goya
```
---

### [POST] 対象ユーザが飢餓貢献ポイントと商品を交換する(ログイン必須)

#### エンドポイント
```
POST http://54.238.92.95:8080/api/v1/trade-items
```


#### Request Parameters
場所|パラメータ名|指定する値|説明|
|:-|:-|:-|:-|
|Body|user_id|文字列|対象ユーザーのID|
|Body|trade_item_id|数値|対象レシピのID|


#### Example Request

```
curl -F "user_id=goya" -F "trade_item_id=1" http://54.238.92.95:8080/api/v1/trade-items
```
---


### [GET] 飢餓貢献ポイントで交換した履歴情報を取得する

#### エンドポイント
```
GET http://54.238.92.95:8080/api/v1/trade-items-Histories/users/:user_id
```

#### Request Parameters
場所|パラメータ名|指定する値|説明|
|:-|:-|:-|:-|
|Path|user_id|文字列|対象ユーザーのID|


#### Example Response

+ Response 200 (application/json)

    <details>
    <summary>Body</summary>
    <pre>
    <code>
    [
        {
            "date": "2019-06-28",
            "trade_items": [
                {
                    "trade_name": "人をダメにするソファ",
                    "trade_image_url": "http://urx.blue/V5dk",
                    "trade_point": 2000
                }
            ]
        },
        {
            "date": "2019-06-27",
            "trade_items": [
                {
                    "trade_name": "スコッティ フラワーパック",
                    "trade_image_url": "http://ur0.work/W8jw",
                    "trade_point": 300
                },
                {
                    "trade_name": "クリネックス ティシュー",
                    "trade_image_url": "http://ur0.work/UR5I",
                    "trade_point": 400
                }
            ]
        }
    ]
    </code>
    </pre>
    </details>


#### Example Request

```
curl http://54.238.92.95:8080/api/v1/trade-items-Histories/users/goya
```
---