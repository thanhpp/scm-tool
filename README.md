# scm

## Tài liệu

- https://dbdiagram.io/d/61efb53a7cf3fc0e7c621c2d
- https://drive.google.com/drive/folders/1EV529MDBP5CvLXkzccr4G1FMFUlmbNe7?usp=sharing
- https://www.figma.com/file/iLSOm3MXdVA3Z1zx1Z2ncd/scm?node-id=0%3A1

## Tính năng

- [x] Quản lý kho
  - [x] Thêm kho
  - [x] Sửa kho
  - [x] Lấy thông tin kho
  - [x] Lấy danh sách kho
- [ ] Quản lý sản phẩm
  - [x] Tạo sản phẩm
  - [x] Lưu ảnh sản phẩm
  - [x] Tao kieu san pham
  - [ ] Sửa sản phẩm
  - [x] Lấy thông tin sản phẩm
  - [x] Lấy danh sách sản phẩm
  - [x] Lấy danh sách kiểu sản phẩm 
- [ ] Quản lý nhập xuất
  - [x] **Tạo phiếu nhập**
  - [x] Tạo số seri
  - [x] Truy xuất phiếu nhập
  - [x] Lấy thông tin phiếu nhập
  - [ ] **Tạo phiếu xuất**
  - [ ] Lấy thông tin phiếu xuất
  - [x] Tạo nhà cung cấp
  - [x] Lấy danh sách nhà cung cấp
  - [x] Sửa thông tin nhà cung cấp 
- [ ] Xác thực/Phân quyền
  - [x] Đăng nhập
  - [x] Tạo tài khoản
  - [ ] Sửa tài khoản
  - [ ] Quên mật khẩu
  - [ ] Phân quyền
- [x] Tạo số seri -> blockchain
  - [x] MintNFT blockchain 
  - [x] Upload lên IPFS
  - [x] Server để lưu trữ và mint NFT
- [x] Tra cứu thông tin theo số seri

## Overleaf

- https://www.overleaf.com/read/bxyvbfxqdztg
- https://www.overleaf.com/project/62c687809aee882527f0c10a

## Diagrams

### Create new item

```mermaid
sequenceDiagram
  actor NhanVien
  activate NhanVien
  participant Interface
  participant HTTPServer
  participant Application
  participant Factory
  participant Database
  participant LocalStorage

  NhanVien ->>+ Interface: Create new item request
  Interface ->>+ HTTPServer: Create new item request
  HTTPServer ->> HTTPServer: Unmarshal request
  HTTPServer ->>+ Application: Create new item information
  Application ->>+ Database: Get item type by ID
  alt Item type not exst
    Database -->> HTTPServer: Item type not exist
    HTTPServer ->> HTTPServer: Marshal response
    HTTPServer -->> Interface: Error response
    Interface -->> NhanVien: Error response
  else
    Database -->>- HTTPServer: Item type information
    Application ->>+ LocalStorage: Save Images
    LocalStorage -->>- Application: Image paths
    Application ->>+ Factory: Create new item
    alt Invalid new item information
      Factory -->> Application: Error message
      Application ->> LocalStorage: Delete images by paths
      LocalStorage -->> Application: Response
      Application -->> HTTPServer: Error message
      HTTPServer ->> HTTPServer: Marshal response
      HTTPServer -->> Interface: Error response
      Interface -->> NhanVien: Error response
    else
      Factory -->>- Application: New item
      Application ->>+ Database: Save new item
      Database -->>- Application: Response
      Application -->>- HTTPServer: Success message
      HTTPServer ->> HTTPServer: Marshal response
      HTTPServer -->>- Interface: Success message
      Interface -->>- NhanVien: Success message
    end
  end
  
  deactivate NhanVien
```

### Generate new serials
```mermaid
sequenceDiagram
  actor NhanVien
  participant Interface
  participant HTTPServer
  participant Application
  participant Factory
  participant Database

  activate NhanVien
  
  NhanVien ->>+ Interface: Generate new serials request
  Interface ->>+ HTTPServer: Generate new serials request
  HTTPServer ->> HTTPServer: Unmarshal request
  HTTPServer ->>+ Application: Generate new serials by import ticket ID
  Application ->>+ Database: Query import ticket by ID
  alt Import ticket not exist
    Database -->> Application: Import ticket not found error
    Application -->> HTTPServer: Error message
    HTTPServer ->> HTTPServer: Marshal response
    HTTPServer -->> Interface: Error response
    Interface -->> NhanVien: Error response
  else
    Database -->>- Application: Import ticket information
    Application ->>+ Factory: Generate serials
    Factory ->> Factory: Validate import ticket status
    alt Import ticket is invalid for generating new serials
      Factory -->> Application: Invalid import ticket error
      Application -->> HTTPServer: Error message
      HTTPServer ->> HTTPServer: Marshal response
      HTTPServer -->> Interface: Error response
      Interface -->> NhanVien: Error response
    else
      Factory -->>- Application: Serials information
      Application ->>+ Database: Batch insert serials
      Database -->>- Application: Response
      Application -->>- HTTPServer: Serials information
      HTTPServer ->> HTTPServer: Marshal response
      HTTPServer -->>- Interface: Serials information
      Interface -->>- NhanVien: Serials information
    end
  end

  deactivate NhanVien
```

