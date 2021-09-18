package service

import (
    "experimental/rohithbalaji123/prototype/authorizer/entity"
    "fmt"
    "testing"
)

func TestHasPermission(t *testing.T) {
    testCases := []struct{
        name string
        permissionHierarchies []entity.PermissionHierarchyEntity
        resourceHierarchies []entity.ResourceHierarchyEntity
        permissionBindings []entity.PermissionBindingEntity
        inputQuery entity.HasPermissionQuery
        expected bool
    } {
        {
            name: "direct permission binding exists",
            permissionHierarchies: []entity.PermissionHierarchyEntity {
                entity.NewPermissionHierarchyEntity("read", "project", "read", "team"),
                entity.NewPermissionHierarchyEntity("read", "project", "write", "project"),
                entity.NewPermissionHierarchyEntity("write", "project", "write", "team"),
            },
            resourceHierarchies: []entity.ResourceHierarchyEntity {
                entity.NewResourceHierarchyEntity("project-1", "project", "team-1", "team"),
            },
            permissionBindings: []entity.PermissionBindingEntity {
                entity.NewPermissionBindingEntity("write", "team-1", "team", "user-1"),
                entity.NewPermissionBindingEntity("read", "team-1", "team", "user-1"),
                entity.NewPermissionBindingEntity("read", "project-1", "project", "user-2"),
                entity.NewPermissionBindingEntity("write", "team-1", "team", "user-3"),
            },
            inputQuery: entity.NewHasPermissionQuery("read", "team-1", "team", "user-1"),
            expected: true,
        },
        {
            name: "direct permission binding does not exist, indirect permission binding does not exist",
            permissionHierarchies: []entity.PermissionHierarchyEntity {
                entity.NewPermissionHierarchyEntity("read", "project", "read", "team"),
                entity.NewPermissionHierarchyEntity("read", "project", "write", "project"),
                entity.NewPermissionHierarchyEntity("write", "project", "write", "team"),
            },
            resourceHierarchies: []entity.ResourceHierarchyEntity {
                entity.NewResourceHierarchyEntity("project-1", "project", "team-1", "team"),
            },
            permissionBindings: []entity.PermissionBindingEntity {
                entity.NewPermissionBindingEntity("write", "team-1", "team", "user-1"),
                entity.NewPermissionBindingEntity("read", "team-1", "team", "user-1"),
                entity.NewPermissionBindingEntity("read", "project-1", "project", "user-2"),
                entity.NewPermissionBindingEntity("write", "team-1", "team", "user-3"),
            },
            inputQuery: entity.NewHasPermissionQuery("read", "team-1", "team", "user-2"),
            expected: false,
        },
        {
            name: "direct permission binding does not exist, indirect permission binding exists for same resource type",
            permissionHierarchies: []entity.PermissionHierarchyEntity {
                entity.NewPermissionHierarchyEntity("read", "project", "read", "team"),
                entity.NewPermissionHierarchyEntity("read", "project", "write", "project"),
                entity.NewPermissionHierarchyEntity("write", "project", "write", "team"),
            },
            resourceHierarchies: []entity.ResourceHierarchyEntity {
                entity.NewResourceHierarchyEntity("project-1", "project", "team-1", "team"),
            },
            permissionBindings: []entity.PermissionBindingEntity {
                entity.NewPermissionBindingEntity("write", "team-1", "team", "user-1"),
                entity.NewPermissionBindingEntity("read", "team-1", "team", "user-1"),
                entity.NewPermissionBindingEntity("read", "project-1", "project", "user-2"),
                entity.NewPermissionBindingEntity("write", "team-1", "team", "user-3"),
            },
            inputQuery: entity.NewHasPermissionQuery("read", "project-1", "project", "user-1"),
            expected: true,
        },
        {
            name: "direct permission binding does not exist, indirect permission binding exists for different resource type in first level",
            permissionHierarchies: []entity.PermissionHierarchyEntity {
                entity.NewPermissionHierarchyEntity("read", "project", "read", "team"),
                entity.NewPermissionHierarchyEntity("read", "project", "write", "project"),
                entity.NewPermissionHierarchyEntity("write", "project", "write", "team"),
            },
            resourceHierarchies: []entity.ResourceHierarchyEntity {
                entity.NewResourceHierarchyEntity("write", "project-1", "project", "user-1"),
            },
            permissionBindings: []entity.PermissionBindingEntity {
                entity.NewPermissionBindingEntity("write", "team-1", "team", "user-1"),
                entity.NewPermissionBindingEntity("read", "team-1", "team", "user-1"),
                entity.NewPermissionBindingEntity("read", "project-1", "project", "user-2"),
                entity.NewPermissionBindingEntity("write", "team-1", "team", "user-3"),
            },
            inputQuery: entity.NewHasPermissionQuery("read", "project-1", "project", "user-1"),
            expected: true,
        },
        {
            name: "direct permission binding does not exist, indirect permission binding exists for different resource type in second level",
            permissionHierarchies: []entity.PermissionHierarchyEntity {
                entity.NewPermissionHierarchyEntity("read", "project", "read", "team"),
                entity.NewPermissionHierarchyEntity("read", "project", "write", "project"),
                entity.NewPermissionHierarchyEntity("write", "project", "write", "team"),
            },
            resourceHierarchies: []entity.ResourceHierarchyEntity {
                entity.NewResourceHierarchyEntity("write", "project-1", "project", "user-1"),
            },
            permissionBindings: []entity.PermissionBindingEntity {
                entity.NewPermissionBindingEntity("write", "team-1", "team", "user-1"),
                entity.NewPermissionBindingEntity("read", "team-1", "team", "user-1"),
                entity.NewPermissionBindingEntity("read", "project-1", "project", "user-2"),
                entity.NewPermissionBindingEntity("write", "team-1", "team", "user-3"),
            },
            inputQuery: entity.NewHasPermissionQuery("read", "project-1", "project", "user-3"),
            expected: true,
        },
    }

    for _, testCase := range testCases {
        testCase := testCase
        t.Run(testCase.name, func(t *testing.T) {
            t.Parallel()
            authorizer := NewAuthorizer()
            for _, resourceHierarchy := range testCase.resourceHierarchies {
                authorizer.ResourceHierarchyRepo.Add(resourceHierarchy)
            }
            for _, permissionHierarchy := range testCase.permissionHierarchies {
                authorizer.PermissionHierarchyRepo.Add(permissionHierarchy)
            }
            for _, permissionBinding := range testCase.permissionBindings {
                authorizer.PermissionBindingRepo.Add(permissionBinding)
            }

            actual := authorizer.HasPermission(testCase.inputQuery)
            fmt.Println(actual == testCase.expected)
        })
    }
}
