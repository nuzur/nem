


-- name: FetchTeamByUUID :many
SELECT * FROM team
WHERE 
     uuid = ?  
    ;


-- name: FetchTeamByUUIDForUpdate :many
SELECT * FROM team
WHERE 
     uuid = ?      
FOR UPDATE;


-- name: FetchTeamByVersion :many
SELECT * FROM team
WHERE 
     version = ?  
    LIMIT ?, ?;



-- name: FetchTeamByStatus :many
SELECT * FROM team
WHERE 
     status = ?  
    LIMIT ?, ?;



-- name: FetchTeamByVersionAndStatus :many
SELECT * FROM team
WHERE 
     version = ? AND status = ?  
    LIMIT ?, ?;





    

    
    
-- name: FetchTeamByVersionOrderedByCreatedAtASC :many
SELECT * FROM team
WHERE 
     version = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchTeamByVersionOrderedByCreatedAtDESC :many
SELECT * FROM team
WHERE 
     version = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchTeamByVersionOrderedByUpdatedAtASC :many
SELECT * FROM team
WHERE 
     version = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchTeamByVersionOrderedByUpdatedAtDESC :many
SELECT * FROM team
WHERE 
     version = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchTeamByStatusOrderedByCreatedAtASC :many
SELECT * FROM team
WHERE 
     status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchTeamByStatusOrderedByCreatedAtDESC :many
SELECT * FROM team
WHERE 
     status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchTeamByStatusOrderedByUpdatedAtASC :many
SELECT * FROM team
WHERE 
     status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchTeamByStatusOrderedByUpdatedAtDESC :many
SELECT * FROM team
WHERE 
     status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchTeamByVersionAndStatusOrderedByCreatedAtASC :many
SELECT * FROM team
WHERE 
     version = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchTeamByVersionAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM team
WHERE 
     version = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchTeamByVersionAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM team
WHERE 
     version = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchTeamByVersionAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM team
WHERE 
     version = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    



-- name: SearchTeam :many
SELECT * FROM team
WHERE 
    name like ? 
    LIMIT ?, ?;








-- name: FetchOrganizationByUUID :many
SELECT * FROM organization
WHERE 
     uuid = ?  
    ;


-- name: FetchOrganizationByUUIDForUpdate :many
SELECT * FROM organization
WHERE 
     uuid = ?      
FOR UPDATE;


-- name: FetchOrganizationByVersion :many
SELECT * FROM organization
WHERE 
     version = ?  
    LIMIT ?, ?;



-- name: FetchOrganizationByStatus :many
SELECT * FROM organization
WHERE 
     status = ?  
    LIMIT ?, ?;



-- name: FetchOrganizationByVersionAndStatus :many
SELECT * FROM organization
WHERE 
     version = ? AND status = ?  
    LIMIT ?, ?;





    

    
    
-- name: FetchOrganizationByVersionOrderedByCreatedAtASC :many
SELECT * FROM organization
WHERE 
     version = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchOrganizationByVersionOrderedByCreatedAtDESC :many
SELECT * FROM organization
WHERE 
     version = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchOrganizationByVersionOrderedByUpdatedAtASC :many
SELECT * FROM organization
WHERE 
     version = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchOrganizationByVersionOrderedByUpdatedAtDESC :many
SELECT * FROM organization
WHERE 
     version = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchOrganizationByStatusOrderedByCreatedAtASC :many
SELECT * FROM organization
WHERE 
     status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchOrganizationByStatusOrderedByCreatedAtDESC :many
SELECT * FROM organization
WHERE 
     status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchOrganizationByStatusOrderedByUpdatedAtASC :many
SELECT * FROM organization
WHERE 
     status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchOrganizationByStatusOrderedByUpdatedAtDESC :many
SELECT * FROM organization
WHERE 
     status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchOrganizationByVersionAndStatusOrderedByCreatedAtASC :many
SELECT * FROM organization
WHERE 
     version = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchOrganizationByVersionAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM organization
WHERE 
     version = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchOrganizationByVersionAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM organization
WHERE 
     version = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchOrganizationByVersionAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM organization
WHERE 
     version = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    



-- name: SearchOrganization :many
SELECT * FROM organization
WHERE 
    name like ? 
    LIMIT ?, ?;








-- name: FetchProjectByUUID :many
SELECT * FROM project
WHERE 
     uuid = ?  
    ;


-- name: FetchProjectByUUIDForUpdate :many
SELECT * FROM project
WHERE 
     uuid = ?      
FOR UPDATE;


-- name: FetchProjectByVersion :many
SELECT * FROM project
WHERE 
     version = ?  
    LIMIT ?, ?;



-- name: FetchProjectByName :many
SELECT * FROM project
WHERE 
     name = ?  
    LIMIT ?, ?;



-- name: FetchProjectByVersionAndName :many
SELECT * FROM project
WHERE 
     version = ? AND name = ?  
    LIMIT ?, ?;



-- name: FetchProjectByStatus :many
SELECT * FROM project
WHERE 
     status = ?  
    LIMIT ?, ?;



-- name: FetchProjectByVersionAndStatus :many
SELECT * FROM project
WHERE 
     version = ? AND status = ?  
    LIMIT ?, ?;



-- name: FetchProjectByNameAndStatus :many
SELECT * FROM project
WHERE 
     name = ? AND status = ?  
    LIMIT ?, ?;



-- name: FetchProjectByVersionAndNameAndStatus :many
SELECT * FROM project
WHERE 
     version = ? AND name = ? AND status = ?  
    LIMIT ?, ?;





    

    
    
-- name: FetchProjectByVersionOrderedByCreatedAtASC :many
SELECT * FROM project
WHERE 
     version = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchProjectByVersionOrderedByCreatedAtDESC :many
SELECT * FROM project
WHERE 
     version = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchProjectByVersionOrderedByUpdatedAtASC :many
SELECT * FROM project
WHERE 
     version = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchProjectByVersionOrderedByUpdatedAtDESC :many
SELECT * FROM project
WHERE 
     version = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchProjectByNameOrderedByCreatedAtASC :many
SELECT * FROM project
WHERE 
     name = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchProjectByNameOrderedByCreatedAtDESC :many
SELECT * FROM project
WHERE 
     name = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchProjectByNameOrderedByUpdatedAtASC :many
SELECT * FROM project
WHERE 
     name = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchProjectByNameOrderedByUpdatedAtDESC :many
