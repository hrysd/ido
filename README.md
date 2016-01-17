# Ido

`ido` is cli tool that post stdout [Idobata](https://idobata.io)

## Usage

You need to store hook endpoint in `~/.ido` like following

```json
[
    {
        "Name": "hrysd",
        "Token": "[HOOK_ENDPOINT]"
    },
    {
        "Name": "hoi",
        "Token": "[HOOK_ENDPOINT]" 
    }
]
```

then just redirect stdout to `ido` command.

```
$ cat ~/.vimrc | ido hoi
```