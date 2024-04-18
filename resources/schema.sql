create table product
(
    id          serial       not null primary key,
    name        varchar(255) not null,
    description varchar(255) not null,
    price       numeric      not null,
    height      numeric      not null,
    width       numeric      not null,
    length      numeric      not null,
    weight      numeric      not null,
    created_at  timestamp    not null DEFAULT CURRENT_TIMESTAMP,
    updated_at  timestamp    not null DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

create table coupon
(
    id              serial      not null primary key,
    code            varchar(36) not null,
    percentage      numeric     not null,
    expiration_date timestamp   not null,
    created_at      timestamp   not null DEFAULT CURRENT_TIMESTAMP,
    updated_at      timestamp   not null DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

create table `order`
(
    id         serial          not null primary key,
    code       varchar(36)     not null,
    cpf        varchar(11)     not null,
    status     varchar(20)     not null,
    total      numeric         not null,
    coupon_id  bigint unsigned null,
    created_at timestamp       not null DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp       not null DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (coupon_id) REFERENCES coupon (id)
);

create table item
(
    id         serial          not null primary key,
    order_id   bigint unsigned not null,
    product_id bigint unsigned not null,
    amount     int             not null,
    price      numeric         not null,
    created_at timestamp       not null DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp       not null DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (product_id) REFERENCES product (id),
    FOREIGN KEY (order_id) REFERENCES `order` (id)
);

create table freight
(
    id         serial          not null primary key,
    order_id   bigint unsigned not null,
    code       varchar(36)     not null,
    price      numeric         not null,
    duration   numeric         not null,
    eta        timestamp       not null,
    FOREIGN KEY (order_id) REFERENCES `order` (id),
    created_at timestamp       not null DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp       not null DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);


insert into product (name, description, price, height, width, length, weight)
values ('Smartwatch', 'Amazfit series 10', 300, 10.00, 10.00, 10.00, 10.00);

insert into coupon (code, percentage, expiration_date)
values ('APP_10', 10, '2023-12-31 23:59:59');
