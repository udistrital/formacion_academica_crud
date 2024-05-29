-- object: formacion_academica | type: SCHEMA --
-- DROP SCHEMA IF EXISTS formacion_academica CASCADE;
CREATE SCHEMA formacion_academica;
-- ddl-end --
COMMENT ON SCHEMA formacion_academica IS 'Esquema para el módulo de formación académica del campus';
-- ddl-end --

SET search_path TO pg_catalog,public,formacion_academica;
-- ddl-end --

-- object: formacion_academica.dato_adicional_formacion_academica | type: TABLE --
-- DROP TABLE IF EXISTS formacion_academica.dato_adicional_formacion_academica CASCADE;
CREATE TABLE formacion_academica.dato_adicional_formacion_academica (
	id serial NOT NULL,
	tipo_dato_adicional integer NOT NULL,
	valor text NOT NULL,
	formacion_academica integer NOT NULL,
	activo boolean NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	CONSTRAINT pk_dato_adicional_formacion_academica PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE formacion_academica.dato_adicional_formacion_academica IS 'Tabla que guarda el valor de los datos adicionales de la formación académica';
-- ddl-end --
COMMENT ON COLUMN formacion_academica.dato_adicional_formacion_academica.id IS 'Identificador de la tabla dato_adicional_formacion_academica';
-- ddl-end --
COMMENT ON COLUMN formacion_academica.dato_adicional_formacion_academica.tipo_dato_adicional IS 'Tipo de dato adicional para la formación académica, referencia a la tabla tipo_dato_adicional del core_new';
-- ddl-end --
COMMENT ON COLUMN formacion_academica.dato_adicional_formacion_academica.valor IS 'Almacena el valor del dato adicional de la formacion academica';
-- ddl-end --
COMMENT ON COLUMN formacion_academica.dato_adicional_formacion_academica.activo IS 'Campo para indicar el estado del registro';
-- ddl-end --
COMMENT ON COLUMN formacion_academica.dato_adicional_formacion_academica.fecha_creacion IS 'Fecha de creación del registro';
-- ddl-end --
COMMENT ON COLUMN formacion_academica.dato_adicional_formacion_academica.fecha_modificacion IS 'Fecha de la última modificación del registro';
-- ddl-end --
COMMENT ON CONSTRAINT pk_dato_adicional_formacion_academica ON formacion_academica.dato_adicional_formacion_academica  IS 'Llave primaria de la tabla dato_adicional_formacion_academica';
-- ddl-end --

-- object: formacion_academica.formacion_academica | type: TABLE --
-- DROP TABLE IF EXISTS formacion_academica.formacion_academica CASCADE;
CREATE TABLE formacion_academica.formacion_academica (
	id serial NOT NULL,
	persona integer NOT NULL,
	fecha_inicio date NOT NULL,
	fecha_finalizacion date,
	titulacion integer,
	activo boolean NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	CONSTRAINT pk_formacion_academica PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE formacion_academica.formacion_academica IS 'Formación académica de la persona';
-- ddl-end --
COMMENT ON COLUMN formacion_academica.formacion_academica.id IS 'Identificador de la tabla formacion_academica';
-- ddl-end --
COMMENT ON COLUMN formacion_academica.formacion_academica.persona IS 'Referencia a la tabla ente';
-- ddl-end --
COMMENT ON COLUMN formacion_academica.formacion_academica.fecha_inicio IS 'Fecha de inicio de la formación académica';
-- ddl-end --
COMMENT ON COLUMN formacion_academica.formacion_academica.fecha_finalizacion IS 'Fecha de finalización de la formación académica';
-- ddl-end --
COMMENT ON COLUMN formacion_academica.formacion_academica.titulacion IS 'Referencia a la tabla programa_academico del esquema academica';
-- ddl-end --
COMMENT ON COLUMN formacion_academica.formacion_academica.activo IS 'Indica el estado del registro';
-- ddl-end --
COMMENT ON COLUMN formacion_academica.formacion_academica.fecha_creacion IS 'Fecha de creación del registro';
-- ddl-end --
COMMENT ON COLUMN formacion_academica.formacion_academica.fecha_modificacion IS 'Fecha de la última modificación del registro';
-- ddl-end --
COMMENT ON CONSTRAINT pk_formacion_academica ON formacion_academica.formacion_academica  IS 'Llave primaria de la tabla formacion_academica';
-- ddl-end --

-- object: formacion_academica.soporte_formacion_academica | type: TABLE --
-- DROP TABLE IF EXISTS formacion_academica.soporte_formacion_academica CASCADE;
CREATE TABLE formacion_academica.soporte_formacion_academica (
	id serial NOT NULL,
	documento integer NOT NULL,
	descripcion character varying(250),
	formacion_academica integer NOT NULL,
	activo boolean NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	CONSTRAINT pk_soporte_formacion_academica PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE formacion_academica.soporte_formacion_academica IS 'Tabla que almacena los soportes correspondientes a una formacion académica';
-- ddl-end --
COMMENT ON COLUMN formacion_academica.soporte_formacion_academica.id IS 'Identificador de la tabla soporte_formacion_academica';
-- ddl-end --
COMMENT ON COLUMN formacion_academica.soporte_formacion_academica.documento IS 'Referencia a la tabla de documento del microservicio de documentos';
-- ddl-end --
COMMENT ON COLUMN formacion_academica.soporte_formacion_academica.descripcion IS 'Descripción del soporte asociado a la formación académica';
-- ddl-end --
COMMENT ON COLUMN formacion_academica.soporte_formacion_academica.activo IS 'Indica el estado del registro';
-- ddl-end --
COMMENT ON COLUMN formacion_academica.soporte_formacion_academica.fecha_creacion IS 'Fecha de creación del registro';
-- ddl-end --
COMMENT ON COLUMN formacion_academica.soporte_formacion_academica.fecha_modificacion IS 'Fecha de la última modificación del registro';
-- ddl-end --
COMMENT ON CONSTRAINT pk_soporte_formacion_academica ON formacion_academica.soporte_formacion_academica  IS 'Llave primaria de la tabla soporte_formacion_academica';
-- ddl-end --

-- object: fk_soporte_formacion_academica_formacion_academica | type: CONSTRAINT --
-- ALTER TABLE formacion_academica.soporte_formacion_academica DROP CONSTRAINT IF EXISTS fk_soporte_formacion_academica_formacion_academica CASCADE;
ALTER TABLE formacion_academica.soporte_formacion_academica ADD CONSTRAINT fk_soporte_formacion_academica_formacion_academica FOREIGN KEY (formacion_academica)
REFERENCES formacion_academica.formacion_academica (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_dato_adicional_formacion_academica_formacion_academica | type: CONSTRAINT --
-- ALTER TABLE formacion_academica.dato_adicional_formacion_academica DROP CONSTRAINT IF EXISTS fk_dato_adicional_formacion_academica_formacion_academica CASCADE;
ALTER TABLE formacion_academica.dato_adicional_formacion_academica ADD CONSTRAINT fk_dato_adicional_formacion_academica_formacion_academica FOREIGN KEY (formacion_academica)
REFERENCES formacion_academica.formacion_academica (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

GRANT USAGE ON SCHEMA acta_recibido TO desarrollooas;
GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA acta_recibido TO desarrollooas;
GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA acta_recibido TO desarrollooas;
