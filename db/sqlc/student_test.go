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


func createRandomStudent(t *testing.T) Student{
	mentor := createRandomMentor(t)
	args := CreateStudentParams{
		Email      : util.RandomEmail(),
		Password   : util.RandomPassword(18),
		UserName   : util.RandomStringUser(12),
		FirstName  : util.RandomString(12),
		LastName   : util.RandomString(9),
		RollNumber : util.RandomString(10),
		Stream	   : util.RandomStream(),
		Section    : util.RandomString(1),
		Course     : util.RandomBranch(),
		Phone      : util.RandomPhone(),
		Mentor     : sql.NullInt64{Int64: mentor.ID, Valid: mentor.ID!=0},
	}
	student, err := testQueries.CreateStudent(context.Background(), args)
	if(err!=nil){
		log.Fatal("Failed to CreateStudent account:", err)
	}
	require.NotEmpty(t, student)

	require.Equal(t, args.Email, student.Email)
	require.Equal(t, args.Password, student.Password)
	require.Equal(t, args.UserName, student.UserName)
	require.Equal(t, args.FirstName, student.FirstName)
	require.Equal(t, args.LastName, student.LastName)
	require.Equal(t, args.RollNumber, student.RollNumber)
	require.Equal(t, args.Stream, student.Stream)
	require.Equal(t, args.Section, student.Section)
	require.Equal(t, args.Course, student.Course)
	require.Equal(t, args.Phone, student.Phone)
	require.Equal(t, args.Mentor, student.Mentor)
	
	require.NotZero(t, student.ID)
	require.NotZero(t, student.CreatedAt)

	return student
}
func TestCreateStudent(t *testing.T){
	createRandomStudent(t)
}

func TestGetStudent(t *testing.T){
	student1 := createRandomStudent(t)
	student2, err := testQueries.GetStudent(context.Background(), student1.RollNumber)
	if(err!=nil){
		log.Fatal("Failed to get student account-------", err)
	}
	require.NotEmpty(t, student2)

	require.Equal(t, student1.ID, student2.ID)
	require.Equal(t, student1.Email, student2.Email)
	require.Equal(t, student1.Password, student2.Password)
	require.Equal(t, student1.UserName, student2.UserName)
	require.Equal(t, student1.FirstName, student2.FirstName)
	require.Equal(t, student1.LastName, student2.LastName)
	require.Equal(t, student1.RollNumber, student2.RollNumber)
	require.Equal(t, student1.Stream, student2.Stream)
	require.Equal(t, student1.Section, student2.Section)
	require.Equal(t, student1.Course, student2.Course)
	require.Equal(t, student1.Phone, student2.Phone)
	require.Equal(t, student1.Mentor, student2.Mentor)
	require.WithinDuration(t, student1.CreatedAt.Time, student2.CreatedAt.Time, time.Second)
}

// func TestUpdateAdminPhone(t *testing.T){
// 	admin1 := createRandomStudent(t)
// 	arg := UpdateAdminPhoneParams{
// 		IDNumber: admin1.IDNumber,
// 		Phone: util.RandomPhone(),
// 	}
// 	admin2, err := testQueries.UpdateAdminPhone(context.Background(), arg)
// 	if(err!=nil){
// 		log.Fatal("Failed to update Admin account's phone number -------", err)
// 	}
// 	require.NotEmpty(t, admin2)

// 	require.Equal(t, admin1.ID, admin2.ID)
// 	require.Equal(t, admin1.Email, admin2.Email)
// 	require.Equal(t, admin1.Password, admin2.Password)
// 	require.Equal(t, admin1.UserName, admin2.UserName)
// 	require.Equal(t, admin1.FirstName, admin2.FirstName)
// 	require.Equal(t, admin1.LastName, admin2.LastName)
// 	require.Equal(t, admin1.IDNumber, admin2.IDNumber)
// 	require.Equal(t, arg.Phone, admin2.Phone)
// 	require.WithinDuration(t, admin1.CreatedAt.Time, admin2.CreatedAt.Time, time.Second)
// }

func TestDeleteStudent(t *testing.T){
	student1 := createRandomStudent(t)
	_, err := testQueries.DeleteStudent(context.Background(), student1.RollNumber)
	if(err!=nil){
		log.Fatal("Failed to DeleteStudent account-------", err)
	}

	student2, err := testQueries.GetAdmin(context.Background(), student1.RollNumber)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, student2) 
}

func TestListStudents(t *testing.T){
	for i:=0;i<10;i++ {
		createRandomStudent(t)
	}

	arg := ListStudentsParams{
		Limit: 5,
		Offset: 5,
	}

	accounts, err := testQueries.ListStudents(context.Background(), arg)
	if(err!=nil){
		log.Fatal("Failed to get all student account-------", err)
	}
	require.Len(t, accounts, 5)
	for _, student := range accounts{
		require.NotEmpty(t, student)
	}
}	