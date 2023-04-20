# go-file-hash-helper

Quickly generate and verify file sha256

## UUsage
1. Download the latest release from the [Releases](https://github.com/dreamguxiang/go-file-hash-helper/releases).
2. Launch the executable and follow the instructions
```
COMMANDS:
   Generate, g  Generate [filename].fs256
       OPTIONS:
           --path value, -p value  specify path (default: ./)
           --help, -h              show help

   Verify, v    Verify [filename].fs256
       OPTIONS:
           --path value, -p value    specify path (default: ./)
           --remove value, -r value  remove .fs256 file after verify (default: false)
           --help, -h                show help
    
   help, h      Shows a list of commands or help for one command

EXAMPLE:
  .\go_file_hash_helper.exe Generate -p <Path>  ----- Generate hash file

  .\go_file_hash_helper.exe Verify -p <Path> -r <IsRemove> ----- Verify file 

```