SELECT * FROM project
WHERE 
     name = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchProjectByVersionAndNameOrderedByCreatedAtASC :many
SELECT * FROM project
WHERE 
     version = ? AND name = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchProjectByVersionAndNameOrderedByCreatedAtDESC :many
SELECT * FROM project
WHERE 
     version = ? AND name = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchProjectByVersionAndNameOrderedByUpdatedAtASC :many
SELECT * FROM project
WHERE 
     version = ? AND name = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchProjectByVersionAndNameOrderedByUpdatedAtDESC :many
SELECT * FROM project
WHERE 
     version = ? AND name = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchProjectByStatusOrderedByCreatedAtASC :many
SELECT * FROM project
WHERE 
     status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchProjectByStatusOrderedByCreatedAtDESC :many
SELECT * FROM project
WHERE 
     status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchProjectByStatusOrderedByUpdatedAtASC :many
SELECT * FROM project
WHERE 
     status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchProjectByStatusOrderedByUpdatedAtDESC :many
SELECT * FROM project
WHERE 
     status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchProjectByVersionAndStatusOrderedByCreatedAtASC :many
SELECT * FROM project
WHERE 
     version = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchProjectByVersionAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM project
WHERE 
     version = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchProjectByVersionAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM project
WHERE 
     version = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchProjectByVersionAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM project
WHERE 
     version = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchProjectByNameAndStatusOrderedByCreatedAtASC :many
SELECT * FROM project
WHERE 
     name = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchProjectByNameAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM project
WHERE 
     name = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchProjectByNameAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM project
WHERE 
     name = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchProjectByNameAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM project
WHERE 
     name = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchProjectByVersionAndNameAndStatusOrderedByCreatedAtASC :many
SELECT * FROM project
WHERE 
     version = ? AND name = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchProjectByVersionAndNameAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM project
WHERE 
     version = ? AND name = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchProjectByVersionAndNameAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM project
WHERE 
     version = ? AND name = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchProjectByVersionAndNameAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM project
WHERE 
     version = ? AND name = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    



-- name: SearchProject :many
SELECT * FROM project
WHERE 
    name like ? OR
    
    description like ? 
    LIMIT ?, ?;








-- name: FetchExtensionByUUID :many
SELECT * FROM extension
WHERE 
     uuid = ?  
    ;


-- name: FetchExtensionByUUIDForUpdate :many
SELECT * FROM extension
WHERE 
     uuid = ?      
FOR UPDATE;


-- name: FetchExtensionByVersion :many
SELECT * FROM extension
WHERE 
     version = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByIdentifier :many
SELECT * FROM extension
WHERE 
     identifier = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByVersionAndIdentifier :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByRepository :many
SELECT * FROM extension
WHERE 
     repository = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByVersionAndRepository :many
SELECT * FROM extension
WHERE 
     version = ? AND repository = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByIdentifierAndRepository :many
SELECT * FROM extension
WHERE 
     identifier = ? AND repository = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByVersionAndIdentifierAndRepository :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND repository = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByExtensionType :many
SELECT * FROM extension
WHERE 
     extension_type = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByVersionAndExtensionType :many
SELECT * FROM extension
WHERE 
     version = ? AND extension_type = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByIdentifierAndExtensionType :many
SELECT * FROM extension
WHERE 
     identifier = ? AND extension_type = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByVersionAndIdentifierAndExtensionType :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND extension_type = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByRepositoryAndExtensionType :many
SELECT * FROM extension
WHERE 
     repository = ? AND extension_type = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByVersionAndRepositoryAndExtensionType :many
SELECT * FROM extension
WHERE 
     version = ? AND repository = ? AND extension_type = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByIdentifierAndRepositoryAndExtensionType :many
SELECT * FROM extension
WHERE 
     identifier = ? AND repository = ? AND extension_type = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionType :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND repository = ? AND extension_type = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByPublic :many
SELECT * FROM extension
WHERE 
     public = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByVersionAndPublic :many
SELECT * FROM extension
WHERE 
     version = ? AND public = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByIdentifierAndPublic :many
SELECT * FROM extension
WHERE 
     identifier = ? AND public = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByVersionAndIdentifierAndPublic :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND public = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByRepositoryAndPublic :many
SELECT * FROM extension
WHERE 
     repository = ? AND public = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByVersionAndRepositoryAndPublic :many
SELECT * FROM extension
WHERE 
     version = ? AND repository = ? AND public = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByIdentifierAndRepositoryAndPublic :many
SELECT * FROM extension
WHERE 
     identifier = ? AND repository = ? AND public = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByVersionAndIdentifierAndRepositoryAndPublic :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND repository = ? AND public = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByExtensionTypeAndPublic :many
SELECT * FROM extension
WHERE 
     extension_type = ? AND public = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByVersionAndExtensionTypeAndPublic :many
SELECT * FROM extension
WHERE 
     version = ? AND extension_type = ? AND public = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByIdentifierAndExtensionTypeAndPublic :many
SELECT * FROM extension
WHERE 
     identifier = ? AND extension_type = ? AND public = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByVersionAndIdentifierAndExtensionTypeAndPublic :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND extension_type = ? AND public = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByRepositoryAndExtensionTypeAndPublic :many
SELECT * FROM extension
WHERE 
     repository = ? AND extension_type = ? AND public = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublic :many
SELECT * FROM extension
WHERE 
     version = ? AND repository = ? AND extension_type = ? AND public = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndPublic :many
SELECT * FROM extension
WHERE 
     identifier = ? AND repository = ? AND extension_type = ? AND public = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublic :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND repository = ? AND extension_type = ? AND public = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByStatus :many
SELECT * FROM extension
WHERE 
     status = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByVersionAndStatus :many
SELECT * FROM extension
WHERE 
     version = ? AND status = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByIdentifierAndStatus :many
SELECT * FROM extension
WHERE 
     identifier = ? AND status = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByVersionAndIdentifierAndStatus :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND status = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByRepositoryAndStatus :many
SELECT * FROM extension
WHERE 
     repository = ? AND status = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByVersionAndRepositoryAndStatus :many
SELECT * FROM extension
WHERE 
     version = ? AND repository = ? AND status = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByIdentifierAndRepositoryAndStatus :many
