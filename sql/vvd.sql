drop table hospitals;
drop table specs;
drop table doctors;
drop table patients;
drop table sessions;
drop table logs;

-- Таблица медицинских организаций
create table hospitals (
	id integer primary key autoincrement,
	oid varchar(128),
	name varchar(128),
	address varchar(128),
	phone varchar(128)
);


-- Таблица специальностей
create table specs (
	id integer primary key autoincrement,
	code integer,
	name varchar(128)
);

-- Таблица врачей
create table doctors (
	id integer primary key autoincrement,
	hospital_id integer,
	spec_id integer,
	snils varchar(128),
	doctor_code varchar(128),
	name varchar(128),
	foreign key(hospital_id) references hospitals(id),
	foreign key(spec_id) references specs(id)
);

-- Таблица пациентов
create table patients (
	id integer primary key autoincrement,
	hospital_id integer,
	oms varchar(128),
	snils varchar(128),
	first_name varchar(128),
	last_name varchar(128),
	middle_name varchar(128),
	birth_date date,
	sex char,
	email varchar(128),
	phone varchar(128),
	foreign key(hospital_id) references hospitals(id)
)

-- Таблица сессий, для элемента slot_id
create table sessions (
	id integer primary key autoincrement,
	uuid varchar(128),
	patient_id integer,
	create_date datetime,
	expire_date datetime,
	foreign key(patient_id) references patients(id)
);

-- Таблица слотов для записи
create table slots (
	id integer primary key autoincrement,
	doctor_id integer,
	uuid varchar(128),
	visit_date datetime,
	foreign key(doctor_id) references doctors(id)
);

create table logs (
	id integer primary key autoincrement,
	session_id integer,
	create_date datetime,
	host varchar(128),
	url varchar(128),
	method varchar(128),
	action varchar(128),
	request varchar(10000),
	response varchar(10000),
	foreign key(session_id) references sessions(id)
);

-- Добавление медицинских организаций
insert into hospitals (oid, name, address, phone) values ("1.2.643.5.1.13.13.12.2.65.6754", "ГБУЗ Городская поликлиника № 1", "​693010, Южно-Сахалинск, Коммунистический проспект, 11", "8(4242)43-31-54");
insert into hospitals (oid, name, address, phone) values ("1.2.643.5.1.13.13.12.2.65.6742", "ГБУЗ Городская поликлиника № 2", "693010, Южно-Сахалинск, Проспект Мира, 85", "8(4242)30-00-34");

-- Добавление пациентов
insert into patients (hospital_id, oms, snils, first_name, middle_name, last_name, birth_date, sex, email, phone) values (1, "1541981080564920", "85454789361", "Грегори", "Петрочич", "Хауз", "1959-05-15", "M", "g.house@gmail.com", "+79621274565");
insert into patients (hospital_id, oms, snils, first_name, middle_name, last_name, birth_date, sex, email, phone) values (2, "3421989010269800", "24100692415", "Лиза", "Ивановна", "Кадди", "1968-01-25", "F", "l.kaddy@gmail.com", "+79621274767");

-- Добавление специйальностей
insert into specs (code, name) values (109, "врач-терапевт");

-- Добавление врачей
insert into doctors (hospital_id, spec_id, snils, doctor_code, name) values (1, 1, "27377963737", "489744019", "Роберт Михайловч Чейз");
insert into doctors (hospital_id, spec_id, snils, doctor_code, name) values (2, 1, "12830301517", "272543789", "Эллисон Борисовна Кэмерон");



