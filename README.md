# GO --RESTful APIアプリケーション実装例--

## 使用フレームワーク
### Beego

Beegoを使用した理由はMVCモデルに基づいた設計とその汎用性。
標準実装されたORMを始め、構築がとてもしやすい。
GO特有である、ポインタ概念を使用したデータベース上へのアクセスを踏まえ、
MVCを用いた責任分散の簡易化が可能

比較的大規模なアプリケーション使用を想定しています。

## 初期化処理
・app.confの読み込み
・DBアクセス
・ルーティング
・ログ
・コントローラ/サービス DI