SELECT * FROM extension
WHERE 
     identifier = ? AND repository = ? AND status = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByVersionAndIdentifierAndRepositoryAndStatus :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND repository = ? AND status = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByExtensionTypeAndStatus :many
SELECT * FROM extension
WHERE 
     extension_type = ? AND status = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByVersionAndExtensionTypeAndStatus :many
SELECT * FROM extension
WHERE 
     version = ? AND extension_type = ? AND status = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByIdentifierAndExtensionTypeAndStatus :many
SELECT * FROM extension
WHERE 
     identifier = ? AND extension_type = ? AND status = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByVersionAndIdentifierAndExtensionTypeAndStatus :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND extension_type = ? AND status = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByRepositoryAndExtensionTypeAndStatus :many
SELECT * FROM extension
WHERE 
     repository = ? AND extension_type = ? AND status = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByVersionAndRepositoryAndExtensionTypeAndStatus :many
SELECT * FROM extension
WHERE 
     version = ? AND repository = ? AND extension_type = ? AND status = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndStatus :many
SELECT * FROM extension
WHERE 
     identifier = ? AND repository = ? AND extension_type = ? AND status = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndStatus :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND repository = ? AND extension_type = ? AND status = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByPublicAndStatus :many
SELECT * FROM extension
WHERE 
     public = ? AND status = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByVersionAndPublicAndStatus :many
SELECT * FROM extension
WHERE 
     version = ? AND public = ? AND status = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByIdentifierAndPublicAndStatus :many
SELECT * FROM extension
WHERE 
     identifier = ? AND public = ? AND status = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByVersionAndIdentifierAndPublicAndStatus :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND public = ? AND status = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByRepositoryAndPublicAndStatus :many
SELECT * FROM extension
WHERE 
     repository = ? AND public = ? AND status = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByVersionAndRepositoryAndPublicAndStatus :many
SELECT * FROM extension
WHERE 
     version = ? AND repository = ? AND public = ? AND status = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByIdentifierAndRepositoryAndPublicAndStatus :many
SELECT * FROM extension
WHERE 
     identifier = ? AND repository = ? AND public = ? AND status = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByVersionAndIdentifierAndRepositoryAndPublicAndStatus :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND repository = ? AND public = ? AND status = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByExtensionTypeAndPublicAndStatus :many
SELECT * FROM extension
WHERE 
     extension_type = ? AND public = ? AND status = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByVersionAndExtensionTypeAndPublicAndStatus :many
SELECT * FROM extension
WHERE 
     version = ? AND extension_type = ? AND public = ? AND status = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByIdentifierAndExtensionTypeAndPublicAndStatus :many
SELECT * FROM extension
WHERE 
     identifier = ? AND extension_type = ? AND public = ? AND status = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByVersionAndIdentifierAndExtensionTypeAndPublicAndStatus :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND extension_type = ? AND public = ? AND status = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByRepositoryAndExtensionTypeAndPublicAndStatus :many
SELECT * FROM extension
WHERE 
     repository = ? AND extension_type = ? AND public = ? AND status = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublicAndStatus :many
SELECT * FROM extension
WHERE 
     version = ? AND repository = ? AND extension_type = ? AND public = ? AND status = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatus :many
SELECT * FROM extension
WHERE 
     identifier = ? AND repository = ? AND extension_type = ? AND public = ? AND status = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatus :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND repository = ? AND extension_type = ? AND public = ? AND status = ?  
    LIMIT ?, ?;





    

    
    
