### 1. How different is Functional Programming from Object Oriented Programming? Explain.

Lập trình hàm (Functional Programming - FP) và Lập trình hướng đối tượng (Object-Oriented Programming - OOP) là hai hướng tiếp cận khác nhau trong phát triển phần mềm. Mặc dù cả hai đều nhằm giải quyết cùng một vấn đề là viết mã hiệu quả và dễ bảo trì, nhưng chúng có các khác biệt cơ bản như sau:

- Dữ liệu và Trạng thái: Trong OOP, dữ liệu và hành vi được đóng gói cùng nhau trong các đối tượng. Các đối tượng chứa trạng thái nội tại, và các phương thức của chúng thay đổi trạng thái này. Trong FP, dữ liệu và hành vi được tách biệt. Dữ liệu là không thay đổi (immutable), và các hàm hoạt động trên dữ liệu này để tạo ra giá trị mới. FP nhấn mạnh việc sử dụng các hàm thuần túy (pure functions), không có hiệu ứng phụ và luôn tạo ra cùng một kết quả với cùng một đầu vào.

- Sự biến đổi: OOP cho phép sử dụng trạng thái có thể thay đổi (mutable state), có nghĩa là trạng thái nội tại của một đối tượng có thể được sửa đổi theo thời gian. Sự biến đổi này có thể dẫn đến các tương tác phức tạp và các lỗi khó theo dõi. Trong khi đó, FP khuyến khích việc sử dụng tính không thay đổi (immutability), trong đó dữ liệu là không thay đổi và không thể thay đổi sau khi khởi tạo. Điều này giảm hiệu ứng phụ và làm mã nguồn dễ dự đoán và dễ lập luận hơn.

- Luồng điều khiển: Trong OOP, luồng điều khiển thường được quản lý thông qua cấu trúc điều khiển như vòng lặp và điều kiện, cũng như các cơ chế như đa hình và kế thừa. Trong FP, luồng điều khiển được thực hiện thông qua sự kết hợp hàm, đệ quy và các hàm bậc cao. FP tập trung vào việc biểu diễn tính toán dưới dạng sự kết hợp của các hàm thuần túy, cho phép mã nguồn ngắn gọn và khái quát hơn.

- Hiệu ứng phụ: OOP thường dựa vào hiệu ứng phụ (side effects), chẳng hạn như sửa đổi trạng thái có thể thay đổi hoặc tương tác với thế giới bên ngoài (ví dụ: thao tác I/O). FP khuyến khích tránh hiệu ứng phụ, tập trung vào các hàm thuần túy tạo ra kết quả dự đoán được và không ảnh hưởng đến trạng thái chương trình bên ngoài phạm vi của chúng. Bằng cách giảm thiểu hiệu ứng phụ, FP thúc đẩy việc viết mã dễ kiểm thử, gỡ lỗi và lập luận.

- Biến đổi dữ liệu: OOP thường sử dụng phương thức và thông điệp để biến đổi dữ liệu trong các đối tượng, tuân theo nguyên tắc đóng gói (encapsulation). FP xem xét dữ liệu là không thay đổi và sử dụng các hàm bậc cao và biến đổi để xử lý và biến đổi dữ liệu. Lập trình hàm thường sử dụng các khái niệm như map, filter và reduce để hoạt động trên tập hợp dữ liệu.

- Kế thừa và đa hình: OOP phụ thuộc mạnh vào các khái niệm kế thừa và đa hình để tái sử dụng mã và mở rộng chức năng. Kế thừa cho phép đối tượng kế thừa thuộc tính và hành vi từ các lớp cha. Đa hình cho phép các đối tượng khác nhau có thể được xử lý một cách tương đồng. FP tập trung hơn vào việc kết hợp và ưa thích sự kết hợp của các hàm nhỏ hơn so với kế thừa và đa hình.

Cả hai hướng tiếp cận đều có ưu điểm và nhược điểm riêng, và sự lựa chọn giữa chúng phụ thuộc vào yêu cầu cụ thể và ngữ cảnh của dự án. Một số ngôn ngữ, như Scala và Swift, cho phép sử dụng cả lập trình hàm và hướng đối tượng, giúp nhà phát triển tận dụng những ưu điểm của cả hai hướng tiếp cận.

### 2. Tại sao lập trình hàm thường được sử dụng để viết các core packages?

Lập trình hàm được coi là phù hợp để viết core packages (gói lõi) vì có một số lợi ích chính sau:

- Độ tin cậy: Lập trình hàm khuyến khích sử dụng các hàm thuần (pure functions) và tính bất biến (immutability), giúp giảm thiểu các tác động phụ và các bug liên quan đến trạng thái (state-related bugs). Điều này tạo ra các core packages ổn định, đáng tin cậy và dễ kiểm tra.

