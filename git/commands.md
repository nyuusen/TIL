# Gitコマンド集

ペア/モブプログラミングする時はコマンドの方がわかり良いのでコマンドを

## git diff

コミット間の差分を確認する

```
git diff <commit_hash_A> <commit_hash_B>
```

現在ブランチとの差分を確認する

```
git diff <commit_hash>
```

ファイル一覧のみを表示する

```
git diff --name-only
```

## git push 

-uオプションは、--set-upstreamの省略形(upstreamは上流という意味)  
現在ブランチを指定したリモートブランチに紐付けするという意味  

```
// (一度のみ実行)-uでリモートブランチを紐付けする
git push -u origin hogehoge

// 以降はgit pushのみで紐付けしたリモートブランチにプッシュできる
git push
```