-- name: FetchExtensionByVersionOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByVersionOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByIdentifierOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     identifier = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByIdentifierOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     identifier = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByIdentifierOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     identifier = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByIdentifierOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     identifier = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByVersionAndIdentifierOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndIdentifierOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByVersionAndIdentifierOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndIdentifierOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByRepositoryOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     repository = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByRepositoryOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     repository = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByRepositoryOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     repository = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByRepositoryOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     repository = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByVersionAndRepositoryOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND repository = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndRepositoryOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND repository = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByVersionAndRepositoryOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND repository = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndRepositoryOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND repository = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByIdentifierAndRepositoryOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND repository = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByIdentifierAndRepositoryOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND repository = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByIdentifierAndRepositoryOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND repository = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByIdentifierAndRepositoryOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND repository = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByVersionAndIdentifierAndRepositoryOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND repository = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndIdentifierAndRepositoryOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND repository = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByVersionAndIdentifierAndRepositoryOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND repository = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndIdentifierAndRepositoryOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND repository = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByExtensionTypeOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     extension_type = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByExtensionTypeOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     extension_type = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByExtensionTypeOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     extension_type = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByExtensionTypeOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     extension_type = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByVersionAndExtensionTypeOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND extension_type = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndExtensionTypeOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND extension_type = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByVersionAndExtensionTypeOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND extension_type = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndExtensionTypeOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND extension_type = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByIdentifierAndExtensionTypeOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND extension_type = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByIdentifierAndExtensionTypeOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND extension_type = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByIdentifierAndExtensionTypeOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND extension_type = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByIdentifierAndExtensionTypeOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND extension_type = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByVersionAndIdentifierAndExtensionTypeOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND extension_type = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndIdentifierAndExtensionTypeOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND extension_type = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByVersionAndIdentifierAndExtensionTypeOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND extension_type = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndIdentifierAndExtensionTypeOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND extension_type = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByRepositoryAndExtensionTypeOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     repository = ? AND extension_type = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByRepositoryAndExtensionTypeOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     repository = ? AND extension_type = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByRepositoryAndExtensionTypeOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     repository = ? AND extension_type = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByRepositoryAndExtensionTypeOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     repository = ? AND extension_type = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByVersionAndRepositoryAndExtensionTypeOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND repository = ? AND extension_type = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndRepositoryAndExtensionTypeOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND repository = ? AND extension_type = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByVersionAndRepositoryAndExtensionTypeOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND repository = ? AND extension_type = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndRepositoryAndExtensionTypeOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND repository = ? AND extension_type = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByIdentifierAndRepositoryAndExtensionTypeOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND repository = ? AND extension_type = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByIdentifierAndRepositoryAndExtensionTypeOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND repository = ? AND extension_type = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByIdentifierAndRepositoryAndExtensionTypeOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND repository = ? AND extension_type = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByIdentifierAndRepositoryAndExtensionTypeOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND repository = ? AND extension_type = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND repository = ? AND extension_type = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND repository = ? AND extension_type = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND repository = ? AND extension_type = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND repository = ? AND extension_type = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByPublicOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     public = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByPublicOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     public = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByPublicOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     public = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByPublicOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     public = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByVersionAndPublicOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND public = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndPublicOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND public = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByVersionAndPublicOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND public = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndPublicOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND public = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByIdentifierAndPublicOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND public = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByIdentifierAndPublicOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND public = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByIdentifierAndPublicOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND public = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByIdentifierAndPublicOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND public = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByVersionAndIdentifierAndPublicOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND public = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndIdentifierAndPublicOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND public = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByVersionAndIdentifierAndPublicOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND public = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndIdentifierAndPublicOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND public = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByRepositoryAndPublicOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     repository = ? AND public = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByRepositoryAndPublicOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     repository = ? AND public = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByRepositoryAndPublicOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     repository = ? AND public = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByRepositoryAndPublicOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     repository = ? AND public = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByVersionAndRepositoryAndPublicOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND repository = ? AND public = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndRepositoryAndPublicOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND repository = ? AND public = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByVersionAndRepositoryAndPublicOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND repository = ? AND public = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndRepositoryAndPublicOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND repository = ? AND public = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByIdentifierAndRepositoryAndPublicOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND repository = ? AND public = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByIdentifierAndRepositoryAndPublicOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND repository = ? AND public = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByIdentifierAndRepositoryAndPublicOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND repository = ? AND public = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByIdentifierAndRepositoryAndPublicOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND repository = ? AND public = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByVersionAndIdentifierAndRepositoryAndPublicOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND repository = ? AND public = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndIdentifierAndRepositoryAndPublicOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND repository = ? AND public = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByVersionAndIdentifierAndRepositoryAndPublicOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND repository = ? AND public = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndIdentifierAndRepositoryAndPublicOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND repository = ? AND public = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByExtensionTypeAndPublicOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     extension_type = ? AND public = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByExtensionTypeAndPublicOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     extension_type = ? AND public = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByExtensionTypeAndPublicOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     extension_type = ? AND public = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByExtensionTypeAndPublicOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     extension_type = ? AND public = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByVersionAndExtensionTypeAndPublicOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND extension_type = ? AND public = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndExtensionTypeAndPublicOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND extension_type = ? AND public = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByVersionAndExtensionTypeAndPublicOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND extension_type = ? AND public = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndExtensionTypeAndPublicOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND extension_type = ? AND public = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByIdentifierAndExtensionTypeAndPublicOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND extension_type = ? AND public = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByIdentifierAndExtensionTypeAndPublicOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND extension_type = ? AND public = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByIdentifierAndExtensionTypeAndPublicOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND extension_type = ? AND public = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByIdentifierAndExtensionTypeAndPublicOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND extension_type = ? AND public = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByVersionAndIdentifierAndExtensionTypeAndPublicOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND extension_type = ? AND public = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndIdentifierAndExtensionTypeAndPublicOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND extension_type = ? AND public = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByVersionAndIdentifierAndExtensionTypeAndPublicOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND extension_type = ? AND public = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndIdentifierAndExtensionTypeAndPublicOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND extension_type = ? AND public = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByRepositoryAndExtensionTypeAndPublicOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     repository = ? AND extension_type = ? AND public = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByRepositoryAndExtensionTypeAndPublicOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     repository = ? AND extension_type = ? AND public = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByRepositoryAndExtensionTypeAndPublicOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     repository = ? AND extension_type = ? AND public = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByRepositoryAndExtensionTypeAndPublicOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     repository = ? AND extension_type = ? AND public = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublicOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND repository = ? AND extension_type = ? AND public = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublicOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND repository = ? AND extension_type = ? AND public = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublicOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND repository = ? AND extension_type = ? AND public = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublicOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND repository = ? AND extension_type = ? AND public = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndPublicOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND repository = ? AND extension_type = ? AND public = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndPublicOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND repository = ? AND extension_type = ? AND public = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndPublicOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND repository = ? AND extension_type = ? AND public = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndPublicOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND repository = ? AND extension_type = ? AND public = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublicOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND repository = ? AND extension_type = ? AND public = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublicOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND repository = ? AND extension_type = ? AND public = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublicOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND repository = ? AND extension_type = ? AND public = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublicOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND repository = ? AND extension_type = ? AND public = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByStatusOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByStatusOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByStatusOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByStatusOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByVersionAndStatusOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByVersionAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByIdentifierAndStatusOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByIdentifierAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByIdentifierAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByIdentifierAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByVersionAndIdentifierAndStatusOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndIdentifierAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByVersionAndIdentifierAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndIdentifierAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByRepositoryAndStatusOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     repository = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByRepositoryAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     repository = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByRepositoryAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     repository = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByRepositoryAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     repository = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByVersionAndRepositoryAndStatusOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND repository = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndRepositoryAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND repository = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByVersionAndRepositoryAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND repository = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndRepositoryAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND repository = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByIdentifierAndRepositoryAndStatusOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND repository = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByIdentifierAndRepositoryAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND repository = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByIdentifierAndRepositoryAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND repository = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByIdentifierAndRepositoryAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND repository = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByVersionAndIdentifierAndRepositoryAndStatusOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND repository = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndIdentifierAndRepositoryAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND repository = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByVersionAndIdentifierAndRepositoryAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND repository = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndIdentifierAndRepositoryAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND repository = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByExtensionTypeAndStatusOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     extension_type = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByExtensionTypeAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     extension_type = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByExtensionTypeAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     extension_type = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByExtensionTypeAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     extension_type = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByVersionAndExtensionTypeAndStatusOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND extension_type = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndExtensionTypeAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND extension_type = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByVersionAndExtensionTypeAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND extension_type = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndExtensionTypeAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND extension_type = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByIdentifierAndExtensionTypeAndStatusOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND extension_type = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByIdentifierAndExtensionTypeAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND extension_type = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByIdentifierAndExtensionTypeAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND extension_type = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByIdentifierAndExtensionTypeAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND extension_type = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByVersionAndIdentifierAndExtensionTypeAndStatusOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND extension_type = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndIdentifierAndExtensionTypeAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND extension_type = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByVersionAndIdentifierAndExtensionTypeAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND extension_type = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndIdentifierAndExtensionTypeAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND extension_type = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByRepositoryAndExtensionTypeAndStatusOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     repository = ? AND extension_type = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByRepositoryAndExtensionTypeAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     repository = ? AND extension_type = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByRepositoryAndExtensionTypeAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     repository = ? AND extension_type = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByRepositoryAndExtensionTypeAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     repository = ? AND extension_type = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByVersionAndRepositoryAndExtensionTypeAndStatusOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND repository = ? AND extension_type = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndRepositoryAndExtensionTypeAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND repository = ? AND extension_type = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByVersionAndRepositoryAndExtensionTypeAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND repository = ? AND extension_type = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndRepositoryAndExtensionTypeAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND repository = ? AND extension_type = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndStatusOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND repository = ? AND extension_type = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND repository = ? AND extension_type = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND repository = ? AND extension_type = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND repository = ? AND extension_type = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndStatusOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND repository = ? AND extension_type = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND repository = ? AND extension_type = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND repository = ? AND extension_type = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND repository = ? AND extension_type = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByPublicAndStatusOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     public = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByPublicAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     public = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByPublicAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     public = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByPublicAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     public = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByVersionAndPublicAndStatusOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND public = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndPublicAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND public = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByVersionAndPublicAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND public = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndPublicAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND public = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByIdentifierAndPublicAndStatusOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND public = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByIdentifierAndPublicAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND public = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByIdentifierAndPublicAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND public = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByIdentifierAndPublicAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND public = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByVersionAndIdentifierAndPublicAndStatusOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND public = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndIdentifierAndPublicAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND public = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByVersionAndIdentifierAndPublicAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND public = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndIdentifierAndPublicAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND public = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByRepositoryAndPublicAndStatusOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     repository = ? AND public = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByRepositoryAndPublicAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     repository = ? AND public = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByRepositoryAndPublicAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     repository = ? AND public = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByRepositoryAndPublicAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     repository = ? AND public = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByVersionAndRepositoryAndPublicAndStatusOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND repository = ? AND public = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndRepositoryAndPublicAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND repository = ? AND public = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByVersionAndRepositoryAndPublicAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND repository = ? AND public = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndRepositoryAndPublicAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND repository = ? AND public = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByIdentifierAndRepositoryAndPublicAndStatusOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND repository = ? AND public = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByIdentifierAndRepositoryAndPublicAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND repository = ? AND public = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByIdentifierAndRepositoryAndPublicAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND repository = ? AND public = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByIdentifierAndRepositoryAndPublicAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND repository = ? AND public = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByVersionAndIdentifierAndRepositoryAndPublicAndStatusOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND repository = ? AND public = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndIdentifierAndRepositoryAndPublicAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND repository = ? AND public = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByVersionAndIdentifierAndRepositoryAndPublicAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND repository = ? AND public = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndIdentifierAndRepositoryAndPublicAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND repository = ? AND public = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByExtensionTypeAndPublicAndStatusOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     extension_type = ? AND public = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByExtensionTypeAndPublicAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     extension_type = ? AND public = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByExtensionTypeAndPublicAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     extension_type = ? AND public = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByExtensionTypeAndPublicAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     extension_type = ? AND public = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByVersionAndExtensionTypeAndPublicAndStatusOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND extension_type = ? AND public = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndExtensionTypeAndPublicAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND extension_type = ? AND public = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByVersionAndExtensionTypeAndPublicAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND extension_type = ? AND public = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndExtensionTypeAndPublicAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND extension_type = ? AND public = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByIdentifierAndExtensionTypeAndPublicAndStatusOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND extension_type = ? AND public = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByIdentifierAndExtensionTypeAndPublicAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND extension_type = ? AND public = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByIdentifierAndExtensionTypeAndPublicAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND extension_type = ? AND public = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByIdentifierAndExtensionTypeAndPublicAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND extension_type = ? AND public = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByVersionAndIdentifierAndExtensionTypeAndPublicAndStatusOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND extension_type = ? AND public = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndIdentifierAndExtensionTypeAndPublicAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND extension_type = ? AND public = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByVersionAndIdentifierAndExtensionTypeAndPublicAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND extension_type = ? AND public = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndIdentifierAndExtensionTypeAndPublicAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND extension_type = ? AND public = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByRepositoryAndExtensionTypeAndPublicAndStatusOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     repository = ? AND extension_type = ? AND public = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByRepositoryAndExtensionTypeAndPublicAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     repository = ? AND extension_type = ? AND public = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByRepositoryAndExtensionTypeAndPublicAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     repository = ? AND extension_type = ? AND public = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByRepositoryAndExtensionTypeAndPublicAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     repository = ? AND extension_type = ? AND public = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublicAndStatusOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND repository = ? AND extension_type = ? AND public = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublicAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND repository = ? AND extension_type = ? AND public = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublicAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND repository = ? AND extension_type = ? AND public = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublicAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND repository = ? AND extension_type = ? AND public = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatusOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND repository = ? AND extension_type = ? AND public = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND repository = ? AND extension_type = ? AND public = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND repository = ? AND extension_type = ? AND public = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND repository = ? AND extension_type = ? AND public = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatusOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND repository = ? AND extension_type = ? AND public = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND repository = ? AND extension_type = ? AND public = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND repository = ? AND extension_type = ? AND public = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND repository = ? AND extension_type = ? AND public = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    



