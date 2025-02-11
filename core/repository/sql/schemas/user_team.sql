
CREATE TABLE IF NOT EXISTS `user_team` ( 
  `uuid` CHAR(36) NOT NULL  ,
  `user_uuid` CHAR(36) NOT NULL  ,
  `user_email` VARCHAR(255) NOT NULL  ,
  `team_uuid` CHAR(36) NOT NULL  ,
  `roles` JSON NOT NULL  ,
  `status` INT NOT NULL  ,
  `created_at` DATETIME NOT NULL  default CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL  default CURRENT_TIMESTAMP,
  `created_by_uuid` CHAR(36) NOT NULL  ,
  `updated_by_uuid` CHAR(36) NOT NULL  ,  
  PRIMARY KEY (`uuid`),
  INDEX `user_uuid` (`user_uuid` ASC),
  INDEX `user_email` (`user_email` ASC),
  INDEX `team_uuid` (`team_uuid` ASC),
  INDEX `status` (`status` ASC),
  INDEX `created_at` (`created_at` ASC),
  INDEX `updated_at` (`updated_at` ASC))  
ENGINE = InnoDB;
