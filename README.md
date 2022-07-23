# `hidi`: Command line tool to obfuscate AWS files of any sort


## Usage:
```bash
go run . < original.txt > scrambled.txt 
```

## Why?

I need a tool that allows me to safely publish any data coming from an AWS
accounts without the worry of leaking sensitive information

List of AWS data containing potential sensitive information:

- billing files like the
  [AWS CUR](https://docs.aws.amazon.com/cur/latest/userguide/what-is-cur.html)
- [placebo](https://github.com/garnaat/placebo) files used for testing, as they
  are recordings of actual api calls to AWS, and could leak account sensitive
  informations
- ?

### How AWS sensitive data looks like?

Even if it seems things like

- Account ID
- Resources ID (like `i-10bf43c2be699e77a`, `ami-b0c6444b`, ...) are NOT
  considered sensitive (read
  [here](https://www.lastweekinaws.com/blog/are-aws-account-ids-sensitive-information/)
  (2022-02-16) )

I would prefer to keep certain things not disclosed to avoid issues and respect
sone sort of "least disclosed" principle (read
[here](https://rhinosecuritylabs.com/aws/assume-worst-aws-assume-role-enumeration/)
(2018-08-29)) and let YOU decide if it's necessary to share certain info with
the rest of the world.

## How?

The idea is to have a brute file parser that scans files line by line applying
heuristics (probably regexes) to find sensitive data and scramble it,
maintaining it compatible with AWS standards.


