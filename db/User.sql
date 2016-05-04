-- Table: public."User"

-- DROP TABLE public."User";

CREATE TABLE public."User"
(
  "Id" integer NOT NULL DEFAULT nextval('"User_Id_seq"'::regclass),
  "Username" character varying(100),
  "Password" character varying(255),
  "Created" date,
  CONSTRAINT "User_pkey" PRIMARY KEY ("Id"),
  CONSTRAINT username_unq UNIQUE ("Username")
)
WITH (
  OIDS=FALSE
);
ALTER TABLE public."User"
  OWNER TO postgres;

-- Index: public."Username_idx"

-- DROP INDEX public."Username_idx";

CREATE INDEX "Username_idx"
  ON public."User"
  USING btree
  ("Username" COLLATE pg_catalog."default");

