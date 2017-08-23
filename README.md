# go-crud

### 実行方法

##### 1. glide up を実行（※ glide コマンドインストール済みであることが前提）

###### ※Mac or Linux環境の場合、下記コマンドでインストール可能

###### curl https://glide.sh/get | sh

###### 参考：「https://github.com/Masterminds/glide」

##### 2. 以下３種の環境変数をセット

###### 「MYSQL_USER」=MySQL接続ユーザ

###### 「MYSQL_PASS」=MySQL接続パスワード

###### 「MYSQL_SCHEMA」=MySQL接続スキーマ

##### 3. cmd/config.toml 内を任意の設定に変更

##### 4. cmd/ に移動し、 go run main.go を実行
