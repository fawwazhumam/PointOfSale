package app

import "github.com/gofiber/fiber/v3"

func SetupRoutes(app *fiber.App, container *Container) {
	api := app.Group("/api/v1")

	roles := api.Group("/roles")
	roles.Post("/", container.RoleController.CreateRole)
	roles.Get("/", container.RoleController.GetAllRoles)
	roles.Get("/:id", container.RoleController.GetRoleByID)
	roles.Put("/:id", container.RoleController.UpdateRole)
	roles.Delete("/:id", container.RoleController.DeleteRole)

	users := api.Group("/users")
	users.Post("/", container.UserController.CreateUser)
	users.Get("/", container.UserController.GetAllUsers)
	users.Get("/email/:email", container.UserController.GetUserByEmail)
	users.Get("/role/:roleName", container.UserController.GetUserByRoleName)
	users.Get("/:id", container.UserController.GetUserByID)
	users.Put("/:id", container.UserController.UpdateUser)
	users.Delete("/:id", container.UserController.DeleteUser)

	userRoles := api.Group("/user-roles")
	userRoles.Post("/", container.UserController.AssignUserToRole)
	userRoles.Get("/", container.UserController.GetAllUserRoles)
	userRoles.Get("/:id", container.UserController.GetUserRoleByID)
	userRoles.Put("/:id", container.UserController.EditAssignUserToRole)

	auth := api.Group("/auth")
	auth.Post("/login", container.AuthController.Login)

	upload := api.Group("/upload")
	upload.Post("/photo", container.UploadController.UploadPhoto)
}