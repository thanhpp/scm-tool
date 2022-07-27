# Frontend feature


| Features                | APIs                     | Method |              BE               |              FE               | Note |
| :---------------------- | :----------------------- | ------ | :---------------------------: | :---------------------------: | :--: |
| Create item             | /item                    | POST   |    :ballot_box_with_check:    |    :ballot_box_with_check:    |      |
| Get items               | /item                    | GET    |    :ballot_box_with_check:    |    :ballot_box_with_check:    |      |
| Get item                | /item/:sku               | GET    |    :ballot_box_with_check:    |    :ballot_box_with_check:    |      |
| show item image         | /file/:filename          | GET    |    :ballot_box_with_check:    |    :ballot_box_with_check:    |      |
| Update item             | /item/:sku               | PUT    |    :ballot_box_with_check:    |    :ballot_box_with_check:    |      |
| Create item type        | /item-type               | POST   |    :ballot_box_with_check:    |    :ballot_box_with_check:    |      |
| Get all item type       | /item-type               | GET    |    :ballot_box_with_check:    |    :ballot_box_with_check:    |      |
| Update item type        | /item-type/:item-type-id | PUT    |    :ballot_box_with_check:    |    :ballot_box_with_check:    |      |
| Create storage          | /storage                 | POST   |    :ballot_box_with_check:    |    :ballot_box_with_check:    |      |
| Get storage list        | /storage                 | GET    |    :ballot_box_with_check:    |    :ballot_box_with_check:    |      |
| Get storage             | /storage/:storage-id     | GET    |    :ballot_box_with_check:    |    :ballot_box_with_check:    |      |
| Update storage          | /storage/:storage-id     | PUT    |    :ballot_box_with_check:    |    :ballot_box_with_check:    |      |
| Create supplier         | /supplier                | POST   |    :ballot_box_with_check:    |    :ballot_box_with_check:    |      |
| Get list supplier       | /supplier                | GET    |    :ballot_box_with_check:    |    :ballot_box_with_check:    |      |
| Get supplier            | /supplier/:supplier-id   | GET    |    :ballot_box_with_check:    |    :ballot_box_with_check:    |      |
| Update supplier         | /supplier/:supplier-id   | PUT    |    :ballot_box_with_check:    |    :ballot_box_with_check:    |      |
| Login                   | /login                   | POST   |    :ballot_box_with_check:    | :negative_squared_cross_mark: |      |
| Get user accounts       | /users                   | GET    |    :ballot_box_with_check:    | :negative_squared_cross_mark: |      |
| Reset user password     | /users/password          | PATCH  |    :ballot_box_with_check:    | :negative_squared_cross_mark: |      |
| Create import tickets   | /import_ticket           | POST   |    :ballot_box_with_check:    | :negative_squared_cross_mark: |      |
| Get list import tickets | /import_ticket           | GET    | :negative_squared_cross_mark: | :negative_squared_cross_mark: |      |
| Get import ticket       | /import_ticket/:id       | GET    | :negative_squared_cross_mark: | :negative_squared_cross_mark: |      |
| Generate serials        | /import_ticket/serials   | POST   |    :ballot_box_with_check:    | :negative_squared_cross_mark: |      |
| Get serial info         | /serial/:seri            | GET    |    :ballot_box_with_check:    | :negative_squared_cross_mark: |      |
