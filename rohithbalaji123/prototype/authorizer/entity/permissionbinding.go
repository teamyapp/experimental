package entity

// HasPermissionQuery In future, we might have groups, teams and organisations which will require us to maintain a graph
// of user groups. Hence the user id(maybe, group id) is required in this query.
type HasPermissionQuery struct {
    PermissionType string
    ResourceId string
    ResourceType string
    UserId string
}

func NewHasPermissionQuery(permissionType, resourceId, resourceType, userId string) HasPermissionQuery {
    return HasPermissionQuery{
        PermissionType: permissionType,
        ResourceId:     resourceId,
        ResourceType:   resourceType,
        UserId:         userId,
    }
}

type PermissionBindingEntity struct {
    PermissionType string
    ResourceId string
    ResourceType string
    UserId string
}

func NewPermissionBindingEntity(permissionType, resourceId, resourceType, userId string) PermissionBindingEntity {
    return PermissionBindingEntity{
        PermissionType: permissionType,
        ResourceId: resourceId,
        ResourceType: resourceType,
        UserId: userId,
    }
}

func (p PermissionBindingEntity) Equals(request HasPermissionQuery) bool {
    return p.PermissionType == request.PermissionType &&
            p.ResourceId == request.ResourceId &&
            p.ResourceType == request.ResourceType &&
            p.UserId == request.UserId
}
