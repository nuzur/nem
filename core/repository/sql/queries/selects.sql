


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



-- name: FetchExtensionByPro :many
SELECT * FROM extension
WHERE 
     pro = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByVersionAndPro :many
SELECT * FROM extension
WHERE 
     version = ? AND pro = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByIdentifierAndPro :many
SELECT * FROM extension
WHERE 
     identifier = ? AND pro = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByVersionAndIdentifierAndPro :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND pro = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByRepositoryAndPro :many
SELECT * FROM extension
WHERE 
     repository = ? AND pro = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByVersionAndRepositoryAndPro :many
SELECT * FROM extension
WHERE 
     version = ? AND repository = ? AND pro = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByIdentifierAndRepositoryAndPro :many
SELECT * FROM extension
WHERE 
     identifier = ? AND repository = ? AND pro = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByVersionAndIdentifierAndRepositoryAndPro :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND repository = ? AND pro = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByExtensionTypeAndPro :many
SELECT * FROM extension
WHERE 
     extension_type = ? AND pro = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByVersionAndExtensionTypeAndPro :many
SELECT * FROM extension
WHERE 
     version = ? AND extension_type = ? AND pro = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByIdentifierAndExtensionTypeAndPro :many
SELECT * FROM extension
WHERE 
     identifier = ? AND extension_type = ? AND pro = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByVersionAndIdentifierAndExtensionTypeAndPro :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND extension_type = ? AND pro = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByRepositoryAndExtensionTypeAndPro :many
SELECT * FROM extension
WHERE 
     repository = ? AND extension_type = ? AND pro = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPro :many
SELECT * FROM extension
WHERE 
     version = ? AND repository = ? AND extension_type = ? AND pro = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndPro :many
SELECT * FROM extension
WHERE 
     identifier = ? AND repository = ? AND extension_type = ? AND pro = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPro :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND repository = ? AND extension_type = ? AND pro = ?  
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



-- name: FetchExtensionByProAndPublic :many
SELECT * FROM extension
WHERE 
     pro = ? AND public = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByVersionAndProAndPublic :many
SELECT * FROM extension
WHERE 
     version = ? AND pro = ? AND public = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByIdentifierAndProAndPublic :many
SELECT * FROM extension
WHERE 
     identifier = ? AND pro = ? AND public = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByVersionAndIdentifierAndProAndPublic :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND pro = ? AND public = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByRepositoryAndProAndPublic :many
SELECT * FROM extension
WHERE 
     repository = ? AND pro = ? AND public = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByVersionAndRepositoryAndProAndPublic :many
SELECT * FROM extension
WHERE 
     version = ? AND repository = ? AND pro = ? AND public = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByIdentifierAndRepositoryAndProAndPublic :many
SELECT * FROM extension
WHERE 
     identifier = ? AND repository = ? AND pro = ? AND public = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByVersionAndIdentifierAndRepositoryAndProAndPublic :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND repository = ? AND pro = ? AND public = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByExtensionTypeAndProAndPublic :many
SELECT * FROM extension
WHERE 
     extension_type = ? AND pro = ? AND public = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByVersionAndExtensionTypeAndProAndPublic :many
SELECT * FROM extension
WHERE 
     version = ? AND extension_type = ? AND pro = ? AND public = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByIdentifierAndExtensionTypeAndProAndPublic :many
SELECT * FROM extension
WHERE 
     identifier = ? AND extension_type = ? AND pro = ? AND public = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByVersionAndIdentifierAndExtensionTypeAndProAndPublic :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND extension_type = ? AND pro = ? AND public = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByRepositoryAndExtensionTypeAndProAndPublic :many
SELECT * FROM extension
WHERE 
     repository = ? AND extension_type = ? AND pro = ? AND public = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByVersionAndRepositoryAndExtensionTypeAndProAndPublic :many
SELECT * FROM extension
WHERE 
     version = ? AND repository = ? AND extension_type = ? AND pro = ? AND public = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndProAndPublic :many
SELECT * FROM extension
WHERE 
     identifier = ? AND repository = ? AND extension_type = ? AND pro = ? AND public = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndProAndPublic :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND repository = ? AND extension_type = ? AND pro = ? AND public = ?  
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



-- name: FetchExtensionByProAndStatus :many
SELECT * FROM extension
WHERE 
     pro = ? AND status = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByVersionAndProAndStatus :many
SELECT * FROM extension
WHERE 
     version = ? AND pro = ? AND status = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByIdentifierAndProAndStatus :many
SELECT * FROM extension
WHERE 
     identifier = ? AND pro = ? AND status = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByVersionAndIdentifierAndProAndStatus :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND pro = ? AND status = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByRepositoryAndProAndStatus :many
SELECT * FROM extension
WHERE 
     repository = ? AND pro = ? AND status = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByVersionAndRepositoryAndProAndStatus :many
SELECT * FROM extension
WHERE 
     version = ? AND repository = ? AND pro = ? AND status = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByIdentifierAndRepositoryAndProAndStatus :many
SELECT * FROM extension
WHERE 
     identifier = ? AND repository = ? AND pro = ? AND status = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByVersionAndIdentifierAndRepositoryAndProAndStatus :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND repository = ? AND pro = ? AND status = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByExtensionTypeAndProAndStatus :many
SELECT * FROM extension
WHERE 
     extension_type = ? AND pro = ? AND status = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByVersionAndExtensionTypeAndProAndStatus :many
SELECT * FROM extension
WHERE 
     version = ? AND extension_type = ? AND pro = ? AND status = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByIdentifierAndExtensionTypeAndProAndStatus :many
SELECT * FROM extension
WHERE 
     identifier = ? AND extension_type = ? AND pro = ? AND status = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByVersionAndIdentifierAndExtensionTypeAndProAndStatus :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND extension_type = ? AND pro = ? AND status = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByRepositoryAndExtensionTypeAndProAndStatus :many
SELECT * FROM extension
WHERE 
     repository = ? AND extension_type = ? AND pro = ? AND status = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByVersionAndRepositoryAndExtensionTypeAndProAndStatus :many
SELECT * FROM extension
WHERE 
     version = ? AND repository = ? AND extension_type = ? AND pro = ? AND status = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndProAndStatus :many
SELECT * FROM extension
WHERE 
     identifier = ? AND repository = ? AND extension_type = ? AND pro = ? AND status = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndProAndStatus :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND repository = ? AND extension_type = ? AND pro = ? AND status = ?  
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



-- name: FetchExtensionByProAndPublicAndStatus :many
SELECT * FROM extension
WHERE 
     pro = ? AND public = ? AND status = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByVersionAndProAndPublicAndStatus :many
SELECT * FROM extension
WHERE 
     version = ? AND pro = ? AND public = ? AND status = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByIdentifierAndProAndPublicAndStatus :many
SELECT * FROM extension
WHERE 
     identifier = ? AND pro = ? AND public = ? AND status = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByVersionAndIdentifierAndProAndPublicAndStatus :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND pro = ? AND public = ? AND status = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByRepositoryAndProAndPublicAndStatus :many
