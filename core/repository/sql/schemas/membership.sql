
CREATE TABLE IF NOT EXISTS `membership` ( 
  `uuid` CHAR(36) NOT NULL  ,
  `owner_uuid` CHAR(36) NOT NULL  ,
  `type` INT NOT NULL  ,
  `start_date` DATE NOT NULL  default '2022-02-02',
  `billing_metadata` JSON NOT NULL  ,
  `status` INT NOT NULL  ,
  `created_at` DATETIME NOT NULL  default CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL  default CURRENT_TIMESTAMP,
  `created_by_uuid` CHAR(36) NOT NULL  ,
  `updated_by_uuid` CHAR(36) NOT NULL  ,  
  PRIMARY KEY (`uuid`),
  INDEX `owner_uuid` (`owner_uuid` ASC),
  INDEX `type` (`type` ASC),
  INDEX `status` (`status` ASC),
  INDEX `created_at` (`created_at` ASC),
  INDEX `updated_at` (`updated_at` ASC))  
ENGINE = InnoDB;
