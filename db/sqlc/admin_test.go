package db

import (
	"context"
	"database/sql"
	"log"
	"testing"
	"time"

	"github.com/bHARATHaALETI018/go-backend/util"
	"github.com/stretchr/testify/require"
)


func createRandomAdmin(t *testing.T) Admin{
	args := CreateAdminParams{
		Email      : util.RandomEmail(),
		Password   : util.RandomPassword(18),
		UserName   : util.RandomStringUser(12),
		FirstName  : util.RandomString(12),
		LastName   : util.RandomString(9),
		IDNumber   : util.RandomStringUser(10),
		Phone      : util.RandomPhone(),
	}
	admin, err := testQueries.CreateAdmin(context.Background(), args)
	if(err!=nil){
		log.Fatal("Failed to create Admin account:", err)
	}
	require.NotEmpty(t, admin)

	require.Equal(t, args.Email, admin.Email)
	require.Equal(t, args.Password, admin.Password)
	require.Equal(t, args.UserName, admin.UserName)
	require.Equal(t, args.FirstName, admin.FirstName)
	require.Equal(t, args.LastName, admin.LastName)
	require.Equal(t, args.IDNumber, admin.IDNumber)
	require.Equal(t, args.Phone, admin.Phone)
	
	require.NotZero(t, admin.ID)
	require.NotZero(t, admin.CreatedAt)

	return admin
}
func TestCreateAdmin(t *testing.T){
	createRandomAdmin(t)
}

func TestGetAdmin(t *testing.T){
	admin1 := createRandomAdmin(t)
	admin2, err := testQueries.GetAdmin(context.Background(), admin1.IDNumber)
	if(err!=nil){
		log.Fatal("Failed to get Admin account-------", err)
	}
	require.NotEmpty(t, admin2)

	require.Equal(t, admin1.ID, admin2.ID)
	require.Equal(t, admin1.Email, admin2.Email)
	require.Equal(t, admin1.Password, admin2.Password)
	require.Equal(t, admin1.UserName, admin2.UserName)
	require.Equal(t, admin1.FirstName, admin2.FirstName)
	require.Equal(t, admin1.LastName, admin2.LastName)
	require.Equal(t, admin1.IDNumber, admin2.IDNumber)
	require.Equal(t, admin1.Phone, admin2.Phone)
	require.WithinDuration(t, admin1.CreatedAt.Time, admin2.CreatedAt.Time, time.Second)
}

func TestUpdateAdminPhone(t *testing.T){
	admin1 := createRandomAdmin(t)
	arg := UpdateAdminPhoneParams{
		IDNumber: admin1.IDNumber,
		Phone: util.RandomPhone(),
	}
	admin2, err := testQueries.UpdateAdminPhone(context.Background(), arg)
	if(err!=nil){
		log.Fatal("Failed to update Admin account's phone number -------", err)
	}
	require.NotEmpty(t, admin2)

	require.Equal(t, admin1.ID, admin2.ID)
	require.Equal(t, admin1.Email, admin2.Email)
	require.Equal(t, admin1.Password, admin2.Password)
	require.Equal(t, admin1.UserName, admin2.UserName)
	require.Equal(t, admin1.FirstName, admin2.FirstName)
	require.Equal(t, admin1.LastName, admin2.LastName)
	require.Equal(t, admin1.IDNumber, admin2.IDNumber)
	require.Equal(t, arg.Phone, admin2.Phone)
	require.WithinDuration(t, admin1.CreatedAt.Time, admin2.CreatedAt.Time, time.Second)
}

func TestDeleteAdmin(t *testing.T){
	admin1 := createRandomAdmin(t)
	_, err := testQueries.DeleteAdmin(context.Background(), admin1.IDNumber)
	if(err!=nil){
		log.Fatal("Failed to delete Admin account-------", err)
	}

	admin2, err := testQueries.GetAdmin(context.Background(), admin1.IDNumber)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, admin2) 
}

func TestListAdmins(t *testing.T){
	for i:=0;i<10;i++ {
		createRandomAdmin(t)
	}

	arg := ListAdminsParams{
		Limit: 5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAdmins(context.Background(), arg)
	if(err!=nil){
		log.Fatal("Failed to delete Admin account-------", err)
	}
	require.Len(t, accounts, 5)
	for _, admin := range accounts{
		require.NotEmpty(t, admin)
	}
}	