SELECT * FROM extension
WHERE 
     repository = ? AND pro = ? AND public = ? AND status = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByVersionAndRepositoryAndProAndPublicAndStatus :many
SELECT * FROM extension
WHERE 
     version = ? AND repository = ? AND pro = ? AND public = ? AND status = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByIdentifierAndRepositoryAndProAndPublicAndStatus :many
SELECT * FROM extension
WHERE 
     identifier = ? AND repository = ? AND pro = ? AND public = ? AND status = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByVersionAndIdentifierAndRepositoryAndProAndPublicAndStatus :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND repository = ? AND pro = ? AND public = ? AND status = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByExtensionTypeAndProAndPublicAndStatus :many
SELECT * FROM extension
WHERE 
     extension_type = ? AND pro = ? AND public = ? AND status = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByVersionAndExtensionTypeAndProAndPublicAndStatus :many
SELECT * FROM extension
WHERE 
     version = ? AND extension_type = ? AND pro = ? AND public = ? AND status = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByIdentifierAndExtensionTypeAndProAndPublicAndStatus :many
SELECT * FROM extension
WHERE 
     identifier = ? AND extension_type = ? AND pro = ? AND public = ? AND status = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByVersionAndIdentifierAndExtensionTypeAndProAndPublicAndStatus :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND extension_type = ? AND pro = ? AND public = ? AND status = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByRepositoryAndExtensionTypeAndProAndPublicAndStatus :many
SELECT * FROM extension
WHERE 
     repository = ? AND extension_type = ? AND pro = ? AND public = ? AND status = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByVersionAndRepositoryAndExtensionTypeAndProAndPublicAndStatus :many
SELECT * FROM extension
WHERE 
     version = ? AND repository = ? AND extension_type = ? AND pro = ? AND public = ? AND status = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndProAndPublicAndStatus :many
SELECT * FROM extension
WHERE 
     identifier = ? AND repository = ? AND extension_type = ? AND pro = ? AND public = ? AND status = ?  
    LIMIT ?, ?;



-- name: FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndProAndPublicAndStatus :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND repository = ? AND extension_type = ? AND pro = ? AND public = ? AND status = ?  
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

        
    

    
    
-- name: FetchExtensionByProOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     pro = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByProOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     pro = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByProOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     pro = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByProOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     pro = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByVersionAndProOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND pro = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndProOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND pro = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByVersionAndProOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND pro = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndProOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND pro = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByIdentifierAndProOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND pro = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByIdentifierAndProOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND pro = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByIdentifierAndProOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND pro = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByIdentifierAndProOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND pro = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByVersionAndIdentifierAndProOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND pro = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndIdentifierAndProOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND pro = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByVersionAndIdentifierAndProOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND pro = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndIdentifierAndProOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND pro = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByRepositoryAndProOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     repository = ? AND pro = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByRepositoryAndProOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     repository = ? AND pro = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByRepositoryAndProOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     repository = ? AND pro = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByRepositoryAndProOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     repository = ? AND pro = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByVersionAndRepositoryAndProOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND repository = ? AND pro = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndRepositoryAndProOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND repository = ? AND pro = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByVersionAndRepositoryAndProOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND repository = ? AND pro = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndRepositoryAndProOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND repository = ? AND pro = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByIdentifierAndRepositoryAndProOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND repository = ? AND pro = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByIdentifierAndRepositoryAndProOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND repository = ? AND pro = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByIdentifierAndRepositoryAndProOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND repository = ? AND pro = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByIdentifierAndRepositoryAndProOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND repository = ? AND pro = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByVersionAndIdentifierAndRepositoryAndProOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND repository = ? AND pro = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndIdentifierAndRepositoryAndProOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND repository = ? AND pro = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByVersionAndIdentifierAndRepositoryAndProOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND repository = ? AND pro = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndIdentifierAndRepositoryAndProOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND repository = ? AND pro = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByExtensionTypeAndProOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     extension_type = ? AND pro = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByExtensionTypeAndProOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     extension_type = ? AND pro = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByExtensionTypeAndProOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     extension_type = ? AND pro = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByExtensionTypeAndProOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     extension_type = ? AND pro = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByVersionAndExtensionTypeAndProOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND extension_type = ? AND pro = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndExtensionTypeAndProOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND extension_type = ? AND pro = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByVersionAndExtensionTypeAndProOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND extension_type = ? AND pro = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndExtensionTypeAndProOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND extension_type = ? AND pro = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByIdentifierAndExtensionTypeAndProOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND extension_type = ? AND pro = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByIdentifierAndExtensionTypeAndProOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND extension_type = ? AND pro = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByIdentifierAndExtensionTypeAndProOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND extension_type = ? AND pro = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByIdentifierAndExtensionTypeAndProOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND extension_type = ? AND pro = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByVersionAndIdentifierAndExtensionTypeAndProOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND extension_type = ? AND pro = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndIdentifierAndExtensionTypeAndProOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND extension_type = ? AND pro = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByVersionAndIdentifierAndExtensionTypeAndProOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND extension_type = ? AND pro = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndIdentifierAndExtensionTypeAndProOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND extension_type = ? AND pro = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByRepositoryAndExtensionTypeAndProOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     repository = ? AND extension_type = ? AND pro = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByRepositoryAndExtensionTypeAndProOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     repository = ? AND extension_type = ? AND pro = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByRepositoryAndExtensionTypeAndProOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     repository = ? AND extension_type = ? AND pro = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByRepositoryAndExtensionTypeAndProOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     repository = ? AND extension_type = ? AND pro = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByVersionAndRepositoryAndExtensionTypeAndProOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND repository = ? AND extension_type = ? AND pro = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndRepositoryAndExtensionTypeAndProOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND repository = ? AND extension_type = ? AND pro = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByVersionAndRepositoryAndExtensionTypeAndProOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND repository = ? AND extension_type = ? AND pro = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndRepositoryAndExtensionTypeAndProOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND repository = ? AND extension_type = ? AND pro = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndProOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND repository = ? AND extension_type = ? AND pro = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndProOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND repository = ? AND extension_type = ? AND pro = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndProOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND repository = ? AND extension_type = ? AND pro = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndProOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND repository = ? AND extension_type = ? AND pro = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndProOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND repository = ? AND extension_type = ? AND pro = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndProOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND repository = ? AND extension_type = ? AND pro = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndProOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND repository = ? AND extension_type = ? AND pro = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndProOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND repository = ? AND extension_type = ? AND pro = ?  
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

        
    

    
    
-- name: FetchExtensionByProAndPublicOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     pro = ? AND public = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByProAndPublicOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     pro = ? AND public = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByProAndPublicOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     pro = ? AND public = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByProAndPublicOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     pro = ? AND public = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByVersionAndProAndPublicOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND pro = ? AND public = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndProAndPublicOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND pro = ? AND public = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByVersionAndProAndPublicOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND pro = ? AND public = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndProAndPublicOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND pro = ? AND public = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByIdentifierAndProAndPublicOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND pro = ? AND public = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByIdentifierAndProAndPublicOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND pro = ? AND public = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByIdentifierAndProAndPublicOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND pro = ? AND public = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByIdentifierAndProAndPublicOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND pro = ? AND public = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByVersionAndIdentifierAndProAndPublicOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND pro = ? AND public = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndIdentifierAndProAndPublicOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND pro = ? AND public = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByVersionAndIdentifierAndProAndPublicOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND pro = ? AND public = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndIdentifierAndProAndPublicOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND pro = ? AND public = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByRepositoryAndProAndPublicOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     repository = ? AND pro = ? AND public = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByRepositoryAndProAndPublicOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     repository = ? AND pro = ? AND public = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByRepositoryAndProAndPublicOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     repository = ? AND pro = ? AND public = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByRepositoryAndProAndPublicOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     repository = ? AND pro = ? AND public = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByVersionAndRepositoryAndProAndPublicOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND repository = ? AND pro = ? AND public = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndRepositoryAndProAndPublicOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND repository = ? AND pro = ? AND public = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByVersionAndRepositoryAndProAndPublicOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND repository = ? AND pro = ? AND public = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndRepositoryAndProAndPublicOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND repository = ? AND pro = ? AND public = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByIdentifierAndRepositoryAndProAndPublicOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND repository = ? AND pro = ? AND public = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByIdentifierAndRepositoryAndProAndPublicOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND repository = ? AND pro = ? AND public = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByIdentifierAndRepositoryAndProAndPublicOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND repository = ? AND pro = ? AND public = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByIdentifierAndRepositoryAndProAndPublicOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND repository = ? AND pro = ? AND public = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByVersionAndIdentifierAndRepositoryAndProAndPublicOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND repository = ? AND pro = ? AND public = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndIdentifierAndRepositoryAndProAndPublicOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND repository = ? AND pro = ? AND public = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByVersionAndIdentifierAndRepositoryAndProAndPublicOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND repository = ? AND pro = ? AND public = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndIdentifierAndRepositoryAndProAndPublicOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND repository = ? AND pro = ? AND public = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByExtensionTypeAndProAndPublicOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     extension_type = ? AND pro = ? AND public = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByExtensionTypeAndProAndPublicOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     extension_type = ? AND pro = ? AND public = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByExtensionTypeAndProAndPublicOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     extension_type = ? AND pro = ? AND public = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByExtensionTypeAndProAndPublicOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     extension_type = ? AND pro = ? AND public = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByVersionAndExtensionTypeAndProAndPublicOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND extension_type = ? AND pro = ? AND public = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndExtensionTypeAndProAndPublicOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND extension_type = ? AND pro = ? AND public = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByVersionAndExtensionTypeAndProAndPublicOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND extension_type = ? AND pro = ? AND public = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndExtensionTypeAndProAndPublicOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND extension_type = ? AND pro = ? AND public = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByIdentifierAndExtensionTypeAndProAndPublicOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND extension_type = ? AND pro = ? AND public = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByIdentifierAndExtensionTypeAndProAndPublicOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND extension_type = ? AND pro = ? AND public = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByIdentifierAndExtensionTypeAndProAndPublicOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND extension_type = ? AND pro = ? AND public = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByIdentifierAndExtensionTypeAndProAndPublicOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND extension_type = ? AND pro = ? AND public = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByVersionAndIdentifierAndExtensionTypeAndProAndPublicOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND extension_type = ? AND pro = ? AND public = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndIdentifierAndExtensionTypeAndProAndPublicOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND extension_type = ? AND pro = ? AND public = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByVersionAndIdentifierAndExtensionTypeAndProAndPublicOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND extension_type = ? AND pro = ? AND public = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndIdentifierAndExtensionTypeAndProAndPublicOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND extension_type = ? AND pro = ? AND public = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByRepositoryAndExtensionTypeAndProAndPublicOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     repository = ? AND extension_type = ? AND pro = ? AND public = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByRepositoryAndExtensionTypeAndProAndPublicOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     repository = ? AND extension_type = ? AND pro = ? AND public = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByRepositoryAndExtensionTypeAndProAndPublicOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     repository = ? AND extension_type = ? AND pro = ? AND public = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByRepositoryAndExtensionTypeAndProAndPublicOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     repository = ? AND extension_type = ? AND pro = ? AND public = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByVersionAndRepositoryAndExtensionTypeAndProAndPublicOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND repository = ? AND extension_type = ? AND pro = ? AND public = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndRepositoryAndExtensionTypeAndProAndPublicOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND repository = ? AND extension_type = ? AND pro = ? AND public = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByVersionAndRepositoryAndExtensionTypeAndProAndPublicOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND repository = ? AND extension_type = ? AND pro = ? AND public = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndRepositoryAndExtensionTypeAndProAndPublicOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND repository = ? AND extension_type = ? AND pro = ? AND public = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndProAndPublicOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND repository = ? AND extension_type = ? AND pro = ? AND public = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndProAndPublicOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND repository = ? AND extension_type = ? AND pro = ? AND public = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndProAndPublicOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND repository = ? AND extension_type = ? AND pro = ? AND public = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndProAndPublicOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND repository = ? AND extension_type = ? AND pro = ? AND public = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndProAndPublicOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND repository = ? AND extension_type = ? AND pro = ? AND public = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndProAndPublicOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND repository = ? AND extension_type = ? AND pro = ? AND public = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndProAndPublicOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND repository = ? AND extension_type = ? AND pro = ? AND public = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndProAndPublicOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND repository = ? AND extension_type = ? AND pro = ? AND public = ?  
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

        
    

    
    
-- name: FetchExtensionByProAndStatusOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     pro = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByProAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     pro = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByProAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     pro = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByProAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     pro = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByVersionAndProAndStatusOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND pro = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndProAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND pro = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByVersionAndProAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND pro = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndProAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND pro = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByIdentifierAndProAndStatusOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND pro = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByIdentifierAndProAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND pro = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByIdentifierAndProAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND pro = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByIdentifierAndProAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND pro = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByVersionAndIdentifierAndProAndStatusOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND pro = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndIdentifierAndProAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND pro = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByVersionAndIdentifierAndProAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND pro = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndIdentifierAndProAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND pro = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByRepositoryAndProAndStatusOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     repository = ? AND pro = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByRepositoryAndProAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     repository = ? AND pro = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByRepositoryAndProAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     repository = ? AND pro = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByRepositoryAndProAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     repository = ? AND pro = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByVersionAndRepositoryAndProAndStatusOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND repository = ? AND pro = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndRepositoryAndProAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND repository = ? AND pro = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByVersionAndRepositoryAndProAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND repository = ? AND pro = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndRepositoryAndProAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND repository = ? AND pro = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByIdentifierAndRepositoryAndProAndStatusOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND repository = ? AND pro = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByIdentifierAndRepositoryAndProAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND repository = ? AND pro = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByIdentifierAndRepositoryAndProAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND repository = ? AND pro = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByIdentifierAndRepositoryAndProAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND repository = ? AND pro = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByVersionAndIdentifierAndRepositoryAndProAndStatusOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND repository = ? AND pro = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndIdentifierAndRepositoryAndProAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND repository = ? AND pro = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByVersionAndIdentifierAndRepositoryAndProAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND repository = ? AND pro = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndIdentifierAndRepositoryAndProAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND repository = ? AND pro = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByExtensionTypeAndProAndStatusOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     extension_type = ? AND pro = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByExtensionTypeAndProAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     extension_type = ? AND pro = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByExtensionTypeAndProAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     extension_type = ? AND pro = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByExtensionTypeAndProAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     extension_type = ? AND pro = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByVersionAndExtensionTypeAndProAndStatusOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND extension_type = ? AND pro = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndExtensionTypeAndProAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND extension_type = ? AND pro = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByVersionAndExtensionTypeAndProAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND extension_type = ? AND pro = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndExtensionTypeAndProAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND extension_type = ? AND pro = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByIdentifierAndExtensionTypeAndProAndStatusOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND extension_type = ? AND pro = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByIdentifierAndExtensionTypeAndProAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND extension_type = ? AND pro = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByIdentifierAndExtensionTypeAndProAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND extension_type = ? AND pro = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByIdentifierAndExtensionTypeAndProAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND extension_type = ? AND pro = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByVersionAndIdentifierAndExtensionTypeAndProAndStatusOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND extension_type = ? AND pro = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndIdentifierAndExtensionTypeAndProAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND extension_type = ? AND pro = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByVersionAndIdentifierAndExtensionTypeAndProAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND extension_type = ? AND pro = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndIdentifierAndExtensionTypeAndProAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND extension_type = ? AND pro = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByRepositoryAndExtensionTypeAndProAndStatusOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     repository = ? AND extension_type = ? AND pro = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByRepositoryAndExtensionTypeAndProAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     repository = ? AND extension_type = ? AND pro = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByRepositoryAndExtensionTypeAndProAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     repository = ? AND extension_type = ? AND pro = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByRepositoryAndExtensionTypeAndProAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     repository = ? AND extension_type = ? AND pro = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByVersionAndRepositoryAndExtensionTypeAndProAndStatusOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND repository = ? AND extension_type = ? AND pro = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndRepositoryAndExtensionTypeAndProAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND repository = ? AND extension_type = ? AND pro = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByVersionAndRepositoryAndExtensionTypeAndProAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND repository = ? AND extension_type = ? AND pro = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndRepositoryAndExtensionTypeAndProAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND repository = ? AND extension_type = ? AND pro = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndProAndStatusOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND repository = ? AND extension_type = ? AND pro = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndProAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND repository = ? AND extension_type = ? AND pro = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndProAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND repository = ? AND extension_type = ? AND pro = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndProAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND repository = ? AND extension_type = ? AND pro = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndProAndStatusOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND repository = ? AND extension_type = ? AND pro = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndProAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND repository = ? AND extension_type = ? AND pro = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndProAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND repository = ? AND extension_type = ? AND pro = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndProAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND repository = ? AND extension_type = ? AND pro = ? AND status = ?  
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

        
    

    
    
