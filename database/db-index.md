## Overview

- Index cần vùng nhớ riêng, tách biệt với dữ liệu gốc => db tạo ra một bảng copy cho index (hay meta data).
- Index được implement ở tầng storage engine.

## Pros

- Tăng tốc độ truy vấn (Select).
- Giảm tải server và thao tác I/O xuống disk cho truy vấn dữ liệu lớn.
- Tìm kiếm dữ liệu sẽ thực hiện trên index trước, sau đó mới tìm kiếm trên dữ liệu gốc.

## Cons

- Tốn không gian lưu trữ (disk, memory)
- Insert, update, delete thường thì chậm hơn.
- Đánh index vô tội vạ sẽ gây chậm db do phải maintain index.
- Index có thể gây nhiễu trong plan execute query và nhầm lẫn khi lựa chọn index.

## Các kiểu chia index

- Data structure (B-tree, Hash, R-tree, Bitmap)
- Phisically (Clustered index, Non-clustered index)
- Index field (single, compound, unique, partial)

# Cấu trúc dữ liệu của index

## B-tree Index
- Là cấu trúc dữ liệu dạng cây cân bằng Balance Tree.
- Mỗi node chứa n giá trị và chứa con trỏ tới các node dưới.
- Dữ liệu các node được sắp xếp từ lớn tới bé
- Mỗi node đều chứa con trỏ -> trỏ tới dữ liệu gốc
- Giữa các node lá là liên kết dạng double link list
- Mỗi node lá được lưu trữ trong 1 block(page). Đơn vị lưu trữ nhỏ nhất của DB.
- Block sẽ phụ thuộc vào storage engine. Trong mysql(InnoDB) block size của index mặc định là 16kb.

Example: đánh index user_id và order_id từ trái sang phải.

[x] where user_id > 20 and order_id = 23.

Like operator (Left mode prefix)
[v] Where name like 'WI%ND'
[x] Where name like ''%WIND%''

