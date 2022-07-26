# Frontend feature


| Features                | APIs                     | Method |          BE          |          FE          | Note |
| :---------------------- | :----------------------- | ------ | :------------------: | :------------------: | :--: |
| Create item             | /item                    | POST   |  :heavy_check_mark:  | :white_large_square: |      |
| Get items               | /item                    | GET    |  :heavy_check_mark:  | :white_large_square: |      |
| Get item                | /item/:sku               | GET    | :heavy_check_mark: | :white_large_square: |      |
| show item image         | /file/:filename          | GET    |  :heavy_check_mark:  | :white_large_square: |      |
| Update item             | /item/:sku               | PUT    |  :white_large_square:  | :white_large_square: |      |
| Create item type        | /item-type               | POST   |  :heavy_check_mark:  | :white_large_square: |      |
| Get all item type       | /item-type               | GET    |  :heavy_check_mark:  | :white_large_square: |      |
| Update item type        | /item-type/:item-type-id | PUT    | :white_large_square: | :white_large_square: |      |
| Create storage          | /storage                 | POST   |  :heavy_check_mark:  | :white_large_square: |      |
| Get storage             | /storage                 | GET    |  :heavy_check_mark:  | :white_large_square: |      |
| Get storage             | /storage/:storage-id     | GET    | :white_large_square: | :white_large_square: |      |
| Update storage          | /storage/:storage-id     | PUT    |  :heavy_check_mark:  | :white_large_square: |      |
| Create supplier         | /supplier                | POST   |  :heavy_check_mark:  | :white_large_square: |      |
| Get supplier            | /supplier                | GET    |  :heavy_check_mark:  | :white_large_square: |      |
| Get supplier            | /supplier/:supplier-id   | GET    | :white_large_square: | :white_large_square: |      |
| Update supplier         | /supplier/:supplier-id   | PUT    |  :heavy_check_mark:  | :white_large_square: |      |
| Login                   | /login                   | POST   |  :heavy_check_mark:  | :white_large_square: |      |
| Get user accounts       | /users                   | GET    | :white_large_square: | :white_large_square: |      |
| Reset user password     | /users/password          | PATCH  | :white_large_square: | :white_large_square: |      |
| Create import tickets   | /import_ticket           | POST   |  :heavy_check_mark:  | :white_large_square: |      |
| Get list import tickets | /import_ticket           | GET    | :white_large_square: | :white_large_square: |      |
| Get import ticket       | /import_ticket/:id       | GET    | :white_large_square: | :white_large_square: |      |
| Generate serials        | /import_ticket/serials   | POST   |  :heavy_check_mark:  | :white_large_square: |      |
| Get serial info         | /serial/:seri            | GET    |  :heavy_check_mark:  | :white_large_square: |      |
