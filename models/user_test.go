package models_test

import "helpdesk_lite/models"

func (ms *ModelSuite) Test_User_Create_OK() {
	count, err := ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(0, count)

	u := &models.User{
		Email: uniqueEmail("user"),
		PasswordHash: "hashed_password",
		Role: "user",
	}

	verrs, err := ms.DB.ValidateAndCreate(u)
	ms.NoError(err)
	ms.False(verrs.HasAny())
	ms.NotZero(u.ID)

	count, err = ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(1, count)
}

func (ms *ModelSuite) Test_User_Create_ValidationErrors() {
	count, err := ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(0, count)

	u := &models.User{
		Email: "",
		PasswordHash: "hashed_password",
		Role: "superadmin",
	}

	verrs, err := ms.DB.ValidateAndCreate(u)
	ms.NoError(err)
	ms.True(verrs.HasAny())

	count, err = ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(0, count)
}

func (ms *ModelSuite) Test_User_Create_EmailAlreadyExists() {
	email := uniqueEmail("duplicate-user")

	u1 := &models.User{
		Email: email,
		PasswordHash: "hashed_password",
		Role: "user",
	}

	verrs, err := ms.DB.ValidateAndCreate(u1)
	ms.NoError(err)
	ms.False(verrs.HasAny())
	ms.NotZero(u1.ID)
	
	count, err := ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(1, count)

	u2 := &models.User{
		Email: email,
		PasswordHash: "another-hash",
		Role: "admin",
	}

	verrs, err = ms.DB.ValidateAndCreate(u2)
	ms.NoError(err)
	ms.True(verrs.HasAny())

	count, err = ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(1, count)
}

func (ms *ModelSuite) Test_User_Update_SameEmail_OK() {
	email := uniqueEmail("same-email-update")
	u := createUser(ms, email, "user")

	u.Role = "admin"
	verrs, err := ms.DB.ValidateAndUpdate(u)
	ms.NoError(err)
	ms.False(verrs.HasAny())

	reloaded := &models.User{}
	err = ms.DB.Find(reloaded, u.ID)
	ms.NoError(err)
	ms.Equal("admin", reloaded.Role)
	ms.Equal(email, reloaded.Email)
}

func (ms *ModelSuite) Test_User_Update_InvalidRole_Fails() {
	u := createUser(ms, uniqueEmail("invalide-role-update"), "user")

	u.Role = "owner"

	verrs, err := ms.DB.ValidateAndUpdate(u)
	ms.NoError(err)
	ms.True(verrs.HasAny())

	reloaded := &models.User{}
	err = ms.DB.Find(reloaded, u.ID)
	ms.NoError(err)
	ms.Equal("user", reloaded.Role)
}