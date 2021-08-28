package main

import (
    "experimental/rohithbalaji123/permission-hierarchy-prototype/service"
    "fmt"
)

func main() {
    permissionCheckerService := service.NewPermissionCheckerService()

    fmt.Println(permissionCheckerService.HasPermission("read", "team-1", "team", "user-1"))
    fmt.Println(permissionCheckerService.HasPermission("read", "team-1", "team", "user-2"))
    fmt.Println(permissionCheckerService.HasPermission("read", "project-1", "project", "user-2"))
    fmt.Println(permissionCheckerService.HasPermission("write", "project-1", "project", "user-2"))
    fmt.Println(permissionCheckerService.HasPermission("read", "project-1", "project", "user-1"))
    fmt.Println(permissionCheckerService.HasPermission("write", "project-1", "project", "user-1"))
    fmt.Println(permissionCheckerService.HasPermission("read", "project-1", "project", "user-3"))
}
