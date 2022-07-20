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

### Auto mint and update NFT

```mermaid
sequenceDiagram
  participant Application
  loop Auto update loop
    activate Application
    Application ->>+ Database: Get serials with empty token ID
    Database -->>- Application: Serials information
    
    loop Each serial
      Application ->>+ NFTService: Mint NFT request
      NFTService -->>- Application: Mint NFT response
    end

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

```

## Ref

- https://github.com/INFURA/ipfs-upload-client/blob/master/main.go
- https://ethereum.org/en/developers/tutorials/how-to-view-nft-in-metamask/
- https://goethereumbook.org/en/smart-contract-write/
- https://dev.to/rounakbanik/writing-an-nft-collectible-smart-contract-2nh8
