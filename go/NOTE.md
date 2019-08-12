# go lang note

## go env

| env var | description |
| ------- | ----------- |
|GOARCH|コンパイルした実行ファイルが実行可能なCPUアーキテクチャを指定します。|amd64|
|GOBIN|$GOBINを設定すると、`go get`で取得した時のコマンドのインストール先になります。|
|GOCACHE|/root/.cache/go-build|
|GOEXE||
|GOFLAGS||
|GOHOSTARCH|amd64|
|GOHOSTOS|linux|
|GOOS|コンパイルした実行ファイルを動作させるOSを指定します。|linux|
|GOPATH|ワーキングディレクトリを指定します。|/go|
|GOPROXY||
|GORACE||
|GOROOT|開発ツールのインストールパスを指定します。|/usr/local/go|
|GOTMPDIR||
|GOTOOLDIR|/usr/local/go/pkg/tool/linux_amd64|
|GCCGO|gccgo|
|CC|gcc|
|CXX|g++|
|CGO_ENABLED|1|
|GOMOD||
|CGO_CFLAGS|-g -O2|
|CGO_CPPFLAGS||
|CGO_CXXFLAGS|-g -O2|
|CGO_FFLAGS|-g -O2|
|CGO_LDFLAGS|-g -O2|
|PKG_CONFIG|pkg-config|
|GOGCCFLAGS|-fPIC -m64 -pthread -fno-caret-diagnostics -Qunused-arguments -fmessage-length=0 -fdebug-prefix-map=/tmp/go-build089226586=/tmp/go-build -gno-record-gcc-switches|


## 依存関係の管理

https://budougumi0617.github.io/2019/02/15/go-modules-on-go112/

### dep

- Gopkg.toml マニフェストファイルで、プロジェクトの依存関係をvendorディレクトリで管理する
- Gopkg.lock にてパッケージのバージョンをより厳密に管理します

### go get

- `$GOPATH` を中心としてパッケージ取得管理を行う
- go 1.Xを使っているとき、go getで取得されるコードはまずgo1 branch or tagから取得される
- branchにもtagにもgo1が存在しなかったときにmasterのコードが取得される

### go mod

