-- Database generated with pgModeler (PostgreSQL Database Modeler).
-- pgModeler version: 1.1.0-beta1
-- PostgreSQL version: 16.0
-- Project Site: pgmodeler.io
-- Model Author: Deskx
-- object: dbadmin | type: ROLE --
-- DROP ROLE IF EXISTS dbadmin;
-- CREATE ROLE dbadmin WITH 
-- 	SUPERUSER
-- 	CREATEDB
-- 	CREATEROLE
-- 	INHERIT
-- 	LOGIN
-- 	REPLICATION
-- 	BYPASSRLS
-- 	 ENCRYPTED PASSWORD '1010aa';
-- ddl-end --

-- object: agroetec | type: ROLE --
-- DROP ROLE IF EXISTS agroetec;
-- CREATE ROLE agroetec WITH 
-- 	CREATEROLE
-- 	 ENCRYPTED PASSWORD '1010aa';
-- ddl-end --


-- Database creation must be performed outside a multi lined SQL file. 
-- These commands were put in this file only as a convenience.
-- 
-- object: agroetec | type: DATABASE --
-- DROP DATABASE IF EXISTS agroetec;
-- CREATE DATABASE agroetec
-- OWNER = agroetec;
-- ddl-end --


-- -- object: apigw | type: SCHEMA --
-- -- DROP SCHEMA IF EXISTS apigw CASCADE;
-- CREATE SCHEMA apigw;
-- -- ddl-end --
-- ALTER SCHEMA apigw OWNER TO agroetec;
-- -- ddl-end --

-- -- object: business | type: SCHEMA --
-- -- DROP SCHEMA IF EXISTS business CASCADE;
-- CREATE SCHEMA business;
-- -- ddl-end --
-- ALTER SCHEMA business OWNER TO agroetec;
-- -- ddl-end --

-- -- object: iot | type: SCHEMA --
-- -- DROP SCHEMA IF EXISTS iot CASCADE;
-- CREATE SCHEMA iot;
-- -- ddl-end --
-- ALTER SCHEMA iot OWNER TO agroetec;
-- -- ddl-end --

-- SET search_path TO pg_catalog,public,apigw,business,iot;

-- COPY pessoa
-- FROM '/docker-entrypoint-initdb.d/pessoa.csv'
-- DELIMITER ','
-- CSV HEADER;