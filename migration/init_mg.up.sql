CREATE TABLE orders
(
    id                 serial primary key,
    order_uid          varchar(255) unique,
    track_number       varchar(255) not null,
    entry              varchar(255) not null,
    locale             varchar(255) not null,
    internal_signature varchar(255),
    customer_id        varchar(255) not null,
    delivery_service   varchar(255) not null,
    shardkey           varchar(255) not null,
    sm_id              int not null,
    date_created       timestamp not null,
    oof_shard          varchar(255) not null
);

CREATE TABLE deliveries
(
    id       serial primary key,
    order_id int references orders (id),
    name     varchar(255) not null,
    phone    varchar(255) not null,
    zip      varchar(255) not null,
    city     varchar(255) not null,
    address  varchar(255) not null,
    region   varchar(255) not null,
    email    varchar(255) not null
);

CREATE TABLE payments
(
    id            serial primary key,
    order_id      int references orders (id),
    transaction   varchar(255) not null,
    request_id    varchar(255),
    currency      varchar(255) not null,
    provider      varchar(255) not null,
    amount        int not null,
    payment_dt    int not null,
    bank          varchar(255) not null,
    delivery_cost int not null,
    goods_total   int not null,
    custom_fee    int not null
);

CREATE TABLE items
(
    id           serial primary key,
    order_id     int references orders (id),
    chrt_id      int not null,
    track_number varchar(255) not null,
    price        int not null,
    rid          varchar(255) not null,
    name         varchar(255) not null,
    sale         int not null,
    size         varchar(255) not null,
    total_price  int not null,
    nm_id        int not null,
    brand        varchar(255) not null,
    status       int not null
);