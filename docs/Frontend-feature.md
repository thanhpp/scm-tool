# Frontend feature


| Features                | APIs                     | Method |          BE          |          FE          | Note |
| :---------------------- | :----------------------- | ------ | :------------------: | :------------------: | :--: |
| Create item             | /item                    | POST   |  :heavy_check_mark:  | :white_large_square: |      |
| Get items               | /item                    | GET    |  :heavy_check_mark:  | :white_large_square: |      |
| Get item                | /item/:sku               | GET    |  :heavy_check_mark:  | :white_large_square: |      |
| show item image         | /file/:filename          | GET    |  :heavy_check_mark:  | :white_large_square: |      |
| Update item             | /item/:sku               | PUT    |  :heavy_check_mark:  | :white_large_square: |      |
| Create item type        | /item-type               | POST   |  :heavy_check_mark:  |  :heavy_check_mark:  |      |
| Get all item type       | /item-type               | GET    |  :heavy_check_mark:  |  :heavy_check_mark:  |      |
| Update item type        | /item-type/:item-type-id | PUT    |  :heavy_check_mark:  |  :heavy_check_mark:  |      |
| Create storage          | /storage                 | POST   |  :heavy_check_mark:  |  :heavy_check_mark:  |      |
| Get storage             | /storage                 | GET    |  :heavy_check_mark:  |  :heavy_check_mark:  |      |
| Get storage             | /storage/:storage-id     | GET    | :white_large_square: |  :heavy_check_mark:  |      |
| Update storage          | /storage/:storage-id     | PUT    |  :heavy_check_mark:  |  :heavy_check_mark:  |      |
| Create supplier         | /supplier                | POST   |  :heavy_check_mark:  |  :heavy_check_mark:  |      |
| Get list supplier       | /supplier                | GET    |  :heavy_check_mark:  |  :heavy_check_mark:  |      |
| Get supplier            | /supplier/:supplier-id   | GET    | :white_large_square: |  :heavy_check_mark:  |      |
| Update supplier         | /supplier/:supplier-id   | PUT    |  :heavy_check_mark:  |  :heavy_check_mark:  |      |
| Login                   | /login                   | POST   |  :heavy_check_mark:  | :white_large_square: |      |
| Get user accounts       | /users                   | GET    |  :heavy_check_mark:  | :white_large_square: |      |
| Reset user password     | /users/password          | PATCH  |  :heavy_check_mark:  | :white_large_square: |      |
| Create import tickets   | /import_ticket           | POST   |  :heavy_check_mark:  | :white_large_square: |      |
| Get list import tickets | /import_ticket           | GET    | :white_large_square: | :white_large_square: |      |
| Get import ticket       | /import_ticket/:id       | GET    | :white_large_square: | :white_large_square: |      |
| Generate serials        | /import_ticket/serials   | POST   |  :heavy_check_mark:  | :white_large_square: |      |
| Get serial info         | /serial/:seri            | GET    |  :heavy_check_mark:  | :white_large_square: |      |
