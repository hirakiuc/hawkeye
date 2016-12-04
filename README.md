# hawkeye

`hawkeye` is a cli tool to list file paths from the current directory with considering .gitignore file if exists.

## Dependency

`hawkeye` depends on [libgit2/git2go](https://github.com/libgit2/git2go) library to recognize gitignore file.

So you should install `libgit2` before install `hawkeye`.

NOTE: Now, I use `hawkeye` with `libgit2` v0.24.3 on MacOS. :smile:

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

