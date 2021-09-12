package repo

import (
    "experimental/rohithbalaji123/prototype/authorizer/entity"
)

type PermissionBindingRepo struct {
    entities []entity.PermissionBindingEntity
}

func NewPermissionBindingRepo() PermissionBindingRepo {
    p := PermissionBindingRepo{
        entities: make([]entity.PermissionBindingEntity, 0),
    }
    return p
}

func (p PermissionBindingRepo) PermissionBindingExists(request entity.HasPermissionQuery) bool {
    for _, permissionBindingEntity := range p.entities {
        if permissionBindingEntity.Equals(request) {
            return true
        }
    }
    return false
}

func (p PermissionBindingRepo) Add(bindingEntity entity.PermissionBindingEntity) {
    p.entities = append(p.entities, bindingEntity)
}
