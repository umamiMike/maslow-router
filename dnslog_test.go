package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var mockLog = `Jun 10 17:50:00 dnsmasq[21796]: query[A] isatap.lan from 115.34.22.160
Jun 10 17:50:00 dnsmasq[21796]: cached isatap.lan is NXDOMAIN-IPv4
Jun 10 17:50:21 dnsmasq[21796]: query[A] isatap.lan from 115.34.22.160
Jun 10 17:50:21 dnsmasq[21796]: cached isatap.lan is NXDOMAIN-IPv4
Jun 10 17:50:31 dnsmasq[21796]: query[A] isatap.lan from 115.34.22.160
Jun 10 17:50:31 dnsmasq[21796]: cached isatap.lan is NXDOMAIN-IPv4
Jun 10 17:50:37 dnsmasq[21796]: query[A] isatap.lan from 115.34.22.160
Jun 10 17:50:37 dnsmasq[21796]: cached isatap.lan is NXDOMAIN-IPv4
Jun 10 17:50:40 dnsmasq[21796]: query[A] zyx.qq.com from 115.34.22.160
Jun 10 17:50:40 dnsmasq[21796]: forwarded zyx.qq.com to 114.114.114.114
Jun 10 17:50:40 dnsmasq[21796]: forwarded zyx.qq.com to 223.5.5.5
Jun 10 17:50:40 dnsmasq[21796]: reply zyx.qq.com is 123.151.43.51
Jun 10 17:50:40 dnsmasq[21796]: reply zyx.qq.com is 183.60.62.158
Jun 10 17:50:40 dnsmasq[21796]: reply zyx.qq.com is 113.108.1.90
Jun 10 17:50:42 dnsmasq[21796]: query[A] isatap.lan from 115.34.22.160
Jun 10 17:50:42 dnsmasq[21796]: cached isatap.lan is NXDOMAIN-IPv4
Jun 10 17:50:52 dnsmasq[21796]: query[A] isatap.lan from 115.34.22.160
Jun 10 17:50:52 dnsmasq[21796]: cached isatap.lan is NXDOMAIN-IPv4
Jun 10 17:50:58 dnsmasq[21796]: query[A] ic.wps.cn from 115.34.22.160 `

func TestReadAndParseDNSNoFile(t *testing.T) {
	//make a file in the file system (remember to remove it)
	nonexistantfile := "nonexistant"
	_, err := ReadDNS(nonexistantfile)
	if err != nil {
		expected := fmt.Sprintf("open %v: no such file or directory", nonexistantfile)
		//t.Errorf("%v", err)
		assert.Equal(t, expected, err.Error(), "the error message is not correct")
	}
}

func TestReadAndParseDNS(t *testing.T) {
	testfile := "./testdata/logs/dnsmasq.log"
	ret, _ := ReadDNS(testfile)
	fmt.Println(len(ret["pixiedust.buzzfeed.com"]))
}

//testing parseLog used in the readAndParseDns fn
func TestParseLog(t *testing.T) {

	mocklinefromfile := `Dec 31 16:01:44 dnsmasq[1301]: reading /tmp/resolv.dnsmasq `
	key, val, err := parseLog(mocklinefromfile)
	fmt.Printf("the val of key is: %v", key)

	assert.Nil(t, err, "there should be no error")
	assert.Empty(t, key, "the key should be empty")
	assert.Empty(t, val, "the val should be empty")
}

//test parseLog to deal with encountering CNAMEs
func TestParseLogCname(t *testing.T) {

	mocklinefromfile := `Jan 13 19:49:44 dnsmasq[1301]: reply raw.githubusercontent.com is <CNAME>`
	key, val, err := parseLog(mocklinefromfile)
	fmt.Printf("the val of key is: %v", key)

	assert.Nil(t, err, "there should be no error")
	assert.Empty(t, key, "the key should be empty")
	assert.Empty(t, val, "the val should be empty")
}

func TestParseLogBadLine(t *testing.T) {
	badline := `I have no colon, so I am not logged line`
	_, _, err := parseLog(badline)
	assert.Error(t, err, "the line should not be readable")
}

func TestParseLogReply(t *testing.T) {
	//edge case of reply
	line := `Dec 31 16:01:45 dnsmasq[1301]: reply 2.pool.ntp.org is 45.127.112.2`

	key, val, err := parseLog(line)

	assert.Nil(t, err, "if reply exists")
	assert.Equal(t, key, "2.pool.ntp.org", "key is a dns site")
	assert.Equal(t, val, "45.127.112.2", "val is an ip")
}
func BenchmarkReadAndParse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testfile := "./testdata/logs/dnsmasq.log"
		ReadDNS(testfile)
	}

}
func BenchmarkParseLog(b *testing.B) {
	for i := 0; i < b.N; i++ {
		line := `Dec 31 16:01:45 dnsmasq[1301]: reply 2.pool.ntp.org is 45.127.112.2`
		parseLog(line)
	}

}