-- name: FetchExtensionByProAndPublicAndStatusOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     pro = ? AND public = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByProAndPublicAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     pro = ? AND public = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByProAndPublicAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     pro = ? AND public = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByProAndPublicAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     pro = ? AND public = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByVersionAndProAndPublicAndStatusOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND pro = ? AND public = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndProAndPublicAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND pro = ? AND public = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByVersionAndProAndPublicAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND pro = ? AND public = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndProAndPublicAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND pro = ? AND public = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByIdentifierAndProAndPublicAndStatusOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND pro = ? AND public = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByIdentifierAndProAndPublicAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND pro = ? AND public = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByIdentifierAndProAndPublicAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND pro = ? AND public = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByIdentifierAndProAndPublicAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND pro = ? AND public = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByVersionAndIdentifierAndProAndPublicAndStatusOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND pro = ? AND public = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndIdentifierAndProAndPublicAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND pro = ? AND public = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByVersionAndIdentifierAndProAndPublicAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND pro = ? AND public = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndIdentifierAndProAndPublicAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND pro = ? AND public = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByRepositoryAndProAndPublicAndStatusOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     repository = ? AND pro = ? AND public = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByRepositoryAndProAndPublicAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     repository = ? AND pro = ? AND public = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByRepositoryAndProAndPublicAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     repository = ? AND pro = ? AND public = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByRepositoryAndProAndPublicAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     repository = ? AND pro = ? AND public = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByVersionAndRepositoryAndProAndPublicAndStatusOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND repository = ? AND pro = ? AND public = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndRepositoryAndProAndPublicAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND repository = ? AND pro = ? AND public = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByVersionAndRepositoryAndProAndPublicAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND repository = ? AND pro = ? AND public = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndRepositoryAndProAndPublicAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND repository = ? AND pro = ? AND public = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByIdentifierAndRepositoryAndProAndPublicAndStatusOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND repository = ? AND pro = ? AND public = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByIdentifierAndRepositoryAndProAndPublicAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND repository = ? AND pro = ? AND public = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByIdentifierAndRepositoryAndProAndPublicAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND repository = ? AND pro = ? AND public = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByIdentifierAndRepositoryAndProAndPublicAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND repository = ? AND pro = ? AND public = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByVersionAndIdentifierAndRepositoryAndProAndPublicAndStatusOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND repository = ? AND pro = ? AND public = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndIdentifierAndRepositoryAndProAndPublicAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND repository = ? AND pro = ? AND public = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByVersionAndIdentifierAndRepositoryAndProAndPublicAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND repository = ? AND pro = ? AND public = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndIdentifierAndRepositoryAndProAndPublicAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND repository = ? AND pro = ? AND public = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByExtensionTypeAndProAndPublicAndStatusOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     extension_type = ? AND pro = ? AND public = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByExtensionTypeAndProAndPublicAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     extension_type = ? AND pro = ? AND public = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByExtensionTypeAndProAndPublicAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     extension_type = ? AND pro = ? AND public = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByExtensionTypeAndProAndPublicAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     extension_type = ? AND pro = ? AND public = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByVersionAndExtensionTypeAndProAndPublicAndStatusOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND extension_type = ? AND pro = ? AND public = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndExtensionTypeAndProAndPublicAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND extension_type = ? AND pro = ? AND public = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByVersionAndExtensionTypeAndProAndPublicAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND extension_type = ? AND pro = ? AND public = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndExtensionTypeAndProAndPublicAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND extension_type = ? AND pro = ? AND public = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByIdentifierAndExtensionTypeAndProAndPublicAndStatusOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND extension_type = ? AND pro = ? AND public = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByIdentifierAndExtensionTypeAndProAndPublicAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND extension_type = ? AND pro = ? AND public = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByIdentifierAndExtensionTypeAndProAndPublicAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND extension_type = ? AND pro = ? AND public = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByIdentifierAndExtensionTypeAndProAndPublicAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND extension_type = ? AND pro = ? AND public = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByVersionAndIdentifierAndExtensionTypeAndProAndPublicAndStatusOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND extension_type = ? AND pro = ? AND public = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndIdentifierAndExtensionTypeAndProAndPublicAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND extension_type = ? AND pro = ? AND public = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByVersionAndIdentifierAndExtensionTypeAndProAndPublicAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND extension_type = ? AND pro = ? AND public = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndIdentifierAndExtensionTypeAndProAndPublicAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND extension_type = ? AND pro = ? AND public = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByRepositoryAndExtensionTypeAndProAndPublicAndStatusOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     repository = ? AND extension_type = ? AND pro = ? AND public = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByRepositoryAndExtensionTypeAndProAndPublicAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     repository = ? AND extension_type = ? AND pro = ? AND public = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByRepositoryAndExtensionTypeAndProAndPublicAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     repository = ? AND extension_type = ? AND pro = ? AND public = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByRepositoryAndExtensionTypeAndProAndPublicAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     repository = ? AND extension_type = ? AND pro = ? AND public = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByVersionAndRepositoryAndExtensionTypeAndProAndPublicAndStatusOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND repository = ? AND extension_type = ? AND pro = ? AND public = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndRepositoryAndExtensionTypeAndProAndPublicAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND repository = ? AND extension_type = ? AND pro = ? AND public = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByVersionAndRepositoryAndExtensionTypeAndProAndPublicAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND repository = ? AND extension_type = ? AND pro = ? AND public = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndRepositoryAndExtensionTypeAndProAndPublicAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND repository = ? AND extension_type = ? AND pro = ? AND public = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndProAndPublicAndStatusOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND repository = ? AND extension_type = ? AND pro = ? AND public = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndProAndPublicAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND repository = ? AND extension_type = ? AND pro = ? AND public = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndProAndPublicAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND repository = ? AND extension_type = ? AND pro = ? AND public = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndProAndPublicAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     identifier = ? AND repository = ? AND extension_type = ? AND pro = ? AND public = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndProAndPublicAndStatusOrderedByCreatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND repository = ? AND extension_type = ? AND pro = ? AND public = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndProAndPublicAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND repository = ? AND extension_type = ? AND pro = ? AND public = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndProAndPublicAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND repository = ? AND extension_type = ? AND pro = ? AND public = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndProAndPublicAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM extension
WHERE 
     version = ? AND identifier = ? AND repository = ? AND extension_type = ? AND pro = ? AND public = ? AND status = ?  
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



-- name: FetchChangeRequestByChangeType :many
SELECT * FROM change_request
WHERE 
     change_type = ?  
    LIMIT ?, ?;



-- name: FetchChangeRequestByVersionAndChangeType :many
SELECT * FROM change_request
WHERE 
     version = ? AND change_type = ?  
    LIMIT ?, ?;



-- name: FetchChangeRequestByReviewStatus :many
SELECT * FROM change_request
WHERE 
     review_status = ?  
    LIMIT ?, ?;



-- name: FetchChangeRequestByVersionAndReviewStatus :many
SELECT * FROM change_request
WHERE 
     version = ? AND review_status = ?  
    LIMIT ?, ?;



-- name: FetchChangeRequestByChangeTypeAndReviewStatus :many
SELECT * FROM change_request
WHERE 
     change_type = ? AND review_status = ?  
    LIMIT ?, ?;



-- name: FetchChangeRequestByVersionAndChangeTypeAndReviewStatus :many
SELECT * FROM change_request
WHERE 
     version = ? AND change_type = ? AND review_status = ?  
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



-- name: FetchChangeRequestByChangeTypeAndStatus :many
SELECT * FROM change_request
WHERE 
     change_type = ? AND status = ?  
    LIMIT ?, ?;



-- name: FetchChangeRequestByVersionAndChangeTypeAndStatus :many
SELECT * FROM change_request
WHERE 
     version = ? AND change_type = ? AND status = ?  
    LIMIT ?, ?;



-- name: FetchChangeRequestByReviewStatusAndStatus :many
SELECT * FROM change_request
WHERE 
     review_status = ? AND status = ?  
    LIMIT ?, ?;



-- name: FetchChangeRequestByVersionAndReviewStatusAndStatus :many
SELECT * FROM change_request
WHERE 
     version = ? AND review_status = ? AND status = ?  
    LIMIT ?, ?;



-- name: FetchChangeRequestByChangeTypeAndReviewStatusAndStatus :many
SELECT * FROM change_request
WHERE 
     change_type = ? AND review_status = ? AND status = ?  
    LIMIT ?, ?;



-- name: FetchChangeRequestByVersionAndChangeTypeAndReviewStatusAndStatus :many
SELECT * FROM change_request
WHERE 
     version = ? AND change_type = ? AND review_status = ? AND status = ?  
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

        
    

    
    
-- name: FetchChangeRequestByChangeTypeOrderedByCreatedAtASC :many
SELECT * FROM change_request
WHERE 
     change_type = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchChangeRequestByChangeTypeOrderedByCreatedAtDESC :many
SELECT * FROM change_request
WHERE 
     change_type = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchChangeRequestByChangeTypeOrderedByUpdatedAtASC :many
SELECT * FROM change_request
WHERE 
     change_type = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchChangeRequestByChangeTypeOrderedByUpdatedAtDESC :many
SELECT * FROM change_request
WHERE 
     change_type = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchChangeRequestByVersionAndChangeTypeOrderedByCreatedAtASC :many
SELECT * FROM change_request
WHERE 
     version = ? AND change_type = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchChangeRequestByVersionAndChangeTypeOrderedByCreatedAtDESC :many
SELECT * FROM change_request
WHERE 
     version = ? AND change_type = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchChangeRequestByVersionAndChangeTypeOrderedByUpdatedAtASC :many
SELECT * FROM change_request
WHERE 
     version = ? AND change_type = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchChangeRequestByVersionAndChangeTypeOrderedByUpdatedAtDESC :many
SELECT * FROM change_request
WHERE 
     version = ? AND change_type = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchChangeRequestByReviewStatusOrderedByCreatedAtASC :many
