package repo

import "experimental/rohithbalaji123/permission-hierarchy-prototype/entity"

type PermissionHierarchyRepo struct {
    entities []entity.PermissionHierarchyEntity
}

func NewPermissionHierarchyRepo() PermissionHierarchyRepo {
    p := PermissionHierarchyRepo{
        entities: make([]entity.PermissionHierarchyEntity, 0),
    }
    p.entities = append(p.entities, entity.NewPermissionHierarchyEntity("read", "project", "read", "team"))
    p.entities = append(p.entities, entity.NewPermissionHierarchyEntity("read", "project", "write", "project"))
    p.entities = append(p.entities, entity.NewPermissionHierarchyEntity("write", "project", "write", "team"))

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
