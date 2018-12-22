GoSDDL (Security Descriptor Definition Language)
===============================================
[![Build Status](https://travis-ci.org/MonaxGT/gosddl.svg?branch=master)](https://travis-ci.org/MonaxGT/gosddl)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/70d6bf54dd2547d894ee7ba7a9247285)](https://app.codacy.com/app/MonaxGT/gosddl?utm_source=github.com&utm_medium=referral&utm_content=MonaxGT/gosddl&utm_campaign=Badge_Grade_Dashboard)
[![Maintainability](https://api.codeclimate.com/v1/badges/69e05e119408b9f830d4/maintainability)](https://codeclimate.com/github/MonaxGT/gosddl/maintainability)

Converter from SDDL-string to user-friendly JSON. SDDL consist of four part: Owner, Primary Group, DACL, SACL.

This converter works with two mode:
1) Direct
2) API

Installing
------------------------------------------------

To start using gosddl, install Go and run go get:

```go
$ go get -u github.com/MonaxGT/gosddl
```

Direct usage example
------------------------------------------------

```go
go run gosddl.go "D:(A;;GA;;;S-1-5-21-111111111-1111111111-1111111111-11111)(A;;GA;;;SY)(A;;GXGR;;;S-1-5-5-1-1111111111)(A;;GA;;;BA)"

{"owner":"","primary":"","dacl":[{"accountsid":"S-1-5-21-111111111-1111111111-1111111111-11111","aceType":"ACCESS ALLOWED","aceflags":[""],"rights":["GENERIC_ALL"],"objectguid":"","InheritObjectGuid":""},{"accountsid":"Local system","aceType":"ACCESS ALLOWED","aceflags":[""],"rights":["GENERIC_ALL"],"objectguid":"","InheritObjectGuid":""},{"accountsid":"S-1-5-5-1-1111111111","aceType":"ACCESS ALLOWED","aceflags":[""],"rights":["GENERIC_EXECUTE","GENERIC_READ"],"objectguid":"","InheritObjectGuid":""},{"accountsid":"Built-in administrators","aceType":"ACCESS ALLOWED","aceflags":[""],"rights":["GENERIC_ALL"],"objectguid":"","InheritObjectGuid":""}],"daclInheritFlags":null,"sacl":null,"saclInheritFlags":null}
```

API usage example
------------------------------------------------

```go
go run gosddl.go -api

curl 'http://127.0.0.1:8000/sddl/D:(A;;GA;;;S-1-5-21-111111111-1111111111-1111111111-11111)(A;;GA;;;SY)(A;;GXGR;;;S-1-5-5-1-1111111111)(A;;GA;;;BA)'
{"owner":"","primary":"","dacl":[{"accountsid":"S-1-5-21-111111111-1111111111-1111111111-11111","aceType":"ACCESS ALLOWED","aceflags":[""],"rights":["GENERIC_ALL"],"objectguid":"","InheritObjectGuid":""},{"accountsid":"Local system","aceType":"ACCESS ALLOWED","aceflags":[""],"rights":["GENERIC_ALL"],"objectguid":"","InheritObjectGuid":""},{"accountsid":"S-1-5-5-1-1111111111","aceType":"ACCESS ALLOWED","aceflags":[""],"rights":["GENERIC_EXECUTE","GENERIC_READ"],"objectguid":"","InheritObjectGuid":""},{"accountsid":"Built-in administrators","aceType":"ACCESS ALLOWED","aceflags":[""],"rights":["GENERIC_ALL"],"objectguid":"","InheritObjectGuid":""}],"daclInheritFlags":null,"sacl":null,"saclInheritFlags":null}
```

Additionally you can use Docker
------------------------------------------------

```docker
docker build -t gosddl .
docker run -d -p 8000:8000 gosddl -api
docker run --rm -it gosddl "O:BAG:SYD:(D;;GA;;;AN)(D;;GA;;;BG)(A;;GA;;;SY)(A;;GA;;;BA)S:ARAI(AU;SAFA;DCLCRPCRSDWDWO;;;WD)"
```

Links:

[Source](https://docs.microsoft.com/en-us/windows/desktop/secauthz/security-descriptor-definition-language)
