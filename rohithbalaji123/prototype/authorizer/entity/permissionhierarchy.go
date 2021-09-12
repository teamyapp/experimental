package entity

type PermissionHierarchyEntity struct {
    ChildPermissionType  string
    ChildResourceType    string
    ParentPermissionType string
    ParentResourceType   string
}

func NewPermissionHierarchyEntity(childPermissionType, childResourceType, parentPermissionType, parentResourceType string) PermissionHierarchyEntity {
    return PermissionHierarchyEntity{
        ChildPermissionType:  childPermissionType,
        ChildResourceType:    childResourceType,
        ParentPermissionType: parentPermissionType,
        ParentResourceType:   parentResourceType,
    }
}

func (p PermissionHierarchyEntity) ChildTypeEquals(childPermissionType, childResourceType string) bool {
    return p.ChildPermissionType == childPermissionType && p.ChildResourceType == childResourceType
}
