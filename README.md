patign
======

Align / format streams by pattern.  "patign" is a portmanteau of "pattern" and "align".

# install

```
go get github.com/nathanleclaire/patign
```

# usage

Right now only piping from <STDIN> is supported, though that will change soon.

Imagine in file `test.file` you have something like :

```sh
"aws_resource" "droplet" {
    "ip" = "54.23.62.11"
    "foo" = "bar"
    "quuuuuuuuuuuuuux" = "spam"
    "amazon_secret_key" = "ASDFASDFASDFASDFASDFASDF"
    "amazon_id" = "asdffad,fdsa?"
}
```

(this example inspired by [terraform](http://terraform.io) configuration)

Wouldn't it be nice if all those `=` lined up?

Well, with `patign`, they can.

```sh
$ cat test.file | patign "=" 
"aws_resource" "droplet" {
    "foo"               = "bar"
    "quuuuuuuuuuuuuux"  = "spam"
    "amazon_secret_key" = "ASDFASDFASDFASDFASDFASDF"
    "amazon_id"         = "asdffad,fdsa?"
}
```

Another use case that I ran into the other day: let's say you have renamed a bunch of files in git.

```sh
$ git status
# On branch master
# Your branch is ahead of 'origin/master' by 1 commit.
#   (use "git push" to publish your local commits)
#
# Changes to be committed:
#   (use "git reset HEAD <file>..." to unstage)
#
#   renamed:    test/a.txt -> test/b.txt
#   renamed:    test/barquux.java -> test/c.java
#   renamed:    test/foo.java -> test/d.java
#   renamed:    test/foobarwidgetfactory.java -> test/e.java
#   renamed:    test/randomlongname.java -> test/f.java
#   renamed:    test/somethingelesethatsstilldifferent.javca -> test/g.java
```

It's hard to parse the modified files.  Can't they be more nicely formatted?  With `patign`, they can.

```sh
$ git status | ./patign "-"
# On branch master
# Your branch is ahead of 'origin/master' by 1 commit.
#   (use "git push" to publish your local commits)
#
# Changes to be committed:
#   (use "git reset HEAD <file>..." to unstage)
#
#   renamed:    test/barquux.java                            -> test/c.java
#   renamed:    test/foo.java                                -> test/d.java
#   renamed:    test/foobarwidgetfactory.java                -> test/e.java
#   renamed:    test/randomlongname.java                     -> test/f.java
#   renamed:    test/somethingelesethatsstilldifferent.javca -> test/g.java
```

# disclaimer

Naturally it's still early days for `patign` and so there are rampant bugs, issues, etc.  I'd love it if you'd contribute!

# license

MIT
