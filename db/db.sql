CREATE SEQUENCE IF NOT EXISTS user_id_seq;
CREATE TABLE "public"."user" (
    "user_name" varchar(50) NOT NULL,
    "password" varchar(120) NOT NULL,
    "full_name" varchar(50) NOT NULL,
    "address" varchar(50) NOT NULL,
    "email" varchar(50) NOT NULL,
    "created_on" int8,
    "last_login" int8,
    "id" int4 NOT NULL DEFAULT nextval('user_id_seq'::regclass)
);

CREATE SEQUENCE IF NOT EXISTS rectangle_id_seq;
CREATE TABLE "public"."rectangle" (
    "id" int4 NOT NULL DEFAULT nextval('rectangle_id_seq'::regclass),
    "name" varchar(50) NOT NULL,
    "length" int4 NOT NULL,
    "width" int4 NOT NULL
);

CREATE SEQUENCE IF NOT EXISTS square_id_seq;
CREATE TABLE "public"."square" (
    "id" int4 NOT NULL DEFAULT nextval('square_id_seq'::regclass),
    "name" varchar(50) NOT NULL,
    "side" int4 NOT NULL
);

CREATE SEQUENCE IF NOT EXISTS triangle_id_seq;
CREATE TABLE "public"."triangle" (
    "id" int4 NOT NULL DEFAULT nextval('triangle_id_seq'::regclass),
    "name" varchar(50) NOT NULL,
    "side1" int4 NOT NULL,
    "side2" int4 NOT NULL,
    "base" int4 NOT NULL,
    "height" int4 NOT NULL
);

CREATE SEQUENCE IF NOT EXISTS diamond_id_seq;
CREATE TABLE "public"."diamond" (
    "id" int4 NOT NULL DEFAULT nextval('diamond_id_seq'::regclass),
    "name" varchar(50) NOT NULL,
    "side" int4 NOT NULL,
    "diagonal1" int4 NOT NULL,
    "diagonal2" int4 NOT NULL
);