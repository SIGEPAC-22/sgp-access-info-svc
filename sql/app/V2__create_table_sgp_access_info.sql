create table ${flyway:database}.std_state_data
(
    std_id int auto_increment primary key,
    std_state_name varchar(45) not null,
    std_creation_date timestamp null,
    std_created_by varchar(45) null,
    std_modification_date timestamp null,
    std_modified_by varchar(45) null
);

-- trigger psu_state_data
create trigger ${flyway:database}.psi_state_data_insert_aud
    BEFORE INSERT ON ${flyway:database}.std_state_data
    FOR EACH ROW
    set NEW.std_created_by=USER(),
        NEW.std_creation_date=now();

create trigger ${flyway:database}.std_state_data_update_aud
    BEFORE UPDATE ON ${flyway:database}.std_state_data
    FOR EACH ROW
    set NEW.std_modification_date=now(),
        NEW.std_modified_by=USER();



create table ${flyway:database}.pss_personal_sex
(
    pss_id int auto_increment primary key,
    pss_gender_name varchar(50) not null,
    pss_state_data_id int not null default 1,
    pss_creation_date timestamp null,
    pss_created_by varchar(45) null,
    pss_modification_date timestamp null,
    pss_modified_by varchar(45) null,
    CONSTRAINT fk_state_data_personal FOREIGN KEY (pss_state_data_id) REFERENCES std_state_data (std_id)
);

-- trigger pss_personal_sex
create trigger ${flyway:database}.pss_sex_insert_aud
    BEFORE INSERT ON ${flyway:database}.pss_personal_sex
    FOR EACH ROW
    set NEW.pss_created_by=USER(),
        NEW.pss_creation_date=now();

create trigger ${flyway:database}.pss_sex_update_aud
    BEFORE UPDATE ON ${flyway:database}.pss_personal_sex
    FOR EACH ROW
    set NEW.pss_modification_date=now(),
        NEW.pss_modified_by=USER();


create table ${flyway:database}.dct_document_type
(
    dct_id int auto_increment primary key,
    dct_document_name varchar(45) not null,
    dct_state_data_id int not null default 1,
    dct_creation_date timestamp null,
    dct_created_by varchar(45) null,
    dct_modification_date timestamp null,
    dct_modified_by varchar(45) null,
    CONSTRAINT fk_state_data_document_type FOREIGN KEY (dct_state_data_id) REFERENCES std_state_data (std_id)
);

-- psi_document_type
create trigger ${flyway:database}.dct_document_ty_insert_aud
    BEFORE INSERT ON ${flyway:database}.dct_document_type
    FOR EACH ROW
    set NEW.dct_created_by=USER(),
        NEW.dct_creation_date=now();

create trigger ${flyway:database}.dct_document_ty_update_aud
    BEFORE UPDATE ON ${flyway:database}.dct_document_type
    FOR EACH ROW
    set NEW.dct_modification_date=now(),
        NEW.dct_modified_by=USER();



create table ${flyway:database}.tur_type_user
(
    tur_id int auto_increment primary key,
    tur_name_type varchar(45) not null,
    tur_description_type varchar(100) null,
    tur_state_data_id int not null default 1,
    tur_creation_date timestamp null,
    tur_created_by varchar(45) null,
    tur_modification_date timestamp null,
    tur_modified_by varchar(45) null,
    CONSTRAINT fk_state_data_type_user FOREIGN KEY (tur_state_data_id) REFERENCES std_state_data (std_id)
);

-- tur_type_user
create trigger ${flyway:database}.tur_type_user_insert_aud
    BEFORE INSERT ON ${flyway:database}.tur_type_user
    FOR EACH ROW
    set NEW.tur_created_by=USER(),
        NEW.tur_creation_date=now();

create trigger ${flyway:database}.tur_type_user_update_aud
    BEFORE UPDATE ON ${flyway:database}.tur_type_user
    FOR EACH ROW
    set NEW.tur_modification_date=now(),
        NEW.tur_modified_by=USER();



create table ${flyway:database}.psi_personal_information
(
    psi_id int auto_increment primary key not null,
    psi_first_name varchar(50) not null,
    psi_second_name varchar(50) null,
    psi_first_last_name varchar(50) not null,
    psi_second_last_name varchar(50) null,
    psi_sex_id int not null,
    psi_date_of_birth date not null,
    psi_document_type_id int not null,
    psi_document_number varchar(45) unique null,
    psi_user varchar(45) not null,
    psi_password varchar(45) not null,
    psi_type_user int not null,
    psi_account_creation_date date not null,
    psi_data_of_last_use date null,
    psi_state_data_id int not null default 1,
    psi_creation_date timestamp null,
    psi_created_by varchar(45) null,
    psi_modification_date timestamp null,
    psi_modified_by varchar(45) null,
    CONSTRAINT fk_sex FOREIGN KEY (psi_sex_id) REFERENCES pss_personal_sex (pss_id),
    CONSTRAINT fk_document_type FOREIGN KEY (psi_document_type_id) REFERENCES dct_document_type (dct_id),
    CONSTRAINT fk_state_data FOREIGN KEY (psi_state_data_id) REFERENCES std_state_data (std_id),
    CONSTRAINT fk_type_user FOREIGN KEY (psi_type_user) REFERENCES tur_type_user (tur_id)
);

-- trigger psi_personal_information
create trigger ${flyway:database}.psi_personal_inf_insert_aud
    BEFORE INSERT ON ${flyway:database}.psi_personal_information
    FOR EACH ROW
    set NEW.psi_created_by=USER(),
        NEW.psi_creation_date=now();

create trigger ${flyway:database}.psi_personal_inf_update_aud
    BEFORE UPDATE ON ${flyway:database}.psi_personal_information
    FOR EACH ROW
    set NEW.psi_modification_date=now(),
        NEW.psi_modified_by=USER();