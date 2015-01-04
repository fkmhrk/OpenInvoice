= API Document

== ログイン

ユーザー名/パスワードでアクセストークンを取得する。

//emlist{
 POST /api/v1/token

 {
   "username" : "fkm",
   "password" : "123456"
 }

 response
 
 HTTP 200 OK
 {
   "id" : "user1234",
   "access_token" : "token1234",
   "token_type" : "bearer"
 }
//}

== ユーザー作成

管理者トークンでのみ作成可能。

//emlist{
 Authorization:bearer {access_token}
 POST /api/v1/users

 {
   "username" : "fkm",
   "password" : "123456"
 }

 response

 HTTP 200
 {
   "id" : "user1234"
 }
//}


== 取引レコード取得

tradingテーブルの内容を取得する。

//emlist{
  Authorization:bearer {access_token}
  GET /api/v1/tradings

  response

  {
    "tradings" : [
      {
        "id" : "20150101",
        "company_id" : "会社ID",
        "subject" : "件名",
        "work_from" : 1122,
        "work_to" : 2233,
        "assignee" : "担当者ID",
        "product" : "成果物"
      },
      {
        "id" : "20150102",
        "company_id" : "会社ID2",
        "subject" : "件名",
        "work_from" : 1133,
        "work_to" : 2288,
        "assignee" : "担当者ID",
        "product" : "成果物"      
      }
    ]
  }
//}

== 取引レコード作成

tradingレコードを作成する。assigneeは権限のあるユーザーのみ指定可能。権限が無い場合はログイン中のユーザー。

//emlist{
  Authorization:bearer {access_token}
  PUT /api/v1/tradings/{取引ID}

  {
    "company_id" : "会社ID",
    "subject" : "件名",
    "work_from" : 1122,
    "work_to" : 2233,
    "assignee" : "担当者ID",
    "product" : "成果物"
  }

  response

  HTTP 204 No body
//}

== 取引の項目取得

指定した取引の項目を取得する。

//emlist{
  Authorization:bearer {access_token}
  GET /api/v1/tradings/{取引ID}/items

  response

  HTTP 200 OK
  {
    "items" : [
      {
        "id" : "item1122",
        "subject" : "件名",
        "unit_price" : 20000,
        "amount" : 2,
        "degree" : "人月",
        "tax_type" : 1,
        "memo" : "備考"
      },
      {
        "id" : "item2233",
        "subject" : "件名",
        "unit_price" : 30000,
        "amount" : 1,
        "degree" : "人月",
        "tax_type" : 2,
        "memo" : "備考"      
      }
    ]
  }
//}


== 取引の項目追加

指定した取引に項目を追加する。

//emlist{
  Authorization:bearer {access_token}
  POST /api/v1/tradings/{取引ID}/items

  {
    "subject" : "件名",
    "unit_price" : 20000,
    "amount" : 2,
    "degree" : "人月",
    "tax_type" : 1,
    "memo" : "備考"
  }

  response

  HTTP 201 Added
  {
    "id" : "item1122"
  }
//}

== 取引の項目修正

指定した取引項目を修正する。

//emlist{
  Authorization:bearer {access_token}
  PUT /api/v1/tradings/{取引ID}/items/{項目ID}

  {
    "subject" : "件名",
    "unit_price" : 20000,
    "amount" : 2,
    "degree" : "人月",
    "tax_type" : 1,
    "memo" : "備考"
  }

  response
  
  HTTP 204 No body
//}

== 取引の項目削除

指定した取引項目を削除する。

//emlist{
  Authorization:bearer {access_token}
  DELETE /api/v1/tradings/{取引ID}/items/{項目ID}

  response
  
  HTTP 204 No body
//}