-- name: SearchExtension :many
SELECT * FROM extension
WHERE 
    identifier like ? OR
    
    description like ? 
    LIMIT ?, ?;








-- name: FetchExtensionVersionByUUID :many
SELECT * FROM extension_version
WHERE 
     uuid = ?  
    ;


-- name: FetchExtensionVersionByUUIDForUpdate :many
SELECT * FROM extension_version
WHERE 
     uuid = ?      
FOR UPDATE;


-- name: FetchExtensionVersionByVersion :many
SELECT * FROM extension_version
WHERE 
     version = ?  
    LIMIT ?, ?;



-- name: FetchExtensionVersionByStatus :many
SELECT * FROM extension_version
WHERE 
     status = ?  
    LIMIT ?, ?;



-- name: FetchExtensionVersionByVersionAndStatus :many
SELECT * FROM extension_version
WHERE 
     version = ? AND status = ?  
    LIMIT ?, ?;





    

    
    
-- name: FetchExtensionVersionByVersionOrderedByCreatedAtASC :many
SELECT * FROM extension_version
WHERE 
     version = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionVersionByVersionOrderedByCreatedAtDESC :many
SELECT * FROM extension_version
WHERE 
     version = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionVersionByVersionOrderedByUpdatedAtASC :many
SELECT * FROM extension_version
WHERE 
     version = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionVersionByVersionOrderedByUpdatedAtDESC :many
