
CREATE TABLE IF NOT EXISTS `extension_execution` ( 
  `uuid` CHAR(36) NOT NULL  ,
  `extension_uuid` CHAR(36) NOT NULL  ,
  `extension_version_uuid` CHAR(36) NOT NULL  ,
  `project_extension_uuid` CHAR(36)   ,
  `project_uuid` CHAR(36) NOT NULL  ,
  `project_version_uuid` CHAR(36) NOT NULL  ,
  `executed_by_uuid` CHAR(36) NOT NULL  ,
  `metadata` JSON   ,
  `status` INT NOT NULL  ,
  `status_msg` VARCHAR(255)   ,
  `created_at` DATETIME NOT NULL  default CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL  default CURRENT_TIMESTAMP,  
  PRIMARY KEY (`uuid`),
  INDEX `extension_uuid` (`extension_uuid` ASC),
  INDEX `extension_version_uuid` (`extension_version_uuid` ASC),
  INDEX `project_uuid` (`project_uuid` ASC),
  INDEX `project_version_uuid` (`project_version_uuid` ASC),
  INDEX `executed_by_uuid` (`executed_by_uuid` ASC),
  INDEX `status` (`status` ASC),
  INDEX `created_at` (`created_at` ASC),
  INDEX `updated_at` (`updated_at` ASC))  
ENGINE = InnoDB;
