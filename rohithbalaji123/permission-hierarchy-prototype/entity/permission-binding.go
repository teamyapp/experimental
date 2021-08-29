package entity

type HasPermissionRequest struct {
    PermissionType string
    ResourceId string
    ResourceType string
    UserId string
}

func NewHasPermissionRequest(permissionType, resourceId, resourceType, userId string) HasPermissionRequest {
    return HasPermissionRequest{
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

func (p PermissionBindingEntity) Equals(request HasPermissionRequest) bool {
    return p.PermissionType == request.PermissionType &&
            p.ResourceId == request.ResourceId &&
            p.ResourceType == request.ResourceType &&
            p.UserId == request.UserId
}