SELECT * FROM change_request
WHERE 
     review_status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchChangeRequestByReviewStatusOrderedByCreatedAtDESC :many
SELECT * FROM change_request
WHERE 
     review_status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchChangeRequestByReviewStatusOrderedByUpdatedAtASC :many
SELECT * FROM change_request
WHERE 
     review_status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchChangeRequestByReviewStatusOrderedByUpdatedAtDESC :many
SELECT * FROM change_request
WHERE 
     review_status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchChangeRequestByVersionAndReviewStatusOrderedByCreatedAtASC :many
SELECT * FROM change_request
WHERE 
     version = ? AND review_status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchChangeRequestByVersionAndReviewStatusOrderedByCreatedAtDESC :many
SELECT * FROM change_request
WHERE 
     version = ? AND review_status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchChangeRequestByVersionAndReviewStatusOrderedByUpdatedAtASC :many
SELECT * FROM change_request
WHERE 
     version = ? AND review_status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchChangeRequestByVersionAndReviewStatusOrderedByUpdatedAtDESC :many
SELECT * FROM change_request
WHERE 
     version = ? AND review_status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchChangeRequestByChangeTypeAndReviewStatusOrderedByCreatedAtASC :many
SELECT * FROM change_request
WHERE 
     change_type = ? AND review_status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchChangeRequestByChangeTypeAndReviewStatusOrderedByCreatedAtDESC :many
SELECT * FROM change_request
WHERE 
     change_type = ? AND review_status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchChangeRequestByChangeTypeAndReviewStatusOrderedByUpdatedAtASC :many
SELECT * FROM change_request
WHERE 
     change_type = ? AND review_status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchChangeRequestByChangeTypeAndReviewStatusOrderedByUpdatedAtDESC :many
SELECT * FROM change_request
WHERE 
     change_type = ? AND review_status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchChangeRequestByVersionAndChangeTypeAndReviewStatusOrderedByCreatedAtASC :many
SELECT * FROM change_request
WHERE 
     version = ? AND change_type = ? AND review_status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchChangeRequestByVersionAndChangeTypeAndReviewStatusOrderedByCreatedAtDESC :many
SELECT * FROM change_request
WHERE 
     version = ? AND change_type = ? AND review_status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchChangeRequestByVersionAndChangeTypeAndReviewStatusOrderedByUpdatedAtASC :many
SELECT * FROM change_request
WHERE 
     version = ? AND change_type = ? AND review_status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchChangeRequestByVersionAndChangeTypeAndReviewStatusOrderedByUpdatedAtDESC :many
SELECT * FROM change_request
WHERE 
     version = ? AND change_type = ? AND review_status = ?  
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

        
    

    
    
-- name: FetchChangeRequestByChangeTypeAndStatusOrderedByCreatedAtASC :many
SELECT * FROM change_request
WHERE 
     change_type = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchChangeRequestByChangeTypeAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM change_request
WHERE 
     change_type = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchChangeRequestByChangeTypeAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM change_request
WHERE 
     change_type = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchChangeRequestByChangeTypeAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM change_request
WHERE 
     change_type = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchChangeRequestByVersionAndChangeTypeAndStatusOrderedByCreatedAtASC :many
SELECT * FROM change_request
WHERE 
     version = ? AND change_type = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchChangeRequestByVersionAndChangeTypeAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM change_request
WHERE 
     version = ? AND change_type = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchChangeRequestByVersionAndChangeTypeAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM change_request
WHERE 
     version = ? AND change_type = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchChangeRequestByVersionAndChangeTypeAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM change_request
WHERE 
     version = ? AND change_type = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchChangeRequestByReviewStatusAndStatusOrderedByCreatedAtASC :many
SELECT * FROM change_request
WHERE 
     review_status = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchChangeRequestByReviewStatusAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM change_request
WHERE 
     review_status = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchChangeRequestByReviewStatusAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM change_request
WHERE 
     review_status = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchChangeRequestByReviewStatusAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM change_request
WHERE 
     review_status = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchChangeRequestByVersionAndReviewStatusAndStatusOrderedByCreatedAtASC :many
SELECT * FROM change_request
WHERE 
     version = ? AND review_status = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchChangeRequestByVersionAndReviewStatusAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM change_request
WHERE 
     version = ? AND review_status = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchChangeRequestByVersionAndReviewStatusAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM change_request
WHERE 
     version = ? AND review_status = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchChangeRequestByVersionAndReviewStatusAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM change_request
WHERE 
     version = ? AND review_status = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchChangeRequestByChangeTypeAndReviewStatusAndStatusOrderedByCreatedAtASC :many
SELECT * FROM change_request
WHERE 
     change_type = ? AND review_status = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchChangeRequestByChangeTypeAndReviewStatusAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM change_request
WHERE 
     change_type = ? AND review_status = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchChangeRequestByChangeTypeAndReviewStatusAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM change_request
WHERE 
     change_type = ? AND review_status = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchChangeRequestByChangeTypeAndReviewStatusAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM change_request
WHERE 
     change_type = ? AND review_status = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchChangeRequestByVersionAndChangeTypeAndReviewStatusAndStatusOrderedByCreatedAtASC :many
SELECT * FROM change_request
WHERE 
     version = ? AND change_type = ? AND review_status = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchChangeRequestByVersionAndChangeTypeAndReviewStatusAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM change_request
WHERE 
     version = ? AND change_type = ? AND review_status = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchChangeRequestByVersionAndChangeTypeAndReviewStatusAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM change_request
WHERE 
     version = ? AND change_type = ? AND review_status = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchChangeRequestByVersionAndChangeTypeAndReviewStatusAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM change_request
WHERE 
     version = ? AND change_type = ? AND review_status = ? AND status = ?  
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



-- name: FetchUserTeamByRole :many
SELECT * FROM user_team
WHERE 
     role = ?  
    LIMIT ?, ?;



-- name: FetchUserTeamByUserEmailAndRole :many
SELECT * FROM user_team
WHERE 
     user_email = ? AND role = ?  
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



-- name: FetchUserTeamByRoleAndStatus :many
SELECT * FROM user_team
WHERE 
     role = ? AND status = ?  
    LIMIT ?, ?;



-- name: FetchUserTeamByUserEmailAndRoleAndStatus :many
SELECT * FROM user_team
WHERE 
     user_email = ? AND role = ? AND status = ?  
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

        
    

    
    
-- name: FetchUserTeamByRoleOrderedByCreatedAtASC :many
SELECT * FROM user_team
WHERE 
     role = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchUserTeamByRoleOrderedByCreatedAtDESC :many
SELECT * FROM user_team
WHERE 
     role = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchUserTeamByRoleOrderedByUpdatedAtASC :many
SELECT * FROM user_team
WHERE 
     role = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchUserTeamByRoleOrderedByUpdatedAtDESC :many
SELECT * FROM user_team
WHERE 
     role = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchUserTeamByUserEmailAndRoleOrderedByCreatedAtASC :many
SELECT * FROM user_team
WHERE 
     user_email = ? AND role = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchUserTeamByUserEmailAndRoleOrderedByCreatedAtDESC :many
