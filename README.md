GoSDDL (Security Descriptor Definition Language)
------------------------------

Converter from SDDL-string to user-friendly JSON. SDDL consist of four part: Owner, Primary Group, DACL, SACL.

This converter works with two mode:
1) Direct
2) API

## Installing
To start using gosddl, install Go and run go get:

```
$ go get -u github.com/MonaxGT/gosddl
```

## Direct usage example

```
go run gosddl.go "D:(A;;GA;;;S-1-5-21-111111111-1111111111-1111111111-11111)(A;;GA;;;SY)(A;;GXGR;;;S-1-5-5-1-1111111111)(A;;GA;;;BA)"

{"owner":"","primary":"","dacl":[{"accountsid":"S-1-5-21-111111111-1111111111-1111111111-11111","aceType":"ACCESS ALLOWED","aceflags":[""],"rights":["GENERIC_ALL"],"objectguid":"","InheritObjectGuid":""},{"accountsid":"Local system","aceType":"ACCESS ALLOWED","aceflags":[""],"rights":["GENERIC_ALL"],"objectguid":"","InheritObjectGuid":""},{"accountsid":"S-1-5-5-1-1111111111","aceType":"ACCESS ALLOWED","aceflags":[""],"rights":["GENERIC_EXECUTE","GENERIC_READ"],"objectguid":"","InheritObjectGuid":""},{"accountsid":"Built-in administrators","aceType":"ACCESS ALLOWED","aceflags":[""],"rights":["GENERIC_ALL"],"objectguid":"","InheritObjectGuid":""}],"daclInheritFlags":null,"sacl":null,"saclInheritFlags":null}

```

## API usage example

```
go run gosddl.go -api

curl 'http://127.0.0.1:8000/sddl/D:(A;;GA;;;S-1-5-21-111111111-1111111111-1111111111-11111)(A;;GA;;;SY)(A;;GXGR;;;S-1-5-5-1-1111111111)(A;;GA;;;BA)'
{"owner":"","primary":"","dacl":[{"accountsid":"S-1-5-21-111111111-1111111111-1111111111-11111","aceType":"ACCESS ALLOWED","aceflags":[""],"rights":["GENERIC_ALL"],"objectguid":"","InheritObjectGuid":""},{"accountsid":"Local system","aceType":"ACCESS ALLOWED","aceflags":[""],"rights":["GENERIC_ALL"],"objectguid":"","InheritObjectGuid":""},{"accountsid":"S-1-5-5-1-1111111111","aceType":"ACCESS ALLOWED","aceflags":[""],"rights":["GENERIC_EXECUTE","GENERIC_READ"],"objectguid":"","InheritObjectGuid":""},{"accountsid":"Built-in administrators","aceType":"ACCESS ALLOWED","aceflags":[""],"rights":["GENERIC_ALL"],"objectguid":"","InheritObjectGuid":""}],"daclInheritFlags":null,"sacl":null,"saclInheritFlags":null}
```

## Additionally you can use Docker:

```
docker build -t gosddl .

docker run -d -p 8000:8000 gosddl -api
docker run --rm -it gosddl "O:BAG:SYD:(D;;GA;;;AN)(D;;GA;;;BG)(A;;GA;;;SY)(A;;GA;;;BA)S:ARAI(AU;SAFA;DCLCRPCRSDWDWO;;;WD)"
```

Links:

[Source](https://docs.microsoft.com/en-us/windows/desktop/secauthz/security-descriptor-definition-language)