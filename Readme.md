# PPL Lab Packer

This program packs and encrypts your codes for submission.

More specifically, when you run `pack` in directory `proj`, 
it will only pack the following files and directories, if they exist.

* files
  * `proj/dune-project`
  * `proj/*.opam`
* directories
  * `proj/lib`
  * `proj/bin`
  * `proj/test`

Such restrictions make uploads slim and without redundant files.

It produces a PGP message `stu-code.pgp.txt`.