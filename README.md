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
|Body|user_id|文字列|ログインしたいユーザーID|
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
|Body|user_id|文字列|ログインしたいユーザーID|

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

### [POST] 材料IDに結びつく材料データをカートに追加する(ログイン必要)

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

### [GET] 与えられたキーを元に検索した結果のレシピデータを取得する

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




