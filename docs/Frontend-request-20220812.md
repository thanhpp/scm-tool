# Yêu cầu chỉnh sửa giao diện

## Cao

### 1. Tạo được nhiều sản phẩm trong cùng 1 import ticket

### 2. Tất cả các API đều hiển thị kết quả trả về

1. Có thể thông qua alert
1. Ít nhất hiển thị được thông tin trong trường error.message
1. **CẦN CÓ**: Trả về thông tin serial khi tạo số seri 

## Trung bình

### 1. Giao diện /login

1. Bỏ nút signup
1. Dành cho khách hàng => Customer section
1. Quét mã sản phẩm => Get product information 
1. Log in => Login
1. Username => Username

### 2. Giao diện /serial

1. Quét mã sản phẩm => Product serial
1. Hiện đầy đủ thông tin các level, đặc biệt là nft_info.*
```json
{
  "error": {
    "code": 200,
    "message": "OK"
  },
  "data": {
    "seri": "9998114561101101075111097115561124850101559848103",
    "import_ticket_id": 25,
    "storage": {
      "name": "test after edit",
      "location": "testLocation"
    },
    "supplier": {
        "name": "name",
    },
    "item": {
      "name": "name1",
      "desc": "desc1"
    },
    "nft_info": {
      "seri": "9998114561101101075111097115561124850101559848103",
      "tx_hash": "0xdc28ea790897e081c262b6479e09ecbc9c38ea4027f323658f0877b5a6a770e4",
      "ipfs_cid": "QmVk8UQ7uryENsAG1dirCueBy2XdZw1DxXD9BtuBjf2BrT",
      "token_id": 368
    }
  }
}
```

### Lấy thông tin số seri theo import ticket id (GET /import-ticket/:id/serials)

## Thấp

### 1. Đổi tên section Homepage => User

### 2. Trang quản lý user tên là Warehouses => User

### 3. Hiển thị danh sách chọn tại các trường nhập ID

1. Tự mapping name => ID khi gửi request

### 4. Thay đổi size 1 page = 20

## Q & A:

1. Đổi order của get list có ảnh hưởng FE không?