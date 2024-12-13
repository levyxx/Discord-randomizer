# Discord-randomizer

v1.0(Release:2024/12/12)
任意のn面ダイスを複数個同時に振ることができます。(/random)

v1.1(Release:2024/12/13)
複数の要素から1つをランダムに出力することができるようになりました。(/select)

## インストール

### DiscordBotの作成

以下の手順に従ってDiscordBotを作成する。

1. https://discord.com/developers/applications で`New Application`をクリック。

2. 適当な名前をつけて`Create`をクリック

3. `SETTINGS`カテゴリの`Installation`を選択し、画面最下部の`Guild Install`から`SCOPES`に`bot`を追加し、`PERMISSIONS`に`Send Messages`を追加する。

4.  `SETTINGS`カテゴリの`Bot`を選択し、`MESSAGE CONTENT INTENT`を有効にする。

5. `SETTINGS`カテゴリの`OAuth2`を選択し、`OAuth2 URL Generator`から`bot`を選択する。

6. 画面下部に`BOT PERMISSIONS`が表示されるので、`View Channels`, `Send Messages`, `Read Message History`にチェックを入れる。 

7. 下に表示されている`GENERATED URL`に表示されているURLをコピーして適当なWEBで開ければDiscordBotをサーバーに追加できる。

8. `SETTINGS`カテゴリの`Bot`を選択し、表示されている`TOKEN`を保存しておく。なければ`Reset Token`をクリックして再生成する。

### .envファイルの設定

以下の`.env.sample`を参考に`mainパッケージ`があるディレクトリの直下に`.env`ファイルを作成する。

```
DISCORD_TOKEN=xxx
```

|名称|内容|
|--|--|
|DISCORD_TOKEN|DiscordBotのトークン。DiscordBot作成手順の8.で保存したもの。|

## 使い方

上記の手順に従ってDiscordBotと`.env`ファイルの設定を完了させる。

この状態で`mainパッケージ`があるディレクトリから以下のコマンドを実行する。

```
nohup go run . &
```

これでバックグラウンドでコードを実行できるので、この間Botが使用可能になる。

### 運用

`WSL`を使用している場合、`sudo nano /etc/resolv.conf`でDNSサーバーを以下のように設定する。

```
nameserver 8.8.8.8
nameserver 8.8.4.4
```

`WSL`を再起動しても変更されないように、`/etc/wsl.conf`ファイルに以下の内容を追記する。

```
[network]
generateResolvConf = false
```

設定を反映させるために`WSL`を再起動する。

この設定を行うことによって一度実行を止めてもすぐに再開することが可能になる。

以後WSLからのログアウト後に実行したい場合は``sudo nano /etc/resolv.conf`で開いたファイルに以下の内容を書き込んでから実行すれば良い。

```
nameserver 8.8.8.8
nameserver 8.8.4.4
```

## 機能と動作

サービスの機能と動作についての説明。

### `/random mdn`

n面ダイスをm回振ることができるコマンドです。また、各ダイスの出目の総和も出力されます。

### `/select v1 v2 ... vn`

v1 v2 ... vnのうち1つをランダムに出力するコマンドです。

### `/help`

機能やコマンドの説明を表示するコマンドです。

### 他の使用方法

上記コマンドの`/`の部分を`!`に変更してもコマンドを使用することが可能です。