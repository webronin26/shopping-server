# shopping-server
這是一個模擬購物商城的伺服器端
配合 [shopping-android版本](https://github.com/webronin26/shopping) 的 client 端來使用

## 目前的功能

| API  |
| ------------- |
| (GET) /main |
- 回傳首頁的圖片url / 影片 / 首頁宣傳資訊 

| API  |
| ------------- |
| (GET) /search?q=keywords |
- 回傳夾有關鍵字名稱的產品 

| API  |
| ------------- |
| (GET) /query/product/{product_id} |
- 回傳某產品詳細資訊

| API  |
| ------------- |
| (GET) /query/product/{product_id} |
- 回傳某產品詳細資訊

| API  |
| ------------- |
|(POST) /login|
| (Request body) email : 使用者email / password : 使用者密碼|
- 使用email和密碼登入會員，驗證成功之後回傳 token 值

| API  |
| ------------- |
|(POST) /member/buy|
|(Header) Authorization : 使用者的token|
|(Request body) product_id : 產品的id / number : 要購買的數量|
- 購買產品

| API  |
| ------------- |
|(DELETE) /member/logout |
|(Header) Authorization : 使用者的 token|
- 登出

## 目前的Table

accounts 表紀錄用戶「enail」，「姓名」，「密碼」和「雜湊密鑰」

| Field        |  Type       |   Null       |   Key        |   Default  |      Extra  |
| ------------- | ------------- | ------------- | ------------- | ------------- | ------------- |
| id       | int(11)      | NO   | PRI | NULL    | auto_increment |
| email    | varchar(255) | NO   |  UNI   | NULL    |        .        |
| name     | varchar(255) | NO   | .    | NULL    |      .          |
| password | varchar(255) | NO   | .    | NULL    |   .             |
| token    | varchar(255) | YES  | UNI | NULL    |   .             |

products 表紀錄「產品名稱」，「產品價錢」，「產品資訊」，「圖片網址」，「產品擁有者」，「最大一次購買量」，「目前產品數量」，「開放購買的時間」

| Field        |  Type       |   Null       |   Key        |   Default  |      Extra  |
| ------------- | ------------- | ------------- | ------------- | ------------- | ------------- |
| id             | int(11)      | NO   | PRI | NULL    | auto_increment |
| name           | varchar(255) | NO   |   .  | NULL    |        .        |
| price          | int(11)      | NO   |  .   | NULL    |             .   |
| product_info   | varchar(255) | YES  |  .   | NULL    |         .       |
| image_url      | varchar(255) | YES  |  .   | NULL    |          .      |
| owner          | varchar(255) | NO   |  .   | NULL    |           .     |
| max_buy        | int(11)      | NO   | .    | NULL    |             .   |
| item_number    | int(11)      | NO   |    . | NULL    |          .      |
| available_time | timestamp    | YES  |   .  | NULL    |          .      |

records 表紀錄「目前被購買的產品序號」，「目前購買者的使用者序號」，「購買的數量」，「購買的時間」

| Field        |  Type       |   Null       |   Key        |   Default  |      Extra  |
| ------------- | ------------- | ------------- | ------------- | ------------- | ------------- |
| id          | int(11)   | NO   | PRI | NULL    | auto_increment |
| product_id  | int(11)   | YES  |    . | NULL    |     .           |
| customer_id | int(11)   | YES  |   .  | NULL    |    .            |
| number      | int(11)   | YES  |    . | NULL    |      .          |
| date        | timestamp | YES  |  .   | NULL    |      .          |