SELECT * FROM user_team
WHERE 
     user_email = ? AND role = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchUserTeamByUserEmailAndRoleOrderedByUpdatedAtASC :many
SELECT * FROM user_team
WHERE 
     user_email = ? AND role = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchUserTeamByUserEmailAndRoleOrderedByUpdatedAtDESC :many
SELECT * FROM user_team
WHERE 
     user_email = ? AND role = ?  
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

        
    

    
    
-- name: FetchUserTeamByRoleAndStatusOrderedByCreatedAtASC :many
SELECT * FROM user_team
WHERE 
     role = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchUserTeamByRoleAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM user_team
WHERE 
     role = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchUserTeamByRoleAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM user_team
WHERE 
     role = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchUserTeamByRoleAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM user_team
WHERE 
     role = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchUserTeamByUserEmailAndRoleAndStatusOrderedByCreatedAtASC :many
SELECT * FROM user_team
WHERE 
     user_email = ? AND role = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchUserTeamByUserEmailAndRoleAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM user_team
WHERE 
     user_email = ? AND role = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchUserTeamByUserEmailAndRoleAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM user_team
WHERE 
     user_email = ? AND role = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchUserTeamByUserEmailAndRoleAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM user_team
WHERE 
     user_email = ? AND role = ? AND status = ?  
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

        
    










-- name: FetchUserProjectByUUID :many
SELECT * FROM user_project
WHERE 
     uuid = ?  
    ;


-- name: FetchUserProjectByUUIDForUpdate :many
SELECT * FROM user_project
WHERE 
     uuid = ?      
FOR UPDATE;


-- name: FetchUserProjectByUserEmail :many
SELECT * FROM user_project
WHERE 
     user_email = ?  
    LIMIT ?, ?;



-- name: FetchUserProjectByRole :many
SELECT * FROM user_project
WHERE 
     role = ?  
    LIMIT ?, ?;



-- name: FetchUserProjectByUserEmailAndRole :many
SELECT * FROM user_project
WHERE 
     user_email = ? AND role = ?  
    LIMIT ?, ?;



-- name: FetchUserProjectByStatus :many
SELECT * FROM user_project
WHERE 
     status = ?  
    LIMIT ?, ?;



-- name: FetchUserProjectByUserEmailAndStatus :many
SELECT * FROM user_project
WHERE 
     user_email = ? AND status = ?  
    LIMIT ?, ?;



-- name: FetchUserProjectByRoleAndStatus :many
SELECT * FROM user_project
WHERE 
     role = ? AND status = ?  
    LIMIT ?, ?;



-- name: FetchUserProjectByUserEmailAndRoleAndStatus :many
SELECT * FROM user_project
WHERE 
     user_email = ? AND role = ? AND status = ?  
    LIMIT ?, ?;





    

    
    
-- name: FetchUserProjectByUserEmailOrderedByCreatedAtASC :many
SELECT * FROM user_project
WHERE 
     user_email = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchUserProjectByUserEmailOrderedByCreatedAtDESC :many
SELECT * FROM user_project
WHERE 
     user_email = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchUserProjectByUserEmailOrderedByUpdatedAtASC :many
SELECT * FROM user_project
WHERE 
     user_email = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchUserProjectByUserEmailOrderedByUpdatedAtDESC :many
SELECT * FROM user_project
WHERE 
     user_email = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchUserProjectByRoleOrderedByCreatedAtASC :many
SELECT * FROM user_project
WHERE 
     role = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchUserProjectByRoleOrderedByCreatedAtDESC :many
SELECT * FROM user_project
WHERE 
     role = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchUserProjectByRoleOrderedByUpdatedAtASC :many
SELECT * FROM user_project
WHERE 
     role = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchUserProjectByRoleOrderedByUpdatedAtDESC :many
SELECT * FROM user_project
WHERE 
     role = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchUserProjectByUserEmailAndRoleOrderedByCreatedAtASC :many
SELECT * FROM user_project
WHERE 
     user_email = ? AND role = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchUserProjectByUserEmailAndRoleOrderedByCreatedAtDESC :many
SELECT * FROM user_project
WHERE 
     user_email = ? AND role = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchUserProjectByUserEmailAndRoleOrderedByUpdatedAtASC :many
SELECT * FROM user_project
WHERE 
     user_email = ? AND role = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchUserProjectByUserEmailAndRoleOrderedByUpdatedAtDESC :many
SELECT * FROM user_project
WHERE 
     user_email = ? AND role = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchUserProjectByStatusOrderedByCreatedAtASC :many
SELECT * FROM user_project
WHERE 
     status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchUserProjectByStatusOrderedByCreatedAtDESC :many
SELECT * FROM user_project
WHERE 
     status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchUserProjectByStatusOrderedByUpdatedAtASC :many
SELECT * FROM user_project
WHERE 
     status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchUserProjectByStatusOrderedByUpdatedAtDESC :many
SELECT * FROM user_project
WHERE 
     status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchUserProjectByUserEmailAndStatusOrderedByCreatedAtASC :many
SELECT * FROM user_project
WHERE 
     user_email = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchUserProjectByUserEmailAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM user_project
WHERE 
     user_email = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchUserProjectByUserEmailAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM user_project
WHERE 
     user_email = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchUserProjectByUserEmailAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM user_project
WHERE 
     user_email = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchUserProjectByRoleAndStatusOrderedByCreatedAtASC :many
SELECT * FROM user_project
WHERE 
     role = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchUserProjectByRoleAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM user_project
WHERE 
     role = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchUserProjectByRoleAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM user_project
WHERE 
     role = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchUserProjectByRoleAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM user_project
WHERE 
     role = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchUserProjectByUserEmailAndRoleAndStatusOrderedByCreatedAtASC :many
SELECT * FROM user_project
WHERE 
     user_email = ? AND role = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchUserProjectByUserEmailAndRoleAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM user_project
WHERE 
     user_email = ? AND role = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchUserProjectByUserEmailAndRoleAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM user_project
WHERE 
     user_email = ? AND role = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchUserProjectByUserEmailAndRoleAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM user_project
WHERE 
     user_email = ? AND role = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    










-- name: FetchMembershipByUUID :many
SELECT * FROM membership
WHERE 
     uuid = ?  
    ;


-- name: FetchMembershipByUUIDForUpdate :many
SELECT * FROM membership
WHERE 
     uuid = ?      
FOR UPDATE;


-- name: FetchMembershipByType :many
SELECT * FROM membership
WHERE 
     type = ?  
    LIMIT ?, ?;



-- name: FetchMembershipByStatus :many
SELECT * FROM membership
WHERE 
     status = ?  
    LIMIT ?, ?;



-- name: FetchMembershipByTypeAndStatus :many
SELECT * FROM membership
WHERE 
     type = ? AND status = ?  
    LIMIT ?, ?;





    

    
    
-- name: FetchMembershipByTypeOrderedByStartDateASC :many
SELECT * FROM membership
WHERE 
     type = ?  
    ORDER BY start_date ASC
    LIMIT ?, ?;

