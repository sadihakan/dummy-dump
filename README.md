# DummyDump

Simple Database Backup Tool (zero dependencies)

## Getting Started

### Requirements

* Go >= 1.15

### Export

* How to run the program

```
go run ./cmd -source=<postgres/mysql> -export -user=<User Name> -db=<Database Name>
```
## Import

* How to run the program

```
go run ./cmd -source <postgres/mysql> -import -user=<User Name> -file=<File Path> 
```

## Usage

- *import*  
Flag: -import  
Type: bool  
Default: false  

- *export*  
Flag: -export  
Type: bool  
Default: false 

- *source*  
Flag: -source  
Type: string  
Default: null  
Values: **mysql**, **postgres** 

- *user*  
Flag: -user  
Type: string  
Default: null 

- *path*  
Flag: -path  
Type: string  
Default: null 

- *db*  
Flag: -db  
Type: string  
Default: null

- *binaryPath*  
Flag: -binaryPath  
Type: string  
Default: null  

Contributors names and contact info

ex. [@sadihakan](https://github.com/sadihakan/)    
ex. [@onurcevik](https://github.com/onurcevik/)



## License

This project is licensed under the sadihakan License - see the LICENSE.md file for details


