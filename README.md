# mercari-gopherdojo-02

# ex00

## 概要

タイピングゲーム

## build

```
go mod tidy
go build
```

## usage

```
./typing_game
```

## 注釈

* ゲーム時間は30秒です

* https://dragonquest.fandom.com/wiki/List_of_monsters_in_Dragon_Quest_Monsters:_Terry%27s_Wonderland_3D?action=edit からデータをいただきました。

# ex01

## 概要

分割ダウンローダ

## build

```
go mod tidy
go build
```

## usage

```
./download [url]
```

## 注釈

* range-accessを3回送信して高速化します。
* ファイルの容量があまり多くない場合にはcurlコマンドと大差ないか、遅くなる可能性があります。
