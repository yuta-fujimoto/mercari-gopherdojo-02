with open('words.txt') as f:
    # ファイルの内容を読み出す
    data = f.read()  
    # data = data.lower()  # 小文字にするならこのタイミングが楽


# 単語カウント-------------------------
# 単語を数える辞書を作成
words = {}
# split()でスペースと改行で分割したリストから単語を取り出す
for word in data.split():
    # 単語をキーとして値に1を足していく。
    # 辞書に単語がない、すなわち初めて辞書に登録するときは0+1になる。
    words[word] = words.get(word, 0) + 1  #

# リストに取り出して単語の出現回数でソート
d = [(v, k) for k, v in words.items()]
d.sort()
d.reverse()

# 標準出力-------------------------
for count, word in d:
	if count > 10 and len(word) > 5:
	    print(count, word)
