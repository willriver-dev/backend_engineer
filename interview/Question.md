### 1. How different is Functional Programming from Object Oriented Programming? Explain.

Functional Programming (FP) and Object-Oriented Programming (OOP) are two different paradigms in software development. Here are the key differences between the two:
    
- Data and State:
    OOP: In OOP, data and behavior are encapsulated within objects. Objects store state and have methods to manipulate that state.
    FP: In FP, data and behavior are separated. Data is immutable, and functions operate on this data to produce new values. FP promotes the use of pure functions that have no side effects and always produce the same result for the same inputs.



- Mutability:
    OOP: OOP allows mutable state, meaning objects can change their state during execution.
    FP: FP encourages immutability. Data is not changed after it is created, and there is no change in state. Instead, functions create new values based on the original data.


- Control Flow:
    OOP: OOP uses control structures like loops and conditionals to control the program flow. Mechanisms such as inheritance and polymorphism are also used to manage control flow.
    FP: FP uses function composition, recursion, and higher-order functions to control program flow. FP focuses on representing computations by combining functions, rather than using traditional control structures.


- Side Effects:
    OOP: OOP often involves side effects, such as modifying shared state or performing I/O operations. This can lead to complex interactions and make the code difficult to test and reason about.
    FP: FP encourages avoiding side effects and emphasizes the use of pure functions. Pure functions have no side effects and always produce predictable results, without affecting external state.

### 2. Why use functional program to write core?

Functional programming is considered suitable for writing core packages due to several key benefits:

- Reliability: Functional programming encourages the use of pure functions and immutability, which minimize side effects and state-related bugs. This results in stable, reliable core packages that are easy to test.

- Maintainability and Reusability: Functional programming emphasizes modularity and composition, allowing core packages to be built by combining small functions into larger functional blocks. This helps separate concerns and creates maintainable and reusable core packages.

- Performance and Optimization: Some functional programming languages provide automatic optimization mechanisms, such as lazy evaluation or memoization. These mechanisms improve the performance of core packages by avoiding unnecessary computations and caching previously computed results for reuse.

- Parallel and Concurrent Programming: Functional programming provides a good foundation for parallel programming and concurrent programming. Due to immutability and separation of concerns, core packages can be safely used in multithreaded or multiprocessing environments without encountering issues related to shared state.

- Clarity and Understandability: Functional programming focuses on representing logic through pure functions, eliminating side effects and hidden state. This makes core packages easier to understand and explain, allowing developers to have a clear understanding of the behavior and interaction of functions.

`However, not every situation may be suitable for functional programming. There are certain specific requirements or scenarios that may lead to choosing other programming styles such as object-oriented programming or procedural programming.`

### 3. What are the difference between Views and Tables?
Views and tables are two different concepts in database management systems. Here are the main differences between them:

1. **Definition**:

    Table: A table is a basic database object that represents a collection of structured data. It consists of rows and columns, where each row represents a record or an entity, and each column represents a specific attribute or field of that entity.
    View: A view is a virtual table that is derived from one or more tables or other views. It does not physically store data but provides a way to present the data in a customized or filtered manner. A view is defined by a query that specifies the columns and rows to be included in the view.

2. **Storage**:
    
    Table: Tables store data physically in the database. The data is persisted and can be directly modified, inserted, or deleted.
    View: Views do not store data themselves. They are a logical representation of data that is dynamically generated based on the underlying tables or views they are built upon. The data shown in a view is retrieved on-the-fly when the view is accessed.

3. **Data Manipulation**:

    Table: Tables allow direct manipulation of data using SQL operations such as INSERT, UPDATE, and DELETE. Changes made to a table affect the actual data stored in the database.
    View: Views are primarily used for querying and presenting data. While some views allow limited data manipulation, such as updating specific columns or rows, they generally have restrictions on modifications and are primarily read-only.

4. **Structure**:

    Table: Tables have a fixed structure defined by the table schema, including the columns, data types, and constraints.
    View: Views can have a structure that differs from the underlying tables. They can combine columns from multiple tables, apply filters or calculations, and present a subset or transformation of the original data.
5. **Data Abstraction**:

    Table: Tables provide a concrete representation of data, holding the actual values and records.
    View: Views provide a level of abstraction by presenting a virtual table that can hide certain columns, rows, or complex queries. They allow users to work with a simplified or customized view of the data without exposing the underlying tables' details.
6. **Dependency**:

    Table: Tables are independent database objects and can exist on their own.
    View: Views depend on the underlying tables or views. If the underlying structure changes, the view may need to be updated or recreated.

`In summary, tables store physical data, allow direct data manipulation, and have a fixed structure, while views provide a virtual representation of data, are primarily used for querying, and can have a customized structure and data presentation.`

### 4. What data tables do an inventory management system require?

An inventory management system typically requires several data tables to effectively track and manage inventory. Here are some commonly used tables in an inventory management system:

1. **Products Table**: This table stores information about the products in inventory. It typically includes fields such as product ID, name, description, category, price, quantity, supplier information, and other product-specific attributes.

2. **Inventory Table**: This table tracks the stock levels of each product in inventory. It includes fields such as product ID, current quantity, minimum stock level, maximum stock level, reorder point, and other inventory-related information. This table helps in managing stock levels and generating alerts for low stock or stock replenishment.

3. **Locations Table**: If the inventory is stored in multiple locations or warehouses, a locations table is used to store information about each location. It may include fields like location ID, name, address, capacity, and other relevant details. This table helps in tracking the physical location of inventory items.

4. **Orders Table**: This table records information about purchase orders or sales orders related to inventory. It typically includes fields like order ID, date, customer/supplier details, order status, and other relevant information. This table helps in tracking and managing orders, including order processing, fulfillment, and delivery.

5. **Transactions Table**: This table records all inventory-related transactions, such as stock inflows (e.g., purchases, returns) and outflows (e.g., sales, transfers). It includes fields like transaction ID, date, product ID, quantity, transaction type, and other transaction-specific details. This table provides a transaction history for auditing, reporting, and inventory reconciliation purposes.

6. **Suppliers and Customers** Tables: These tables store information about suppliers and customers, respectively. They include fields like ID, name, contact details, and other relevant information. These tables are used for managing relationships with suppliers and customers, tracking purchases, sales, and communication details.

7. **Categories Table**: This table categorizes products into different categories or groups. It includes fields like category ID, name, description, and other relevant attributes. This table helps in organizing and classifying products for better inventory management and reporting.

8. **Units of Measurement Table**: This table defines various units of measurement used for products (e.g., pieces, kilograms, liters). It includes fields like unit ID, name, symbol, and conversion factors if applicable. This table ensures consistent measurement standards and supports conversions between different units.

These are some of the key data tables commonly found in an inventory management system. Depending on the specific requirements of the system, additional tables or variations of these tables may be included to accommodate specific business needs and workflows.



