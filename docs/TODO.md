# Todo 

## Quyền đồ án

- [x] Lời cam kết
- [x] Lời cảm ơn
- [x] Tóm tắt nội dung
- [ ] Tóm tắt nội dung T.Anh
- [x] Viết tắt
- [x] Thuật ngữ
- [x] Giới thiệu
  - [x] Đặt vấn đề  
  - [x] Mục tiêu và phạm vi
  - [x] Định hướng giải pháp
  - [x] Bố cục đồ án
- [x] Khảo sát
  - [x] Khảo sát hiện trạng
  - [x] Tổng quan chức năng
    - [x] Usecase tổng quát
    - [x] Phân rã usecase
      - [x] Usecase quản lý sản phẩm
      - [x] Usecase quản lý phiếu nhập
      - [x] Usecase quản lý tài khoản
      - [x] Usecase tra cứu theo số seri
    - [x] Quy trình nghiệp vụ
      - [x] Nghiệp vụ tạo phiếu nhập
      - [x] Nghiệp vụ tra cứu số seri
    - [x] Đặc tả chức năng (4-7 usecases)
      - [x] Tạo sản phẩm
      - [x] Quản lý tài khoản
      - [x] Sinh số seri
      - [x] Tra cứu thông tin sản phẩm
      - [ ] [MAYBE] Quản lý Nhà cung cấp/Kho
    - [x] Yêu cầu phi chức năng   
- [x] Công nghệ
  - [x] Blockchain
  - [x] Smart contracts
  - [x] NFT
  - [x] IPFS
  - [x] Docker 
  - [x] Golang
- [x] Kết quả thực nghiệm
  - [x] Thiết kế kiến trúc (microservice)
    - [x] Biểu đồ tổng quan
    - [x] Giải thích kiến trúc 
  - [x] Thiết kế tổng quan (UML package diagram)
    - [x] scm service
    - [x] nft service
    - [x] giải thích packages
  - [x] Thiết kế chi tiết gói   
    - [x] scm service
    - [x] nft service
  - [x] Thiết kế chi tiết
    - [x] Thiết kế giao diện
    - [x] Thiết kế lớp (Class diagram, Sequence Diagram)
      - [x] Class Diagram
      - [x] Biểu đồ trình tự tạo sản phẩm
      - [x] Biểu đồ trình tự sinh số seri
      - [x] Biểu đồ trình tự tự động tạo NFT và cập nhật thông tin
      - [x] Biểu đồ trình tự tạo NFT tại serivce quản lý NFT
    - [x] Thiết kế cơ sở dữ liệu
      - [x] ER Diagram
      - [x] Thiet ke DB (postgresql)
        - [x] import_ticket
        - [x] import_ticket_bill_image
        - [x] import_ticket_product_image
        - [x] import_ticket_details
        - [ ] import_ticket_id_seq
        - [x] item
        - [x] item_image
        - [x] item_type
        - [ ] item_type_id_seq
        - [x] serial
        - [x] storage
        - [ ] storage_id_seq
        - [ ] supplier
        - [ ] supplier_id_seq
        - [x] user
        - [ ] user_id_seq
        - [x] serial_nft
  - [ ] Xây dựng ứng dụng
    - [x] Thư viện và công cụ sử dụng
      - [x] OpenZeppelin
      - [x] Gin
      - [x] Postgres
      - [x] go-ethereum
    - [x] Kết quả đạt được
      - [x] So lieu ve phan mem
    - [ ] **Minh hoa chuc nang chinh**
    - [x] Kiem thu
    - [x] Trien khai
  - [x] Giải pháp đóng góp
    - [x] Vấn đề truy xuất nguồn gốc
    - [ ] **don't know**
  - [x] Kết luận  
    - [x] Kết quả sản phẩm
    - [x] Hướng phát triển  
  - [x] Tài liệu tham khảo
  - [ ] Phụ lục
    - [ ] Đặc tả usecases  
      - [ ] Lấy danh sách sản phẩm
      - [ ] Lấy thông tin sản phẩm
      - [ ] Lấy danh sách nhà cung cấp
      - [ ] Lấy thông tin nhà cung cấp
      - [ ] Cập nhật thông tin nhà cung cấp
      - [ ] Lấy danh sách kho
      - [ ] Lấy thông tin kho
      - [ ] Cập nhật thông tin kho
      - [ ] Tạo kiểu sản phẩm
      - [ ] Lấy danh sách phiếu nhập
      - [ ] Lấy thông tin phiếu nhập

## Yêu cầu phụ

- [ ] Đổi tên bảng lên đầu

## Phần mềm minh hoạ

### Phải làm

- [x] Service mint NFT
  - [x] Nhận data để mint NFT
  - [x] Upload lên IPFS
  - [x] Mint NFT
  - [x] Database
  - [x] Server APIs
- [x] Integrate services
  - [x] Gửi data để mint NFT
  - [x] Query seri đính kèm NFT
- [ ] Tài khoản
  - [ ] Reset mật khẩu
  - [ ] Phân quyền
- [ ] Lay san pham them ma Storage

### Sẽ làm nếu còn thời gian

- [ ] Tạo phiếu xuất
- [ ] Tính năng chuyển NFT cho người khác
- [ ] Cập nhật trạng thái phiếu nhập
