# Chương 5: Thiết kế phần mềm

## 5.1 Nhóm các loại dữ liệu khác nhau

- Điều quan trọng phải nhớ rằng trong golang, các khái niệm về việc kế thừ hoặc tạo lớp con thực sự không tồn tại. Chúng ta nên tránh các design pattern này.

- Khi nào nên tạo một type mới?
    - Khai báo kiểu mới khi nó đại diện cho một thứ gì đó mới hoặc độc đáo.
    - Đừng tạo aliases chỉ vì nó khiến dễ đọc hơn.
    - Embed types không phải vì quan tâm đến trạng thái (state), mà là quan tâm đến hành vi.
    - Nếu không suy nghĩ về hành vi (behavior), bạn sẽ tự khoá mình trong bản thiết kế mà bạn không thể phát triển trong tương lai, mà không cần thay đổi code liên tục.

## 5.2 Đừng thiết kế bằng interfaces.

- Khi bạn là 1 developer, bạn nên có 2 modes (chế độ). 1 là lập trình viên (programer), 2 là một kỹ sư (engineer)
- Khi bạn lập trình, bạn hãy tập trung vào việc làm sao khiến code chạy được. Trước tiên là giải quyết vấn đề và loại bỏ những rào cản. Hãy chứng minh những ý tưởng của bạn nó có thể hoạt động. Đó là tất cả những gì bạn cần quan tâm khi ở mode lập trình viên. Việc này phải được hoàn thiện một cách cụ thể, nhưng nó sẽ không bao giờ sẵn sàng cho go live.
- Sau khi đã có prototype of code (mẫu code) có thể giải quyết được vấn đề, bạn cần chuyển sang mode engineer. Lúc này bạn cần tập trung vào code ở micro-level về ngữ nghĩa dữ liệu (value, semantic), và tính dễ đọc (readability) của code. Sau đó ở cấp độ macro-level để đáp ứng các mô hình tư duy (mental-model), khả năng bảo trì. Bạn cũng cần tập trung vào errors và failure states (trạng thái lỗi)
- Đây là 1 chu trình tái cấu trúc (Refactoring). Refactoring để dễ đọc (readability), hiệu quả (efficiency), trừu tượng hay khái quát hoá (abtraction), và có khả năng kiểm thử (testability). Cách hiệu quả nhất để áp dụng là bạn bắt đầu với những đoạn code cụ thể, hãy KHÁM PHÁ những interfaces bạn thực sự cần. Không cần apply abstraction trừ khi nó thực sự cần thiết.
- Mọi vấn đề bạn cần giải quyết thường liên quan đến chuyện biến đổi dữ liệu (data transformations). Nếu bạn không hiểu về dữ liệu, bạn sẽ không hiểu được vấn đề là gì. Nếu bạn không hiểu vấn đề là gì, bạn sẽ không thể viết được code. Bắt đầu bằng một giải pháp cụ thể dựa trên cấu trúc dữ liệu cụ thể là rất quan trọng.
- "Dữ liệu thống trị tất cả, nếu bạn chọn đúng cấu trúc dữ liệu và sắp xếp mọi thứ hợp lý thì mọi thuật toán hầu như sẽ luôn là hiển nhiên" - Rob Pike
- Khi nào cần sự trừu tượng (abstraction)? Khi chúng ta thấy một nơi trong code mà dữ liệu có thể sẽ thay đổi và chúng ta cần giảm thiểu sự thay đổi liên tục của code. Lúc đó ta có thể sử dụng abstraction để giúp code dễ test. Nhưng tốt hơn là bạn nên tránh điều này nếu có thể. Testable functions tốt nhất là khi input là raw data thì output cũng là raw data. Dữ liệu đến từ đâu hoặc đi đến đâu không quan trọng.
- "Đừng thiết kế bằng interfaces, hãy khám phá chúng". - Rob Pike

## 5.3 Composition

- Cách tốt nhất để tận dụng tính năng embedding là thông qua mẫu thiết kế hợp thành (compositional design pattern). Ý tưởng là tổ hợp các largre types từ small types và tập trung vào việc nhóm hành vi (composition of behavior).

