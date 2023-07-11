package util

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func GetSplit(stringvalues, strseparate, strjoin string) string { //ids := "546,545,2171"
	idsslice := strings.Split(stringvalues, strseparate)

	//fmt.Printf("idsslice : %v", idsslice)
	idsstr := strings.Join(idsslice, strjoin)
	//fmt.Printf("idsstr : %v", idsstr)
	return idsstr
}

// GetSliceInt64 slice of int64 for given string comma separated
func GetSliceInt64(infovalueidsstr string) []int64 {
	idstringSlice := strings.Split(infovalueidsstr, ",")
	idsslice := []int64{}
	for _, v := range idstringSlice {
		//log.Println(i, v) //log.Printf("Value %d of type %T", v, v)
		n, err := strconv.ParseInt(v, 10, 64)
		if err == nil {
			log.Printf("%d of type %T", n, n)
		}
		idsslice = append(idsslice, n)
	}
	return idsslice
}

func GetBucketNumberFileName(id, uniqueIdName string) (string, error) {
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Printf("Convert failed id : %v, error: %v", id, err)
		return "", err
	}

	bucket := i % 10000
	bucketName := fmt.Sprintf("%v/%v%s", bucket, uniqueIdName, ".json")
	return bucketName, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	fmt.Println("HERE IN password:", hash)
	fmt.Println("HERE DB password:", password)
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
