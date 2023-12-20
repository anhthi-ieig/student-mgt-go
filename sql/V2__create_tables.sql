CREATE TABLE IF NOT EXISTS users (
	id int NOT NULL GENERATED ALWAYS AS IDENTITY,
	username text NOT NULL,
	"password" text NULL,
	"name" text NULL,
	"role" text NOT NULL,
	CONSTRAINT users_pk PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS class (
	id int NOT NULL GENERATED ALWAYS AS IDENTITY,
	"name" text NULL,
	"start_date" timestamp NULL,
	end_date timestamp NULL,
	subjects text NULL,
	CONSTRAINT class_pk PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS class_to_student (
	class_id int,
	student_id int,
	CONSTRAINT class_student_fk FOREIGN KEY (class_id) REFERENCES class(id) ON DELETE CASCADE,
	CONSTRAINT student_class_fk FOREIGN KEY (student_id) REFERENCES users(id) ON DELETE CASCADE,
	CONSTRAINT class_student_pk PRIMARY KEY (class_id, student_id)
);

CREATE TABLE IF NOT EXISTS class_to_teacher (
	class_id int,
	teacher_id int,
	CONSTRAINT class_teacher_fk FOREIGN KEY (class_id) REFERENCES class(id) ON DELETE CASCADE,
	CONSTRAINT teacher_class_fk FOREIGN KEY (teacher_id) REFERENCES users(id) ON DELETE CASCADE,
	CONSTRAINT class_teacher_pk PRIMARY KEY (class_id, teacher_id)
);