SELECT * FROM extension_version
WHERE 
     version = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionVersionByStatusOrderedByCreatedAtASC :many
SELECT * FROM extension_version
WHERE 
     status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionVersionByStatusOrderedByCreatedAtDESC :many
SELECT * FROM extension_version
WHERE 
     status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionVersionByStatusOrderedByUpdatedAtASC :many
SELECT * FROM extension_version
WHERE 
     status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionVersionByStatusOrderedByUpdatedAtDESC :many
SELECT * FROM extension_version
WHERE 
     status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionVersionByVersionAndStatusOrderedByCreatedAtASC :many
SELECT * FROM extension_version
WHERE 
     version = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionVersionByVersionAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM extension_version
WHERE 
     version = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionVersionByVersionAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM extension_version
WHERE 
     version = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionVersionByVersionAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM extension_version
WHERE 
     version = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    










-- name: FetchUserByUUID :many
SELECT * FROM user
WHERE 
     uuid = ?  
    ;


-- name: FetchUserByUUIDForUpdate :many
SELECT * FROM user
WHERE 
     uuid = ?      
FOR UPDATE;


-- name: FetchUserByIdentifier :many
SELECT * FROM user
WHERE 
     identifier = ?  
    LIMIT ?, ?;



-- name: FetchUserByEmail :many
SELECT * FROM user
WHERE 
     email = ?  
    LIMIT ?, ?;



-- name: FetchUserByIdentifierAndEmail :many
SELECT * FROM user
WHERE 
     identifier = ? AND email = ?  
    LIMIT ?, ?;



-- name: FetchUserByStatus :many
SELECT * FROM user
WHERE 
     status = ?  
    LIMIT ?, ?;



-- name: FetchUserByIdentifierAndStatus :many
SELECT * FROM user
WHERE 
     identifier = ? AND status = ?  
    LIMIT ?, ?;



-- name: FetchUserByEmailAndStatus :many
SELECT * FROM user
WHERE 
     email = ? AND status = ?  
    LIMIT ?, ?;



-- name: FetchUserByIdentifierAndEmailAndStatus :many
SELECT * FROM user
WHERE 
     identifier = ? AND email = ? AND status = ?  
    LIMIT ?, ?;





    

    
    
-- name: FetchUserByIdentifierOrderedByCreatedAtASC :many
SELECT * FROM user
WHERE 
     identifier = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchUserByIdentifierOrderedByCreatedAtDESC :many
SELECT * FROM user
WHERE 
     identifier = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchUserByIdentifierOrderedByUpdatedAtASC :many
SELECT * FROM user
WHERE 
     identifier = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchUserByIdentifierOrderedByUpdatedAtDESC :many
SELECT * FROM user
WHERE 
     identifier = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchUserByEmailOrderedByCreatedAtASC :many
SELECT * FROM user
WHERE 
     email = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchUserByEmailOrderedByCreatedAtDESC :many
SELECT * FROM user
WHERE 
     email = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchUserByEmailOrderedByUpdatedAtASC :many
SELECT * FROM user
WHERE 
     email = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchUserByEmailOrderedByUpdatedAtDESC :many
SELECT * FROM user
WHERE 
     email = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchUserByIdentifierAndEmailOrderedByCreatedAtASC :many
SELECT * FROM user
WHERE 
     identifier = ? AND email = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchUserByIdentifierAndEmailOrderedByCreatedAtDESC :many
SELECT * FROM user
WHERE 
     identifier = ? AND email = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchUserByIdentifierAndEmailOrderedByUpdatedAtASC :many
SELECT * FROM user
WHERE 
     identifier = ? AND email = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchUserByIdentifierAndEmailOrderedByUpdatedAtDESC :many
SELECT * FROM user
WHERE 
     identifier = ? AND email = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchUserByStatusOrderedByCreatedAtASC :many
SELECT * FROM user
WHERE 
     status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchUserByStatusOrderedByCreatedAtDESC :many
SELECT * FROM user
WHERE 
     status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchUserByStatusOrderedByUpdatedAtASC :many
SELECT * FROM user
WHERE 
     status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchUserByStatusOrderedByUpdatedAtDESC :many
SELECT * FROM user
WHERE 
     status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchUserByIdentifierAndStatusOrderedByCreatedAtASC :many
SELECT * FROM user
WHERE 
     identifier = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchUserByIdentifierAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM user
WHERE 
     identifier = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchUserByIdentifierAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM user
WHERE 
     identifier = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchUserByIdentifierAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM user
WHERE 
     identifier = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchUserByEmailAndStatusOrderedByCreatedAtASC :many
SELECT * FROM user
WHERE 
     email = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchUserByEmailAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM user
WHERE 
     email = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchUserByEmailAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM user
WHERE 
     email = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchUserByEmailAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM user
WHERE 
     email = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchUserByIdentifierAndEmailAndStatusOrderedByCreatedAtASC :many
SELECT * FROM user
WHERE 
     identifier = ? AND email = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchUserByIdentifierAndEmailAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM user
