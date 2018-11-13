# ga2cw
Google Analyticsで取得したPVをチャットワークに自動で送信する<br><br>
## 使い方<br>
[Google API Console - Google Cloud Console](https://console.cloud.google.com/apis)でAnalytics APIあたりを有効にし、サービスアカウントキーを取得し、secret.jsonにリネームした後、後述するga2cwディレクトリに配置する。<br>
この内容に関しては下記の記事に詳しい説明があります。<br>
[【最新版】GoogleAnalytics API 認証設定について画像多めで説明します。 | 東京上野のWeb制作会社LIG](https://liginc.co.jp/356517)<br><br>
チャットワークのサイトを開き、アクセストークンの取得と自動送信先のルームIDを調べます。方法は下記の記事に詳しい説明があります。<br>
[PHPでチャットワークAPIを介してメッセージを投稿してみる - saitodev.co](https://saitodev.co/article/1695)<br>
チャットワークの各値を取得したら、config.json.sampleをconfig.jsonにリネームし、各値を挿入します。<br>
最後にGoogle Analyticsを開き、プロファイルIDを調べて、config.jsonに挿入すれば設定は終了です。<br>
プロファイルIDについては下記の記事に詳しい説明があります。<br>
[Go言語でGoogle Analyticsのデータをチャットワークに送ってみる - saitodev.co](https://saitodev.co/article/2061)<br><br>
設定終了後は下記のコマンドを参考に動作するかお試しください。<br><br>
```
go get golang.org/x/oauth2
go get golang.org/x/oauth2/google
go get google.golang.org/api/analytics/v3
cd サンプルコードのクローンを配置したいディレクトリ
git clone https://github.com/ryoko-saito/ga2cw.git
cd ga2cw
GOOGLE_APPLICATION_CREDENTIALS=secret.json go run main.go
```
