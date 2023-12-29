package util

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-.@_"
func init(){
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

func RandonInt(min, max int64) int64{
	return min+rand.Int63n(max-min+1)
}

func RandomString(n int) string{
	var sb strings.Builder
	k:= len(alphabet)-4
	for i:=0;i<n;i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomStringUser(n int) string{
	var sb strings.Builder
	k:= len(alphabet)-2
	for i:=0;i<n;i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomPassword(n int) string{
	var sb strings.Builder
	k:= len(alphabet)
	for i:=0;i<n;i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomEmail() string{
	var addon = "@gmail.com"
	return RandomString(9)+addon
}

func RandomPhone() string{
	num := RandonInt(1111111111, 9999999999)
	return strconv.Itoa(int(num))
}

func RandomBranch() string{
	branches := []string{"CSE", "ECE", "ME", "EEE", "CE"}
	return branches[rand.Intn(len(branches))]
}

func RandomStream() string{
	branches := []string{"BTECH", "BBA", "MBA", "MTECH", "CE"}
	return branches[rand.Intn(len(branches))]
}