package repo

import (
    "experimental/rohithbalaji123/permission-hierarchy-prototype/entity"
)

type PermissionBindingRepo struct {
    entities []entity.PermissionBindingEntity
}

func NewPermissionBindingRepo() PermissionBindingRepo {
    p := PermissionBindingRepo{
        entities: make([]entity.PermissionBindingEntity, 0),
    }
    p.entities = append(p.entities, entity.NewPermissionBindingEntity("write", "team-1", "team", "user-1"))
    p.entities = append(p.entities, entity.NewPermissionBindingEntity("read", "team-1", "team", "user-1"))
    p.entities = append(p.entities, entity.NewPermissionBindingEntity("read", "project-1", "project", "user-2"))
    p.entities = append(p.entities, entity.NewPermissionBindingEntity("write", "project-1", "project", "user-3"))

    return p
}

func (p PermissionBindingRepo) PermissionBindingExists(request entity.HasPermissionRequest) bool {
    for _, permissionBindingEntity := range p.entities {
        if permissionBindingEntity.Equals(request) {
            return true
        }
    }
    return false
}
