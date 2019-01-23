CREATE TABLE userinfo
(
	    uid serial NOT NULL,
	    username character varying(100) NOT NULL,
	    departname character varying(500) NOT NULL,
	    Created date,
	    CONSTRAINT userinfo_pkey PRIMARY KEY (uid)
)
WITH (OIDS=FALSE);
