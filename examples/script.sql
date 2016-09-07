CREATE SCHEMA `productcatalog_00` DEFAULT CHARACTER SET utf8 ;

CREATE TABLE `productcatalog_00`.`top_category` (
  `id` INT NOT NULL COMMENT '',
  `tenant_id` INT NOT NULL COMMENT '',
  `code` VARCHAR(20) NOT NULL COMMENT '',
  `name` VARCHAR(50) NOT NULL COMMENT '',
  PRIMARY KEY (`id`)  COMMENT '',
  UNIQUE INDEX `code_UNIQUE` (`tenant_id` ASC, `code` ASC)  COMMENT '',
  UNIQUE INDEX `name_UNIQUE` (`tenant_id` ASC, `name` ASC)  COMMENT '');

CREATE TABLE `productcatalog_00`.`sub_category` (
  `id` INT NOT NULL COMMENT '',
  `tenant_id` INT NOT NULL COMMENT '',
  `top_category_id` INT NOT NULL COMMENT '',
  `code` VARCHAR(20) NOT NULL COMMENT '',
  `name` VARCHAR(50) NOT NULL COMMENT '',
  `description` VARCHAR(200) NULL COMMENT '',
  PRIMARY KEY (`id`)  COMMENT '',
  UNIQUE INDEX `code_UNIQUE` (`tenant_id` ASC, `code` ASC)  COMMENT '',
  UNIQUE INDEX `name_UNIQUE` (`tenant_id` ASC, `name` ASC)  COMMENT '');

CREATE SCHEMA `productcatalog_01` DEFAULT CHARACTER SET utf8 ;

CREATE TABLE `productcatalog_01`.`top_category` (
  `id` INT NOT NULL COMMENT '',
  `tenant_id` INT NOT NULL COMMENT '',
  `code` VARCHAR(20) NOT NULL COMMENT '',
  `name` VARCHAR(50) NOT NULL COMMENT '',
  PRIMARY KEY (`id`)  COMMENT '',
  UNIQUE INDEX `code_UNIQUE` (`tenant_id` ASC, `code` ASC)  COMMENT '',
  UNIQUE INDEX `name_UNIQUE` (`tenant_id` ASC, `name` ASC)  COMMENT '');

CREATE TABLE `productcatalog_01`.`sub_category` (
  `id` INT NOT NULL COMMENT '',
  `tenant_id` INT NOT NULL COMMENT '',
  `top_category_id` INT NOT NULL COMMENT '',
  `code` VARCHAR(20) NOT NULL COMMENT '',
  `name` VARCHAR(50) NOT NULL COMMENT '',
  `description` VARCHAR(200) NULL COMMENT '',
  PRIMARY KEY (`id`)  COMMENT '',
  UNIQUE INDEX `code_UNIQUE` (`tenant_id` ASC, `code` ASC)  COMMENT '',
  UNIQUE INDEX `name_UNIQUE` (`tenant_id` ASC, `name` ASC)  COMMENT '');
