# 他己紹介サービス - バックエンド

他己紹介サービスのバックエンド処理全般

## 開発を始める

1. **Go1.13**をインストール
2. `cd $GOPATH/src`
3. `git clone https://github.com/~~`

### VSCodeの設定

https://qiita.com/ochipin/items/cae787d75ae91247c722

## 使用するライブラリ

- webフレームワーク：echo

- aws系ライブラリ： [aws](https://github.com/aws)/[aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2)


## パッケージ構成について

[今すぐ「レイヤードアーキテクチャ+DDD」を理解しよう。（golang）](https://qiita.com/tono-maron/items/345c433b86f74d314c8d) ←これがわかりやすいか
[Go のサーバーサイド実装におけるレイヤ設計と実装](https://www.slideshare.net/pospome/go-80591000)  
[Go の package 構成と開発のベタープラクティス](https://engineer.recruit-lifestyle.co.jp/techblog/2018-03-16-go-ddd/)  
[Go のパッケージ構成の失敗遍歴と現状確認](https://medium.com/@timakin/go%E3%81%AE%E3%83%91%E3%83%83%E3%82%B1%E3%83%BC%E3%82%B8%E6%A7%8B%E6%88%90%E3%81%AE%E5%A4%B1%E6%95%97%E9%81%8D%E6%AD%B4%E3%81%A8%E7%8F%BE%E7%8A%B6%E7%A2%BA%E8%AA%8D-fc6a4369337)  

- interface
  - APIのハンドラなどを定義
- domain
  - model
    - ドメイン情報（≒DBデータの構造体）の定義とドメインロジックを配置。
  - repository
    - データの永続化処理の**抽象**(`interface≒メソッド定義` を記述するだけ)を配置。
- infra
  - DB 等へのデータ永続化処理の**実装**を配置。(`domain/repository`で定義したメソッドを持つ構造体を定義する。このメソッドが、クエリ・ファイルUL/DLなどの処理の実態である)
- usecase
  - ハンドラ から呼ばれる Usecase を配置。
  - ビジネスロジックが実装される。
  - トランザクション（ある場合）はこのレイヤーで制御。
- di
  - 依存関係の解決用関数（要はNewしまくる感じ）を配置。



## 実装について

interfaceの目的は、mockを使ったテストをしやすくするため

- DBアクセス層とUsecase層を切り離すことでUsecase層のテストがやりやすくなる等

本当は全てにinterfaceをかましてあげるのが良いが、今回は時間短縮のためにデータ永続化処理（DBアクセス）のみに適応させる。

処理の呼び出し順的には

```
main.go(webフレームワークechoが常駐)
↓
interface/router.go(handler直接関数呼び出し)
↓
interface/handler(usecase直接関数呼び出し)
↓
usecase(domain/repositoryを経由して、infra層のメソッドを呼び出し)
↓
domain/repository
⬆︎
infra(domain/repository、domain/modelに依存)➡︎domain/model
```

### 実装手順（参考）

1. `domain/model`で構造体を定義する。`interface/router.go`にURLルーティングを定義する
2. `domain/repository`でメソッドをinterfaceで定義する（Save() 、Delete()、Update()、Find()など ）
3. `infra`に②の実装を記載する。この時、関数ではなく②の定義を満たすメソッドとして記載する。
4. `usecase`にビジネスロジックを記載する。（ここはinterface定義しなくて良い）
5. `handler`に、httpのGET、POSTなどを受け取ってusecase の処理を呼び出す関数を記載する（ここはinterface定義しなくて良い）
6. `di/injector.go`に依存性解決関数を記載する（ここが一番詰まるかも）

##　階層イメージ

```sh
.
├── README.md
├── domain
│   ├── model
│   │      ├── user.go
│   │      └── post.go	etc...
│   └── repository
│          ├── user_repo.go
│          └── post_repo.go	etc...
├── infra
│   ├── s3.go
│   └── dynamodb.go	etc...
├── interface
│   ├── handler
│   │      ├── user_handler.go
│   │      └── post_handler.go	etc...
│   └── router.go
├── usecase
│   ├── post.go
│   ├── view.go
│   └── register.go	etc...
├── di
│   └── injector.go
└── main.go
```



## テストについて

**usecase層に対しての単体テストは記載する。**

​	例：usecase の関数を引数与えて呼び出したら想定通りの値が返ってくることを確認する

基本的なテストの書き方はここらへん参照

- [goでmockを使ったテストをする](https://qiita.com/marnie_ms4/items/5925f136d23c8a0b4a4c)

VSCodeでは、関数のテスト雛形を自動で作ってくれる機能ある。

- [VS CodeのGo言語テストコード生成ツールを使ってみたらめちゃくちゃ便利だった話とか](http://kdnakt.hatenablog.com/entry/2019/01/03/080000)