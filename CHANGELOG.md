## Добавление баз данных
### OpenSearch
Просто так к нему не подключишься
```
docker run -d --name opensearch-dashboards \
  -p 5601:5601 \
  -e OPENSEARCH_HOSTS=http://localhost:9200 \
  -e OPENSEARCH_USERNAME=admin \
  -e OPENSEARCH_PASSWORD=Opensearch1_ \
  opensearchproject/opensearch-dashboards:latest
```

... 
По итогу получилось впендюрить все ноды нормально, OpenSearch использует самоподписанные сертификаты вместе с OpenSearch Dashboards

## Postgres Tables
+ В ./db_datagen лежит файлик создания всех таблиц из логической схемы


## Разработка кода
+ Все компоненты для баз данных валяются в ./src/go.mod



