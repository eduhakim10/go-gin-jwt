CREATE TABLE company (
    id SERIAL,
    c_name character varying(255),
    c_phone character varying(255),
    c_email character varying(255),
    c_address character varying(255),
    c_fax character varying(255),
    province_id bigint,
    regenci_id bigint,
    c_postal_code character varying(255),
    c_image character varying(255),
    c_pic_name character varying(255),
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone,
    status smallint DEFAULT 1,
    del smallint DEFAULT 0,
    created_by character varying(50),
    updated_by character varying(50),
    c_code character varying(20),
    c_img character varying(200),
    c_img_dir character varying(200),
    c_website character varying(200),
    app_name character varying(200),
    api_code character varying(50),
    is_non_vendor character varying(20),
    midtrans_id character varying(200) DEFAULT NULL::character varying
);
CREATE TABLE teacher (
    id bigint NOT NULL,
    company_id integer,
    nik character varying(255),
    name character varying(255),
    departement character varying(255),
    "position" character varying(255),
    created_by character varying(255),
    updated_by character varying(255),
    status integer,
    del integer,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone,
    departement_id integer,
    phone character varying,
    gender character varying,
    address text,
    image_ character varying(255),
    email character varying(200)
);
CREATE TABLE master_study (
    id SERIAL,
    company_id smallint,
    study_name character varying(200) DEFAULT NULL::character varying,
    code character varying(200) DEFAULT NULL::character varying,
    atas_nama character varying(200) DEFAULT NULL::character varying,
    status smallint DEFAULT 1,
    created_at timestamp(0) without time zone,
    created_by character varying(30) DEFAULT NULL::character varying,
    updated_by timestamp(0) without time zone,
    updated_at timestamp(0) without time zone DEFAULT NULL::timestamp without time zone,
    del smallint DEFAULT 0
);
CREATE TABLE teacher_study (
    id SERIAL,
    company_id smallint,
    teacher_id integer,
    code_teacher character varying(200) DEFAULT NULL::character varying,
    study_id integer,
    code_study character varying(200) DEFAULT NULL::character varying,
    status smallint DEFAULT 1,
    created_at timestamp(0) without time zone,
    created_by character varying(30) DEFAULT NULL::character varying,
    updated_by timestamp(0) without time zone,
    updated_at timestamp(0) without time zone DEFAULT NULL::timestamp without time zone,
    del smallint DEFAULT 0
);
CREATE TABLE student (
    id SERIAL,
    company_id integer,
    code character varying(255),
    name character varying(255),
    departement character varying(255),
    "position" character varying(255),
    created_by character varying(255),
    updated_by character varying(255),
    status integer,
    del integer,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone,
    email character varying(200)
);

CREATE TABLE student_class (
    id SERIAL,
    company_id integer,
    class_id integer,
    code_class character varying(255),
    name character varying(255),
    "class_name" character varying(255),
    created_by character varying(255),
    updated_by character varying(255),
    status integer,
    del integer,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone,
    email character varying(200)
);
CREATE TABLE master_class (
    id SERIAL,
    class_name character varying(30),
    code character varying,
    status character varying(30),
    created_at timestamp without time zone,
    created_by character varying(100),
    updated_at timestamp without time zone,
    updated_by character varying(100),
    del integer DEFAULT 0,
    class_desc text
);

CREATE TABLE master_bank (
    id SERIAL,
    company_id smallint,
    bank_id smallint,
    bank_name character varying(200) DEFAULT NULL::character varying,
    no_rek character varying(200) DEFAULT NULL::character varying,
    atas_nama character varying(200) DEFAULT NULL::character varying,
    status smallint DEFAULT 1,
    created_at timestamp(0) without time zone,
    created_by character varying(30) DEFAULT NULL::character varying,
    updated_by timestamp(0) without time zone,
    updated_at timestamp(0) without time zone DEFAULT NULL::timestamp without time zone,
    del smallint DEFAULT 0
);
CREATE TABLE payment_gateway (
    id SERIAL,
    bank_id smallint,
    bank_name character varying(200) DEFAULT NULL::character varying,
    no_rek character varying(200) DEFAULT NULL::character varying,
    atas_nama character varying(200) DEFAULT NULL::character varying,
    status smallint DEFAULT 1,
    created_at timestamp(0) without time zone,
    created_by character varying(30) DEFAULT NULL::character varying,
    updated_by timestamp(0) without time zone,
    updated_at timestamp(0) without time zone DEFAULT NULL::timestamp without time zone,
    del smallint DEFAULT 0
);
CREATE TABLE payment_config (
    id SERIAL,
    api_code character varying(200) DEFAULT NULL::character varying,
    status smallint DEFAULT 1,
    created_at timestamp(0) without time zone,
    created_by character varying(30) DEFAULT NULL::character varying,
    updated_by timestamp(0) without time zone,
    updated_at timestamp(0) without time zone DEFAULT NULL::timestamp without time zone,
    del smallint DEFAULT 0
);
CREATE TABLE payment_purpose (
    id SERIAL,
    purp_code character varying(200) DEFAULT NULL::character varying,
    purp_name character varying(200) DEFAULT NULL::character varying,
    cost character varying(200) DEFAULT NULL::character varying,
    status smallint DEFAULT 1,
    created_at timestamp(0) without time zone,
    created_by character varying(30) DEFAULT NULL::character varying,
    updated_by timestamp(0) without time zone,
    updated_at timestamp(0) without time zone DEFAULT NULL::timestamp without time zone,
    del smallint DEFAULT 0
);
CREATE TABLE trx_payment (
    id SERIAL,
    bank_id smallint,
    bank_name character varying(200) DEFAULT NULL::character varying,
    no_rek character varying(200) DEFAULT NULL::character varying,
    atas_nama character varying(200) DEFAULT NULL::character varying,
    code_student character varying(30) DEFAULT NULL::character varying,
    name_student character varying(150) DEFAULT NULL::character varying,
    purp_code character varying(30) DEFAULT NULL::character varying,
    payment_name character varying(30) DEFAULT NULL::character varying,
    payment_desc character varying(30) DEFAULT NULL::character varying,
    status smallint DEFAULT 1,
    expired_date timestamp(0) without time zone,
    created_at timestamp(0) without time zone,
    created_by character varying(30) DEFAULT NULL::character varying,
    updated_by timestamp(0) without time zone,
    updated_at timestamp(0) without time zone DEFAULT NULL::timestamp without time zone,
    del smallint DEFAULT 0
);

