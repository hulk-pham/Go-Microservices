curl --location 'http://localhost:4080/api/index' \
--header 'Authorization: Basic YWRtaW46Q29tcGxleHBhc3MjMTIz' \
--header 'Content-Type: application/json' \
--data '{
    "name": "users",
    "storage_type": "disk",
    "shard_num": 1,
    "mappings": {
        "properties": {
            "id": {
                "type": "numeric",
                "index": true,
                "store": true,
                "highlightable": true
            },
            "first_name": {
                "type": "text",
                "index": true,
                "store": true,
                "highlightable": true
            },
            "last_name": {
                "type": "text",
                "index": true,
                "store": true,
                "highlightable": true
            },
            "email": {
                "type": "text",
                "index": true,
                "store": true,
                "highlightable": true
            },
            "hobby": {
                "type": "text",
                "index": true,
                "store": true,
                "highlightable": true
            },
            "phone_number": {
                "type": "text",
                "index": true,
                "store": true,
                "highlightable": true
            },
            "address": {
                "type": "text",
                "index": true,
                "store": true,
                "highlightable": true
            },
            "status": {
                "type": "keyword",
                "index": true,
                "sortable": true,
                "aggregatable": true
            },
            "dob": {
                "type": "date",
                "format": "2006-01-02T15:04:05Z07:00",
                "index": true,
                "sortable": true,
                "aggregatable": true
            },
            "avatar": {
                "type": "text",
                "index": true,
                "store": true,
                "highlightable": true
            },
            "created_at": {
                "type": "date",
                "format": "2006-01-02T15:04:05Z07:00",
                "index": true,
                "sortable": true,
                "aggregatable": true
            },
            "updated_at": {
                "type": "date",
                "format": "2006-01-02T15:04:05Z07:00",
                "index": true,
                "sortable": true,
                "aggregatable": true
            }
        }
    }
}'