CREATE TABLE folders (
    id SERIAL,
    parent_id INT,
    name VARCHAR(150) NOT NULL,
    created_at TIMESTAMP default current_timestamp,
    modified_at TIMESTAMP NOT NULL,
    deleted BOOL NOT NULL DEFAULT false,
    PRIMARY KEY (id),
    CONSTRAINT fk_parent
        FOREIGN KEY (parent_id)
            REFERENCES folders(id)
)