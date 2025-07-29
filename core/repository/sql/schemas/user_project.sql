
CREATE TABLE IF NOT EXISTS `user_project` ( 
  `uuid` CHAR(36) NOT NULL  ,
  `user_uuid` CHAR(36)   ,
  `user_email` VARCHAR(255)   ,
  `project_uuid` CHAR(36) NOT NULL  ,
  `role` INT NOT NULL  ,
  `review_required_structure` TINYINT(1) NOT NULL  ,
  `review_required_data` TINYINT(1) NOT NULL  ,
  `status` INT NOT NULL  ,
  `created_at` DATETIME NOT NULL  default CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL  default CURRENT_TIMESTAMP,
  `created_by_uuid` CHAR(36) NOT NULL  ,
  `updated_by_uuid` CHAR(36) NOT NULL  ,  
  PRIMARY KEY (`uuid`),
  INDEX `user_uuid` (`user_uuid` ASC),
  INDEX `user_email` (`user_email` ASC),
  INDEX `project_uuid` (`project_uuid` ASC),
  INDEX `role` (`role` ASC),
  INDEX `status` (`status` ASC),
  INDEX `created_at` (`created_at` ASC),
  INDEX `updated_at` (`updated_at` ASC))  
ENGINE = InnoDB;