WHERE 
     identifier = ? AND email = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchUserByIdentifierAndEmailAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM user
WHERE 
     identifier = ? AND email = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchUserByIdentifierAndEmailAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM user
WHERE 
     identifier = ? AND email = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    










-- name: FetchChangeRequestByUUID :many
SELECT * FROM change_request
WHERE 
     uuid = ?  
    ;


-- name: FetchChangeRequestByUUIDForUpdate :many
SELECT * FROM change_request
WHERE 
     uuid = ?      
FOR UPDATE;


-- name: FetchChangeRequestByVersion :many
SELECT * FROM change_request
WHERE 
     version = ?  
    LIMIT ?, ?;



-- name: FetchChangeRequestByStatus :many
SELECT * FROM change_request
WHERE 
     status = ?  
    LIMIT ?, ?;



-- name: FetchChangeRequestByVersionAndStatus :many
SELECT * FROM change_request
WHERE 
     version = ? AND status = ?  
    LIMIT ?, ?;





    

    
    
-- name: FetchChangeRequestByVersionOrderedByCreatedAtASC :many
SELECT * FROM change_request
WHERE 
     version = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchChangeRequestByVersionOrderedByCreatedAtDESC :many
SELECT * FROM change_request
WHERE 
     version = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchChangeRequestByVersionOrderedByUpdatedAtASC :many
SELECT * FROM change_request
WHERE 
     version = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchChangeRequestByVersionOrderedByUpdatedAtDESC :many
SELECT * FROM change_request
WHERE 
     version = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchChangeRequestByStatusOrderedByCreatedAtASC :many
SELECT * FROM change_request
WHERE 
     status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchChangeRequestByStatusOrderedByCreatedAtDESC :many
SELECT * FROM change_request
WHERE 
     status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchChangeRequestByStatusOrderedByUpdatedAtASC :many
SELECT * FROM change_request
WHERE 
     status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchChangeRequestByStatusOrderedByUpdatedAtDESC :many
SELECT * FROM change_request
WHERE 
     status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchChangeRequestByVersionAndStatusOrderedByCreatedAtASC :many
SELECT * FROM change_request
WHERE 
     version = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchChangeRequestByVersionAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM change_request
WHERE 
     version = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchChangeRequestByVersionAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM change_request
WHERE 
     version = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchChangeRequestByVersionAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM change_request
WHERE 
     version = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    



-- name: SearchChangeRequest :many
SELECT * FROM change_request
WHERE 
    title like ? OR
    
    description like ? 
    LIMIT ?, ?;








-- name: FetchProjectVersionByUUID :many
SELECT * FROM project_version
WHERE 
     uuid = ?  
    ;


-- name: FetchProjectVersionByUUIDForUpdate :many
SELECT * FROM project_version
WHERE 
     uuid = ?      
FOR UPDATE;


-- name: FetchProjectVersionByVersion :many
SELECT * FROM project_version
WHERE 
     version = ?  
    LIMIT ?, ?;



-- name: FetchProjectVersionByReviewStatus :many
SELECT * FROM project_version
WHERE 
     review_status = ?  
    LIMIT ?, ?;



-- name: FetchProjectVersionByVersionAndReviewStatus :many
SELECT * FROM project_version
WHERE 
     version = ? AND review_status = ?  
    LIMIT ?, ?;



-- name: FetchProjectVersionByStatus :many
SELECT * FROM project_version
WHERE 
     status = ?  
    LIMIT ?, ?;



-- name: FetchProjectVersionByVersionAndStatus :many
SELECT * FROM project_version
WHERE 
     version = ? AND status = ?  
    LIMIT ?, ?;



-- name: FetchProjectVersionByReviewStatusAndStatus :many
SELECT * FROM project_version
WHERE 
     review_status = ? AND status = ?  
    LIMIT ?, ?;



-- name: FetchProjectVersionByVersionAndReviewStatusAndStatus :many
SELECT * FROM project_version
WHERE 
     version = ? AND review_status = ? AND status = ?  
    LIMIT ?, ?;





    

    
    
-- name: FetchProjectVersionByVersionOrderedByCreatedAtASC :many
SELECT * FROM project_version
WHERE 
     version = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchProjectVersionByVersionOrderedByCreatedAtDESC :many
SELECT * FROM project_version
WHERE 
     version = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchProjectVersionByVersionOrderedByUpdatedAtASC :many
SELECT * FROM project_version
WHERE 
     version = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchProjectVersionByVersionOrderedByUpdatedAtDESC :many
SELECT * FROM project_version
WHERE 
     version = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchProjectVersionByReviewStatusOrderedByCreatedAtASC :many
SELECT * FROM project_version
WHERE 
     review_status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchProjectVersionByReviewStatusOrderedByCreatedAtDESC :many
SELECT * FROM project_version
WHERE 
     review_status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchProjectVersionByReviewStatusOrderedByUpdatedAtASC :many
SELECT * FROM project_version
WHERE 
     review_status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchProjectVersionByReviewStatusOrderedByUpdatedAtDESC :many
SELECT * FROM project_version
WHERE 
     review_status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchProjectVersionByVersionAndReviewStatusOrderedByCreatedAtASC :many
SELECT * FROM project_version
WHERE 
     version = ? AND review_status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchProjectVersionByVersionAndReviewStatusOrderedByCreatedAtDESC :many
SELECT * FROM project_version
WHERE 
     version = ? AND review_status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchProjectVersionByVersionAndReviewStatusOrderedByUpdatedAtASC :many
SELECT * FROM project_version
WHERE 
     version = ? AND review_status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchProjectVersionByVersionAndReviewStatusOrderedByUpdatedAtDESC :many
SELECT * FROM project_version
WHERE 
     version = ? AND review_status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchProjectVersionByStatusOrderedByCreatedAtASC :many
SELECT * FROM project_version
WHERE 
     status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchProjectVersionByStatusOrderedByCreatedAtDESC :many
SELECT * FROM project_version
WHERE 
     status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchProjectVersionByStatusOrderedByUpdatedAtASC :many
SELECT * FROM project_version
WHERE 
     status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchProjectVersionByStatusOrderedByUpdatedAtDESC :many
SELECT * FROM project_version
WHERE 
     status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchProjectVersionByVersionAndStatusOrderedByCreatedAtASC :many
SELECT * FROM project_version
WHERE 
     version = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchProjectVersionByVersionAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM project_version
