# Discord-randomizer

v1.0

任意のn面ダイスを複数個同時に振ることができます。

## インストール

### DiscordBotの作成

以下の手順に従ってDiscordBotを作成する。

1. https://discord.com/developers/applications で`New Application`をクリック。

2. 適当な名前をつけて`Create`をクリック

3. `SETTINGS`カテゴリの`OAuth2`を選択し、`OAuth2 URL Generator`から`bot`を選択する。

4. 画面下部に`BOT PERMISSIONS`が表示されるので、`Send Messages`にチェックを入れる。

5. `SETTINGS`カテゴリの`Installation`を選択し、画面最下部の`Guild Install`から`SCOPES`に`bot`を追加し、`PERMISSIONS`に`Send Messages`を追加する。

6. 下に表示されている`GENERATED URL`に表示されているURLをコピーして適当なWEBで開ければDiscordBotをサーバーに追加できる。

7. `SETTINGS`カテゴリの`Bot`を選択し、表示されている`TOKEN`を保存しておく。なければ`Reset Token`をクリックして再生成する。

### .envファイルの設定

以下の`.env.sample`を参考に`mainパッケージ`があるディレクトリの直下に`.env`ファイルを作成する。

```
DISCORD_TOKEN=xxx
```

|名称|内容|
|--|--|
|DISCORD_TOKEN|DiscordBotのトークン|

### 使い方

上記の手順に従ってDiscordBotと`.env`ファイルの設定を完了させる。

この状態で`mainパッケージ`があるディレクトリから以下のコマンドを実行する。

```
nohup go run . &
```

これでバックグラウンドでコードを実行できるので、この間Botが使用可能になる。

## 機能と動作

サービスの機能と動作についての説明。

### `/random mdn`

n面ダイスをm回振ることができるコマンドです。また、各ダイスの出目の総和も出力されます。

### `/help`

機能やコマンドの説明を表示するコマンドです。

## 注意点

一度コードの実行を止めると5分程度コードの実行が不能になります。