-- MySQL
CREATE TABLE `examples_echo` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `content` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='examples_echo table';


-- PostgresSQL
CREATE TABLE examples_echo (
	id serial NOT NULL,
	content varchar(233) NULL,
	CONSTRAINT examples_echo_pkey PRIMARY KEY (id)
);