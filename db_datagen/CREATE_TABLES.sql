CREATE TABLE STUDENTS (
    student_id VARCHAR PRIMARY KEY,
    group_id VARCHAR,
    first_name VARCHAR,
    last_name VARCHAR,
    middle_name VARCHAR,
    snils VARCHAR,
    birth_year DATE,
    auth_email VARCHAR,
    auth_hash VARCHAR,
    FOREIGN KEY (group_id) REFERENCES GROUPS(group_id)
);

CREATE TABLE UNIVERSITY (
    university_id SERIAL PRIMARY KEY,
    name VARCHAR,
    location VARCHAR
);

CREATE TABLE GROUPS (
    group_id VARCHAR PRIMARY KEY,
    cafedra_id VARCHAR,
    year INT,
    major_id INT,
    FOREIGN KEY (cafedra_id) REFERENCES CAFEDRA(cafedra_id),
    FOREIGN KEY (major_id) REFERENCES MAJOR(major_id)
);

CREATE TABLE CAFEDRA (
    cafedra_id VARCHAR PRIMARY KEY,
    cafedra_name VARCHAR,
    institute_id INT,
    FOREIGN KEY (institute_id) REFERENCES INSTITUTE(institute_id)
);

CREATE TABLE MAJOR (
    major_id SERIAL PRIMARY KEY,
    major_name VARCHAR,
    institute_id INT,
    major_type INT,
    major_code VARCHAR,
    FOREIGN KEY (institute_id) REFERENCES INSTITUTE(institute_id)
);

CREATE TABLE INSTITUTE (
    institute_id SERIAL PRIMARY KEY,
    institute_name VARCHAR,
    university_id INT,
    FOREIGN KEY (university_id) REFERENCES UNIVERSITY(university_id)
);

CREATE TABLE DISCIPLINE (
    discipline_id SERIAL PRIMARY KEY,
    discipline_name VARCHAR,
    cafedra_id VARCHAR,
    lector_fio VARCHAR,
    FOREIGN KEY (cafedra_id) REFERENCES CAFEDRA(cafedra_id)
);

CREATE TABLE VISITS (
    lesson_id INT,
    student_id VARCHAR,
    date TIMESTAMP,
    FOREIGN KEY (lesson_id) REFERENCES LESSONS(lesson_id),
    FOREIGN KEY (student_id) REFERENCES STUDENTS(student_id)
);

CREATE TABLE SHEDULE (
    id SERIAL PRIMARY KEY,
    date TIMESTAMP,
    lesson_id INT,
    group_id VARCHAR,
    room INT,
    professor_id INT,
    FOREIGN KEY (lesson_id) REFERENCES LESSONS(lesson_id),
    FOREIGN KEY (group_id) REFERENCES GROUPS(group_id)
);

CREATE TABLE ROOMS_TYPE (
    type_id SERIAL PRIMARY KEY,
    description VARCHAR
);

CREATE TABLE MAP_DISCIPLINE_GROUP (
    group_id VARCHAR,
    discipline_id INT,
    FOREIGN KEY (group_id) REFERENCES GROUPS(group_id),
    FOREIGN KEY (discipline_id) REFERENCES DISCIPLINE(discipline_id)
);

CREATE TABLE LESSON_TYPE (
    type SERIAL PRIMARY KEY,
    type_name VARCHAR
);

CREATE TABLE LESSONS (
    lesson_id SERIAL PRIMARY KEY,
    discipline_id INT,
    lesson_type INT,
    lesson_name VARCHAR,
    required_room_type INT,
    FOREIGN KEY (discipline_id) REFERENCES DISCIPLINE(discipline_id),
    FOREIGN KEY (lesson_type) REFERENCES LESSON_TYPE(type),
    FOREIGN KEY (required_room_type) REFERENCES ROOMS_TYPE(type_id)
);

CREATE TABLE THEMES (
    theme_id SERIAL PRIMARY KEY,
    lesson_id INT,
    name VARCHAR,
    material_name VARCHAR,
    material_path VARCHAR,
    FOREIGN KEY (lesson_id) REFERENCES LESSONS(lesson_id)
);
