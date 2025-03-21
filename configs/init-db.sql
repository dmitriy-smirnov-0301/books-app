CREATE DATABASE books_db;
CREATE USER book_admin WITH PASSWORD 'book_password';
ALTER DATABASE books_db OWNER TO book_admin;
GRANT USAGE ON SCHEMA public TO book_admin;
GRANT CREATE ON SCHEMA public TO book_admin;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO book_admin;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO book_admin;
GRANT ALL PRIVILEGES ON ALL FUNCTIONS IN SCHEMA public TO book_admin;
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT ALL ON TABLES TO book_admin;
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT ALL ON SEQUENCES TO book_admin;
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT ALL ON FUNCTIONS TO book_admin;