### SCM Service Generate NFT
```mermaid
sequenceDiagram
  participant Application
      loop Auto mint NFT interval
          activate Application
          Application ->>+ Database: Get serials with empty token ID
          Database -->>- Application: Serials information

          loop Each serial
              Application -) RabbitMQ: Publish mint NFT message
          end
          deactivate Application
      end
```

### NFT Service Generate NFT
```mermaid
sequenceDiagram
  participant RabbitMQ
  participant Application
  participant IPFSClient
  participant OnchainClient
  participant SmartContractInstance
  participant Factory
  participant LocalStorage
  participant IPFS

    RabbitMQ -) Application: Generate NFT message
    activate Application
    Application ->>+ Database: Seri duplication check
    alt Duplicate seri
      Database -->>- Application: Duplicate seri
      Application ->> Application: skip 
    else
      Application ->> Application: Generate Metadata
      Application ->>+ LocalStorage: Save metadata file
      LocalStorage -->>- Application: Metadata file path
      Application ->>+ IPFSClient: Generate IPFS File from metadata file
      IPFSClient ->>+ IPFS: Generate IPFS file
      IPFS -->>- IPFSClient: IPFS file information
      IPFSClient -->>- Application: IPFS file CID
      Application ->> Application: Generate Token URI
      Application ->>+ OnchainClient: Get gas price
      OnchainClient ->>+ Blockchain: Get gas price
      Blockchain -->>- OnchainClient: Current gas price
      OnchainClient -->>- Application: Gasprice
      Application ->>+ SmartContractInstance: Mint NFT
      SmartContractInstance ->>+ Blockchain: Create mint NFT Transaction
      Blockchain -->>- SmartContractInstance: Transaction information
      SmartContractInstance -->>- Application: Transaction information
      Application ->>+ Factory: Create SerialNFT
      Factory -->>- Application: New SerialNFT
      Application ->>+ Database: Save SerialNFT
      Database -->>- Application: Response
    deactivate Application
    end
```


### SCM Auto update NFT

```mermaid
sequenceDiagram
  participant Application
    par Rest update
        loop Auto update loop
        activate Application
        Application ->>+ Database: Get serials with empty token ID
        Database -->>- Application: Serials information

        loop Each serial
            Application ->>+ NFTService: Get NFT information by seri
            NFTService -->>- Application: Response
            opt Response with token ID
            Application ->>+ Database: Save token ID by seri
            Database -->>- Application: Response
            end
        end
        deactivate Application
        end
    and RabbitMQ update
        activate Application
        RabbitMQ -) Application: NFT Information
        Application ->>+ Database: Update NFT Information by seri
        Database -->>- Application: Response
        deactivate Application
  end
```

### NFT Update transaction information
```mermaid
sequenceDiagram
  participant Application
  participant Database
  participant OnchainClient
  loop Auto update transaction interval
    activate Application
    Application ->>+ Database: Get NFT information without token ID
    Database -->>- Application: NFT information
    loop Each NFT
      Application ->>+ OnchainClient: Get transaction information by hash
      OnchainClient ->>+ Blockchain: Get transaction receipt
      Blockchain -->>- OnchainClient: Transaction receipt
      OnchainClient ->> OnchainClient: Get token ID from receipt logs
      OnchainClient -->>- Application: Transaction information 
      Application ->>+ Database: Update NFT information
      Database -->>- Application: Response
      Application -) RabbitMQ: Publish update NFT information message
    end
    deactivate Application
  end
```
### ERDiagram

```mermaid
erDiagram 
  User {
    int     id
    string  username
    string  name
    string  password
  }

  Item {
    string  sku
    string  name
    string  desc
    float   sell_price
  }

  ItemImage {
    string path
  }

  ItemType {
    int     item_type_id
    string  name
    string  desc
  }

  ImportTicket {
    int         id
    int         status
    timestamp   send_time
    timestamp   receive_time
    float       fee
  }

  ImportTicketProductImage {
    string  product_image_path
  }

  ImportTicketBillImage {
    int     import_ticket_id
    string  bill_image_path
  }

  ImportTicketDetail {
    int     import_ticket_id
    string  item_sku
    int     buy_quantity
    int     receive_quantity
    float   buy_price
  }

  NFT {
    string  transaction_hash
    string  ipfs_hash
    string  metadata
    int     token_id
  }

  Serial {
    string    seri
    int       status
    int       token_id
  }

  Storage {
    int     id
    string  name
    string  location
    string  desc
  }

  Supplier {
    int     id
    string  name
    string  phone
    string  email
  }

  Item ||--o{ ItemImage : has
  Item ||--|| ItemType: is
  ImportTicket ||--o{ ImportTicketProductImage: has
  ImportTicket ||--o{ ImportTicketBillImage: has
  ImportTicket ||--|{ ImportTicketDetail: has
  ImportTicketDetail ||--|| Item: is
  ImportTicket ||--|| Supplier: from
  ImportTicket ||--|| Storage: to
  Serial }o--|| ImportTicket: from
  Serial }o--|| Item: is
  Serial ||--|| NFT: has
```

## Ref

- https://github.com/INFURA/ipfs-upload-client/blob/master/main.go
- https://ethereum.org/en/developers/tutorials/how-to-view-nft-in-metamask/
- https://goethereumbook.org/en/smart-contract-write/
- https://dev.to/rounakbanik/writing-an-nft-collectible-smart-contract-2nh8
