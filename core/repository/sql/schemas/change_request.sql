
CREATE TABLE IF NOT EXISTS `change_request` ( 
  `uuid` CHAR(36) NOT NULL  ,
  `version` INT NOT NULL  ,
  `title` VARCHAR(255) NOT NULL  ,
  `description` VARCHAR(255)   ,
  `review_type` INT NOT NULL  ,
  `data_changes` JSON   ,
  `version_changes` JSON   ,
  `reviews` JSON   ,
  `owner_uuid` CHAR(36) NOT NULL  ,
  `status` INT NOT NULL  ,
  `created_at` DATETIME NOT NULL  default CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL  default CURRENT_TIMESTAMP,
  `created_by_uuid` CHAR(36) NOT NULL  ,
  `updated_by_uuid` CHAR(36) NOT NULL  ,  
  PRIMARY KEY (`uuid`),
  INDEX `version` (`version` ASC),
  INDEX `status` (`status` ASC),
  INDEX `created_at` (`created_at` ASC),
  INDEX `updated_at` (`updated_at` ASC))  
ENGINE = InnoDB;
