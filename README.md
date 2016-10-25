##src
####stringutil.go
  - library/package
  - this is a package being used in the hello.go
  - to test if the package compiles with `go build`
      - `go build path/to/file/directory`
      - if successful nothing will print
      - To create outputs, you have to type `go install` which will place the package object inside the pkg directory of the workspace
  - then, the new built package file will be created automatically in the pkg directory, and available for use and import in files in the src folder

##pkg
####Creating Packages
  - package names:
  `package [packageName]` must be the first line of the file
      - this will be the default name for imports
  *importing packages*
    - package name is the last element of the import path
    - executable commands must always use package `main`
    - import paths must be unique, but package names do not have to do 
