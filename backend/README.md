## CREATE TABLE

First step is to create table called "item" which is going to be used for storing data. Simply copy the command.

```sql
CREATE TABLE item(
    id SERIAL PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    description TEXT,
    price INT NOT NULL,
    quantity INT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);
```

## Create store procedures

### Procedure for inserting fields inside the table "item"

```sql
CREATE OR REPLACE PROCEDURE insert_item(
in_item_title VARCHAR(100),
in_item_description TEXT,
in_item_price INT,
in_item_quantity INT
)
LANGUAGE plpgsql
AS $$
BEGIN
-- Insert a new record into the item table
INSERT INTO item (title, description, price, quantity, created_at, updated_at)
VALUES (in_item_title, in_item_description, in_item_price, in_item_quantity, NOW(), NOW());
END;
$$
;
```

### Procedure for getting all items

```sql
CREATE OR REPLACE FUNCTION get_all_items()
    RETURNS TABLE (
                      id INT,
                      title VARCHAR,
                      description TEXT,
                      price INT,
                      quantity INT,
                      created_at TIMESTAMP,
                      updated_at TIMESTAMP
                  ) AS $$
BEGIN
    RETURN QUERY
        SELECT
            item.id,
            item.title,
            item.description,
            item.price,
            item.quantity,
            item.created_at,
            item.updated_at
        FROM item;
END;
$$ LANGUAGE plpgsql;
```

### Procedure for getting data for one item via id

```sql
CREATE OR REPLACE FUNCTION get_item_by_id(p_id INT)
    RETURNS TABLE (
                      id INT,
                      title VARCHAR,
                      description TEXT,
                      price INT,
                      quantity INT,
                      created_at TIMESTAMP,
                      updated_at TIMESTAMP
                  ) AS $$
BEGIN
    RETURN QUERY
        SELECT
            item.id,
            item.title,
            item.description,
            item.price,
            item.quantity,
            item.created_at,
            item.updated_at
        FROM item
        WHERE item.id = p_id;
END;
$$ LANGUAGE plpgsql;
```

### Procedure for updating data

```sql
CREATE OR REPLACE FUNCTION update_item(
    item_id INT,
    item_title TEXT,
    item_description TEXT,
    item_price INT,
    item_quantity INT
) RETURNS VOID AS $$
BEGIN
    UPDATE items
    SET title = item_title,
        description = item_description,
        price = item_price,
        quantity = item_quantity
    WHERE id = item_id;
END;
$$ LANGUAGE plpgsql;
```