-- name: FetchMembershipByTypeOrderedByStartDateDESC :many
SELECT * FROM membership
WHERE 
     type = ?  
    ORDER BY start_date DESC
    LIMIT ?, ?;

        
-- name: FetchMembershipByTypeOrderedByCreatedAtASC :many
SELECT * FROM membership
WHERE 
     type = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchMembershipByTypeOrderedByCreatedAtDESC :many
SELECT * FROM membership
WHERE 
     type = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchMembershipByTypeOrderedByUpdatedAtASC :many
SELECT * FROM membership
WHERE 
     type = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchMembershipByTypeOrderedByUpdatedAtDESC :many
SELECT * FROM membership
WHERE 
     type = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchMembershipByStatusOrderedByStartDateASC :many
SELECT * FROM membership
WHERE 
     status = ?  
    ORDER BY start_date ASC
    LIMIT ?, ?;

-- name: FetchMembershipByStatusOrderedByStartDateDESC :many
SELECT * FROM membership
WHERE 
     status = ?  
    ORDER BY start_date DESC
    LIMIT ?, ?;

        
-- name: FetchMembershipByStatusOrderedByCreatedAtASC :many
SELECT * FROM membership
WHERE 
     status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchMembershipByStatusOrderedByCreatedAtDESC :many
SELECT * FROM membership
WHERE 
     status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchMembershipByStatusOrderedByUpdatedAtASC :many
SELECT * FROM membership
WHERE 
     status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchMembershipByStatusOrderedByUpdatedAtDESC :many
SELECT * FROM membership
WHERE 
     status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchMembershipByTypeAndStatusOrderedByStartDateASC :many
SELECT * FROM membership
WHERE 
     type = ? AND status = ?  
    ORDER BY start_date ASC
    LIMIT ?, ?;

-- name: FetchMembershipByTypeAndStatusOrderedByStartDateDESC :many
SELECT * FROM membership
WHERE 
     type = ? AND status = ?  
    ORDER BY start_date DESC
    LIMIT ?, ?;

        
-- name: FetchMembershipByTypeAndStatusOrderedByCreatedAtASC :many
SELECT * FROM membership
WHERE 
     type = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchMembershipByTypeAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM membership
WHERE 
     type = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchMembershipByTypeAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM membership
WHERE 
     type = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchMembershipByTypeAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM membership
WHERE 
     type = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    










-- name: FetchAiUsageByUUID :many
SELECT * FROM ai_usage
WHERE 
     uuid = ?  
    ;


-- name: FetchAiUsageByUUIDForUpdate :many
SELECT * FROM ai_usage
WHERE 
     uuid = ?      
FOR UPDATE;


-- name: FetchAiUsageByContext :many
SELECT * FROM ai_usage
WHERE 
     context = ?  
    LIMIT ?, ?;



-- name: FetchAiUsageByProvider :many
SELECT * FROM ai_usage
WHERE 
     provider = ?  
    LIMIT ?, ?;



-- name: FetchAiUsageByContextAndProvider :many
SELECT * FROM ai_usage
WHERE 
     context = ? AND provider = ?  
    LIMIT ?, ?;



-- name: FetchAiUsageByStatus :many
SELECT * FROM ai_usage
WHERE 
     status = ?  
    LIMIT ?, ?;



-- name: FetchAiUsageByContextAndStatus :many
SELECT * FROM ai_usage
WHERE 
     context = ? AND status = ?  
    LIMIT ?, ?;



-- name: FetchAiUsageByProviderAndStatus :many
SELECT * FROM ai_usage
WHERE 
     provider = ? AND status = ?  
    LIMIT ?, ?;



-- name: FetchAiUsageByContextAndProviderAndStatus :many
SELECT * FROM ai_usage
WHERE 
     context = ? AND provider = ? AND status = ?  
    LIMIT ?, ?;





    

    
    
-- name: FetchAiUsageByContextOrderedByCreatedAtASC :many
SELECT * FROM ai_usage
WHERE 
     context = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchAiUsageByContextOrderedByCreatedAtDESC :many
SELECT * FROM ai_usage
WHERE 
     context = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchAiUsageByContextOrderedByUpdatedAtASC :many
SELECT * FROM ai_usage
WHERE 
     context = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchAiUsageByContextOrderedByUpdatedAtDESC :many
SELECT * FROM ai_usage
WHERE 
     context = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchAiUsageByProviderOrderedByCreatedAtASC :many
SELECT * FROM ai_usage
WHERE 
     provider = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchAiUsageByProviderOrderedByCreatedAtDESC :many
SELECT * FROM ai_usage
WHERE 
     provider = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchAiUsageByProviderOrderedByUpdatedAtASC :many
SELECT * FROM ai_usage
WHERE 
     provider = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchAiUsageByProviderOrderedByUpdatedAtDESC :many
SELECT * FROM ai_usage
WHERE 
     provider = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchAiUsageByContextAndProviderOrderedByCreatedAtASC :many
SELECT * FROM ai_usage
WHERE 
     context = ? AND provider = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchAiUsageByContextAndProviderOrderedByCreatedAtDESC :many
SELECT * FROM ai_usage
WHERE 
     context = ? AND provider = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchAiUsageByContextAndProviderOrderedByUpdatedAtASC :many
SELECT * FROM ai_usage
WHERE 
     context = ? AND provider = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchAiUsageByContextAndProviderOrderedByUpdatedAtDESC :many
SELECT * FROM ai_usage
WHERE 
     context = ? AND provider = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchAiUsageByStatusOrderedByCreatedAtASC :many
SELECT * FROM ai_usage
WHERE 
     status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchAiUsageByStatusOrderedByCreatedAtDESC :many
SELECT * FROM ai_usage
WHERE 
     status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchAiUsageByStatusOrderedByUpdatedAtASC :many
SELECT * FROM ai_usage
WHERE 
     status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchAiUsageByStatusOrderedByUpdatedAtDESC :many
SELECT * FROM ai_usage
WHERE 
     status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchAiUsageByContextAndStatusOrderedByCreatedAtASC :many
SELECT * FROM ai_usage
WHERE 
     context = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchAiUsageByContextAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM ai_usage
WHERE 
     context = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchAiUsageByContextAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM ai_usage
WHERE 
     context = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchAiUsageByContextAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM ai_usage
WHERE 
     context = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchAiUsageByProviderAndStatusOrderedByCreatedAtASC :many
SELECT * FROM ai_usage
WHERE 
     provider = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchAiUsageByProviderAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM ai_usage
WHERE 
     provider = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchAiUsageByProviderAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM ai_usage
WHERE 
     provider = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchAiUsageByProviderAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM ai_usage
WHERE 
     provider = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchAiUsageByContextAndProviderAndStatusOrderedByCreatedAtASC :many
SELECT * FROM ai_usage
WHERE 
     context = ? AND provider = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchAiUsageByContextAndProviderAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM ai_usage
WHERE 
     context = ? AND provider = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchAiUsageByContextAndProviderAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM ai_usage
WHERE 
     context = ? AND provider = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchAiUsageByContextAndProviderAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM ai_usage
WHERE 
     context = ? AND provider = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    