- Dễ bảo trì và tái sử dụng: Lập trình hàm khuyến khích sử dụng tính rời rạc (modularity) và tổ hợp (composition), cho phép xây dựng core packages bằng cách kết hợp các hàm nhỏ thành các khối chức năng lớn hơn. Điều này giúp tách riêng logic và logic xử lý, tạo ra các core packages dễ bảo trì và tái sử dụng.

- Hiệu suất và tối ưu hóa: Một số ngôn ngữ lập trình hàm cung cấp cơ chế tối ưu hóa tự động, như lập trình chính quy (lazy evaluation) hoặc tối ưu hóa thực thể (memoization). Những cơ chế này giúp cải thiện hiệu suất của core packages bằng cách tránh tính toán không cần thiết và lưu trữ kết quả tính toán trước đó để sử dụng lại.

- Paralell và concurrent programming: Lập trình hàm tạo điều kiện tốt cho lập trình song song (parallel programming) và đồng thời (concurrent programming). Do tính bất biến và tính rời rạc, các core packages có thể được sử dụng một cách an toàn trong môi trường đa luồng hoặc đa tiến trình mà không gặp các vấn đề liên quan đến trạng thái chia sẻ.

- Độ rõ ràng và dễ hiểu: Lập trình hàm tập trung vào việc biểu diễn logic bằng các hàm thuần, loại bỏ các tác động phụ và trạng thái ẩn. Điều này làm cho core packages dễ hiểu và dễ giải thích, giúp người phát triển hiểu rõ hơn về hành vi và tương tác của các chức năng.

Tuy nhiên, không phải mọi tình huống đều phù hợp với lập trình hàm. Có một số vấn đề hoặc yêu cầu đặc biệt có thể dẫn đến việc chọn các phong cách lập trình khác, như lập trình hướng đối tượng hoặc lập trình cấu trúc.

### 3. Sự khác biệt giữa table và view?

Dưới đây là những khác biệt chính giữa Views và Tables trong hệ quản trị cơ sở dữ liệu:

1. **Định nghĩa**:

    Table: Table là một đối tượng cơ bản trong cơ sở dữ liệu, đại diện cho một tập hợp dữ liệu có cấu trúc. Nó bao gồm các hàng và cột, trong đó mỗi hàng đại diện cho một bản ghi hoặc một đối tượng, và mỗi cột đại diện cho một thuộc tính hoặc trường cụ thể của đối tượng đó.
    View: View là một bảng ảo được tạo ra từ một hoặc nhiều bảng hoặc view khác. Nó không lưu trữ dữ liệu một cách vật lý, mà cung cấp một cách để hiển thị dữ liệu theo một cách tùy chỉnh hoặc được lọc. View được xác định bởi một truy vấn chỉ định các cột và hàng được bao gồm trong view.

2. **Lưu trữ**:

    Table: Bảng lưu trữ dữ liệu một cách vật lý trong cơ sở dữ liệu. Dữ liệu được lưu trữ và có thể được thay đổi, chèn hoặc xóa trực tiếp.
    View: View không lưu trữ dữ liệu. Nó chỉ là một biểu diễn logic của dữ liệu được tạo ra động dựa trên các bảng hoặc view gốc. Dữ liệu được hiển thị trong view được truy xuất khi view được truy cập.

3. **Thao tác dữ liệu**:

    Table: Bảng cho phép thao tác dữ liệu trực tiếp bằng các phép toán SQL như INSERT, UPDATE và DELETE. Các thay đổi được thực hiện trên bảng ảnh hưởng đến dữ liệu thực tế được lưu trữ trong cơ sở dữ liệu.
    View: View chủ yếu được sử dụng để truy vấn và hiển thị dữ liệu. Mặc dù một số view cho phép thay đổi dữ liệu hạn chế như cập nhật các cột hoặc hàng cụ thể, nhưng chúng chủ yếu là chỉ đọc.

4. **Cấu trúc**:

    Table: Bảng có cấu trúc cố định được xác định bởi schema của bảng, bao gồm các cột, kiểu dữ liệu và ràng buộc.
    View: View có thể có cấu trúc khác so với các bảng gốc. Nó có thể kết hợp các cột từ nhiều bảng, áp dụng bộ lọc hoặc tính toán và hiển thị một phần hoặc biến đổi dữ liệu gốc.

5. **Trừu tượng hóa dữ liệu**:

    Table: Bảng cung cấp một biểu diễn cụ thể của dữ liệu, chứa các giá trị và bản ghi thực tế.
    View: View cung cấp một mức độ trừu tượng bằng cách hiển thị một bảng ảo, có thể ẩn đi các cột, hàng hoặc truy vấn phức tạp. Nó cho phép người dùng làm việc với một cái nhìn đơn giản hoặc được tùy chỉnh của dữ liệu mà không tiết lộ chi tiết về các bảng gốc.

