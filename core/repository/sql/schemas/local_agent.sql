
CREATE TABLE IF NOT EXISTS `local_agent` ( 
  `uuid` CHAR(36) NOT NULL  ,
  `user_uuid` CHAR(36) NOT NULL  ,
  `token_hash` VARCHAR(255)   ,
  `machine_name` VARCHAR(255)   ,
  `os` VARCHAR(255)   ,
  `cli_version` VARCHAR(255)   ,
  `connections` JSON   ,
  `status` INT NOT NULL  ,
  `last_seen_at` DATETIME   default CURRENT_TIMESTAMP,
  `revoked_at` DATETIME   default CURRENT_TIMESTAMP,
  `created_at` DATETIME NOT NULL  default CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL  default CURRENT_TIMESTAMP,
  `created_by_uuid` CHAR(36) NOT NULL  ,
  `updated_by_uuid` CHAR(36) NOT NULL  ,  
  PRIMARY KEY (`uuid`),
  INDEX `status` (`status` ASC),
  INDEX `created_at` (`created_at` ASC),
  INDEX `updated_at` (`updated_at` ASC))  
ENGINE = InnoDB;
