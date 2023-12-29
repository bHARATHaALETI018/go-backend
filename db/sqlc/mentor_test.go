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


func createRandomMentor(t *testing.T) Mentor{
	args := CreateMentorParams{
		Email      : util.RandomEmail(),
		Password   : util.RandomPassword(18),
		UserName   : util.RandomStringUser(12),
		FirstName  : util.RandomString(12),
		LastName   : util.RandomString(9),
		Phone      : util.RandomPhone(),
		IDNumber   : util.RandomStringUser(10),
	}
	account, err := testQueries.CreateMentor(context.Background(), args)
	if(err!=nil){
		log.Fatal("Failed to create Mentor account --- ", err)
	}
	require.NotEmpty(t, account)

	require.Equal(t, args.Email, account.Email)
	require.Equal(t, args.Password, account.Password)
	require.Equal(t, args.UserName, account.UserName)
	require.Equal(t, args.FirstName, account.FirstName)
	require.Equal(t, args.LastName, account.LastName)
	require.Equal(t, args.IDNumber, account.IDNumber)
	require.Equal(t, args.Phone, account.Phone)
	
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}
func TestCreateMentor(t *testing.T){
	createRandomMentor(t)
}

func TestGetMentor(t *testing.T){
	account1 := createRandomMentor(t)
	account2, err := testQueries.GetMentor(context.Background(), account1.IDNumber)
	if(err!=nil){
		log.Fatal("Failed to get mentor account-------", err)
	}
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Email, account2.Email)
	require.Equal(t, account1.Password, account2.Password)
	require.Equal(t, account1.UserName, account2.UserName)
	require.Equal(t, account1.FirstName, account2.FirstName)
	require.Equal(t, account1.LastName, account2.LastName)
	require.Equal(t, account1.IDNumber, account2.IDNumber)
	require.Equal(t, account1.Phone, account2.Phone)
	require.WithinDuration(t, account1.CreatedAt.Time, account2.CreatedAt.Time, time.Second)
}

func TestUpdateMentorPhone(t *testing.T){
	account1 := createRandomMentor(t)
	arg := UpdateMentorPhoneParams{
		IDNumber: account1.IDNumber,
		Phone: util.RandomPhone(),
	}
	account2, err := testQueries.UpdateMentorPhone(context.Background(), arg)
	if(err!=nil){
		log.Fatal("Failed to update mentor account's phone number -------", err)
	}
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Email, account2.Email)
	require.Equal(t, account1.Password, account2.Password)
	require.Equal(t, account1.UserName, account2.UserName)
	require.Equal(t, account1.FirstName, account2.FirstName)
	require.Equal(t, account1.LastName, account2.LastName)
	require.Equal(t, account1.IDNumber, account2.IDNumber)
	require.Equal(t, arg.Phone, account2.Phone)
	require.WithinDuration(t, account1.CreatedAt.Time, account2.CreatedAt.Time, time.Second)
}

func TestDeleteMentor(t *testing.T){
	account1 := createRandomMentor(t)
	_, err := testQueries.DeleteMentor(context.Background(), account1.IDNumber)
	if(err!=nil){
		log.Fatal("Failed to delete mentor account-------", err)
	}

	account2, err := testQueries.GetMentor(context.Background(), account1.IDNumber)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account2) 
}

func TestListMentors(t *testing.T){
	for i:=0;i<10;i++ {
		createRandomMentor(t)
	}

	arg := ListMentorsParams{
		Limit: 5,
		Offset: 5,
	}

	accounts, err := testQueries.ListMentors(context.Background(), arg)
	if(err!=nil){
		log.Fatal("Failed to delete Admin account-------", err)
	}
	require.Len(t, accounts, 5)
	for _, account := range accounts{
		require.NotEmpty(t, account)
	}
}	