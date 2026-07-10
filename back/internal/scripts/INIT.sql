/*--------------------------------------------TABLAS AUTENTICACIÓN--------------------------------------------*/
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE rol (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) UNIQUE NOT NULL, -- admin, seller, buyer, etc
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE permission (
    id SERIAL PRIMARY KEY,
    name VARCHAR(150) UNIQUE NOT NULL, -- ej: "create_product"
    description TEXT
);

CREATE TABLE user_rol (
    user_id INT NOT NULL,
    role_id INT NOT NULL,

    PRIMARY KEY (user_id, role_id),

    FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE,

    FOREIGN KEY (role_id)
        REFERENCES rol(id)
        ON DELETE CASCADE
);

CREATE TABLE role_permission (
    role_id INT NOT NULL,
    permission_id INT NOT NULL,

    PRIMARY KEY (role_id, permission_id),

    CONSTRAINT fk_role_permission_role
        FOREIGN KEY (role_id)
        REFERENCES rol(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_role_permission_permission
        FOREIGN KEY (permission_id)
        REFERENCES permission(id)
        ON DELETE CASCADE
);
/*--------------------------------------------TABLAS SISTEMA--------------------------------------------*/
CREATE TABLE client (
    id SERIAL PRIMARY KEY,
    user_id INT UNIQUE, -- 1 a 1 opcional
    company_id INT,
    first_name VARCHAR(100),
    last_name VARCHAR(100),
    email VARCHAR(255),
    phone VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    status INT DEFAULT 1, -- 1 activo, 0 inactivo

    CONSTRAINT fk_user
        FOREIGN KEY(user_id)
        REFERENCES users(id)
        ON DELETE SET NULL
);

CREATE TABLE category_company (
    id SERIAL PRIMARY KEY,
    name VARCHAR(80) NOT NULL UNIQUE,
    description TEXT
);

CREATE TABLE company (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    category_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    status INT DEFAULT 1, -- 1 activo, 0 inactivo
    description TEXT,
    logo_path TEXT,
    cover_image_path TEXT,

    CONSTRAINT fk_company_category
        FOREIGN KEY (category_id)
        REFERENCES category_company(id)
        ON DELETE RESTRICT
);
ALTER TABLE client
ADD CONSTRAINT fk_company
FOREIGN KEY (company_id)
REFERENCES company(id)
ON DELETE SET NULL;

CREATE TABLE category_product (
    id SERIAL PRIMARY KEY,
    name VARCHAR(80) NOT NULL UNIQUE,
    description TEXT
);

CREATE TABLE product (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    type VARCHAR(20), -- product | service
    category_id INT NOT NULL,
    price NUMERIC(10,2) NOT NULL,
    stock INT DEFAULT 0,
    company_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    status INT DEFAULT 1, -- 1 activo, 0 inactivo

    CONSTRAINT fk_product_company
        FOREIGN KEY (company_id)
        REFERENCES company(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_product_category
        FOREIGN KEY (category_id)
        REFERENCES category_product(id)
        ON DELETE RESTRICT
);

CREATE INDEX IF NOT EXISTS idx_company_category_id ON company(category_id);
CREATE INDEX IF NOT EXISTS idx_product_category_id ON product(category_id);

CREATE TABLE invoice (
    id SERIAL PRIMARY KEY,
    buyer_client_id INT NOT NULL,
    seller_company_id INT NOT NULL,
    created_by_user_id INT NOT NULL,
    total_amount NUMERIC(10,2) NOT NULL,
    status_invoice VARCHAR(50) NOT NULL, -- pending, paid, cancelled
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    status INT DEFAULT 1, -- 1 activo, 0 inactivo

    CONSTRAINT fk_invoice_buyer
        FOREIGN KEY (buyer_client_id)
        REFERENCES client(id),

    CONSTRAINT fk_invoice_seller_company
        FOREIGN KEY (seller_company_id)
        REFERENCES company(id),

    CONSTRAINT fk_invoice_created_by
        FOREIGN KEY (created_by_user_id)
        REFERENCES users(id)
);

CREATE TABLE invoice_item (
    id SERIAL PRIMARY KEY,
    invoice_id INT NOT NULL,
    product_id INT NOT NULL,
    quantity INT NOT NULL,
    price NUMERIC(10,2) NOT NULL,

    CONSTRAINT fk_invoice
        FOREIGN KEY(invoice_id)
        REFERENCES invoice(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_product
        FOREIGN KEY(product_id)
        REFERENCES product(id)
);

CREATE INDEX idx_client_user_id ON client(user_id);
CREATE INDEX idx_product_company_id ON product(company_id);
CREATE INDEX idx_invoice_buyer ON invoice(buyer_client_id);

ALTER TABLE user_rol
ADD CONSTRAINT user_role_unique UNIQUE (user_id, role_id);

ALTER TABLE users
ADD COLUMN status SMALLINT DEFAULT 1,
ADD COLUMN deleted_at TIMESTAMP NULL;

ALTER TABLE invoice
ADD COLUMN subtotal NUMERIC(10,2) DEFAULT 0,
ADD COLUMN taxes NUMERIC(10,2) DEFAULT 0,
ADD COLUMN paid_amount NUMERIC(10,2) DEFAULT 0,
ADD COLUMN updated_at TIMESTAMP NULL;

ALTER TABLE invoice_item
ADD COLUMN product_name VARCHAR(255) NOT NULL DEFAULT '',
ADD COLUMN subtotal NUMERIC(10,2) DEFAULT 0;

CREATE TABLE invoice_payment (
    id SERIAL PRIMARY KEY,

    invoice_id INT NOT NULL,
    amount NUMERIC(10,2) NOT NULL,

    payment_method VARCHAR(50),

    paid_by_user_id INT NOT NULL,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_payment_invoice
        FOREIGN KEY(invoice_id)
        REFERENCES invoice(id),

    CONSTRAINT fk_payment_user
        FOREIGN KEY(paid_by_user_id)
        REFERENCES users(id)
);

-- STOCK RESERVATION ENGINE PRO
ALTER TABLE product
ADD COLUMN reserved_stock INT NOT NULL DEFAULT 0;

ALTER TABLE product
ADD COLUMN image_path TEXT;

ALTER TABLE invoice_item
ADD COLUMN status SMALLINT DEFAULT 1,
ADD COLUMN deleted_at TIMESTAMP NULL,
ADD COLUMN created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
ADD COLUMN updated_at TIMESTAMP NULL;




