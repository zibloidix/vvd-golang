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
	mo_id integer,
	name varchar(128),
	foreign key(mo_id) references hospitals(id)
);

-- Таблица врачей
create table doctors (
	id integer primary key autoincrement,
	hospital_id integer,
	spec_id integer,
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



