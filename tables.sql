CREATE TABLE learn_user (
	user_id                 bigserial ,
    username                varchar(50),
	password                varchar(50),

	version                 bigint,
	create_user_id          bigint,
	update_user_id          bigint,
	create_datetime         varchar(14),
	update_datetime         varchar(14),

    active                  varchar(1),
    active_datetime         varchar(14),
    non_active_datetime     varchar(14),

	CONSTRAINT learn_user_pkey PRIMARY KEY(user_id)
);

CREATE TABLE learn_product (
	product_id              bigserial ,
    product_code            varchar(50),
	product_name            varchar(50),
    price                   numeric,

	version                 bigint,
	create_user_id          bigint,
	update_user_id          bigint,
	create_datetime         varchar(14),
	update_datetime         varchar(14),

    active                  varchar(1),
    active_datetime         varchar(14),
    non_active_datetime     varchar(14),

	CONSTRAINT learn_product_pkey PRIMARY KEY(product_id)
);

CREATE TABLE learn_penjualan (
	penjualan_id            bigserial ,
    total_penjualan         numeric,
	total_pembayaran        numeric,
    total_kembalian         numeric,

	version                 bigint,
	create_user_id          bigint,
	update_user_id          bigint,
	create_datetime         varchar(14),
	update_datetime         varchar(14),

	CONSTRAINT learn_penjualan_pkey PRIMARY KEY(penjualan_id)
);

CREATE TABLE learn_penjualan_item (
	penjualan_item_id       bigserial ,
    penjualan_id            bigint,
	product_id              bigint,
    qty                     numeric,
    price                   numeric,

	version                 bigint,
	create_user_id          bigint,
	update_user_id          bigint,
	create_datetime         varchar(14),
	update_datetime         varchar(14),

	CONSTRAINT learn_penjualan_item_pkey PRIMARY KEY(penjualan_item_id)
);

INSERT INTO learn_user (
	username,
	password,
	version,
	create_user_id,
	update_user_id,
	create_datetime,
	update_datetime,

	active,
	active_datetime,
	non_active_datetime
) VALUES
	('sts', 'sts123', 0, NULL, NULl, TO_CHAR(NOW(), 'YYYYMMDDHH24MI'), TO_CHAR(NOW(), 'YYYYMMDDHH24MI'), 'Y', TO_CHAR(NOW(), 'YYYYMMDDHH24MI'), TO_CHAR(NOW(), 'YYYYMMDDHH24MI'));

INSERT INTO learn_product (
	product_code,
	product_name,
	price,
	version,
	create_user_id,
	update_user_id,
	create_datetime,
	update_datetime,
	active,
	active_datetime,
	non_active_datetime
) VALUES 
	('P001', 'Beras Pandan Wangi', 			120000, 0, NULL, NULL, TO_CHAR(NOW(), 'YYYYMMDDHH24MI'), TO_CHAR(NOW(), 'YYYYMMDDHH24MI'), 'Y', TO_CHAR(NOW(), 'YYYYMMDDHH24MI'), NULL),
	('P002', 'Gula Merah Organik', 			30000, 	0, NULL, NULL, TO_CHAR(NOW(), 'YYYYMMDDHH24MI'), TO_CHAR(NOW(), 'YYYYMMDDHH24MI'), 'Y', TO_CHAR(NOW(), 'YYYYMMDDHH24MI'), NULL),
	('P003', 'Minyak Goreng Kemasan', 		20000, 	0, NULL, NULL, TO_CHAR(NOW(), 'YYYYMMDDHH24MI'), TO_CHAR(NOW(), 'YYYYMMDDHH24MI'), 'Y', TO_CHAR(NOW(), 'YYYYMMDDHH24MI'), NULL),
	('P004', 'Tepung Terigu Segitiga Biru', 10000, 	0, NULL, NULL, TO_CHAR(NOW(), 'YYYYMMDDHH24MI'), TO_CHAR(NOW(), 'YYYYMMDDHH24MI'), 'Y', TO_CHAR(NOW(), 'YYYYMMDDHH24MI'), NULL),
	('P005', 'Telur Ayam Ras', 				27000, 	0, NULL, NULL, TO_CHAR(NOW(), 'YYYYMMDDHH24MI'), TO_CHAR(NOW(), 'YYYYMMDDHH24MI'), 'Y', TO_CHAR(NOW(), 'YYYYMMDDHH24MI'), NULL),
	('P006', 'Kacang Tanah Panggang', 		12000, 	0, NULL, NULL, TO_CHAR(NOW(), 'YYYYMMDDHH24MI'), TO_CHAR(NOW(), 'YYYYMMDDHH24MI'), 'Y', TO_CHAR(NOW(), 'YYYYMMDDHH24MI'), NULL),
	('P007', 'Ikan Asin Jambal Roti:', 		35000, 	0, NULL, NULL, TO_CHAR(NOW(), 'YYYYMMDDHH24MI'), TO_CHAR(NOW(), 'YYYYMMDDHH24MI'), 'Y', TO_CHAR(NOW(), 'YYYYMMDDHH24MI'), NULL),
	('P008', 'Bawang Putih Organik', 		25000, 	0, NULL, NULL, TO_CHAR(NOW(), 'YYYYMMDDHH24MI'), TO_CHAR(NOW(), 'YYYYMMDDHH24MI'), 'Y', TO_CHAR(NOW(), 'YYYYMMDDHH24MI'), NULL),
	('P009', 'Cabai Rawit Merah', 			15000, 	0, NULL, NULL, TO_CHAR(NOW(), 'YYYYMMDDHH24MI'), TO_CHAR(NOW(), 'YYYYMMDDHH24MI'), 'Y', TO_CHAR(NOW(), 'YYYYMMDDHH24MI'), NULL),
	('P010', 'Kecap Manis Tradisional', 	10000, 	0, NULL, NULL, TO_CHAR(NOW(), 'YYYYMMDDHH24MI'), TO_CHAR(NOW(), 'YYYYMMDDHH24MI'), 'Y', TO_CHAR(NOW(), 'YYYYMMDDHH24MI'), NULL);

