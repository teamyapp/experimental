package service

import (
    "experimental/rohithbalaji123/permission-hierarchy-prototype/entity"
    "experimental/rohithbalaji123/permission-hierarchy-prototype/repo"
)

type PermissionCheckerService struct {
    permissionRepo repo.PermissionBindingRepo
    permissionHierarchyRepo repo.PermissionHierarchyRepo
    resourceHierarchyRepo repo.ResourceHierarchyRepo
}

func NewPermissionCheckerService() PermissionCheckerService {
    return PermissionCheckerService{
        permissionRepo: repo.NewPermissionBindingRepo(),
        permissionHierarchyRepo: repo.NewPermissionHierarchyRepo(),
        resourceHierarchyRepo: repo.NewResourceHierarchyRepo(),
    }
}

func (p PermissionCheckerService) HasPermission(permissionType, resourceId, resourceType, userId string) bool {
    queue := make([]entity.HasPermissionRequest, 0)
    queue = append(queue, entity.NewHasPermissionRequest(permissionType, resourceId, resourceType, userId))

    for len(queue) != 0 {
        curPermissionCheck := queue[0]
        queue = queue[1:]

        if p.permissionRepo.PermissionBindingExists(curPermissionCheck) {
            return true
        }

        parentPermissions := p.permissionHierarchyRepo.FindParentPermissions(curPermissionCheck.PermissionType, curPermissionCheck.ResourceType)
        for _, parentPermission := range parentPermissions {
            parentResourceIds := p.resourceHierarchyRepo.FindParentResourceIds(parentPermission.ParentResourceType, curPermissionCheck.ResourceId, curPermissionCheck.ResourceType)
            for _, parentResourceId := range parentResourceIds {
                queue = append(queue, entity.NewHasPermissionRequest(parentPermission.ParentPermissionType, parentResourceId, parentPermission.ParentResourceType, curPermissionCheck.UserId))
            }
        }
    }

    return false
}
