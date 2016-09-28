CREATE SCHEMA `productcatalog_00` DEFAULT CHARACTER SET utf8 ;
CREATE SCHEMA `productcatalog_01` DEFAULT CHARACTER SET utf8 ;
CREATE SCHEMA `productcatalog_02` DEFAULT CHARACTER SET utf8 ;
CREATE SCHEMA `productcatalog_03` DEFAULT CHARACTER SET utf8 ;
CREATE SCHEMA `productcatalog_04` DEFAULT CHARACTER SET utf8 ;
CREATE SCHEMA `productcatalog_05` DEFAULT CHARACTER SET utf8 ;
CREATE SCHEMA `productcatalog_06` DEFAULT CHARACTER SET utf8 ;
CREATE SCHEMA `productcatalog_07` DEFAULT CHARACTER SET utf8 ;


CREATE TABLE `top_category` (
  `id` INT NOT NULL COMMENT '',
  `tenant_id` INT NOT NULL COMMENT '',
  `code` VARCHAR(20) NOT NULL COMMENT '',
  `name` VARCHAR(50) NOT NULL COMMENT '',
  PRIMARY KEY (`id`),
  UNIQUE INDEX `code_UNIQUE` (`tenant_id` ASC, `code` ASC),
  UNIQUE INDEX `name_UNIQUE` (`tenant_id` ASC, `name` ASC));

CREATE TABLE `sub_category` (
  `id` INT NOT NULL COMMENT '',
  `tenant_id` INT NOT NULL COMMENT '',
  `top_category_id` INT NOT NULL COMMENT '',
  `code` VARCHAR(20) NOT NULL COMMENT '',
  `name` VARCHAR(50) NOT NULL COMMENT '',
  `description` VARCHAR(200) NULL COMMENT '',
  PRIMARY KEY (`id`),
  UNIQUE INDEX `code_UNIQUE` (`tenant_id` ASC, `code` ASC),
  UNIQUE INDEX `name_UNIQUE` (`tenant_id` ASC, `name` ASC));

-- Generate tables config for saashard
/*
    -
        name : top_category

    -
        name : sub_category

*/
select concat('    -
        name : ' , `TABLE_NAME`,'
') `config` from `information_schema`.`tables` where `table_schema`='productcatalog_00' and `Table_Type`='BASE TABLE';
