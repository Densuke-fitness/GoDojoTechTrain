# GoDojoTechTrain

## Launching the Application
Make sure that the godojotechtrain project is in the root folder and enter the following command .
```
docker compose up -d
```
### Check the mysql database.
```
docker exec -it go_db bash

mysql -u root -ppassw0rd

use techtraindb

#check users table .
select * from users;
```

### Check the api specification.
Search in a browser → `http://localhost:10000`

### Check the container logs.
Search in a browser → `http://localhost:9999`

## Trying out the API behavior by using "curl" command 

### CreateUser
You can throw the name in the request and get the token as a response .
```
curl -X POST -H "Content-Type: application/json" -d '{"name":"TestUser"}' localhost:8080/user/create
```
[![Image from Gyazo](https://i.gyazo.com/8dd2ceacfb56fdce9c861bf0a79847c8.png)](https://gyazo.com/8dd2ceacfb56fdce9c861bf0a79847c8)

### GetUser
Insert the token obtained by CreateUser into the x-auth-token in the http header section and throw a request to obtain the name　as a response　.
```
curl -H "X-Auth-Token:<token-key>" http://localhost:8080/user/get
```
[![Image from Gyazo](https://i.gyazo.com/83c265f9c1c243773058a3161e5a4c3f.png)](https://gyazo.com/83c265f9c1c243773058a3161e5a4c3f)
### UpdateUser
Insert the token obtained by CreateUser into the x-auth-token in the http header and throw a new name as a request to update the name .

```
curl -XPUT  -H "X-Auth-Token:<token-key>" -d '{"name":"TestNewUser"}' http://localhost:8080/user/update 
```

[![Image from Gyazo](https://i.gyazo.com/7b337faf2602fc828cbac398ff666fdf.png)](https://gyazo.com/7b337faf2602fc828cbac398ff666fdf)

### DrawGacha

Insert the token obtained by CreateUser into the x-auth-token in the http header and throw drawing gacha times as a request to draw gacha .
```
curl -X POST -H "X-Auth-Token:<token-key>" -d '{"times":10}' localhost:8080/gacha/draw
```

### GetCharacterList
Insert the token obtained by CreateUser into the x-auth-token in the http header section and throw a request to obtain the characters you registerd　as a response　.
```
curl -H "X-Auth-Token:<token-key>" http://localhost:8080/character/list
```

## Testing

It will be launched for CI test execution. It will not work in a local environment without the go environment.

```
sh testing.sh 
```

## ERD

### written by draw.io
- [TechTrainERD.drawio](https://drive.google.com/file/d/1eJ9z5BW58hKN4MKJkjldIqqWgjXTSEEe/view?usp=sharing)

### 
If you want to get ERD by using  schemaspy, please enter the following command .

```
docker run -v "$PWD/schema:/output" --net="host" schemaspy/schemaspy:snapshot \
 -t mysql -host 0.0.0.0:3306 -db techtraindb -u root -p passw0rd -s  techtraindb 
 ```

### Reference
- [オンライン版　CA Tech Dojo サーバサイド (Go)編](https://techbowl.co.jp/techtrain/missions/12)
- [DeNA Codelabs: テスタビリティの高いGoのAPIサーバを開発しよう（その1 ~準備&E2E実装編~）](https://dena.github.io/codelabs/testable-architecture-with-go-part1/#0)
- [GolangCI-LintとCircleCIを利用して静的解析チェックを自動化する](https://blog.mmmcorp.co.jp/blog/2021/01/10/golangci-lint-circleci/)
- [Github ActionsでGoのCI環境を作成する](https://blog.mmmcorp.co.jp/blog/2021/01/10/golangci-lint-circleci/)
- [Go言語で理解するJWT認証 実装ハンズオン](https://qiita.com/po3rin/items/740445d21487dfcb5d9f)
- [Go database/sql tutorial](http://go-database-sql.org/index.html)
- [コンテナで動くDBとアプリとの疎通をGithub Actionsでテストする](https://times.hrbrain.co.jp/entry/test-using-container-on-github-actions)
- [構造体のフィールドをJSON Schemaでvalidationする](https://qiita.com/nownabe/items/e62bbedd2e1687d1ff70)
- [Goでサーバを立ち上げてE2Eテストを実施するCI用のテストコードを書く](https://budougumi0617.github.io/2020/03/27/http-test-in-go/)
- [Goでhttpリクエストを送信する方法](https://qiita.com/taizo/items/c397dbfed7215969b0a5)
- [schemaspy/schemaspy](https://hub.docker.com/r/schemaspy/schemaspy/)
- [DockerでサクッとDBからER図を作成する](https://qiita.com/kamukiriri/items/ab1669c19831c18db9ee)
- [Docker でサクッと MySQL 8 からER図を作成する](https://qiita.com/ngyuki/items/4efa0734e8d8582bfc16)
