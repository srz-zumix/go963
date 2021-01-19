# go963

勤怠管理用 Google カレンダー・スケジュールツール

## 想定

 * Slack などのチャットツールからコマンド実行され、共通のカレンダーに勤怠を書き込む
 * 共通のカレンダーの通知機能やツール連携を利用してその日の勤怠連絡を流す
 * go963 でリスト取得可能なのでそれ使って通知しても良い

## 使い方

```
google calendar tool

Usage:
  go963 [command]

Available Commands:
  createEvents create event from events.json
  createToken  create OAuth 2.0 cache token
  delKintai    delete Kintai
  eventJson    event json format
  help         Help about any command
  listKintai   list Kintai
  setKintai    create / edit Kintai

Flags:
      --cacheToken              cache the OAuth 2.0 token (default true)
  -t, --cacheTokenFile string   cache the OAuth 2.0 token path. if empty, auto-generate
  -c, --calendarid string       Google Calendar Id.
      --config string           config file (default is $HOME/.go963.yaml)
  -h, --help                    help for go963
  -p, --prefix string           go963 google calendar event prefix (default "[Go963]")
  -s, --secretfile string       OAuth 2.0 Client Secret JSON. Default client_secret.json. (default "client_secret.json")
      --strict                  if true, control go963 created event only

Use "go963 [command] --help" for more information about a command.

```

使うコマンドは setKintai/delKintai/listKintai を想定してる。
Google Calendar API を使うためのクライアントシークレット json が必要。
最初に認証トークンを取得するためにブラウザが開く。一度取得したトークンを使えばトークン取得はスキップ可能。
上記2つをシークレットファイルとして用意すればたぶん OK。

## 構成例

