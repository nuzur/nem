
CREATE TABLE IF NOT EXISTS `extension_version` ( 
  `uuid` CHAR(36) NOT NULL  ,
  `version` INT NOT NULL  ,
  `extension_uuid` CHAR(36) NOT NULL  ,
  `display_version` VARCHAR(255)   ,
  `description` VARCHAR(255)   ,
  `repository_tag` VARCHAR(255) NOT NULL  ,
  `configuration_entity` JSON   ,
  `execution_mode` JSON NOT NULL  ,
  `review_status` INT NOT NULL  ,
  `status` INT NOT NULL  ,
  `created_at` DATETIME NOT NULL  default CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL  default CURRENT_TIMESTAMP,
  `created_by_uuid` CHAR(36) NOT NULL  ,
  `updated_by_uuid` CHAR(36) NOT NULL  ,  
  PRIMARY KEY (`uuid`),
  INDEX `version` (`version` ASC),
  INDEX `extension_uuid` (`extension_uuid` ASC),
  INDEX `status` (`status` ASC),
  INDEX `created_at` (`created_at` ASC),
  INDEX `updated_at` (`updated_at` ASC))  
ENGINE = InnoDB;
