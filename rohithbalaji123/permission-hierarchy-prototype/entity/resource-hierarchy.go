package entity

type ResourceHierarchyEntity struct {
    ChildId           string
    ChildResourceType  string
    ParentId           string
    ParentResourceType string
}

func NewResourceHierarchyEntity(childId, childResourceType, parentId, parentResourceType string) ResourceHierarchyEntity {
    return ResourceHierarchyEntity{
        ChildId:            childId,
        ChildResourceType:  childResourceType,
        ParentId:           parentId,
        ParentResourceType: parentResourceType,
    }
}

func (r ResourceHierarchyEntity) HasParentResource(parentResourceType, childId, childResourceType string) bool {
    return r.ParentResourceType == parentResourceType &&
        r.ChildId == childId &&
        r.ChildResourceType == childResourceType
}
