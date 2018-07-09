# Qiitter
Qiita + Twitter Bot

## Feature

特定のQiitaタグに関するQiita記事の投稿をつぶやくBotです。
環境変数の設定で好きなtwitterBotがつくれます。
Golang + AWS Lambda + Qiita API + Twitter API

## Quick Start

下記の環境変数を設定

```bash
export accessSecret= {アクセスキー}
export accessToken= {アクセストークン}
export consumerKey= {カスタムキー}
export consumerSecret= {カスタムシークレット}
export TARGET_TAG= {Qiitaのターゲットにするタグ名}
export HASH_TAG= {twitterにつけるハッシュタグ}
```

and execute !

```
$ dep ensure
$ go run main.go
```

## Twitter Bot Example

環境変数を"Go"にした際のサンプルは下記

|index|value|
|---|---|
|ID|@Po3rinB|
|URL|https://twitter.com/Po3rinB|
|Qiita Target Tag|"Go"|
