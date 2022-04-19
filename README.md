- get : query
- set: query
- không dùng ommitempty, struct tags chỉ có: json (để response), bson (để làm việc với db)
- khai báo stric để chứa config enviroment (.env)
- os.Getenv(...) // khong dung

- method get luôn trả về kết quả cho client, và status code = 200, Get vẫn có 404
- các method còn lại mới ó status >= 300:, 400, 401, 500, 504,..
- hạn chế repeat code

- gom các api vào 1 router, trong package router

- check lỗi trước khi type convert
  // a := b, (string)
  // a, ok := b. (string)
  // if !ok {} // handle eror

- phân trang đi từ 0
- dùng struct để bind query (struct tag `query:"page")

- validation dùng thử viện

Ứng dụng có: - Người dùng: tên, email, tuổi, thành phố,... - Sản phẩm: mã sản phẩm, tên, giá, hình ảnh, số lượng - Đơn hàng: mã đơn, sản phẩm, số tiền, ngày mua,...

Yêu cầu: + có trang quản trị người dùng + Người dùng được đăng ký, đăng nhập + Xem list sản phẩm + Mua sản phẩm + Checkout + Xem list đơn hàng + Xem chi tiết đơn hàng + ...

Mục tiêu (excel): + Phân tích + Schema + Model + API (body, response data)

1. Xem xét hệ thống, tổng hợp những chức năng chính
2. Thiết kế kiến trúc hệ thống (Module, DB, API)
3. Kiến trúc chung các công nghệ nên sử dụng
4. Plan để hoàn thành, estimate thời gian bao lâu

---

1. Tự chọn công nghệ front-end, back-end, DB,...
2. Deploy hệ thống chạy thực tiết
3. Quy trình code, deploy
