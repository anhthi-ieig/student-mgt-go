CREATE TABLE IF NOT EXISTS student (
	id int4 NOT NULL GENERATED ALWAYS AS IDENTITY,
	"name" text NULL,
	created_at timestamp NULL DEFAULT now(),
	CONSTRAINT student_pk PRIMARY KEY (id)
);