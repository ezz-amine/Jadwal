CREATE TABLE todo_table (
    id INTEGER PRIMARY KEY,
    title VARCHAR(35) NOT NULL
);

CREATE TABLE todo_entry (
    id INTEGER PRIMARY KEY,
    content VARCHAR(255) NOT NULL,
    is_done BOOLEAN,
    is_archived BOOLEAN,
    table_id INTEGER NOT NULL,

    FOREIGN KEY (table_id) REFERENCES todo_table (id)
);
