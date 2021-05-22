/**
 * @Author: frankstar
 * @Email: frankstar007@163.com
 * @Date: 2020-07-22 15:24
 * @Project : go-improve
 * @Desc: go 连接数据库
*/

package models

import (
	"crypto/rand"
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"
)

var Db *sql.DB

func init()  {
	var err error
	Db, err = sql.Open("mysql", "root:frankstar@tcp(127.0.0.1:3306)/frank_study?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}
}

func createUUID() (uuid string)  {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		log.Fatal("cannot generate uuid")
	}
	// 0x40 is reserved variant from RFC 4122
	u[8] = (u[8] | 0x40) & 0x7F
	// Set the four most significant bits (bits 12 through 15) of the
	// time_hi_and_version field to the 4-bit version number.
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return

}

func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte (plaintext)))
	return

}
