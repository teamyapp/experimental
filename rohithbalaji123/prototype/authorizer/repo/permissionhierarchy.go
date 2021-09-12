package repo

import "experimental/rohithbalaji123/prototype/authorizer/entity"

type PermissionHierarchyRepo struct {
    entities []entity.PermissionHierarchyEntity
}

func NewPermissionHierarchyRepo() PermissionHierarchyRepo {
    p := PermissionHierarchyRepo{
        entities: make([]entity.PermissionHierarchyEntity, 0),
    }
    return p
}

func (p PermissionHierarchyRepo) FindParentPermissions(childPermissionType, childResourceType string) []entity.PermissionHierarchyEntity {
    response := make([]entity.PermissionHierarchyEntity, 0)
    for _, permissionHierarchyEntity := range p.entities {
        if permissionHierarchyEntity.ChildTypeEquals(childPermissionType, childResourceType) {
            response = append(response, permissionHierarchyEntity)
        }
    }
    return response[:]
}

func (p PermissionHierarchyRepo) Add(hierarchyEntity entity.PermissionHierarchyEntity) {
    p.entities = append(p.entities, hierarchyEntity)
}
