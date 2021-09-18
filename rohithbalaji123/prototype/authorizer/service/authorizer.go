package service

import (
    "experimental/rohithbalaji123/prototype/authorizer/entity"
    "experimental/rohithbalaji123/prototype/authorizer/repo"
)

type Authorizer struct {
    PermissionBindingRepo   repo.PermissionBindingRepo
    PermissionHierarchyRepo repo.PermissionHierarchyRepo
    ResourceHierarchyRepo   repo.ResourceHierarchyRepo
}

func NewAuthorizer() Authorizer {
    return Authorizer{
        PermissionBindingRepo:   repo.NewPermissionBindingRepo(),
        PermissionHierarchyRepo: repo.NewPermissionHierarchyRepo(),
        ResourceHierarchyRepo:   repo.NewResourceHierarchyRepo(),
    }
}

func (p Authorizer) HasPermission(query entity.HasPermissionQuery) bool {
    queue := make([]entity.HasPermissionQuery, 0)
    queue = append(queue, query)

    for len(queue) != 0 {
        curPermissionCheck := queue[0]
        queue = queue[1:]

        if p.PermissionBindingRepo.PermissionBindingExists(curPermissionCheck, userId) {
            return true
        }

        parentPermissions := p.PermissionHierarchyRepo.FindParentPermissions(curPermissionCheck.PermissionType, curPermissionCheck.ResourceType)
        for _, parentPermission := range parentPermissions {
            parentResourceIds := p.ResourceHierarchyRepo.FindParentResourceIds(parentPermission.ParentResourceType, curPermissionCheck.ResourceId, curPermissionCheck.ResourceType)
            for _, parentResourceId := range parentResourceIds {
                queue = append(queue, entity.NewHasPermissionQuery(parentPermission.ParentPermissionType, parentResourceId, parentPermission.ParentResourceType, curPermissionCheck.UserId))
            }
        }
    }

    return false
}
