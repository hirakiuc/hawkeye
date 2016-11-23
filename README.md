# hawkeye

`hawkeye` is a cli tool to list file paths from the current directory with considering .gitignore file if exists.

# HowToUse

Just call `hawkeye` command like this.

```
$ hawkeye
```

## What purpose of this command ?

I want to use this command in vim, like this.

```
function! PecoOpen()
  for filename in split(system("hawkeye | peco"), "\n")
    execute "e" filename
  endfor
endfunction
nnoremap <Leader>op :call PecoOpen()<CR>
```

![HowToUse](https://github.com/hirakiuc/hawkeye/blob/master/screen.gif)

# TODO

See [Issues](https://github.com/hirakiuc/hawkeye/issues).

# LICENSE

MIT

