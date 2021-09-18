package repo

import "experimental/rohithbalaji123/prototype/authorizer/entity"

type ResourceHierarchyRepo struct {
    entities []entity.ResourceHierarchyEntity
}

func NewResourceHierarchyRepo() ResourceHierarchyRepo {
    r := ResourceHierarchyRepo{
        entities: make([]entity.ResourceHierarchyEntity, 0),
    }
    return r
}

func (r ResourceHierarchyRepo) FindParentResourceIds(parentResourceType, childId, childResourceType string) []string {
    response := make([]string, 0)

    // edge case to cover dependent permissions for the same resource. For eg: write permission to a resource implies read permission to it
    if childResourceType == parentResourceType {
        response = append(response, childId)
    }

    for _, resourceHierarchyEntity := range r.entities {
        if resourceHierarchyEntity.HasParentResource(parentResourceType, childId, childResourceType) {
            response = append(response, resourceHierarchyEntity.ParentId)
        }
    }
    return response[:]
}

func (r ResourceHierarchyRepo) Add(hierarchyEntity entity.ResourceHierarchyEntity) {
    r.entities = append(r.entities, hierarchyEntity)
}