- Go Modules(vgo)はGo1.11から導入され始めたGoの新しいバージョン管理
- Go1.12ではまだ有効にはなっていない(Go1.13からはデフォルトで有効になる）
- Go ModulesはSemantic Versioningに基づいてモジュールの管理を行なう
- Go Modulesは可能な限り古いバージョンを採用するのが特徴的
- go.mod, go.sum というファイルにて管理される
- `GO111MODULE=(on|off|auto)` 環境変数により動作が変わる

## ビルド

https://qiita.com/mumoshu/items/0d2f2a13c6e9fc8da2a4

- go build -ldflags "-s -w"


## 開発前に読んでおきたい

https://qiita.com/bushiyama/items/d7b8f466c5a9d33a7230

EN: https://github.com/golang/go/wiki/CodeReviewComments
JA: https://qiita.com/knsh14/items/8b73b31822c109d4c497 

https://golang.org/doc/effective_go.html

## 開発支援ツール

```shell
brew install golangci-lint
go get golang.org/x/tools/cmd/goimports
```


## Emacs

```shell
go get -u -v golang.org/x/tools/cmd/...
```

[SPACEMACS - Go layer](http://develop.spacemacs.org/layers/+lang/go/README.html)


| keybind   | command                              |
| -------   | -------                              |
| SPC c C   | compile                              |
| SPC c r   | recompile                            |
| M-RET x x | spacemacs/go-run-main                |
| SPC e L   | spacemacs/goto-flycheck-error-list   |
| SPC e l   | spacemacs/toggle-flycheck-error-list |
|           | projectile-run-project               |

### spacemacs-go-mode-map

| keybind     | command                                   | description |
| -------     | -------                                   | ----------- |
| M-RET x x   | spacemacs/go-run-main                     |             |
| M-RET i a   | go-import-add                             |             |
| M-RET i g   | go-goto-imports                           |             |
| M-RET i r   | go-remove-unused-imports                  |             |
| M-RET h h   | godoc-at-point                            |             |
| M-RET e b   | go-play-buffer                            |             |
| M-RET e d   | go-download-play                          |             |
| M-RET e r   | go-play-region                            |             |
| M-RET f <   | go-guru-callers                           |             |
| M-RET f >   | go-guru-callees                           |             |
| M-RET f c   | go-guru-peers                             |             |
| M-RET f d   | go-guru-describe                          |             |
| M-RET f e   | go-guru-whicherrs                         |             |
| M-RET f f   | go-guru-freevars                          |             |
| M-RET f i   | go-guru-implements                        |             |
| M-RET f j   | go-guru-definition                        |             |
| M-RET f o   | go-guru-set-scope                         |             |
| M-RET f p   | go-guru-pointsto                          |             |
| M-RET f r   | go-guru-referrers                         |             |
| M-RET f s   | go-guru-callstack                         |             |
| M-RET t P   | spacemacs/go-run-package-tests-nested     |             |
| M-RET t g   | Prefix Command                            |             |
| M-RET t p   | spacemacs/go-run-package-tests            |             |
| M-RET t s   | spacemacs/go-run-test-current-suite       |             |
| M-RET t t   | spacemacs/go-run-test-current-function    |             |
| M-RET t v   | spacemacs/toggle-go-test-verbose          |             |
| M-RET r F   | go-tag-remove                             |             |
| M-RET r N   | go-rename                                 |             |
| M-RET r d   | godoctor-godoc                            |             |
| M-RET r e   | godoctor-extract                          |             |
| M-RET r f   | go-tag-add                                |             |
| M-RET r i   | go-impl                                   |             |
| M-RET r n   | godoctor-rename                           |             |
| M-RET r s   | go-fill-struct                            |             |
| M-RET r t   | godoctor-toggle                           |             |
| M-RET g G   | spacemacs/jump-to-definition-other-window |             |
| M-RET g a   | ff-find-other-file                        |             |
| M-RET g c   | go-coverage                               |             |
| M-RET g g   | spacemacs/jump-to-definition              |             |
| M-RET t g F | go-gen-test-all                           |             |
| M-RET t g f | go-gen-test-exported                      |             |
| M-RET t g g | go-gen-test-dwim                          |             |

### spacemacs-lsp-mode-map

| keybind   | command                                    | description |
| -------   | -------                                    | ----------- |
| M-RET F a | lsp-workspace-folders-add                  |             |
| M-RET F r | lsp-workspace-folders-remove               |             |
| M-RET F s | lsp-workspace-folders-switch               |             |
| M-RET T F | spacemacs/lsp-ui-doc-func                  |             |
| M-RET T I | spacemacs/lsp-ui-sideline-ignore-duplicate |             |
| M-RET T S | spacemacs/lsp-ui-sideline-symb             |             |
| M-RET T d | lsp-ui-doc-mode                            |             |
| M-RET T l | lsp-lens-mode                              |             |
| M-RET T s | lsp-ui-sideline-mode                       |             |
| M-RET r r | lsp-rename                                 |             |
| M-RET b a | lsp-execute-code-action                    |             |
| M-RET b d | lsp-describe-session                       |             |
| M-RET b r | lsp-restart-workspace                      |             |
| M-RET b s | lsp-shutdown-workspace                     |             |
| M-RET h h | lsp-describe-thing-at-point                |             |
| M-RET = b | lsp-format-buffer                          |             |
| M-RET = r | lsp-format-region                          |             |
| M-RET G d | lsp-ui-peek-find-definitions               |             |
| M-RET G e | lsp-ui-flycheck-list                       |             |
| M-RET G i | lsp-ui-peek-find-implementation            |             |
| M-RET G n | lsp-ui-peek-jump-forward                   |             |
| M-RET G p | lsp-ui-peek-jump-backward                  |             |
| M-RET G r | lsp-ui-peek-find-references                |             |
| M-RET G s | lsp-ui-peek-find-workspace-symbol          |             |
| M-RET g K | spacemacs/lsp-avy-goto-symbol              |             |
| M-RET g M | lsp-ui-imenu                               |             |
| M-RET g S | helm-lsp-global-workspace-symbol           |             |
| M-RET g d | xref-find-definitions                      |             |
| M-RET g e | lsp-treemacs-errors-list                   |             |
| M-RET g i | lsp-find-implementation                    |             |
| M-RET g k | spacemacs/lsp-avy-goto-word                |             |
| M-RET g p | xref-pop-marker-stack                      |             |
| M-RET g r | xref-find-references                       |             |
| M-RET g s | helm-lsp-workspace-symbol                  |             |
| M-RET g t | lsp-find-type-definition                   |             |

- https://github.com/weijiangan/flycheck-golangci-lint
- https://github.com/masasam/dotfiles/blob/master/.emacs.d/inits/23go.el
- https://solist.work/blog/posts/language-server-protocol/
- http://glassonion.hatenablog.com/entry/2019/05/11/134135

### dap-mode

+ `dap-go-setup` で [VSCode Go Extension](https://marketplace.visualstudio.com/items?itemName=ms-vscode.Go)をインストール
+ `node.js` をインストール
+ `go get -u github.com/go-delve/delve/cmd/dlv` delveのインストール
+ `dap-debug` を実行

`dap-hydra` が便利

## playground

- [calcurate tax example](https://play.golang.org/p/v5mP-_KEce)


- [【C言語講座】第3回 小数と入力のはなし](https://ch.nicovideo.jp/nikosai/blomaga/ar878343)