WHERE 
     version = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchProjectVersionByVersionAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM project_version
WHERE 
     version = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchProjectVersionByVersionAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM project_version
WHERE 
     version = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchProjectVersionByReviewStatusAndStatusOrderedByCreatedAtASC :many
SELECT * FROM project_version
WHERE 
     review_status = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchProjectVersionByReviewStatusAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM project_version
WHERE 
     review_status = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchProjectVersionByReviewStatusAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM project_version
WHERE 
     review_status = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchProjectVersionByReviewStatusAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM project_version
WHERE 
     review_status = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchProjectVersionByVersionAndReviewStatusAndStatusOrderedByCreatedAtASC :many
SELECT * FROM project_version
WHERE 
     version = ? AND review_status = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchProjectVersionByVersionAndReviewStatusAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM project_version
WHERE 
     version = ? AND review_status = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchProjectVersionByVersionAndReviewStatusAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM project_version
WHERE 
     version = ? AND review_status = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchProjectVersionByVersionAndReviewStatusAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM project_version
WHERE 
     version = ? AND review_status = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    



-- name: SearchProjectVersion :many
SELECT * FROM project_version
WHERE 
    description like ? 
    LIMIT ?, ?;








-- name: FetchUserTeamByUUID :many
SELECT * FROM user_team
WHERE 
     uuid = ?  
    ;


-- name: FetchUserTeamByUUIDForUpdate :many
SELECT * FROM user_team
WHERE 
     uuid = ?      
FOR UPDATE;


-- name: FetchUserTeamByUserEmail :many
SELECT * FROM user_team
WHERE 
     user_email = ?  
    LIMIT ?, ?;



-- name: FetchUserTeamByStatus :many
SELECT * FROM user_team
WHERE 
     status = ?  
    LIMIT ?, ?;



-- name: FetchUserTeamByUserEmailAndStatus :many
SELECT * FROM user_team
WHERE 
     user_email = ? AND status = ?  
    LIMIT ?, ?;





    

    
    
-- name: FetchUserTeamByUserEmailOrderedByCreatedAtASC :many
SELECT * FROM user_team
WHERE 
     user_email = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchUserTeamByUserEmailOrderedByCreatedAtDESC :many
SELECT * FROM user_team
WHERE 
     user_email = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchUserTeamByUserEmailOrderedByUpdatedAtASC :many
SELECT * FROM user_team
WHERE 
     user_email = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchUserTeamByUserEmailOrderedByUpdatedAtDESC :many
SELECT * FROM user_team
WHERE 
     user_email = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchUserTeamByStatusOrderedByCreatedAtASC :many
SELECT * FROM user_team
WHERE 
     status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchUserTeamByStatusOrderedByCreatedAtDESC :many
SELECT * FROM user_team
WHERE 
     status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchUserTeamByStatusOrderedByUpdatedAtASC :many
SELECT * FROM user_team
WHERE 
     status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchUserTeamByStatusOrderedByUpdatedAtDESC :many
SELECT * FROM user_team
WHERE 
     status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchUserTeamByUserEmailAndStatusOrderedByCreatedAtASC :many
SELECT * FROM user_team
WHERE 
     user_email = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchUserTeamByUserEmailAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM user_team
WHERE 
     user_email = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchUserTeamByUserEmailAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM user_team
WHERE 
     user_email = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchUserTeamByUserEmailAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM user_team
WHERE 
     user_email = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    










-- name: FetchExtensionExecutionByUUID :many
SELECT * FROM extension_execution
WHERE 
     uuid = ?  
    ;


-- name: FetchExtensionExecutionByUUIDForUpdate :many
SELECT * FROM extension_execution
WHERE 
     uuid = ?      
FOR UPDATE;


-- name: FetchExtensionExecutionByStatus :many
SELECT * FROM extension_execution
WHERE 
     status = ?  
    LIMIT ?, ?;





    

    
    
-- name: FetchExtensionExecutionByStatusOrderedByCreatedAtASC :many
SELECT * FROM extension_execution
WHERE 
     status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionExecutionByStatusOrderedByCreatedAtDESC :many
SELECT * FROM extension_execution
WHERE 
     status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionExecutionByStatusOrderedByUpdatedAtASC :many
SELECT * FROM extension_execution
WHERE 
     status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionExecutionByStatusOrderedByUpdatedAtDESC :many
SELECT * FROM extension_execution
WHERE 
     status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    










-- name: FetchUserConnectionByUUID :many
SELECT * FROM user_connection
WHERE 
     uuid = ?  
    ;


-- name: FetchUserConnectionByUUIDForUpdate :many
SELECT * FROM user_connection
WHERE 
     uuid = ?      
FOR UPDATE;


-- name: FetchUserConnectionByStatus :many
SELECT * FROM user_connection
WHERE 
     status = ?  
    LIMIT ?, ?;





    

    
    
-- name: FetchUserConnectionByStatusOrderedByCreatedAtASC :many
SELECT * FROM user_connection
WHERE 
     status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchUserConnectionByStatusOrderedByCreatedAtDESC :many
SELECT * FROM user_connection
WHERE 
     status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchUserConnectionByStatusOrderedByUpdatedAtASC :many
SELECT * FROM user_connection
WHERE 
     status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchUserConnectionByStatusOrderedByUpdatedAtDESC :many
SELECT * FROM user_connection
WHERE 
     status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    










-- name: FetchUserProjectVersionByUUID :many
SELECT * FROM user_project_version
WHERE 
     uuid = ?  
    ;


-- name: FetchUserProjectVersionByUUIDForUpdate :many
SELECT * FROM user_project_version
WHERE 
     uuid = ?      
FOR UPDATE;


-- name: FetchUserProjectVersionByStatus :many
SELECT * FROM user_project_version
WHERE 
     status = ?  
    LIMIT ?, ?;





    

    
    
-- name: FetchUserProjectVersionByStatusOrderedByCreatedAtASC :many
SELECT * FROM user_project_version
WHERE 
     status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchUserProjectVersionByStatusOrderedByCreatedAtDESC :many
SELECT * FROM user_project_version
WHERE 
     status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchUserProjectVersionByStatusOrderedByUpdatedAtASC :many
SELECT * FROM user_project_version
WHERE 
     status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchUserProjectVersionByStatusOrderedByUpdatedAtDESC :many
SELECT * FROM user_project_version
WHERE 
     status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    







