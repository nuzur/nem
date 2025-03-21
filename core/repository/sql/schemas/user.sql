
CREATE TABLE IF NOT EXISTS `user` ( 
  `uuid` CHAR(36) NOT NULL  ,
  `identifier` VARCHAR(255) NOT NULL  ,
  `name` VARCHAR(255)   ,
  `last_name` VARCHAR(255)   ,
  `email` VARCHAR(255) NOT NULL  ,
  `user_type` INT NOT NULL  ,
  `country_ios2` VARCHAR(255)   ,
  `locale` VARCHAR(255)   ,
  `metadata` VARCHAR(255)   ,
  `status` INT NOT NULL  ,
  `created_at` DATETIME NOT NULL  default CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL  default CURRENT_TIMESTAMP,  
  PRIMARY KEY (`uuid`),
  INDEX `identifier` (`identifier` ASC),
  INDEX `email` (`email` ASC),
  INDEX `status` (`status` ASC),
  INDEX `created_at` (`created_at` ASC),
  INDEX `updated_at` (`updated_at` ASC))  
ENGINE = InnoDB;
