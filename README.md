# firestore_clean_architecture
firestore + echo + wire + clean architecture をGoで作ってみました

## 起動手順

1. firestoreからCredentialsFileを取得し、firestore_clean_architectureフォルダ直下に配置

2. [こちらのコード部分](https://github.com/ryuto-imai/firestore_clean_architecture/blob/2bb8ce9f2c63de342ddd5a9efbdb4de3f460c909/drivers/database/firestore.go#L12)を上記で取得したファイル名に差し替え

3. `go run .`で起動

## API仕様

http://localhost:8000/users

- GET: 全ユーザーデータの取得

- POST: ユーザーデータの追加
    - Body例：
    ```
    {
        "name": "test",
        "age": 25,
        "address": "tokyo"
    }
    ```