6. **Phụ thuộc**:

    Table: Bảng là các đối tượng cơ sở dữ liệu độc lập và có thể tồn tại độc lập.
    View: View phụ thuộc vào các bảng hoặc view gốc. Nếu cấu trúc gốc thay đổi, view có thể cần được cập nhật hoặc tạo lại.

`Tóm lại, bảng lưu trữ dữ liệu vật lý, cho phép thao tác dữ liệu trực tiếp và có cấu trúc cố định, trong khi view cung cấp một biểu diễn ảo của dữ liệu, chủ yếu được sử dụng để truy vấn và có thể có cấu trúc và hiển thị dữ liệu tùy chỉnh.`

### 4. Tạo database cho một hệ thống quản lý kho hàng.

Một hệ thống quản lý kho thông thường yêu cầu một số bảng dữ liệu để theo dõi và quản lý hiệu quả hàng tồn kho. Dưới đây là một số bảng thông thường được sử dụng trong hệ thống quản lý kho:

1. **Bảng Sản phẩm**: Bảng này lưu trữ thông tin về các sản phẩm trong kho. Thông thường, nó bao gồm các trường như mã sản phẩm, tên, mô tả, danh mục, giá, số lượng, thông tin nhà cung cấp và các thuộc tính cụ thể khác của sản phẩm.

2. **Bảng Tồn kho**: Bảng này theo dõi mức tồn kho của mỗi sản phẩm trong kho. Nó bao gồm các trường như mã sản phẩm, số lượng hiện tại, mức tồn kho tối thiểu, mức tồn kho tối đa, điểm đặt hàng lại và các thông tin khác liên quan đến tồn kho. Bảng này giúp quản lý mức tồn kho và tạo cảnh báo cho mức tồn kho thấp hoặc việc bổ sung hàng hóa.

3. **Bảng Địa điểm**: Nếu hàng tồn kho được lưu trữ tại nhiều địa điểm hoặc kho, một bảng địa điểm được sử dụng để lưu trữ thông tin về mỗi địa điểm. Bảng này có thể bao gồm các trường như mã địa điểm, tên, địa chỉ, công suất và các chi tiết khác. Bảng này giúp theo dõi vị trí vật lý của các mặt hàng trong kho.

4. **Bảng Đơn đặt hàng**: Bảng này ghi lại thông tin về đơn đặt hàng hoặc đơn hàng bán hàng liên quan đến hàng tồn kho. Thông thường, nó bao gồm các trường như mã đơn hàng, ngày, chi tiết khách hàng/nhà cung cấp, trạng thái đơn hàng và các thông tin khác liên quan. Bảng này giúp theo dõi và quản lý đơn hàng, bao gồm xử lý đơn hàng, đáp ứng yêu cầu và giao hàng.

5. **Bảng Giao dịch**: Bảng này ghi lại tất cả các giao dịch liên quan đến hàng tồn kho, chẳng hạn như hàng vào (như mua hàng, trả hàng) và hàng ra (như bán hàng, chuyển hàng). Nó bao gồm các trường như mã giao dịch, ngày, mã sản phẩm, số lượng, loại giao dịch và các chi tiết giao dịch khác. Bảng này cung cấp lịch sử giao dịch để kiểm tra, báo cáo và cân đối kho.

6. **Bảng Nhà cung cấp và Khách hàng**: Đây là các bảng lưu trữ thông tin về nhà cung cấp và khách hàng tương ứng. Chúng bao gồm các trường như ID, tên, thông tin liên hệ và các thông tin khác liên quan. Các bảng này được sử dụng để quản lý mối quan hệ với nhà cung cấp và khách hàng, theo dõi mua hàng, bán hàng và chi tiết liên lạc.

7. **Bảng Danh mục**: Bảng này phân loại các sản phẩm vào các danh mục hoặc nhóm khác nhau. Nó bao gồm các trường như mã danh mục, tên, mô tả và các thuộc tính khác liên quan. Bảng này giúp tổ chức và phân loại sản phẩm để quản lý kho và báo cáo tốt hơn.

8. **Bảng Đơn vị đo lường**: Bảng này xác định các đơn vị đo lường được sử dụng cho các sản phẩm (ví dụ: cái, kilogram, lít). Nó bao gồm các trường như mã đơn vị, tên, ký hiệu và các hệ số chuyển đổi nếu cần thiết. Bảng này đảm bảo tiêu chuẩn đo lường nhất quán và hỗ trợ việc chuyển đổi giữa các đơn vị khác nhau.

Đây là một số bảng dữ liệu chính thường được sử dụng trong hệ thống quản lý kho. Tùy thuộc vào yêu cầu cụ thể của hệ thống, có thể bổ sung thêm bảng hoặc biến thể của các bảng này để phù hợp với nhu cầu và quy trình kinh doanh cụ thể.