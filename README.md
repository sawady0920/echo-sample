# 立ち上げ
docker-compose up -d 

# できたものみる
docker ps

# アプリコンテナの中に入る
docker exec -it echosample_app_1 sh
※echosample_app_1はdocker psで表示される

# mysqlコンテナに入る
docker exec -it echosample_echo-db_1 sh
※echosample_db_1はdocker psで表示される