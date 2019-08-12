# go coding

## panic/recover

- 外部に公開する API はエラーを伝える手段として panic() を使ってはいけない
  - 多値の返り値 w/ error インターフェースを使うこと
  - panic() はパッケージをまたいで伝搬させることがないようにする
- panic/recover 機構はスタックが深くなるような呼び出しをする際に有用
  - 使用することで可読性を向上させることができる

つまり、外部にエラーを伝える手段としては error インターフェースを使うのが正式なやり方で、panic/recover はあくまで内部的に用いるべきもの、ということらしい。

引用： [Golang の defer 文と panic/recover 機構について](https://blog.amedama.jp/entry/2015/10/11/123535)
