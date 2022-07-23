package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/big"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"
)

// Comes from https://stackoverflow.com/a/38915828
// tested with:
// - Ec2 Instances like i-b9b4ffaa
// - AMI like ami-dbcf88b1
// - Volumes like vol-e97db305
const awsIDRegex = "(?i)\\b[a-z]+-[a-z0-9]{8,}"
const awsAccountIDRegex = "\\b[0-9]{12}"

func main() {
	// Let's set a random salt for tis command execution
	rand.Seed(time.Now().UnixNano())
	salt := fmt.Sprint(rand.Intn(100))

	// Start reding Std in
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		line = ScrambleAWSResourceID(line, salt)
		line = ScrambleAWSAccountID(line, salt)
		fmt.Println(line)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

// ScrambleAWSResourceID scans a line of text and replaces all the aws ids with new ones
// keeping their uniqueness characteristics using md5 + salt approach.
// the salt is generated for each run of the command in order to keep the
// 1:1 correspondence between ids in the passed file
func ScrambleAWSResourceID(line, salt string) string {
	r, _ := regexp.Compile(awsIDRegex)
	ids := r.FindAllString(line, -1)
	scrambledLine := line
	for _, id := range ids {
		// A generic AWS resource id like "vol-e97db305"
		// is composed by [resourceType]-[hexSuffix]:
		// resourceType = "vol"
		// hexSuffix = "e97db305"
		resourceType := strings.Split(id, "-")[0]
		hexSuffix := strings.Split(id, "-")[1]
		md5 := GetMD5Hash(hexSuffix + salt)
		// cut the salted MD5 to the same length as the original hexSuffix
		hexSuffixScrambled := string(md5[0:len(hexSuffix)])
		scrambledId := strings.Join([]string{resourceType, hexSuffixScrambled}, "-")
		scrambledLine = strings.Replace(scrambledLine, id, scrambledId, -1)
	}

	return scrambledLine
}

// ScrambleAWSAccountID scrambles AWS account id
func ScrambleAWSAccountID(line, salt string) string {
	r, _ := regexp.Compile(awsAccountIDRegex)
	ids := r.FindAllString(line, -1)
	scrambledLine := line
	for _, id := range ids {
		md5 := GetMD5Hash(id + salt)
		bi := big.NewInt(0)
		bi.SetString(md5, 16)
		scrambledId := bi.String()[0:len(id)]
		scrambledLine = strings.Replace(scrambledLine, id, scrambledId, -1)
	}

	return scrambledLine
}

// GetMD5Hash returns md5 hash of the passed string
// in a string format containing an hex number
func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
