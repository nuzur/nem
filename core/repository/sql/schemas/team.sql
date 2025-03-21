
CREATE TABLE IF NOT EXISTS `team` ( 
  `uuid` CHAR(36) NOT NULL  ,
  `version` INT NOT NULL  ,
  `name` VARCHAR(255) NOT NULL  ,
  `enviorments` JSON   ,
  `review_configs` JSON   ,
  `memberships` JSON   ,
  `stores` JSON   ,
  `connections` JSON   ,
  `object_stores` JSON   ,
  `organization_uuid` CHAR(36)   ,
  `default_entity` JSON   ,
  `status` INT NOT NULL  ,
  `created_at` DATETIME NOT NULL  default CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL  default CURRENT_TIMESTAMP,
  `created_by_uuid` CHAR(36) NOT NULL  ,
  `updated_by_uuid` CHAR(36) NOT NULL  ,  
  PRIMARY KEY (`uuid`),
  INDEX `version` (`version` ASC),
  INDEX `organization_uuid` (`organization_uuid` ASC),
  INDEX `status` (`status` ASC),
  INDEX `created_at` (`created_at` ASC),
  INDEX `updated_at` (`updated_at` ASC),
  INDEX `created_by_uuid` (`created_by_uuid` ASC),
  INDEX `updated_by_uuid` (`updated_by_uuid` ASC))  
ENGINE = InnoDB;
