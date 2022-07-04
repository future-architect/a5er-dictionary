# a5er-dictionary

[![Apache-2.0 license](https://img.shields.io/badge/license-Apache2.0-blue.svg)][license]

[license]: https://github.com/future-architect/a5er-dictionary/blob/main/LICENSE

A5:Mk-2専用 論物変換スクリプトです。

## 使い方

## チュートリアル

1.辞書ファイル（CSV形式)の用意
* dict/dict.txtで辞書ファイルを用意する。
* 下記のように論理名・物理名の順に定義する。

<pre>
ID,id
記事,article
ヘッダ,header
タイトル,title
内容,content
</pre>

2.A5:Mk-2のa5erを用意

3.変換ツールのインストール

4.コマンド実行

変換に失敗した場合は XXX と表示されます。表示がなくなるまで、辞書ファイルに追加します。

### オプション
#### ERD_PATH=xxx（必須）
a5erファイルのパスを指定する。

#### DICT_PATH=xxx
辞書ファイルの位置を指定する。デフォルトはdict/dict.txtになる。

#### OUTPUT_PATH=xxx
物理名を付与したa5erファイルの保存先を指定する。

#### TABLE_PLURAL=(true|false)
テーブル名を複数形にするかどうかを指定する。デフォルトはtrue。
Rails採用の場合はtrue、それ以外の場合はfalseになるのが一般的。

## インストール

```
$ go install github.com/future-architect/a5er-dictionary/cmd@latest
```
