
CREATE TABLE IF NOT EXISTS `user_project_version` ( 
  `uuid` CHAR(36) NOT NULL  ,
  `version` INT NOT NULL  ,
  `project_version_uuid` CHAR(36) NOT NULL  ,
  `user_uuid` CHAR(36) NOT NULL  ,
  `data` JSON   ,
  `status` INT NOT NULL  ,
  `created_at` DATETIME NOT NULL  default CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL  default CURRENT_TIMESTAMP,
  `created_by_uuid` CHAR(36) NOT NULL  ,
  `updated_by_uuid` CHAR(36) NOT NULL  ,  
  PRIMARY KEY (`uuid`),
  INDEX `project_version_uuid` (`project_version_uuid` ASC),
  INDEX `user_uuid` (`user_uuid` ASC),
  INDEX `status` (`status` ASC),
  INDEX `created_at` (`created_at` ASC),
  INDEX `updated_at` (`updated_at` ASC))  
ENGINE = InnoDB;
