# Account Manager

アカウント情報を効率的に検索・管理するためのコマンドラインツールです。

## 概要

Account Managerは、テキストファイルに保存されたアカウント情報を検索するためのインタラクティブなコマンドラインアプリケーションです。キーワード検索とファイル名検索の2つのモードを提供し、パスワードやアカウント情報を素早く見つけることができます。

## 機能

- **キーワード検索**: ファイル内容から特定のキーワードを含むセクションを検索
- **ファイル名検索**: ファイル名のパターンマッチングによる検索
- **インタラクティブUI**: 使いやすいターミナルユーザーインターフェース
- **色付き出力**: 検索結果をハイライト表示
- **セクション単位の表示**: `#`で始まる行でセクションを区切り、関連する情報をまとめて表示

## インストール・ビルド

### 前提条件

- Go 1.24.0以上

### ビルド方法

```bash
# リポジトリをクローン
git clone https://github.com/si-arakaki/account-manager.git
cd account-manager

# 依存関係をダウンロード
go mod tidy

# バイナリをビルド
go build -o ./dist/account ./cmd/account

# または、提供されているビルドスクリプトを使用（macOS ARM64用）
./build.sh
```

## セットアップ

### ディレクトリ構造

アカウント情報は以下のディレクトリに保存します：

- デフォルト: `~/.account/`
- カスタム: `ACCOUNT_HOME` 環境変数で指定

### ファイル形式

アカウント情報ファイルは以下の形式で作成します：

```
# Email Account
username: user@example.com
password: mypassword
server: mail.example.com
port: 993

# Bank Account
account_number: 123456789
routing: 987654321
bank: Example Bank
website: https://bank.example.com

# Social Media
platform: Twitter
username: @myhandle
password: socialpass
```

- `#` で始まる行がセクションヘッダーとなります
- 各セクションには関連するアカウント情報をまとめて記載します
- ファイル名は任意ですが、内容が分かりやすい名前を推奨します

## 使用方法

### 基本的な実行

```bash
# デフォルトディレクトリ（~/.account）を使用
./dist/account

# カスタムディレクトリを指定
ACCOUNT_HOME=/path/to/accounts ./dist/account
```

### モード選択

アプリケーションを実行すると、2つの検索モードから選択できます：

#### 1. キーワード検索モード (keyword)

指定したキーワードを含むセクションを検索します。

**使用例:**
1. モード選択で `keyword` を選択
2. 検索したいキーワードを入力（例: `password`, `bank`, `email`など）
3. 該当するセクションがハイライト付きで表示される

**特徴:**
- 大文字小文字を区別しない検索
- キーワードが含まれるセクション全体を表示
- 検索キーワードが黄色でハイライト
- ファイル名が灰色で表示

#### 2. ファイル名検索モード (filename)

ファイル名のパターンマッチングで検索します。

**使用例:**
1. モード選択で `filename` を選択  
2. ファイル名のパターンを入力（例: `email`, `bank`, `social`など）
3. マッチするファイルの全内容が表示される

**特徴:**
- 大文字小文字を区別しない検索
- 正規表現パターンに対応
- マッチしたファイルの全内容を表示

## 実行例

### キーワード検索の例

```bash
$ ./dist/account
mode?
> keyword
  filename

keyword?
> password
```

**出力例:**
```
/home/user/.account/email.txt
# Email Account
username: user@example.com
password: mypassword
server: mail.example.com

/home/user/.account/social.txt  
# Social Media
platform: Twitter
username: @myhandle
password: socialpass
```

### ファイル名検索の例

```bash
$ ./dist/account
mode?
  keyword
> filename

filename?
> email
```

**出力例:**
```
/home/user/.account/email.txt
# Email Account
username: user@example.com
password: mypassword
server: mail.example.com
port: 993

# Gmail Account
username: myname@gmail.com
password: gmailpass
```

## 環境変数

| 変数名 | 説明 | デフォルト値 |
|--------|------|-------------|
| `ACCOUNT_HOME` | アカウント情報ファイルを格納するディレクトリ | `~/.account` |

## 注意事項

- **セキュリティ**: パスワードなどの機密情報を扱うため、ファイルの権限設定に注意してください
- **バックアップ**: 重要なアカウント情報は定期的にバックアップを取ることを推奨します
- **ファイル形式**: `#` で始まる行をセクション区切りとして使用するため、通常のテキスト内では `#` の使用を避けてください

## トラブルシューティング

### よくある問題

1. **ディレクトリが見つからない**
   ```
   WARN no environment found ACCOUNT_HOME
   WARN using /home/user/.account instead
   ```
   - `~/.account` ディレクトリを作成するか、`ACCOUNT_HOME` 環境変数を設定してください

2. **正規表現エラー** 
   ```
   panic: regexp: Compile(`(?i)(*.txt)`): error parsing regexp
   ```
   - ファイル名検索で無効な正規表現パターンを入力した場合に発生します
   - 正確な正規表現パターンを入力してください（例: `.*\.txt` の代わりに `txt`）

3. **ファイルが見つからない**
   - `ACCOUNT_HOME` ディレクトリ内にテキストファイルが存在することを確認してください
   - ファイルの読み取り権限があることを確認してください

## ライセンス

このプロジェクトのライセンス情報については、リポジトリのライセンスファイルを参照してください。