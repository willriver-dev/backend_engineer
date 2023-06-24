## Show all index of table
```sql
SELECT cls.relname, am.amname, idxes.indexdef
FROM pg_index idx
JOIN pg_class cls ON cls.oid=idx.indexrelid
JOIN pg_class tab ON tab.oid=idx.indrelid
JOIN pg_am am ON am.oid=cls.relam
JOIN pg_indexes idxes ON cls.relname = idxes.indexname
WHERE lower(tab.relname) = '[table_name]';
```

## Create table with partition
### Partition table with range
```sql
CREATE TABLE ENGINEER
(
    id bigserial NOT NULL,
    first_name character varying(255) NOT NULL,
    last_name character varying(255) NOT NULL,
) PARTITION BY RANGE(started_date);

CREATE TABLE ENGINEER_Q1_2020 PARTITION OF ENGINEER FOR VALUES
    FROM ('2020-01-01') TO ('2020-04-01');

CREATE TABLE ENGINEER_Q2_2020 PARTITION OF ENGINEER FOR VALUES
    FROM ('2020-04-01') TO ('2020-07-01');

CREATE TABLE ENGINEER_Q3_2020 PARTITION OF ENGINEER FOR VALUES
    FROM ('2020-07-01') TO ('2020-10-01');

CREATE TABLE ENGINEER_Q4_2020 PARTITION OF ENGINEER FOR VALUES
    FROM ('2020-10-01') TO ('2020-12-31');
```

### Partition table with List
```sql
CREATE TABLE ENGINEER
(
    id bigserial NOT NULL,
    first_name character varying(255) NOT NULL,
    last_name character varying(255) NOT NULL,
    gender smallint NOT NULL,
    country_id bigint NOT NULL,
    title character varying(255) NOT NULL,
    started_date date,
    created timestamp without time zone NOT NULL
) PARTITION BY LIST(title);

CREATE TABLE ENGINEER_ENGINEER PARTITION OF ENGINEER FOR VALUES
    IN ('Backend Engineer', 'Frontend Engineer', 'Fullstack Engineer');

CREATE TABLE ENGINEER_BA_QA PARTITION OF ENGINEER FOR VALUES
    IN ('BA', 'QA');
    
CREATE TABLE ENGINEER_DEFAULT PARTITION OF ENGINEER DEFAULT;
```

### Partition table with Hash
```sql
CREATE TABLE ENGINEER
(
    id bigserial NOT NULL,
    first_name character varying(255) NOT NULL,
    last_name character varying(255) NOT NULL,
    gender smallint NOT NULL,
    country_id bigint NOT NULL,
    title character varying(255) NOT NULL,
    started_date date,
    created timestamp without time zone NOT NULL
) PARTITION BY HASH(country_id);

CREATE TABLE ENGINEER_P1 PARTITION OF ENGINEER
    FOR VALUES WITH (MODULUS 3, REMAINDER 0);

CREATE TABLE ENGINEER_P2 PARTITION OF ENGINEER
    FOR VALUES WITH (MODULUS 3, REMAINDER 1);
    
CREATE TABLE ENGINEER_P3 PARTITION OF ENGINEER
```

`Pros`:
```
Biến một logical table thành nhiều physical table, giảm không gian tìm kiếm, tăng performance cho query trong trường hợp tận dụng được điều kiện WHERE với partition key.
Xóa bỏ các data cũ một cách dễ dàng, nhanh chóng so với cách truyền thống. Ví dụ cần xóa tất cả các record có giới tính nam, nếu chia partition theo giới tính từ đầu. Chỉ cần TRUNCATE/DROP partition đó là ok. Không cần seq scan để DELETE record với WHERE condition.
```

`Cons`
```
Unique constraint, PK constraint không thể thực hiện trên table chính mà phải thực hiện ở partition table.
Các partition table không thể có column nào khác mà không khai báo ở parent table. Nói cách khác, nó kế thừa toàn bộ column và data type của parent table.
Một vài vấn đề liên quan đến trigger. Ví dụ BEFORE ROW trigger ON INSERT. Nó thực hiện trigger function/procedure... trước khi row được insert vào table.
```