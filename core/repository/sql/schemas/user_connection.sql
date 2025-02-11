
CREATE TABLE IF NOT EXISTS `user_connection` ( 
  `uuid` CHAR(36) NOT NULL  ,
  `user_uuid` CHAR(36) NOT NULL  ,
  `project_uuid` CHAR(36) NOT NULL  ,
  `project_version_uuid` CHAR(36) NOT NULL  ,
  `type` INT NOT NULL  ,
  `type_config` JSON NOT NULL  ,
  `db_schema` VARCHAR(255) NOT NULL  ,
  `executions` JSON NOT NULL  ,
  `status` INT NOT NULL  ,
  `created_at` DATETIME NOT NULL  default CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL  default CURRENT_TIMESTAMP,  
  PRIMARY KEY (`uuid`),
  INDEX `user_uuid` (`user_uuid` ASC),
  INDEX `project_uuid` (`project_uuid` ASC),
  INDEX `project_version_uuid` (`project_version_uuid` ASC),
  INDEX `status` (`status` ASC),
  INDEX `created_at` (`created_at` ASC),
  INDEX `updated_at` (`updated_at` ASC))  
ENGINE = InnoDB;
