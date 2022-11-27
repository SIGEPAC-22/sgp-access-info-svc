INSERT INTO ${flyway:database}.std_state_data(std_state_name) VALUES('ACTIVO');
INSERT INTO ${flyway:database}.std_state_data(std_state_name) VALUES('INACTIVO');

INSERT INTO ${flyway:database}.pss_personal_sex(pss_gender_name) VALUES('MASCULINO');
INSERT INTO ${flyway:database}.pss_personal_sex(pss_gender_name) VALUES('FEMENINA');

INSERT INTO ${flyway:database}.dct_document_type(dct_document_name) VALUES('N/A');
INSERT INTO ${flyway:database}.dct_document_type(dct_document_name) VALUES('DUI');
INSERT INTO ${flyway:database}.dct_document_type(dct_document_name) VALUES('NIT');
INSERT INTO ${flyway:database}.dct_document_type(dct_document_name) VALUES('PASAPORTE');

INSERT INTO ${flyway:database}.tur_type_user(tur_name_type, tur_description_type) VALUES('ADMIN', 'ACCESO COMPLETO AL SISTEMA');
INSERT INTO ${flyway:database}.tur_type_user(tur_name_type, tur_description_type) VALUES('USER', 'ACCESO LIMITADO AL SISTEMA');

